// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package releaseTable

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/pkg/errors"

	"github.com/erda-project/erda-infra/base/servicehub"
	"github.com/erda-project/erda-infra/providers/component-protocol/cptype"
	"github.com/erda-project/erda-infra/providers/component-protocol/utils/cputil"
	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/bundle"
	cmpTypes "github.com/erda-project/erda/modules/cmp/component-protocol/types"
	"github.com/erda-project/erda/modules/dop/component-protocol/components/release-manage/access"
	"github.com/erda-project/erda/modules/openapi/component-protocol/components/base"
)

func init() {
	base.InitProviderWithCreator("release-manage", "releaseTable", func() servicehub.Provider {
		return &ComponentReleaseTable{}
	})
}

func (r *ComponentReleaseTable) Render(ctx context.Context, component *cptype.Component, _ cptype.Scenario,
	event cptype.ComponentEvent, gs *cptype.GlobalStateData) error {
	r.InitComponent(ctx)
	if err := r.GenComponentState(component); err != nil {
		return errors.Errorf("failed to gen release table component state, %v", err)
	}

	switch event.Operation {
	case cptype.InitializeOperation:
		r.State.PageNo = 1
		r.State.PageSize = 10
		if err := r.DecodeURLQuery(); err != nil {
			return errors.Errorf("failed to deocode url query for release table component, %v", err)
		}
	case cptype.RenderingOperation, "changePageSize", "changeSort":
		r.State.PageNo = 1
	case "batchSubmit":
		meta, ok := event.OperationData["meta"].(map[string]interface{})
		if !ok {
			return errors.Errorf("invalid meta in operationData")
		}
		operationType, ok := meta["type"].(string)
		if !ok {
			return errors.Errorf("invalid type in meta")
		}
		switch operationType {
		case "formal":
			if err := r.formalReleases(r.State.SelectedRowKeys...); err != nil {
				return errors.Errorf("failed to formal release, %v", err)
			}
		case "delete":
			if err := r.deleteReleases(r.State.SelectedRowKeys...); err != nil {
				return errors.Errorf("failed to delete release, %v", err)
			}
		}
	case "formal":
		id, err := getReleaseID(event.OperationData)
		if err != nil {
			return err
		}
		if err = r.formalReleases(id); err != nil {
			return errors.Errorf("failed to formal release, %v", err)
		}
	case "delete":
		id, err := getReleaseID(event.OperationData)
		if err != nil {
			return err
		}
		if err = r.deleteReleases(id); err != nil {
			return errors.Errorf("failed to delete release, %v", err)
		}
	}
	if err := r.RenderTable(); err != nil {
		return err
	}
	if err := r.EncodeURLQuery(); err != nil {
		return errors.Errorf("failed to encode url query for release table component, %v", err)
	}
	r.SetComponentValue()
	r.Transfer(component)
	return nil
}

func (r *ComponentReleaseTable) InitComponent(ctx context.Context) {
	sdk := cputil.SDK(ctx)
	r.sdk = sdk
	bdl := ctx.Value(cmpTypes.GlobalCtxKeyBundle).(*bundle.Bundle)
	r.bdl = bdl
}

func (r *ComponentReleaseTable) GenComponentState(component *cptype.Component) error {
	if component == nil || component.State == nil {
		return nil
	}
	var state State
	jsonData, err := json.Marshal(component.State)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(jsonData, &state); err != nil {
		return err
	}
	r.State = state
	return nil
}

func (r *ComponentReleaseTable) DecodeURLQuery() error {
	query, ok := r.sdk.InParams["releaseTable__urlQuery"].(string)
	if !ok {
		return nil
	}
	decoded, err := base64.StdEncoding.DecodeString(query)
	if err != nil {
		return err
	}
	urlQuery := make(map[string]interface{})
	if err := json.Unmarshal(decoded, &urlQuery); err != nil {
		return err
	}
	r.State.PageNo = int64(urlQuery["pageNo"].(float64))
	r.State.PageSize = int64(urlQuery["pageSize"].(float64))
	sorter := urlQuery["sorterData"].(map[string]interface{})
	r.State.Sorter.Field, _ = sorter["field"].(string)
	r.State.Sorter.Order, _ = sorter["order"].(string)

	return nil
}

func (r *ComponentReleaseTable) EncodeURLQuery() error {
	urlQuery := make(map[string]interface{})
	urlQuery["pageNo"] = r.State.PageNo
	urlQuery["pageSize"] = r.State.PageSize
	urlQuery["sorterData"] = r.State.Sorter
	jsonData, err := json.Marshal(urlQuery)
	if err != nil {
		return err
	}

	encode := base64.StdEncoding.EncodeToString(jsonData)
	r.State.ReleaseTableURLQuery = encode
	return nil
}

