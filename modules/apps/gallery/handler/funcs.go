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
// limitations under the License.o

package handler

import (
	"encoding/json"
	"sort"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/erda-project/erda-proto-go/apps/gallery/pb"
	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/modules/apps/gallery/dao"
	"github.com/erda-project/erda/modules/apps/gallery/model"
)

func ListOpusTypes() *pb.ListOpusTypesRespData {
	var data pb.ListOpusTypesRespData
	for type_, name := range apistructs.OpusTypeNames {
		data.List = append(data.List, &pb.OpusType{
			Type:        type_.String(),
			Name:        name,
			DisplayName: apistructs.OpusTypeDisplayNames[type_],
		})
	}
	sort.Slice(data.List, func(i, j int) bool {
		return data.List[i].Type > data.List[j].Type
	})
	data.Total = uint32(len(data.List))
	return &data
}

func AdjustPaging(pageSize, pageNo int32) (int, int) {
	var size, no = 10, 1
	if pageSize >= 10 && pageSize <= 1000 {
		size = int(pageSize)
	}
	if pageNo >= 1 {
		no = int(pageNo)
	}
	return size, no
}

func PrepareListOpusesOptions(orgID int64, type_, name string, pageSize, pageNo int) []dao.Option {
	var options = []dao.Option{dao.PageOption(pageSize, pageNo)}
	if type_ != "" {
		options = append(options, dao.WhereOption("type = ?", type_))
	}
	if name != "" {
		options = append(options, dao.WhereOption("name = ?", name))
	}
	options = append(options, dao.WhereOption("org_id = ? OR level = ?", orgID, apistructs.OpusLevelSystem))
	return options
}

func PrepareListOpusesKeywordFilterOption(keyword string, versions []*model.OpusVersion) dao.Option {
	var opusesIDs []string
	for _, version := range versions {
		opusesIDs = append(opusesIDs, version.OpusID)
	}
	keyword = "%" + keyword + "%"
	return dao.WhereOption("name LIKE ? OR display_name LIKE ? OR id IN (?)", keyword, keyword, opusesIDs)
}

func PrepareListVersionsInOpusesIDsOption(opuses []*model.Opus) dao.Option {
	var opusesIDs = make(map[interface{}]struct{})
	for _, opus := range opuses {
		opusesIDs[opus.ID] = struct{}{}
	}
	return dao.InOption("opus_id", opusesIDs)
}

func ComposeListOpusResp(total int64, opuses []*model.Opus, versions []*model.OpusVersion) *pb.ListOpusResp {
	var versionsMap = make(map[string]map[string]*model.OpusVersion)
	for _, version := range versions {
		m, ok := versionsMap[version.OpusID]
		if !ok {
			m = make(map[string]*model.OpusVersion)
		}
		m[version.ID.String] = version
		versionsMap[version.OpusID] = m
	}

	var result = pb.ListOpusResp{Data: &pb.ListOpusRespData{Total: int32(total)}}
	for _, opus := range opuses {
		item := pb.ListOpusRespDataItem{
			Id:          opus.ID.String,
			CreatedAt:   timestamppb.New(opus.CreatedAt),
			UpdatedAt:   timestamppb.New(opus.UpdatedAt),
			OrgID:       opus.OrgID,
			OrgName:     opus.OrgName,
			CreatorID:   opus.CreatorID,
			UpdaterID:   opus.UpdaterID,
			Type:        opus.Type,
			TypeName:    apistructs.OpusTypeNames[apistructs.OpusType(opus.Type)],
			Name:        opus.Name,
			DisplayName: opus.DisplayName,
			Summary:     "",
			Catalog:     opus.Catalog,
			LogoURL:     "",
		}
		if m, ok := versionsMap[opus.ID.String]; ok {
			for k := range m {
				item.Summary = m[k].Summary
				item.LogoURL = m[k].LogoURL
				if k == opus.DefaultVersionID {
					break
				}
			}
		}
		result.Data.List = append(result.Data.List, &item)
	}
	return &result
}

