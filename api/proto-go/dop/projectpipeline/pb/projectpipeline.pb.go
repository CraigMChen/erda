// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: projectpipeline.proto

package pb

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/erda-project/erda-proto-go/common/pb"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListAppRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectID uint64 `protobuf:"varint,1,opt,name=projectID,proto3" json:"projectID,omitempty"`
}

func (x *ListAppRequest) Reset() {
	*x = ListAppRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_projectpipeline_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAppRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAppRequest) ProtoMessage() {}

func (x *ListAppRequest) ProtoReflect() protoreflect.Message {
	mi := &file_projectpipeline_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAppRequest.ProtoReflect.Descriptor instead.
func (*ListAppRequest) Descriptor() ([]byte, []int) {
	return file_projectpipeline_proto_rawDescGZIP(), []int{0}
}

func (x *ListAppRequest) GetProjectID() uint64 {
	if x != nil {
		return x.ProjectID
	}
	return 0
}

type ListAppResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*Application `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ListAppResponse) Reset() {
	*x = ListAppResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_projectpipeline_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAppResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAppResponse) ProtoMessage() {}

func (x *ListAppResponse) ProtoReflect() protoreflect.Message {
	mi := &file_projectpipeline_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAppResponse.ProtoReflect.Descriptor instead.
func (*ListAppResponse) Descriptor() ([]byte, []int) {
	return file_projectpipeline_proto_rawDescGZIP(), []int{1}
}

func (x *ListAppResponse) GetData() []*Application {
	if x != nil {
		return x.Data
	}
	return nil
}

type Application struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID             uint64                 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name           string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	DisplayName    string                 `protobuf:"bytes,3,opt,name=displayName,proto3" json:"displayName,omitempty"`
	Mode           string                 `protobuf:"bytes,4,opt,name=mode,proto3" json:"mode,omitempty"`
	Desc           string                 `protobuf:"bytes,5,opt,name=desc,proto3" json:"desc,omitempty"`
	Logo           string                 `protobuf:"bytes,6,opt,name=logo,proto3" json:"logo,omitempty"`
	IsPublic       bool                   `protobuf:"varint,7,opt,name=isPublic,proto3" json:"isPublic,omitempty"`
	Creator        string                 `protobuf:"bytes,8,opt,name=creator,proto3" json:"creator,omitempty"`
	GitRepo        string                 `protobuf:"bytes,9,opt,name=gitRepo,proto3" json:"gitRepo,omitempty"`
	OrgID          uint64                 `protobuf:"varint,10,opt,name=orgID,proto3" json:"orgID,omitempty"`
	OrgDisplayName string                 `protobuf:"bytes,11,opt,name=orgDisplayName,proto3" json:"orgDisplayName,omitempty"`
	ProjectId      uint64                 `protobuf:"varint,12,opt,name=projectId,proto3" json:"projectId,omitempty"`
	ProjectName    string                 `protobuf:"bytes,13,opt,name=projectName,proto3" json:"projectName,omitempty"`
	IsExternalRepo bool                   `protobuf:"varint,14,opt,name=isExternalRepo,proto3" json:"isExternalRepo,omitempty"`
	CreatedAt      *timestamppb.Timestamp `protobuf:"bytes,15,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt      *timestamppb.Timestamp `protobuf:"bytes,16,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	RunningNum     uint64                 `protobuf:"varint,17,opt,name=runningNum,proto3" json:"runningNum,omitempty"`
	FailedNum      uint64                 `protobuf:"varint,18,opt,name=failedNum,proto3" json:"failedNum,omitempty"`
}

func (x *Application) Reset() {
	*x = Application{}
	if protoimpl.UnsafeEnabled {
		mi := &file_projectpipeline_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Application) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Application) ProtoMessage() {}

func (x *Application) ProtoReflect() protoreflect.Message {
	mi := &file_projectpipeline_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Application.ProtoReflect.Descriptor instead.
func (*Application) Descriptor() ([]byte, []int) {
	return file_projectpipeline_proto_rawDescGZIP(), []int{2}
}

func (x *Application) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Application) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Application) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Application) GetMode() string {
	if x != nil {
		return x.Mode
	}
	return ""
}

func (x *Application) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *Application) GetLogo() string {
	if x != nil {
		return x.Logo
	}
	return ""
}