func (r *ComponentReleaseTable) RenderTable() error {
	userID := r.sdk.Identity.UserID
	orgID := r.sdk.Identity.OrgID
	projectID := r.State.ProjectID
	hasAccess, err := access.HasReadAccess(r.bdl, userID, uint64(projectID))
	if err != nil {
		return errors.Errorf("failed to check access, %v", err)
	}
	if !hasAccess {
		return errors.Errorf("Access denied")
	}

	var startTime, endTime int64 = 0, 0
	if len(r.State.FilterValues.CreatedAtStartEnd) == 2 {
		startTime = r.State.FilterValues.CreatedAtStartEnd[0]
		endTime = r.State.FilterValues.CreatedAtStartEnd[1]
	}

	var appIDs []string
	if r.State.ApplicationID > 0 {
		appIDs = append(appIDs, strconv.FormatInt(r.State.ApplicationID, 10))
	}

	isStable := true

	order := "DESC"
	if r.State.Sorter.Order == "ascend" {
		order = "ASC"
	}

	releaseResp, err := r.bdl.ListReleases(apistructs.ReleaseListRequest{
		Branch:           r.State.FilterValues.BranchID,
		IsStable:         &isStable,
		IsFormal:         &r.State.IsFormal,
		IsProjectRelease: &r.State.IsProjectRelease,
		UserID:           r.State.FilterValues.UserIDs,
		Version:          r.State.VersionValues.Version,
		CommitID:         r.State.FilterValues.CommitID,
		ApplicationID:    appIDs,
		ProjectID:        projectID,
		StartTime:        startTime,
		EndTime:          endTime,
		PageSize:         r.State.PageSize,
		PageNum:          r.State.PageNo,
		OrderBy:          r.State.Sorter.Field,
		Order:            order,
	})
	if err != nil {
		return errors.Errorf("failed to list releases, %v", err)
	}

	r.State.Total = releaseResp.Total

	org, err := r.bdl.GetOrg(orgID)
	if err != nil {
		return errors.Errorf("failed to get org, %v", err)
	}

	var list []Item
	for _, release := range releaseResp.Releases {
		editOperation := Operation{
			Command: Command{
				JumpOut: false,
				Key:     "goto",
				Target:  release.ReleaseID,
			},
			Key:    "gotoDetail",
			Reload: false,
			Text:   r.sdk.I18n("editRelease"),
		}
		formalOperation := Operation{
			Confirm: r.sdk.I18n("confirmFormal"),
			Key:     "formal",
			Reload:  true,
			Text:    r.sdk.I18n("toFormal"),
			Meta: map[string]interface{}{
				"id": release.ReleaseID,
			},
			SuccessMsg: r.sdk.I18n("formalSucceeded"),
		}
		deleteOperation := Operation{
			Confirm: r.sdk.I18n("confirmDelete"),
			Key:     "delete",
			Reload:  true,
			Text:    r.sdk.I18n("deleteRelease"),
			Meta: map[string]interface{}{
				"id": release.ReleaseID,
			},
			SuccessMsg: r.sdk.I18n("deleteSucceeded"),
		}

		downloadPath := fmt.Sprintf("/api/%s/releases/%s/actions/download-yaml", org.Name, release.ReleaseID)

		item := Item{
			ID:          release.ReleaseID,
			Version:     release.Version,
			Application: release.ApplicationName,
			Desc:        release.Desc,
			Creator: Creator{
				RenderType: "userAvatar",
				Value:      []string{release.UserID},
			},
			CreatedAt: release.CreatedAt.Format("2006/01/02 15:04:05"),
			Operations: TableOperations{
				Operations: map[string]interface{}{
					"download": Operation{
						Command: Command{
							JumpOut: true,
							Key:     "goto",
							Target:  downloadPath,
						},
						Key:    "download",
						Reload: false,
						Text:   r.sdk.I18n("downloadDice"),
					},
				},
				RenderType: "tableOperation",
			},
		}
		if !r.State.IsFormal {
			item.Operations.Operations["edit"] = editOperation
			item.Operations.Operations["formal"] = formalOperation
			item.Operations.Operations["delete"] = deleteOperation
			item.BatchOperations = []string{"formal", "delete"}
		}

		list = append(list, item)
	}
	r.Data.List = list
	return nil
}

