// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.12
// source: colab-shield.proto

package protos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Status int32

const (
	Status_OK    Status = 0
	Status_ERROR Status = 1
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "OK",
		1: "ERROR",
	}
	Status_value = map[string]int32{
		"OK":    0,
		"ERROR": 1,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_colab_shield_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_colab_shield_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_colab_shield_proto_rawDescGZIP(), []int{0}
}

type ClaimMode int32

const (
	ClaimMode_UNCLAIMED ClaimMode = 0
	ClaimMode_EXCLUSIVE ClaimMode = 1
	ClaimMode_SHARED    ClaimMode = 2
)

// Enum value maps for ClaimMode.
var (
	ClaimMode_name = map[int32]string{
		0: "UNCLAIMED",
		1: "EXCLUSIVE",
		2: "SHARED",
	}
	ClaimMode_value = map[string]int32{
		"UNCLAIMED": 0,
		"EXCLUSIVE": 1,
		"SHARED":    2,
	}
)

func (x ClaimMode) Enum() *ClaimMode {
	p := new(ClaimMode)
	*p = x
	return p
}

func (x ClaimMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClaimMode) Descriptor() protoreflect.EnumDescriptor {
	return file_colab_shield_proto_enumTypes[1].Descriptor()
}

func (ClaimMode) Type() protoreflect.EnumType {
	return &file_colab_shield_proto_enumTypes[1]
}

func (x ClaimMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClaimMode.Descriptor instead.
func (ClaimMode) EnumDescriptor() ([]byte, []int) {
	return file_colab_shield_proto_rawDescGZIP(), []int{1}
}

// ----------------------------------------
// Health Check
// ----------------------------------------
type HealthCheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status Status `protobuf:"varint,1,opt,name=status,proto3,enum=colabshield.Status" json:"status,omitempty"`
}

func (x *HealthCheckResponse) Reset() {
	*x = HealthCheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_colab_shield_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckResponse) ProtoMessage() {}

func (x *HealthCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_colab_shield_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckResponse.ProtoReflect.Descriptor instead.
func (*HealthCheckResponse) Descriptor() ([]byte, []int) {
	return file_colab_shield_proto_rawDescGZIP(), []int{0}
}

func (x *HealthCheckResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_OK
}

// ----------------------------------------
// Project Messages
// ----------------------------------------
type InitProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId string `protobuf:"bytes,1,opt,name=projectId,proto3" json:"projectId,omitempty"`
}

func (x *InitProjectRequest) Reset() {
	*x = InitProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_colab_shield_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitProjectRequest) ProtoMessage() {}

