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

package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/erda-project/erda-proto-go/apps/gallery/pb"
	commonPb "github.com/erda-project/erda-proto-go/common/pb"
	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/modules/apps/cache"
	"github.com/erda-project/erda/modules/apps/gallery/apierr"
	"github.com/erda-project/erda/modules/apps/gallery/dao"
	"github.com/erda-project/erda/modules/apps/gallery/model"
	"github.com/erda-project/erda/pkg/common/apis"
)

type GalleryHandler struct {
	C *cache.Cache
	L *logrus.Entry
}

func (p *GalleryHandler) ListOpusTypes(_ context.Context, _ *commonPb.VoidRequest) (*pb.ListOpusTypesRespData, error) {
	// todo: i18n
	var data pb.ListOpusTypesRespData
	for type_, name := range apistructs.OpusTypeNames {
		data.List = append(data.List, &pb.OpusType{
			Type:        type_.String(),
			Name:        name,
			DisplayName: apistructs.OpusTypeDisplayNames[type_],
		})
	}
	data.Total = uint32(len(data.List))
	return &data, nil
}

func (p *GalleryHandler) ListOpus(ctx context.Context, req *pb.ListOpusReq) (*pb.ListOpusResp, error) {
	var l = p.L.WithField("func", "ListOpus")

	orgID := apis.GetOrgID(ctx)
	if orgID == "" {
		return nil, apierr.ListOpus.InvalidParameter("invalid orgID")
	}
	var pageNo, pageSize = 1, 10
	if req.GetPageNo() >= 1 {
		pageNo = int(req.GetPageNo())
	}
	if req.GetPageSize() >= 10 && req.GetPageSize() <= 1000 {
		pageSize = int(req.GetPageSize())
	}

	// query opuses by options
	var options []dao.Option
	if req.GetType() != "" {
		options = append(options, dao.WhereOption("type = ?", req.GetType()))
	}
	if req.GetName() != "" {
		options = append(options, dao.WhereOption("name = ?", req.GetName()))
	}
	options = append(options, dao.WhereOption("org_id = ? OR level = ?", orgID, apistructs.OpusLevelSystem))
	total, opuses, err := dao.ListOpuses(dao.Q(), options...)
	if err != nil {
		l.WithError(err).Errorln("failed to Find opuses")
		return nil, apierr.ListOpus.InternalError(err)
	}
	if total == 0 {
		l.Warnln("not found")
		return new(pb.ListOpusResp), nil
	}
	var opusesIDs []string
	for _, opus := range opuses {
		opusesIDs = append(opusesIDs, opus.ID.String)
	}

	if req.GetKeyword() == "" {
		return p.listOpusByIDs(ctx, pageSize, pageNo, opusesIDs)
	}
	return p.listOpusWithKeyword(ctx, pageSize, pageNo, req.GetKeyword(), opusesIDs)
}