func (x *Application) GetIsPublic() bool {
	if x != nil {
		return x.IsPublic
	}
	return false
}

func (x *Application) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

func (x *Application) GetGitRepo() string {
	if x != nil {
		return x.GitRepo
	}
	return ""
}

func (x *Application) GetOrgID() uint64 {
	if x != nil {
		return x.OrgID
	}
	return 0
}

func (x *Application) GetOrgDisplayName() string {
	if x != nil {
		return x.OrgDisplayName
	}
	return ""
}

func (x *Application) GetProjectId() uint64 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

func (x *Application) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

func (x *Application) GetIsExternalRepo() bool {
	if x != nil {
		return x.IsExternalRepo
	}
	return false
}

func (x *Application) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Application) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Application) GetRunningNum() uint64 {
	if x != nil {
		return x.RunningNum
	}
	return 0
}

func (x *Application) GetFailedNum() uint64 {
	if x != nil {
		return x.FailedNum
	}
	return 0
}

type CreateProjectPipelineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectID  uint64 `protobuf:"varint,1,opt,name=projectID,proto3" json:"projectID,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	AppID      uint64 `protobuf:"varint,3,opt,name=appID,proto3" json:"appID,omitempty"`
	SourceType string `protobuf:"bytes,4,opt,name=sourceType,proto3" json:"sourceType,omitempty"`
	Ref        string `protobuf:"bytes,5,opt,name=ref,proto3" json:"ref,omitempty"`
	Path       string `protobuf:"bytes,6,opt,name=path,proto3" json:"path,omitempty"`
	FileName   string `protobuf:"bytes,7,opt,name=fileName,proto3" json:"fileName,omitempty"`
}

func (x *CreateProjectPipelineRequest) Reset() {
	*x = CreateProjectPipelineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_projectpipeline_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProjectPipelineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProjectPipelineRequest) ProtoMessage() {}

func (x *CreateProjectPipelineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_projectpipeline_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProjectPipelineRequest.ProtoReflect.Descriptor instead.
func (*CreateProjectPipelineRequest) Descriptor() ([]byte, []int) {
	return file_projectpipeline_proto_rawDescGZIP(), []int{3}
}

func (x *CreateProjectPipelineRequest) GetProjectID() uint64 {
	if x != nil {
		return x.ProjectID
	}
	return 0
}

func (x *CreateProjectPipelineRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateProjectPipelineRequest) GetAppID() uint64 {
	if x != nil {
		return x.AppID
	}
	return 0
}

func (x *CreateProjectPipelineRequest) GetSourceType() string {
	if x != nil {
		return x.SourceType
	}
	return ""
}

func (x *CreateProjectPipelineRequest) GetRef() string {
	if x != nil {
		return x.Ref
	}
	return ""
}

func (x *CreateProjectPipelineRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *CreateProjectPipelineRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

type CreateProjectPipelineResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectPipeline *ProjectPipeline `protobuf:"bytes,1,opt,name=ProjectPipeline,proto3" json:"ProjectPipeline,omitempty"`
}

func (x *CreateProjectPipelineResponse) Reset() {
	*x = CreateProjectPipelineResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_projectpipeline_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProjectPipelineResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProjectPipelineResponse) ProtoMessage() {}

func (x *CreateProjectPipelineResponse) ProtoReflect() protoreflect.Message {
	mi := &file_projectpipeline_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProjectPipelineResponse.ProtoReflect.Descriptor instead.
func (*CreateProjectPipelineResponse) Descriptor() ([]byte, []int) {
	return file_projectpipeline_proto_rawDescGZIP(), []int{4}
}

func (x *CreateProjectPipelineResponse) GetProjectPipeline() *ProjectPipeline {
	if x != nil {
		return x.ProjectPipeline
	}
	return nil
}

type ProjectPipeline struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID               string                 `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name             string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Creator          string                 `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator,omitempty"`
	Category         string                 `protobuf:"bytes,4,opt,name=category,proto3" json:"category,omitempty"`
	TimeCreated      *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=timeCreated,proto3" json:"timeCreated,omitempty"`
	TimeUpdated      *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=timeUpdated,proto3" json:"timeUpdated,omitempty"`
	SourceType       string                 `protobuf:"bytes,7,opt,name=sourceType,proto3" json:"sourceType,omitempty"`
	Remote           string                 `protobuf:"bytes,8,opt,name=remote,proto3" json:"remote,omitempty"`
	Ref              string                 `protobuf:"bytes,9,opt,name=ref,proto3" json:"ref,omitempty"`
	Path             string                 `protobuf:"bytes,10,opt,name=path,proto3" json:"path,omitempty"`
	FileName         string                 `protobuf:"bytes,11,opt,name=fileName,proto3" json:"fileName,omitempty"`
	PipelineSourceId string                 `protobuf:"bytes,12,opt,name=pipelineSourceId,proto3" json:"pipelineSourceId,omitempty"`
}