func (x *InitProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_colab_shield_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitProjectRequest.ProtoReflect.Descriptor instead.
func (*InitProjectRequest) Descriptor() ([]byte, []int) {
	return file_colab_shield_proto_rawDescGZIP(), []int{1}
}

func (x *InitProjectRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type InitProjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status Status `protobuf:"varint,1,opt,name=status,proto3,enum=colabshield.Status" json:"status,omitempty"`
}

func (x *InitProjectResponse) Reset() {
	*x = InitProjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_colab_shield_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitProjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitProjectResponse) ProtoMessage() {}

func (x *InitProjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_colab_shield_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitProjectResponse.ProtoReflect.Descriptor instead.
func (*InitProjectResponse) Descriptor() ([]byte, []int) {
	return file_colab_shield_proto_rawDescGZIP(), []int{2}
}

func (x *InitProjectResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_OK
}

type ListProjectsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Projects []string `protobuf:"bytes,1,rep,name=projects,proto3" json:"projects,omitempty"`
}

func (x *ListProjectsResponse) Reset() {
	*x = ListProjectsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_colab_shield_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProjectsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProjectsResponse) ProtoMessage() {}

func (x *ListProjectsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_colab_shield_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProjectsResponse.ProtoReflect.Descriptor instead.
func (*ListProjectsResponse) Descriptor() ([]byte, []int) {
	return file_colab_shield_proto_rawDescGZIP(), []int{3}
}

func (x *ListProjectsResponse) GetProjects() []string {
	if x != nil {
		return x.Projects
	}
	return nil
}

// ----------------------------------------
// File Messages
// ----------------------------------------
type ListFilesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId string `protobuf:"bytes,1,opt,name=projectId,proto3" json:"projectId,omitempty"`
}

func (x *ListFilesRequest) Reset() {
	*x = ListFilesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_colab_shield_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFilesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFilesRequest) ProtoMessage() {}

func (x *ListFilesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_colab_shield_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFilesRequest.ProtoReflect.Descriptor instead.
func (*ListFilesRequest) Descriptor() ([]byte, []int) {
	return file_colab_shield_proto_rawDescGZIP(), []int{4}
}

func (x *ListFilesRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type FileInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId     string    `protobuf:"bytes,1,opt,name=fileId,proto3" json:"fileId,omitempty"`
	FileHash   string    `protobuf:"bytes,2,opt,name=fileHash,proto3" json:"fileHash,omitempty"`
	UserIds    []string  `protobuf:"bytes,3,rep,name=userIds,proto3" json:"userIds,omitempty"`
	BranchName string    `protobuf:"bytes,4,opt,name=branchName,proto3" json:"branchName,omitempty"`
	ClaimMode  ClaimMode `protobuf:"varint,5,opt,name=claimMode,proto3,enum=colabshield.ClaimMode" json:"claimMode,omitempty"`
}

func (x *FileInfo) Reset() {
	*x = FileInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_colab_shield_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileInfo) ProtoMessage() {}

func (x *FileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_colab_shield_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileInfo.ProtoReflect.Descriptor instead.
func (*FileInfo) Descriptor() ([]byte, []int) {
	return file_colab_shield_proto_rawDescGZIP(), []int{5}
}

func (x *FileInfo) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *FileInfo) GetFileHash() string {
	if x != nil {
		return x.FileHash
	}
	return ""
}

func (x *FileInfo) GetUserIds() []string {
	if x != nil {
		return x.UserIds
	}
	return nil
}

func (x *FileInfo) GetBranchName() string {
	if x != nil {
		return x.BranchName
	}
	return ""
}

func (x *FileInfo) GetClaimMode() ClaimMode {
	if x != nil {
		return x.ClaimMode
	}
	return ClaimMode_UNCLAIMED
}

type ListFilesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Files []*FileInfo `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *ListFilesResponse) Reset() {
	*x = ListFilesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_colab_shield_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFilesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFilesResponse) ProtoMessage() {}

func (x *ListFilesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_colab_shield_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFilesResponse.ProtoReflect.Descriptor instead.
func (*ListFilesResponse) Descriptor() ([]byte, []int) {
	return file_colab_shield_proto_rawDescGZIP(), []int{6}
}

func (x *ListFilesResponse) GetFiles() []*FileInfo {
	if x != nil {
		return x.Files
	}
	return nil
}

type ClaimFileInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId    string    `protobuf:"bytes,1,opt,name=fileId,proto3" json:"fileId,omitempty"`
	FileHash  string    `protobuf:"bytes,2,opt,name=fileHash,proto3" json:"fileHash,omitempty"`
	ClaimMode ClaimMode `protobuf:"varint,3,opt,name=claimMode,proto3,enum=colabshield.ClaimMode" json:"claimMode,omitempty"`
}

func (x *ClaimFileInfo) Reset() {
	*x = ClaimFileInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_colab_shield_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClaimFileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClaimFileInfo) ProtoMessage() {}

func (x *ClaimFileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_colab_shield_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClaimFileInfo.ProtoReflect.Descriptor instead.
func (*ClaimFileInfo) Descriptor() ([]byte, []int) {
	return file_colab_shield_proto_rawDescGZIP(), []int{7}
}

func (x *ClaimFileInfo) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *ClaimFileInfo) GetFileHash() string {
	if x != nil {
		return x.FileHash
	}
	return ""
}

func (x *ClaimFileInfo) GetClaimMode() ClaimMode {
	if x != nil {
		return x.ClaimMode
	}
	return ClaimMode_UNCLAIMED
}

type ClaimFilesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BranchName string           `protobuf:"bytes,1,opt,name=branchName,proto3" json:"branchName,omitempty"`
	Files      []*ClaimFileInfo `protobuf:"bytes,4,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *ClaimFilesRequest) Reset() {
	*x = ClaimFilesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_colab_shield_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClaimFilesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClaimFilesRequest) ProtoMessage() {}

func (x *ClaimFilesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_colab_shield_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClaimFilesRequest.ProtoReflect.Descriptor instead.
func (*ClaimFilesRequest) Descriptor() ([]byte, []int) {
	return file_colab_shield_proto_rawDescGZIP(), []int{8}
}

func (x *ClaimFilesRequest) GetBranchName() string {
	if x != nil {
		return x.BranchName
	}
	return ""
}

func (x *ClaimFilesRequest) GetFiles() []*ClaimFileInfo {
	if x != nil {
		return x.Files
	}
	return nil
}

type ClaimFilesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status        Status      `protobuf:"varint,1,opt,name=status,proto3,enum=colabshield.Status" json:"status,omitempty"`
	RejectedFiles []*FileInfo `protobuf:"bytes,2,rep,name=rejectedFiles,proto3" json:"rejectedFiles,omitempty"`
}

func (x *ClaimFilesResponse) Reset() {
	*x = ClaimFilesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_colab_shield_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClaimFilesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClaimFilesResponse) ProtoMessage() {}

func (x *ClaimFilesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_colab_shield_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClaimFilesResponse.ProtoReflect.Descriptor instead.
func (*ClaimFilesResponse) Descriptor() ([]byte, []int) {
	return file_colab_shield_proto_rawDescGZIP(), []int{9}
}

func (x *ClaimFilesResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_OK
}

func (x *ClaimFilesResponse) GetRejectedFiles() []*FileInfo {
	if x != nil {
		return x.RejectedFiles
	}
	return nil
}

type FileUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId     string `protobuf:"bytes,1,opt,name=fileId,proto3" json:"fileId,omitempty"`
	FileHash   string `protobuf:"bytes,2,opt,name=fileHash,proto3" json:"fileHash,omitempty"`
	BranchName string `protobuf:"bytes,3,opt,name=branchName,proto3" json:"branchName,omitempty"`
}

func (x *FileUpdateRequest) Reset() {
	*x = FileUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_colab_shield_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileUpdateRequest) ProtoMessage() {}

func (x *FileUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_colab_shield_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileUpdateRequest.ProtoReflect.Descriptor instead.
func (*FileUpdateRequest) Descriptor() ([]byte, []int) {
	return file_colab_shield_proto_rawDescGZIP(), []int{10}
}

func (x *FileUpdateRequest) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *FileUpdateRequest) GetFileHash() string {
	if x != nil {
		return x.FileHash
	}
	return ""
}

func (x *FileUpdateRequest) GetBranchName() string {
	if x != nil {
		return x.BranchName
	}
	return ""
}

type FileUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status Status `protobuf:"varint,1,opt,name=status,proto3,enum=colabshield.Status" json:"status,omitempty"`
}