func (p *GalleryHandler) ListOpusVersions(ctx context.Context, req *pb.ListOpusVersionsReq) (*pb.ListOpusVersionsResp, error) {
	var l = p.L.WithField("func", "listOpusVersions").WithField("opusID", req.GetOpusID())

	orgID, err := apis.GetIntOrgID(ctx)
	if err != nil {
		return nil, apierr.ListOpus.InvalidParameter("invalid orgID")
	}

	lang := apis.GetLang(ctx)
	lang = strings.ToLower(lang)
	lang = strings.ReplaceAll(lang, "-", "_")

	// todo: 鉴权

	// query opus
	opus, ok, err := dao.GetOpusByID(dao.Q(), req.GetOpusID())
	if err != nil {
		l.WithError(err).Errorln("failed to GetOpusByID")
		return nil, apierr.ListOpusVersions.InternalError(err)
	}
	if !ok {
		return nil, apierr.ListOpusVersions.NotFound()
	}

	// query versions
	total, versions, err := dao.ListVersions(dao.Q(), dao.WhereOption("opus_id = ?", req.GetOpusID()))
	if err != nil {
		l.WithError(err).Errorln("failed to Find versions")
		return nil, apierr.ListOpusVersions.InternalError(err)
	}
	if total == 0 {
		l.WithError(err).Errorln("opus's version not found")
		return nil, apierr.ListOpusVersions.NotFound()
	}

	// query presentations
	_, presentations, err := dao.ListPresentations(dao.Q(), dao.WhereOption("opus_id = ?", req.GetOpusID()))
	if err != nil {
		l.WithError(err).Errorln("failed to Find presentations")
		return nil, apierr.ListOpusVersions.InternalError(err)
	}

	// query readmes
	_, readmes, err := dao.ListReadmes(dao.Q(), dao.WhereOption("opus_id = ?", req.GetOpusID()))
	if err != nil {
		l.WithError(err).Errorln("failed to Find readmes")
		return nil, apierr.ListOpusVersions.InternalError(err)
	}

	var (
		presentationMap = make(map[string]*model.OpusPresentation)
		readmesMap      = make(map[string]map[string]*model.OpusReadme)
	)
	for _, item := range presentations {
		presentationMap[item.VersionID] = item
	}
	for _, item := range readmes {
		m, ok := readmesMap[item.VersionID]
		if !ok {
			m = make(map[string]*model.OpusReadme)
		}
		m[item.Lang] = item
		readmesMap[item.VersionID] = m
	}

	var resp = pb.ListOpusVersionsResp{Data: &pb.ListOpusVersionsRespData{
		Id:               opus.ID.String,
		CreatedAt:        timestamppb.New(opus.CreatedAt),
		UpdatedAt:        timestamppb.New(opus.UpdatedAt),
		OrgID:            uint32(orgID),
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
	}}
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
			l.WithError(err).Errorf("failed to Unmarshal version.Labels, labels: %s", version.Labels)
		}
		if pre, ok := presentationMap[version.ID.String]; ok {
			item.Ref = pre.Ref
			item.Desc = pre.Desc
			item.ContactName = pre.ContactName
			item.ContactURL = pre.ContactURL
			item.ContactEmail = pre.ContactEmail
			item.IsOpenSourced = pre.IsOpenSourced
			item.OpensourceURL = pre.OpensourceURL
			item.LicenceName = pre.LicenseName
			item.LicenceURL = pre.LicenseURL
			item.HomepageName = pre.HomepageName
			item.HomepageURL = pre.HomepageURL
			item.HomepageLogoURL = pre.HomepageLogoURL
			item.IsDownloadable = pre.IsDownloadable
			item.DownloadURL = pre.DownloadURL
		}
		if langs, ok := readmesMap[version.ID.String]; ok {
			for k := range langs {
				item.ReadmeLang = langs[k].Lang
				item.ReadmeLangName = langs[k].LangName
				item.Readme = langs[k].Text
				if k == lang {
					break
				}
			}
		}

		resp.Data.Versions = append(resp.Data.Versions, &item)
		resp.UserIDs = append(resp.UserIDs, version.CreatorID)
	}

	return &resp, nil
}