func (x *ProjectPipeline) Reset() {
	*x = ProjectPipeline{}
	if protoimpl.UnsafeEnabled {
		mi := &file_projectpipeline_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectPipeline) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectPipeline) ProtoMessage() {}

func (x *ProjectPipeline) ProtoReflect() protoreflect.Message {
	mi := &file_projectpipeline_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectPipeline.ProtoReflect.Descriptor instead.
func (*ProjectPipeline) Descriptor() ([]byte, []int) {
	return file_projectpipeline_proto_rawDescGZIP(), []int{5}
}

func (x *ProjectPipeline) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *ProjectPipeline) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProjectPipeline) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

func (x *ProjectPipeline) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *ProjectPipeline) GetTimeCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.TimeCreated
	}
	return nil
}

func (x *ProjectPipeline) GetTimeUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.TimeUpdated
	}
	return nil
}

func (x *ProjectPipeline) GetSourceType() string {
	if x != nil {
		return x.SourceType
	}
	return ""
}

func (x *ProjectPipeline) GetRemote() string {
	if x != nil {
		return x.Remote
	}
	return ""
}

func (x *ProjectPipeline) GetRef() string {
	if x != nil {
		return x.Ref
	}
	return ""
}

func (x *ProjectPipeline) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ProjectPipeline) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *ProjectPipeline) GetPipelineSourceId() string {
	if x != nil {
		return x.PipelineSourceId
	}
	return ""
}

var File_projectpipeline_proto protoreflect.FileDescriptor

var file_projectpipeline_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x64, 0x6f,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x1a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x77,
	0x69, 0x74, 0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a, 0x0e,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24,
	0x0a, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x10, 0x00, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x44, 0x22, 0x4c, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x64, 0x6f, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0xb7, 0x04, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x65, 0x73, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63,
	0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6c, 0x6f, 0x67, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x69,
	0x74, 0x52, 0x65, 0x70, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x69, 0x74,
	0x52, 0x65, 0x70, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x44, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x0e, 0x6f, 0x72,
	0x67, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x6f, 0x72, 0x67, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x69, 0x73, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x52, 0x65, 0x70, 0x6f, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x69, 0x73, 0x45, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x4e, 0x75, 0x6d, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0a, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x4e, 0x75, 0x6d, 0x12, 0x1c,
	0x0a, 0x09, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x4e, 0x75, 0x6d, 0x18, 0x12, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x09, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x4e, 0x75, 0x6d, 0x22, 0xf0, 0x01, 0x0a,
	0x1c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a,
	0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x10, 0x00, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05,
	0x61, 0x70, 0x70, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58,
	0x01, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a,
	0x03, 0x72, 0x65, 0x66, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02,
	0x58, 0x01, 0x52, 0x03, 0x72, 0x65, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x22, 0x0a, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2,
	0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x74, 0x0a, 0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x53, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x65, 0x72, 0x64, 0x61,
	0x2e, 0x64, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x70, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x52, 0x0f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0x8d, 0x03, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x12, 0x3c, 0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x3c, 0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x66, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x65, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x70, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x49, 0x64, 0x32, 0x9f, 0x03, 0x0a, 0x16, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0xb4, 0x01, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x36, 0x2e, 0x65, 0x72,
	0x64, 0x61, 0x2e, 0x64, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x70, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x64, 0x6f, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x39, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x17, 0x22, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2d, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0xfa, 0x81, 0xf9, 0x1b, 0x17,
	0x0a, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x70,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0xbb, 0x01, 0x0a, 0x07, 0x4c, 0x69, 0x73, 0x74,
	0x41, 0x70, 0x70, 0x12, 0x28, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x64, 0x6f, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e,
	0x65, 0x72, 0x64, 0x61, 0x2e, 0x64, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28,
	0x12, 0x26, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x70,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x7b, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x7d, 0xfa, 0x81, 0xf9, 0x1b, 0x28, 0x0a, 0x26, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x70, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x44, 0x7d, 0x1a, 0x10, 0xc2, 0xc4, 0xcb, 0x1c, 0x0b, 0x22, 0x03, 0x64, 0x6f,
	0x70, 0x32, 0x04, 0x10, 0x01, 0x20, 0x01, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x72, 0x64, 0x61, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2f, 0x65, 0x72, 0x64, 0x61, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x6f,
	0x2f, 0x64, 0x6f, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x70, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_projectpipeline_proto_rawDescOnce sync.Once
	file_projectpipeline_proto_rawDescData = file_projectpipeline_proto_rawDesc
)