func (x *FileUpdateResponse) Reset() {
	*x = FileUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_colab_shield_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileUpdateResponse) ProtoMessage() {}

func (x *FileUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_colab_shield_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileUpdateResponse.ProtoReflect.Descriptor instead.
func (*FileUpdateResponse) Descriptor() ([]byte, []int) {
	return file_colab_shield_proto_rawDescGZIP(), []int{11}
}

func (x *FileUpdateResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_OK
}

var File_colab_shield_proto protoreflect.FileDescriptor

var file_colab_shield_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x6f, 0x6c, 0x61, 0x62, 0x2d, 0x73, 0x68, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x6f, 0x6c, 0x61, 0x62, 0x73, 0x68, 0x69, 0x65, 0x6c,
	0x64, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42,
	0x0a, 0x13, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6c, 0x61, 0x62, 0x73, 0x68, 0x69,
	0x65, 0x6c, 0x64, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x32, 0x0a, 0x12, 0x49, 0x6e, 0x69, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x22, 0x42, 0x0a, 0x13, 0x49, 0x6e, 0x69, 0x74, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e,
	0x63, 0x6f, 0x6c, 0x61, 0x62, 0x73, 0x68, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x32, 0x0a, 0x14, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x22, 0x30,
	0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64,
	0x22, 0xae, 0x01, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a,
	0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66,
	0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x48, 0x61, 0x73,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x48, 0x61, 0x73,
	0x68, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x62,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x63,
	0x6c, 0x61, 0x69, 0x6d, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16,
	0x2e, 0x63, 0x6f, 0x6c, 0x61, 0x62, 0x73, 0x68, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x43, 0x6c, 0x61,
	0x69, 0x6d, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x09, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x4d, 0x6f, 0x64,
	0x65, 0x22, 0x40, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6c, 0x61, 0x62, 0x73, 0x68, 0x69,
	0x65, 0x6c, 0x64, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x22, 0x79, 0x0a, 0x0d, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x46, 0x69, 0x6c, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x48, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12, 0x34, 0x0a, 0x09, 0x63, 0x6c, 0x61, 0x69,
	0x6d, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x63, 0x6f,
	0x6c, 0x61, 0x62, 0x73, 0x68, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x4d,
	0x6f, 0x64, 0x65, 0x52, 0x09, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x4d, 0x6f, 0x64, 0x65, 0x22, 0x65,
	0x0a, 0x11, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x6c, 0x61, 0x62, 0x73, 0x68, 0x69, 0x65, 0x6c, 0x64,
	0x2e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x22, 0x7e, 0x0a, 0x12, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x46, 0x69,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x63, 0x6f,
	0x6c, 0x61, 0x62, 0x73, 0x68, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3b, 0x0a, 0x0d, 0x72, 0x65, 0x6a, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x63, 0x6f, 0x6c, 0x61, 0x62, 0x73, 0x68, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x46, 0x69, 0x6c, 0x65, 0x73, 0x22, 0x67, 0x0a, 0x11, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69,
	0x6c, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x48, 0x61, 0x73, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1e,
	0x0a, 0x0a, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x41,
	0x0a, 0x12, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6c, 0x61, 0x62, 0x73, 0x68, 0x69, 0x65,
	0x6c, 0x64, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2a, 0x1b, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x06, 0x0a, 0x02, 0x4f,
	0x4b, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x2a, 0x35,
	0x0a, 0x09, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x55,
	0x4e, 0x43, 0x4c, 0x41, 0x49, 0x4d, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x45, 0x58,
	0x43, 0x4c, 0x55, 0x53, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x48, 0x41,
	0x52, 0x45, 0x44, 0x10, 0x02, 0x32, 0x93, 0x03, 0x0a, 0x0b, 0x43, 0x6f, 0x6c, 0x61, 0x62, 0x53,
	0x68, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x49, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x20, 0x2e, 0x63,
	0x6f, 0x6c, 0x61, 0x62, 0x73, 0x68, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x52, 0x0a, 0x0b, 0x49, 0x6e, 0x69, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x1f, 0x2e, 0x63, 0x6f, 0x6c, 0x61, 0x62, 0x73, 0x68, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x49, 0x6e,
	0x69, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x63, 0x6f, 0x6c, 0x61, 0x62, 0x73, 0x68, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x49,
	0x6e, 0x69, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x21, 0x2e, 0x63,
	0x6f, 0x6c, 0x61, 0x62, 0x73, 0x68, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x4c, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x1d,
	0x2e, 0x63, 0x6f, 0x6c, 0x61, 0x62, 0x73, 0x68, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x63, 0x6f, 0x6c, 0x61, 0x62, 0x73, 0x68, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x4a, 0x0a, 0x05, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x12, 0x1e, 0x2e, 0x63, 0x6f, 0x6c, 0x61, 0x62,
	0x73, 0x68, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x46, 0x69, 0x6c, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x63, 0x6f, 0x6c, 0x61, 0x62,
	0x73, 0x68, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x46, 0x69, 0x6c, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x29, 0x5a, 0x27, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6c, 0x69, 0x79, 0x61, 0x6e,
	0x6b, 0x67, 0x2f, 0x63, 0x6f, 0x6c, 0x61, 0x62, 0x2d, 0x73, 0x68, 0x69, 0x65, 0x6c, 0x64, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_colab_shield_proto_rawDescOnce sync.Once
	file_colab_shield_proto_rawDescData = file_colab_shield_proto_rawDesc
)