func (p *GalleryHandler) PutOnArtifacts(ctx context.Context, req *pb.PutOnArtifactsReq) (*pb.PutOnOpusResp, error) {
	l := p.L.WithField("func", "PutOnArtifacts").
		WithField("name", req.GetName()).
		WithField("version", req.GetVersion())

	// check parameters and get org info and user info
	var orgID, err = apis.GetIntOrgID(ctx)
	if err != nil {
		orgID = int64(req.GetOrgID())
	}
	var userID = apis.GetUserID(ctx)
	if userID == "" {
		userID = req.GetUserID()
	}
	if userID == "" {
		return nil, apierr.PutOnArtifacts.InvalidParameter("missing userID")
	}
	orgDto, ok := p.C.GetOrgByOrgID(strconv.FormatInt(orgID, 10))
	if !ok {
		return nil, apierr.PutOnArtifacts.InvalidParameter(fmt.Sprintf("invalid orgID: %s", orgID))
	}
	if req.GetName() == "" {
		return nil, apierr.PutOnArtifacts.InvalidParameter("missing Opus name")
	}
	if req.GetVersion() == "" {
		return nil, apierr.PutOnArtifacts.InvalidParameter("missing Opus version")
	}

	// todo: 鉴权

	// Check if artifacts already exist
	var common = model.Common{
		OrgID:     uint32(orgID),
		OrgName:   orgDto.Name,
		CreatorID: userID,
		UpdaterID: userID,
	}

	tx := dao.Begin()
	defer tx.CommitOrRollback()

	// get the opus by options, if the opus does not exist, create it
	opus, ok, err := dao.GetOpus(tx, dao.MapOption(map[string]interface{}{
		"org_id": orgID,
		"type":   apistructs.OpusTypeArtifactsProject,
		"name":   req.GetName(),
	}))
	if err != nil {
		l.WithError(err).Errorln("failed to Find versions")
		return nil, apierr.PutOnArtifacts.InternalError(err)
	}
	if !ok {
		opus = &model.Opus{
			Common:      common,
			Level:       string(apistructs.OpusLevelOrg),
			Type:        string(apistructs.OpusTypeArtifactsProject),
			Name:        req.GetName(),
			DisplayName: req.GetDisplayName(),
			Catalog:     req.GetCatalog(),
		}
		if err = tx.Create(opus); err != nil {
			l.WithError(err).Errorln("failed to Create opus")
			return nil, apierr.PutOnArtifacts.InternalError(err)
		}
	}

	// get the version by options, if the version exist, return 'already exists' or else create it
	_, ok, err = dao.GetOpusVersion(tx, dao.MapOption(map[string]interface{}{
		"opus_id": opus.ID,
		"version": req.GetVersion(),
	}))
	if err != nil {
		l.WithError(err).Errorln("failed to Find versions")
		return nil, apierr.PutOnArtifacts.InternalError(err)
	}
	if ok {
		l.Warnln("already exists")
		return nil, apierr.PutOnArtifacts.AlreadyExists()
	}
	labels, err := json.Marshal(req.GetLabels())
	if err != nil {
		l.WithError(err).Warnf("failed to json.Marshal labels, labels: %v", req.GetLabels())
	}
	var version = model.OpusVersion{
		Common:  common,
		OpusID:  opus.ID.String,
		Version: req.GetVersion(),
		Summary: req.GetSummary(),
		Labels:  string(labels),
		LogoURL: req.GetLogoURL(),
		IsValid: true,
	}
	if err = tx.Create(&version); err != nil {
		l.WithError(err).Errorln("failed to Create version")
		return nil, apierr.PutOnArtifacts.InternalError(err)
	}

	// create presentation
	var presentation = model.OpusPresentation{
		Common:          common,
		OpusID:          opus.ID.String,
		VersionID:       version.ID.String,
		Desc:            req.GetDesc(),
		ContactName:     req.GetContactName(),
		ContactURL:      req.GetContactURL(),
		ContactEmail:    req.GetContactEmail(),
		IsOpenSourced:   req.GetIsOpenSourced(),
		OpensourceURL:   req.GetOpensourceURL(),
		LicenseName:     req.GetLicenseName(),
		LicenseURL:      req.GetLicenseURL(),
		HomepageName:    req.GetHomepageName(),
		HomepageURL:     req.GetHomepageURL(),
		HomepageLogoURL: req.GetHomepageLogoURL(),
		IsDownloadable:  req.GetIsDownloadable(),
		DownloadURL:     req.GetDownloadURL(),
	}
	if err = tx.Create(&presentation); err != nil {
		l.WithError(err).Errorln("failed to Create presentation")
		return nil, apierr.PutOnArtifacts.InternalError(err)
	}

	// create readmes
	var readmes []*model.OpusReadme
	for _, item := range req.GetReadme() {
		readme := &model.OpusReadme{
			Model:     model.Model{},
			Common:    common,
			OpusID:    opus.ID.String,
			VersionID: version.ID.String,
			Lang:      item.Lang,
			LangName:  item.LangName,
			Text:      item.Text,
		}
		readmes = append(readmes, readme)
	}
	if err = tx.CreateInBatches(readmes, len(readmes)); err != nil {
		l.WithError(err).Errorln("failed to CreateInBatches readmes")
		return nil, apierr.PutOnArtifacts.InternalError(err)
	}

	// create installation
	spec, err := json.Marshal(req.GetInstallation())
	if err != nil {
		return nil, apierr.PutOnArtifacts.InternalError(err)
	}
	var installation = model.OpusInstallation{
		Model:     model.Model{},
		Common:    common,
		OpusID:    opus.ID.String,
		VersionID: version.ID.String,
		Installer: string(apistructs.OpusTypeArtifactsProject),
		Spec:      string(spec),
	}
	if err = tx.Create(&installation); err != nil {
		l.WithError(err).Errorln("failed to Create installation")
		return nil, apierr.PutOnArtifacts.InternalError(err)
	}

	// update opus
	if err = tx.Updates(&opus, map[string]interface{}{
		"updater_id":         userID,
		"default_version_id": version.ID,
		"latest_version_id":  version.ID,
	}, dao.ByIDOption(opus.ID)); err != nil {
		l.WithError(err).Errorln("failed to Updates opus")
		return nil, apierr.PutOnArtifacts.InternalError(err)
	}

	return &pb.PutOnOpusResp{OpusID: version.OpusID, VersionID: version.ID.String}, nil
}