func file_projectpipeline_proto_rawDescGZIP() []byte {
	file_projectpipeline_proto_rawDescOnce.Do(func() {
		file_projectpipeline_proto_rawDescData = protoimpl.X.CompressGZIP(file_projectpipeline_proto_rawDescData)
	})
	return file_projectpipeline_proto_rawDescData
}

var file_projectpipeline_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_projectpipeline_proto_goTypes = []interface{}{
	(*ListAppRequest)(nil),                // 0: erda.dop.projectpipeline.ListAppRequest
	(*ListAppResponse)(nil),               // 1: erda.dop.projectpipeline.ListAppResponse
	(*Application)(nil),                   // 2: erda.dop.projectpipeline.Application
	(*CreateProjectPipelineRequest)(nil),  // 3: erda.dop.projectpipeline.CreateProjectPipelineRequest
	(*CreateProjectPipelineResponse)(nil), // 4: erda.dop.projectpipeline.CreateProjectPipelineResponse
	(*ProjectPipeline)(nil),               // 5: erda.dop.projectpipeline.ProjectPipeline
	(*timestamppb.Timestamp)(nil),         // 6: google.protobuf.Timestamp
}
var file_projectpipeline_proto_depIdxs = []int32{
	2, // 0: erda.dop.projectpipeline.ListAppResponse.data:type_name -> erda.dop.projectpipeline.Application
	6, // 1: erda.dop.projectpipeline.Application.createdAt:type_name -> google.protobuf.Timestamp
	6, // 2: erda.dop.projectpipeline.Application.updatedAt:type_name -> google.protobuf.Timestamp
	5, // 3: erda.dop.projectpipeline.CreateProjectPipelineResponse.ProjectPipeline:type_name -> erda.dop.projectpipeline.ProjectPipeline
	6, // 4: erda.dop.projectpipeline.ProjectPipeline.timeCreated:type_name -> google.protobuf.Timestamp
	6, // 5: erda.dop.projectpipeline.ProjectPipeline.timeUpdated:type_name -> google.protobuf.Timestamp
	3, // 6: erda.dop.projectpipeline.ProjectPipelineService.Create:input_type -> erda.dop.projectpipeline.CreateProjectPipelineRequest
	0, // 7: erda.dop.projectpipeline.ProjectPipelineService.ListApp:input_type -> erda.dop.projectpipeline.ListAppRequest
	4, // 8: erda.dop.projectpipeline.ProjectPipelineService.Create:output_type -> erda.dop.projectpipeline.CreateProjectPipelineResponse
	1, // 9: erda.dop.projectpipeline.ProjectPipelineService.ListApp:output_type -> erda.dop.projectpipeline.ListAppResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_projectpipeline_proto_init() }
func file_projectpipeline_proto_init() {
	if File_projectpipeline_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_projectpipeline_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAppRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_projectpipeline_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAppResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_projectpipeline_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Application); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_projectpipeline_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProjectPipelineRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_projectpipeline_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProjectPipelineResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_projectpipeline_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectPipeline); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_projectpipeline_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_projectpipeline_proto_goTypes,
		DependencyIndexes: file_projectpipeline_proto_depIdxs,
		MessageInfos:      file_projectpipeline_proto_msgTypes,
	}.Build()
	File_projectpipeline_proto = out.File
	file_projectpipeline_proto_rawDesc = nil
	file_projectpipeline_proto_goTypes = nil
	file_projectpipeline_proto_depIdxs = nil
}