func ComposeListOpusVersionRespWithOpus(resp *pb.ListOpusVersionsResp, opus *model.Opus) {
	resp.Data = &pb.ListOpusVersionsRespData{
		Id:               opus.ID.String,
		CreatedAt:        timestamppb.New(opus.CreatedAt),
		UpdatedAt:        timestamppb.New(opus.UpdatedAt),
		OrgID:            opus.OrgID,
		OrgName:          opus.OrgName,
		CreatorID:        opus.CreatorID,
		UpdaterID:        opus.UpdaterID,
		Level:            opus.Level,
		Type:             opus.Type,
		Name:             apistructs.OpusTypeNames[apistructs.OpusType(opus.Type)],
		DisplayName:      opus.DisplayName,
		Catalog:          opus.Catalog,
		DefaultVersionID: opus.DefaultVersionID,
		LatestVersionID:  opus.LatestVersionID,
	}
}

func ComposeListOpusVersionRespWithVersions(resp *pb.ListOpusVersionsResp, versions []*model.OpusVersion) error {
	for _, version := range versions {
		item := pb.ListOpusVersionRespDataVersion{
			Id:        version.ID.String,
			CreatedAt: timestamppb.New(version.CreatedAt),
			UpdatedAt: timestamppb.New(version.UpdatedAt),
			CreatorID: version.CreatorID,
			UpdaterID: version.UpdaterID,
			Version:   version.Version,
			Summary:   version.Summary,
			Labels:    nil,
			LogoURL:   version.LogoURL,
			IsValid:   version.IsValid,
		}
		if err := json.Unmarshal([]byte(version.Labels), &item.Labels); err != nil {
			return errors.Wrapf(err, "failed to Unmarshal version.Labels, labels: %s", version.Labels)
		}
		resp.Data.Versions = append(resp.Data.Versions, &item)
		resp.UserIDs = append(resp.UserIDs, version.CreatorID)
	}
	return nil
}

func ComposeListOpusVersionRespWithPresentations(resp *pb.ListOpusVersionsResp, presentations []*model.OpusPresentation) {
	var presentationMap = make(map[string]*model.OpusPresentation)
	for _, item := range presentations {
		presentationMap[item.VersionID] = item
	}
	for _, version := range resp.Data.Versions {
		if pre, ok := presentationMap[version.GetId()]; ok {
			version.Ref = pre.Ref
			version.Desc = pre.Desc
			version.ContactName = pre.ContactName
			version.ContactURL = pre.ContactURL
			version.ContactEmail = pre.ContactEmail
			version.IsOpenSourced = pre.IsOpenSourced
			version.OpensourceURL = pre.OpensourceURL
			version.LicenceName = pre.LicenseName
			version.LicenceURL = pre.LicenseURL
			version.HomepageName = pre.HomepageName
			version.HomepageURL = pre.HomepageURL
			version.HomepageLogoURL = pre.HomepageLogoURL
			version.IsDownloadable = pre.IsDownloadable
			version.DownloadURL = pre.DownloadURL
		}
	}
}

func ComposeListOpusVersionRespWithReadmes(resp *pb.ListOpusVersionsResp, lang string, readmes []*model.OpusReadme) {
	var readmesMap = make(map[string]map[string]*model.OpusReadme)
	for _, item := range readmes {
		m, ok := readmesMap[item.VersionID]
		if !ok {
			m = make(map[string]*model.OpusReadme)
		}
		m[item.Lang] = item
		readmesMap[item.VersionID] = m
	}
	for _, version := range resp.Data.Versions {
		langs, ok := readmesMap[version.GetId()]
		if !ok {
			continue
		}
		if readme, ok := langs[lang]; ok {
			version.ReadmeLang = readme.Lang
			version.ReadmeLangName = readme.LangName
			version.Readme = readme.Text
			continue
		}
		if readme, ok := langs[apistructs.LangUnkown.String()]; ok {
			version.ReadmeLang = readme.Lang
			version.ReadmeLangName = readme.LangName
			version.Readme = readme.Text
			continue
		}
		for k := range langs {
			version.ReadmeLang = langs[k].Lang
			version.ReadmeLangName = langs[k].LangName
			version.Readme = langs[k].Text
			break
		}
	}
}