func (p *GalleryHandler) PutOffArtifacts(ctx context.Context, req *pb.PutOffArtifactsReq) (*commonPb.VoidResponse, error) {
	l := p.L.WithField("func", "PutOffArtifacts").
		WithField("opus_id", req.GetOpusID()).
		WithField("version_id", req.GetVersionID())

	// check parameters and get org info and user info
	var userID = apis.GetUserID(ctx)
	if userID == "" {
		userID = req.GetUserID()
	}
	if userID == "" {
		return nil, apierr.PutOffArtifacts.InvalidParameter("missing userID")
	}

	// todo: 鉴权

	// query version
	var byIDOption = dao.ByIDOption(req.GetVersionID())
	version, ok, err := dao.GetOpusVersion(dao.Q(), byIDOption)
	if err != nil {
		l.WithError(err).Errorln("failed to GetOpusVersion")
		return nil, apierr.PutOffArtifacts.InternalError(err)
	}
	if !ok {
		l.WithError(err).Warnln("delete version not found")
		return new(commonPb.VoidResponse), nil
	}
	if req.GetOpusID() != version.OpusID {
		return nil, apierr.PutOffArtifacts.InvalidParameter("invalid opusID and versionID")
	}

	tx := dao.Begin()
	defer tx.CommitOrRollback()

	total, _, err := dao.ListVersions(tx, dao.MapOption(map[string]interface{}{"opus_id": req.GetOpusID()}))
	if err != nil {
		l.WithError(err).Errorln("failed to ListVersions")
		return nil, apierr.PutOffArtifacts.InternalError(err)
	}
	if err = tx.Delete(new(model.OpusVersion), byIDOption); err != nil {
		l.WithError(err).Errorln("failed to Delete version")
		return nil, apierr.PutOffArtifacts.InternalError(err)
	}
	var versionIDOption = dao.MapOption(map[string]interface{}{"version_id": req.GetVersionID()})
	if err = tx.Delete(new(model.OpusPresentation), versionIDOption); err != nil {
		l.WithError(err).Errorln("failed to Delete presentation")
		return nil, apierr.PutOffArtifacts.InternalError(err)
	}
	if err = tx.Delete(new(model.OpusReadme), versionIDOption); err != nil {
		l.WithError(err).Errorln("failed to Delete readme")
		return nil, apierr.PutOffArtifacts.InternalError(err)
	}
	if err = tx.Delete(new(model.OpusInstallation), versionIDOption); err != nil {
		l.WithError(err).Errorln("failed to Delete installation")
		return nil, apierr.PutOffArtifacts.InternalError(err)
	}
	if total == 1 {
		if err = tx.Delete(new(model.Opus), dao.ByIDOption(req.GetOpusID())); err != nil {
			l.WithError(err).Errorln("failed to Delete opus")
			return nil, apierr.PutOffArtifacts.InternalError(err)
		}
	}

	return new(commonPb.VoidResponse), nil
}