func file_colab_shield_proto_rawDescGZIP() []byte {
	file_colab_shield_proto_rawDescOnce.Do(func() {
		file_colab_shield_proto_rawDescData = protoimpl.X.CompressGZIP(file_colab_shield_proto_rawDescData)
	})
	return file_colab_shield_proto_rawDescData
}

var file_colab_shield_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_colab_shield_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_colab_shield_proto_goTypes = []interface{}{
	(Status)(0),                  // 0: colabshield.Status
	(ClaimMode)(0),               // 1: colabshield.ClaimMode
	(*HealthCheckResponse)(nil),  // 2: colabshield.HealthCheckResponse
	(*InitProjectRequest)(nil),   // 3: colabshield.InitProjectRequest
	(*InitProjectResponse)(nil),  // 4: colabshield.InitProjectResponse
	(*ListProjectsResponse)(nil), // 5: colabshield.ListProjectsResponse
	(*ListFilesRequest)(nil),     // 6: colabshield.ListFilesRequest
	(*FileInfo)(nil),             // 7: colabshield.FileInfo
	(*ListFilesResponse)(nil),    // 8: colabshield.ListFilesResponse
	(*ClaimFileInfo)(nil),        // 9: colabshield.ClaimFileInfo
	(*ClaimFilesRequest)(nil),    // 10: colabshield.ClaimFilesRequest
	(*ClaimFilesResponse)(nil),   // 11: colabshield.ClaimFilesResponse
	(*FileUpdateRequest)(nil),    // 12: colabshield.FileUpdateRequest
	(*FileUpdateResponse)(nil),   // 13: colabshield.FileUpdateResponse
	(*emptypb.Empty)(nil),        // 14: google.protobuf.Empty
}
var file_colab_shield_proto_depIdxs = []int32{
	0,  // 0: colabshield.HealthCheckResponse.status:type_name -> colabshield.Status
	0,  // 1: colabshield.InitProjectResponse.status:type_name -> colabshield.Status
	1,  // 2: colabshield.FileInfo.claimMode:type_name -> colabshield.ClaimMode
	7,  // 3: colabshield.ListFilesResponse.files:type_name -> colabshield.FileInfo
	1,  // 4: colabshield.ClaimFileInfo.claimMode:type_name -> colabshield.ClaimMode
	9,  // 5: colabshield.ClaimFilesRequest.files:type_name -> colabshield.ClaimFileInfo
	0,  // 6: colabshield.ClaimFilesResponse.status:type_name -> colabshield.Status
	7,  // 7: colabshield.ClaimFilesResponse.rejectedFiles:type_name -> colabshield.FileInfo
	0,  // 8: colabshield.FileUpdateResponse.status:type_name -> colabshield.Status
	14, // 9: colabshield.ColabShield.HealthCheck:input_type -> google.protobuf.Empty
	3,  // 10: colabshield.ColabShield.InitProject:input_type -> colabshield.InitProjectRequest
	14, // 11: colabshield.ColabShield.ListProjects:input_type -> google.protobuf.Empty
	6,  // 12: colabshield.ColabShield.ListFiles:input_type -> colabshield.ListFilesRequest
	10, // 13: colabshield.ColabShield.Claim:input_type -> colabshield.ClaimFilesRequest
	2,  // 14: colabshield.ColabShield.HealthCheck:output_type -> colabshield.HealthCheckResponse
	4,  // 15: colabshield.ColabShield.InitProject:output_type -> colabshield.InitProjectResponse
	5,  // 16: colabshield.ColabShield.ListProjects:output_type -> colabshield.ListProjectsResponse
	8,  // 17: colabshield.ColabShield.ListFiles:output_type -> colabshield.ListFilesResponse
	11, // 18: colabshield.ColabShield.Claim:output_type -> colabshield.ClaimFilesResponse
	14, // [14:19] is the sub-list for method output_type
	9,  // [9:14] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_colab_shield_proto_init() }
func file_colab_shield_proto_init() {
	if File_colab_shield_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_colab_shield_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckResponse); i {
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
		file_colab_shield_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitProjectRequest); i {
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
		file_colab_shield_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitProjectResponse); i {
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
		file_colab_shield_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProjectsResponse); i {
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
		file_colab_shield_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFilesRequest); i {
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
		file_colab_shield_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileInfo); i {
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
		file_colab_shield_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFilesResponse); i {
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
		file_colab_shield_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClaimFileInfo); i {
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
		file_colab_shield_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClaimFilesRequest); i {
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
		file_colab_shield_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClaimFilesResponse); i {
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
		file_colab_shield_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileUpdateRequest); i {
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
		file_colab_shield_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileUpdateResponse); i {
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
			RawDescriptor: file_colab_shield_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_colab_shield_proto_goTypes,
		DependencyIndexes: file_colab_shield_proto_depIdxs,
		EnumInfos:         file_colab_shield_proto_enumTypes,
		MessageInfos:      file_colab_shield_proto_msgTypes,
	}.Build()
	File_colab_shield_proto = out.File
	file_colab_shield_proto_rawDesc = nil
	file_colab_shield_proto_goTypes = nil
	file_colab_shield_proto_depIdxs = nil
}