func (r *ComponentReleaseTable) SetComponentValue() {
	r.Operations = map[string]interface{}{
		"changePageNo": Operation{
			Key:    "changePageNo",
			Reload: true,
		},
		"changePageSize": Operation{
			Key:    "changePageSize",
			Reload: true,
		},
		"formal": Operation{
			Key:    "formal",
			Reload: true,
			Text:   r.sdk.I18n("toFormal"),
		},
		"delete": Operation{
			Key:    "delete",
			Reload: true,
			Text:   r.sdk.I18n("deleteRelease"),
		},
		"changeSort": Operation{
			Key:    "changeSort",
			Reload: true,
		},
	}

	var batchOperations []string
	if !r.State.IsFormal {
		batchOperations = []string{"formal", "delete"}
	}

	columns := []Column{
		{
			DataIndex: "version",
			Title:     r.sdk.I18n("version"),
		},
		{
			DataIndex: "application",
			Title:     r.sdk.I18n("applicationName"),
		},
		{
			DataIndex: "desc",
			Title:     r.sdk.I18n("desc"),
		},
		{
			DataIndex: "creator",
			Title:     r.sdk.I18n("creator"),
		},
		{
			DataIndex: "createdAt",
			Title:     r.sdk.I18n("createdAt"),
			Sorter:    true,
		},
		{
			DataIndex: "operations",
			Title:     r.sdk.I18n("operations"),
		},
	}
	if r.State.IsProjectRelease {
		columns = append(columns[:1], columns[2:]...)
	}

	r.Props = Props{
		BatchOperations: batchOperations,
		Selectable:      true,
		Columns:         columns,
		PageSizeOptions: []string{"10", "20", "50", "100"},
		RowKey:          "id",
	}
}

func (r *ComponentReleaseTable) Transfer(component *cptype.Component) {
	component.Props = cputil.MustConvertProps(r.Props)
	component.Data = map[string]interface{}{
		"list": r.Data.List,
	}
	component.State = map[string]interface{}{
		"releaseTable__urlQuery": r.State.ReleaseTableURLQuery,
		"pageNo":                 r.State.PageNo,
		"pageSize":               r.State.PageSize,
		"total":                  r.State.Total,
		"selectedRowKeys":        r.State.SelectedRowKeys,
		"sorterData":             r.State.Sorter,
		"isProjectRelease":       r.State.IsProjectRelease,
		"projectID":              r.State.ProjectID,
		"isFormal":               r.State.IsFormal,
		"applicationID":          r.State.ApplicationID,
		"versionValues":          r.State.VersionValues,
		"filterValues":           r.State.FilterValues,
	}
	component.Operations = r.Operations
}

func (r *ComponentReleaseTable) formalReleases(releaseID ...string) error {
	userID := r.sdk.Identity.UserID
	orgIDStr := r.sdk.Identity.OrgID
	projectID := r.State.ProjectID
	hasAccess, err := access.HasWriteAccess(r.bdl, userID, uint64(projectID))
	if err != nil {
		return errors.Errorf("failed to check access, %v", err)
	}
	if !hasAccess {
		return errors.Errorf("Access denied")
	}

	orgID, err := strconv.ParseUint(orgIDStr, 10, 64)
	if err != nil {
		return errors.New("invalid org id")
	}
	return r.bdl.ToFormalReleases(orgID, apistructs.ReleasesToFormalRequest{
		ProjectID: projectID,
		ReleaseID: releaseID,
	})
}

func (r *ComponentReleaseTable) deleteReleases(releaseID ...string) error {
	userID := r.sdk.Identity.UserID
	orgIDStr := r.sdk.Identity.OrgID
	projectID := r.State.ProjectID
	hasAccess, err := access.HasWriteAccess(r.bdl, userID, uint64(projectID))
	if err != nil {
		return errors.Errorf("failed to check access, %v", err)
	}
	if !hasAccess {
		return errors.Errorf("Access denied")
	}

	orgID, err := strconv.ParseUint(orgIDStr, 10, 64)
	if err != nil {
		return errors.New("invalid org id")
	}
	return r.bdl.DeleteReleases(orgID, apistructs.ReleasesDeleteRequest{
		ProjectID: projectID,
		ReleaseID: releaseID,
	})
}

func getReleaseID(operationData map[string]interface{}) (string, error) {
	meta, ok := operationData["meta"].(map[string]interface{})
	if !ok {
		return "", errors.New("invalid meta in event.operationData")
	}
	id, ok := meta["id"].(string)
	if !ok {
		return "", errors.New("invalid release id in event.operationData")
	}
	return id, nil
}