func (p *GalleryHandler) PutOnExtensions(ctx context.Context, req *pb.PubOnExtensionsReq) (*pb.PutOnOpusResp, error) {
	l := p.L.WithField("func", "PutOnExtensions").
		WithFields(map[string]interface{}{
			"orgID":   req.GetOrgID(),
			"type":    req.GetType(),
			"name":    req.GetName(),
			"version": req.GetVersion(),
			"level":   req.GetLevel(),
			"mode":    req.GetMode(),
		})

	// check parameters and get org info and user info
	var orgID, err = apis.GetIntOrgID(ctx)
	if err != nil {
		orgID = int64(req.GetOrgID())
	}
	if orgID == 0 && apistructs.OpusLevelOrg.Equal(req.GetLevel()) {
		return nil, apierr.PutOnExtension.InvalidState("missing orgID")
	}
	var userID = apis.GetUserID(ctx)
	if userID == "" {
		userID = req.GetUserID()
	}
	if userID == "" && (apistructs.OpusLevelOrg.Equal(req.GetLevel())) {
		return nil, apierr.PutOnExtension.InvalidParameter("missing userID")
	}
	if apistructs.OpusTypeExtensionAddon.Equal(req.GetType()) && apistructs.OpusTypeExtensionAction.Equal(req.GetType()) {
		return nil, apierr.PutOnExtension.InvalidParameter("invalid type: " + req.GetType())
	}
	if req.GetName() == "" {
		return nil, apierr.PutOnExtension.InvalidParameter("missing name")
	}
	if req.GetVersion() == "" {
		return nil, apierr.PutOnExtension.InvalidParameter("missing version")
	}
	if apistructs.OpusLevelSystem.Equal(req.GetLevel()) && apistructs.OpusLevelOrg.Equal(req.GetLevel()) {
		return nil, apierr.PutOnExtension.InvalidParameter("invalid level: " + req.GetLevel())
	}

	// todo: 鉴权

	// Check if extension already exist
	var (
		where = map[string]interface{}{
			"type":  req.GetType(),
			"name":  req.GetName(),
			"level": req.GetLevel(),
		}
	)
	if apistructs.OpusLevelOrg.Equal(req.GetLevel()) {
		where["org_id"] = orgID
	}
	opus, ok, err := dao.GetOpus(dao.Q(), dao.MapOption(where))
	if err != nil {
		l.WithError(err).WithFields(where).Errorln("failed to First opus")
		return nil, apierr.PutOnExtension.InternalError(err)
	}
	if !ok {
		return p.createExtensions(ctx, l, userID, nil, req)
	}

	where = map[string]interface{}{
		"opus_id": opus.ID,
		"version": req.GetVersion(),
	}
	version, ok, err := dao.GetOpusVersion(dao.Q(), dao.MapOption(where))
	if err != nil {
		l.WithError(err).WithFields(where).Errorln("failed to First version")
		return nil, apierr.PutOnExtension.InternalError(err)
	}
	if !ok {
		return p.createExtensions(ctx, l, userID, opus, req)
	}

	if apistructs.PutOnOpusModeAppend.Equal(req.GetMode()) {
		l.Warnln("failed to put on extension, the version already exists, can not append")
		return nil, apierr.PutOnExtension.AlreadyExists()
	}
	return p.updateExtension(ctx, l, userID, opus, version, req)
}

func (p *GalleryHandler) listOpusByIDs(_ context.Context, pageSize, pageNo int, opusesIDs []string) (*pb.ListOpusResp, error) {
	var l = p.L.WithField("func", "listOpusByIDs")

	total, opuses, err := dao.ListOpuses(
		dao.Q(),
		dao.PageOption(pageSize, pageNo),
		dao.WhereOption("id IN (?)", opusesIDs),
	)
	if err != nil {
		l.WithError(err).Errorln("failed to Find opuses")
		return nil, apierr.ListOpus.InternalError(err)
	}
	if total == 0 {
		return new(pb.ListOpusResp), nil
	}

	_, versions, err := dao.ListVersions(dao.Q(), dao.WhereOption("opus_id IN (?)", opusesIDs))
	if err != nil {
		l.WithError(err).Errorln("failed to Find versions")
		return nil, apierr.ListOpus.InternalError(err)
	}

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
	return &result, nil
}

func (p *GalleryHandler) listOpusWithKeyword(ctx context.Context, pageSize, pageNo int, keyword string, opusesIDs []string) (*pb.ListOpusResp, error) {
	var (
		l        = p.L.WithField("func", "listOpusWithKeyword")
		opuses   []*model.Opus
		versions []*model.OpusVersion
	)

	keyword = "%" + keyword + "%"
	total, versions, err := dao.ListVersions(dao.Q(),
		dao.WhereOption("opus_id IN (?)", opusesIDs),
		dao.WhereOption("summary LIKE ?", keyword))
	if err != nil {
		l.WithError(err).Errorln("failed to Find versions")
		return nil, apierr.ListOpus.InternalError(err)
	}
	if total == 0 {
		l.WithField("len(versions)", len(versions)).Warnln("versions not found")
		return nil, apierr.ListOpus.NotFound()
	}

	total, opuses, err = dao.ListOpuses(dao.Q(),
		dao.WhereOption("id IN (?)", opusesIDs),
		dao.WhereOption("name LIKE ? OR display_name LIKE ?", keyword, keyword))
	if err != nil {
		l.WithError(err).Errorln("failed to Find opuses")
		return nil, apierr.ListOpus.InternalError(err)
	}
	if total == 0 {
		l.WithField("len(opuses)", len(opuses)).Warnln("opuses not found")
		return nil, apierr.ListOpus.NotFound()
	}

	var opusesIDsMap = make(map[string]struct{})
	for _, version := range versions {
		opusesIDsMap[version.OpusID] = struct{}{}
	}
	for _, opus := range opuses {
		opusesIDsMap[opus.ID.String] = struct{}{}
	}

	opusesIDs = nil
	for k := range opusesIDsMap {
		opusesIDs = append(opusesIDs, k)
	}

	return p.listOpusByIDs(ctx, pageSize, pageNo, opusesIDs)
}

func (p *GalleryHandler) updateExtension(ctx context.Context, l *logrus.Entry, userID string, opus *model.Opus, version *model.OpusVersion, req *pb.PubOnExtensionsReq) (*pb.PutOnOpusResp, error) {
	l = l.WithField("func", "updateExtension")

	var err error
	tx := dao.Begin()
	defer tx.CommitOrRollback()

	// update version
	labels, _ := json.Marshal(req.GetLabels())
	if err = tx.Updates(&version, map[string]interface{}{
		"summary":    req.GetSummary(),
		"labels":     string(labels),
		"logo_url":   req.GetLogoURL(),
		"updater_id": userID,
	}, dao.ByIDOption(version.ID)); err != nil {
		l.WithError(err).Errorln("failed to Updates")
		return nil, apierr.PutOnExtension.InternalError(err)
	}

	// update presentation
	if err = tx.Updates(new(model.OpusPresentation), map[string]interface{}{
		"desc":              req.GetDesc(),
		"contact_name":      req.GetContactName(),
		"contact_url":       req.GetContactURL(),
		"contact_email":     req.GetContactEmail(),
		"is_open_sourced":   req.GetIsOpenSourced(),
		"opensource_url":    req.GetOpensourceURL(),
		"licence_name":      req.GetLicenseName(),
		"license_url":       req.GetLicenseURL(),
		"homepage_name":     req.GetHomepageName(),
		"homepage_url":      req.GetHomepageURL(),
		"homepage_logo_url": req.GetHomepageLogoURL(),
		"is_downloadable":   req.GetIsDownloadable(),
		"download_url":      req.GetDownloadURL(),
		"updater_id":        userID,
	}, dao.WhereOption("version_id = ?", version.ID)); err != nil {
		l.WithError(err).Errorln("failed to Updates presentation")
		return nil, apierr.PutOnExtension.InternalError(err)
	}

	// create or update readmes
	for _, item := range req.GetReadme() {
		if item.GetLang() == "" {
			continue
		}
		var where = map[string]interface{}{
			"version_id": version.ID,
			"lang":       item.GetLang(),
		}
		var option = dao.MapOption(where)
		var updates = map[string]interface{}{
			"lang_name": apistructs.LangTypes[apistructs.Lang(item.GetLang())],
			"text":      item.GetText(),
		}
		_, ok, err := dao.GetReadme(tx, option)
		if err != nil {
			l.WithError(err).WithFields(where).Errorln("failed to First readme")
			return nil, apierr.PutOnExtension.InternalError(err)
		}
		if ok {
			if err = tx.Updates(new(model.OpusReadme), updates, option); err != nil {
				l.WithError(err).WithField("lang", item.GetLang()).Errorln("failed to Updates readme")
				return nil, apierr.PutOnExtension.InternalError(err)
			}
		} else {
			if err = tx.Create(&model.OpusReadme{
				Common: model.Common{
					OrgID:     opus.OrgID,
					OrgName:   opus.OrgName,
					CreatorID: userID,
					UpdaterID: userID,
				},
				OpusID:    version.OpusID,
				VersionID: version.ID.String,
				Lang:      item.GetLang(),
				LangName:  apistructs.LangTypes[apistructs.Lang(item.GetLang())],
				Text:      item.GetText(),
			}); err != nil {
				l.WithError(err).WithField("lang", item.GetLang()).Errorln("failed to Create readme")
				return nil, apierr.PutOnExtension.InternalError(err)
			}
		}
	}

	// update opus
	updates := map[string]interface{}{
		"updater_id": userID,
	}
	if req.GetIsDefault() {
		updates["default_version_id"] = version.ID
	}
	if err = tx.Updates(opus, updates, dao.ByIDOption(opus.ID)); err != nil {
		l.WithError(err).Errorln("failed to Updates opus")
		return nil, apierr.PutOnExtension.InternalError(err)
	}

	return &pb.PutOnOpusResp{OpusID: version.OpusID, VersionID: version.ID.String}, nil
}

func (p *GalleryHandler) createExtensions(ctx context.Context, l *logrus.Entry, userID string, opus *model.Opus, req *pb.PubOnExtensionsReq) (*pb.PutOnOpusResp, error) {
	l = l.WithField("func", "createExtensions")

	var err error
	tx := dao.Begin()
	defer tx.CommitOrRollback()

	// create version
	common := model.Common{
		OrgID:     opus.OrgID,
		OrgName:   opus.OrgName,
		CreatorID: userID,
		UpdaterID: userID,
	}

	if opus == nil {
		opus = &model.Opus{
			Model:       model.Model{},
			Common:      common,
			Level:       req.GetLevel(),
			Type:        req.GetType(),
			Name:        req.GetName(),
			DisplayName: req.GetDisplayName(),
			Catalog:     req.GetCatalog(),
		}
		if err = tx.Create(opus); err != nil {
			l.WithError(err).Errorln("failed to Create opus")
			return nil, apierr.PutOnExtension.InternalError(err)
		}
	}
	labels, _ := json.Marshal(req.GetLabels())
	var version = model.OpusVersion{
		Common:  common,
		OpusID:  opus.ID.String,
		Version: req.GetVersion(),
		Summary: req.GetSummary(),
		Labels:  string(labels),
		LogoURL: req.GetLogoURL(),
		IsValid: true,
	}
	if err = tx.Create(&version); err != nil {
		l.WithError(err).Errorln("failed to Create version")
		return nil, apierr.PutOnExtension.InternalError(err)
	}

	// create presentation
	var presentation = model.OpusPresentation{
		Common:          common,
		OpusID:          opus.ID.String,
		VersionID:       version.ID.String,
		Ref:             "",
		Desc:            req.GetDesc(),
		ContactName:     req.GetContactName(),
		ContactURL:      req.GetContactURL(),
		ContactEmail:    req.GetContactEmail(),
		IsOpenSourced:   req.GetIsOpenSourced(),
		OpensourceURL:   req.GetOpensourceURL(),
		LicenseName:     req.GetLicenseName(),
		LicenseURL:      req.GetLicenseURL(),
		HomepageName:    req.GetHomepageName(),
		HomepageURL:     req.GetHomepageURL(),
		HomepageLogoURL: req.GetHomepageLogoURL(),
		IsDownloadable:  req.GetIsDownloadable(),
		DownloadURL:     req.GetDownloadURL(),
	}
	if err = tx.Create(&presentation); err != nil {
		l.WithError(err).Errorln("failed to Create presentation")
		return nil, apierr.PutOnExtension.InternalError(err)
	}

	// create readmes
	var readmes []*model.OpusReadme
	for _, item := range req.GetReadme() {
		if item.GetLang() == "" {
			continue
		}
		readme := model.OpusReadme{
			Common:    common,
			OpusID:    opus.ID.String,
			VersionID: version.ID.String,
			Lang:      item.GetLang(),
			LangName:  apistructs.LangTypes[apistructs.Lang(item.GetLang())],
			Text:      item.GetText(),
		}
		readmes = append(readmes, &readme)
	}
	if err = tx.CreateInBatches(readmes, len(readmes)); err != nil {
		l.WithError(err).Errorln("failed to CreateInBatches readmes")
		return nil, apierr.PutOnExtension.InternalError(err)
	}

	// update opus
	updates := map[string]interface{}{
		"updater_id": userID,
	}
	if req.GetIsDefault() {
		updates["default_version_id"] = version.ID
	}
	if err = tx.Updates(opus, updates, dao.ByIDOption(opus.ID)); err != nil {
		l.WithError(err).Errorln("failed to Updates opus")
		return nil, apierr.PutOnExtension.InternalError(err)
	}

	return &pb.PutOnOpusResp{OpusID: version.OpusID, VersionID: version.ID.String}, nil
}
