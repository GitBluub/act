// Copyright 2024 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: artifact.proto

package artifacts

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateArtifactRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkflowRunBackendId    string                 `protobuf:"bytes,1,opt,name=workflow_run_backend_id,json=workflowRunBackendId,proto3" json:"workflow_run_backend_id,omitempty"`
	WorkflowJobRunBackendId string                 `protobuf:"bytes,2,opt,name=workflow_job_run_backend_id,json=workflowJobRunBackendId,proto3" json:"workflow_job_run_backend_id,omitempty"`
	Name                    string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	ExpiresAt               *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	Version                 int32                  `protobuf:"varint,5,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *CreateArtifactRequest) Reset() {
	*x = CreateArtifactRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateArtifactRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateArtifactRequest) ProtoMessage() {}

func (x *CreateArtifactRequest) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateArtifactRequest.ProtoReflect.Descriptor instead.
func (*CreateArtifactRequest) Descriptor() ([]byte, []int) {
	return file_artifact_proto_rawDescGZIP(), []int{0}
}

func (x *CreateArtifactRequest) GetWorkflowRunBackendId() string {
	if x != nil {
		return x.WorkflowRunBackendId
	}
	return ""
}

func (x *CreateArtifactRequest) GetWorkflowJobRunBackendId() string {
	if x != nil {
		return x.WorkflowJobRunBackendId
	}
	return ""
}

func (x *CreateArtifactRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateArtifactRequest) GetExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiresAt
	}
	return nil
}

func (x *CreateArtifactRequest) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

type CreateArtifactResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok              bool   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	SignedUploadUrl string `protobuf:"bytes,2,opt,name=signed_upload_url,json=signedUploadUrl,proto3" json:"signed_upload_url,omitempty"`
}

func (x *CreateArtifactResponse) Reset() {
	*x = CreateArtifactResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateArtifactResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateArtifactResponse) ProtoMessage() {}

func (x *CreateArtifactResponse) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateArtifactResponse.ProtoReflect.Descriptor instead.
func (*CreateArtifactResponse) Descriptor() ([]byte, []int) {
	return file_artifact_proto_rawDescGZIP(), []int{1}
}

func (x *CreateArtifactResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *CreateArtifactResponse) GetSignedUploadUrl() string {
	if x != nil {
		return x.SignedUploadUrl
	}
	return ""
}

type FinalizeArtifactRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkflowRunBackendId    string                  `protobuf:"bytes,1,opt,name=workflow_run_backend_id,json=workflowRunBackendId,proto3" json:"workflow_run_backend_id,omitempty"`
	WorkflowJobRunBackendId string                  `protobuf:"bytes,2,opt,name=workflow_job_run_backend_id,json=workflowJobRunBackendId,proto3" json:"workflow_job_run_backend_id,omitempty"`
	Name                    string                  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Size                    int64                   `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
	Hash                    *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *FinalizeArtifactRequest) Reset() {
	*x = FinalizeArtifactRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinalizeArtifactRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinalizeArtifactRequest) ProtoMessage() {}

func (x *FinalizeArtifactRequest) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinalizeArtifactRequest.ProtoReflect.Descriptor instead.
func (*FinalizeArtifactRequest) Descriptor() ([]byte, []int) {
	return file_artifact_proto_rawDescGZIP(), []int{2}
}

func (x *FinalizeArtifactRequest) GetWorkflowRunBackendId() string {
	if x != nil {
		return x.WorkflowRunBackendId
	}
	return ""
}

func (x *FinalizeArtifactRequest) GetWorkflowJobRunBackendId() string {
	if x != nil {
		return x.WorkflowJobRunBackendId
	}
	return ""
}

func (x *FinalizeArtifactRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FinalizeArtifactRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *FinalizeArtifactRequest) GetHash() *wrapperspb.StringValue {
	if x != nil {
		return x.Hash
	}
	return nil
}

type FinalizeArtifactResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok         bool  `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	ArtifactId int64 `protobuf:"varint,2,opt,name=artifact_id,json=artifactId,proto3" json:"artifact_id,omitempty"`
}

func (x *FinalizeArtifactResponse) Reset() {
	*x = FinalizeArtifactResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinalizeArtifactResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinalizeArtifactResponse) ProtoMessage() {}

func (x *FinalizeArtifactResponse) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinalizeArtifactResponse.ProtoReflect.Descriptor instead.
func (*FinalizeArtifactResponse) Descriptor() ([]byte, []int) {
	return file_artifact_proto_rawDescGZIP(), []int{3}
}

func (x *FinalizeArtifactResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *FinalizeArtifactResponse) GetArtifactId() int64 {
	if x != nil {
		return x.ArtifactId
	}
	return 0
}

type ListArtifactsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkflowRunBackendId    string                  `protobuf:"bytes,1,opt,name=workflow_run_backend_id,json=workflowRunBackendId,proto3" json:"workflow_run_backend_id,omitempty"`
	WorkflowJobRunBackendId string                  `protobuf:"bytes,2,opt,name=workflow_job_run_backend_id,json=workflowJobRunBackendId,proto3" json:"workflow_job_run_backend_id,omitempty"`
	NameFilter              *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=name_filter,json=nameFilter,proto3" json:"name_filter,omitempty"`
	IdFilter                *wrapperspb.Int64Value  `protobuf:"bytes,4,opt,name=id_filter,json=idFilter,proto3" json:"id_filter,omitempty"`
}

func (x *ListArtifactsRequest) Reset() {
	*x = ListArtifactsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListArtifactsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListArtifactsRequest) ProtoMessage() {}

func (x *ListArtifactsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListArtifactsRequest.ProtoReflect.Descriptor instead.
func (*ListArtifactsRequest) Descriptor() ([]byte, []int) {
	return file_artifact_proto_rawDescGZIP(), []int{4}
}

func (x *ListArtifactsRequest) GetWorkflowRunBackendId() string {
	if x != nil {
		return x.WorkflowRunBackendId
	}
	return ""
}

func (x *ListArtifactsRequest) GetWorkflowJobRunBackendId() string {
	if x != nil {
		return x.WorkflowJobRunBackendId
	}
	return ""
}

func (x *ListArtifactsRequest) GetNameFilter() *wrapperspb.StringValue {
	if x != nil {
		return x.NameFilter
	}
	return nil
}

func (x *ListArtifactsRequest) GetIdFilter() *wrapperspb.Int64Value {
	if x != nil {
		return x.IdFilter
	}
	return nil
}

type ListArtifactsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Artifacts []*ListArtifactsResponse_MonolithArtifact `protobuf:"bytes,1,rep,name=artifacts,proto3" json:"artifacts,omitempty"`
}

func (x *ListArtifactsResponse) Reset() {
	*x = ListArtifactsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListArtifactsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListArtifactsResponse) ProtoMessage() {}

func (x *ListArtifactsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListArtifactsResponse.ProtoReflect.Descriptor instead.
func (*ListArtifactsResponse) Descriptor() ([]byte, []int) {
	return file_artifact_proto_rawDescGZIP(), []int{5}
}

func (x *ListArtifactsResponse) GetArtifacts() []*ListArtifactsResponse_MonolithArtifact {
	if x != nil {
		return x.Artifacts
	}
	return nil
}

type ListArtifactsResponse_MonolithArtifact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkflowRunBackendId    string                 `protobuf:"bytes,1,opt,name=workflow_run_backend_id,json=workflowRunBackendId,proto3" json:"workflow_run_backend_id,omitempty"`
	WorkflowJobRunBackendId string                 `protobuf:"bytes,2,opt,name=workflow_job_run_backend_id,json=workflowJobRunBackendId,proto3" json:"workflow_job_run_backend_id,omitempty"`
	DatabaseId              int64                  `protobuf:"varint,3,opt,name=database_id,json=databaseId,proto3" json:"database_id,omitempty"`
	Name                    string                 `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Size                    int64                  `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
	CreatedAt               *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *ListArtifactsResponse_MonolithArtifact) Reset() {
	*x = ListArtifactsResponse_MonolithArtifact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListArtifactsResponse_MonolithArtifact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListArtifactsResponse_MonolithArtifact) ProtoMessage() {}

func (x *ListArtifactsResponse_MonolithArtifact) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListArtifactsResponse_MonolithArtifact.ProtoReflect.Descriptor instead.
func (*ListArtifactsResponse_MonolithArtifact) Descriptor() ([]byte, []int) {
	return file_artifact_proto_rawDescGZIP(), []int{6}
}

func (x *ListArtifactsResponse_MonolithArtifact) GetWorkflowRunBackendId() string {
	if x != nil {
		return x.WorkflowRunBackendId
	}
	return ""
}

func (x *ListArtifactsResponse_MonolithArtifact) GetWorkflowJobRunBackendId() string {
	if x != nil {
		return x.WorkflowJobRunBackendId
	}
	return ""
}

func (x *ListArtifactsResponse_MonolithArtifact) GetDatabaseId() int64 {
	if x != nil {
		return x.DatabaseId
	}
	return 0
}

func (x *ListArtifactsResponse_MonolithArtifact) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListArtifactsResponse_MonolithArtifact) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListArtifactsResponse_MonolithArtifact) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type GetSignedArtifactURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkflowRunBackendId    string `protobuf:"bytes,1,opt,name=workflow_run_backend_id,json=workflowRunBackendId,proto3" json:"workflow_run_backend_id,omitempty"`
	WorkflowJobRunBackendId string `protobuf:"bytes,2,opt,name=workflow_job_run_backend_id,json=workflowJobRunBackendId,proto3" json:"workflow_job_run_backend_id,omitempty"`
	Name                    string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetSignedArtifactURLRequest) Reset() {
	*x = GetSignedArtifactURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSignedArtifactURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSignedArtifactURLRequest) ProtoMessage() {}

func (x *GetSignedArtifactURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSignedArtifactURLRequest.ProtoReflect.Descriptor instead.
func (*GetSignedArtifactURLRequest) Descriptor() ([]byte, []int) {
	return file_artifact_proto_rawDescGZIP(), []int{7}
}

func (x *GetSignedArtifactURLRequest) GetWorkflowRunBackendId() string {
	if x != nil {
		return x.WorkflowRunBackendId
	}
	return ""
}

func (x *GetSignedArtifactURLRequest) GetWorkflowJobRunBackendId() string {
	if x != nil {
		return x.WorkflowJobRunBackendId
	}
	return ""
}

func (x *GetSignedArtifactURLRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetSignedArtifactURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SignedUrl string `protobuf:"bytes,1,opt,name=signed_url,json=signedUrl,proto3" json:"signed_url,omitempty"`
}

func (x *GetSignedArtifactURLResponse) Reset() {
	*x = GetSignedArtifactURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSignedArtifactURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSignedArtifactURLResponse) ProtoMessage() {}

func (x *GetSignedArtifactURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSignedArtifactURLResponse.ProtoReflect.Descriptor instead.
func (*GetSignedArtifactURLResponse) Descriptor() ([]byte, []int) {
	return file_artifact_proto_rawDescGZIP(), []int{8}
}

func (x *GetSignedArtifactURLResponse) GetSignedUrl() string {
	if x != nil {
		return x.SignedUrl
	}
	return ""
}

type DeleteArtifactRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkflowRunBackendId    string `protobuf:"bytes,1,opt,name=workflow_run_backend_id,json=workflowRunBackendId,proto3" json:"workflow_run_backend_id,omitempty"`
	WorkflowJobRunBackendId string `protobuf:"bytes,2,opt,name=workflow_job_run_backend_id,json=workflowJobRunBackendId,proto3" json:"workflow_job_run_backend_id,omitempty"`
	Name                    string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteArtifactRequest) Reset() {
	*x = DeleteArtifactRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteArtifactRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteArtifactRequest) ProtoMessage() {}

func (x *DeleteArtifactRequest) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteArtifactRequest.ProtoReflect.Descriptor instead.
func (*DeleteArtifactRequest) Descriptor() ([]byte, []int) {
	return file_artifact_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteArtifactRequest) GetWorkflowRunBackendId() string {
	if x != nil {
		return x.WorkflowRunBackendId
	}
	return ""
}

func (x *DeleteArtifactRequest) GetWorkflowJobRunBackendId() string {
	if x != nil {
		return x.WorkflowJobRunBackendId
	}
	return ""
}

func (x *DeleteArtifactRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeleteArtifactResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok         bool  `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	ArtifactId int64 `protobuf:"varint,2,opt,name=artifact_id,json=artifactId,proto3" json:"artifact_id,omitempty"`
}

func (x *DeleteArtifactResponse) Reset() {
	*x = DeleteArtifactResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteArtifactResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteArtifactResponse) ProtoMessage() {}

func (x *DeleteArtifactResponse) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteArtifactResponse.ProtoReflect.Descriptor instead.
func (*DeleteArtifactResponse) Descriptor() ([]byte, []int) {
	return file_artifact_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteArtifactResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *DeleteArtifactResponse) GetArtifactId() int64 {
	if x != nil {
		return x.ArtifactId
	}
	return 0
}

var File_artifact_proto protoreflect.FileDescriptor

var file_artifact_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xf5, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x17, 0x77, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x72, 0x75, 0x6e, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x77, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x75, 0x6e, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x49,
	0x64, 0x12, 0x3c, 0x0a, 0x1b, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x6a, 0x6f,
	0x62, 0x5f, 0x72, 0x75, 0x6e, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x4a, 0x6f, 0x62, 0x52, 0x75, 0x6e, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x54, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02,
	0x6f, 0x6b, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f, 0x75, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73,
	0x69, 0x67, 0x6e, 0x65, 0x64, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72, 0x6c, 0x22, 0xe8,
	0x01, 0x0a, 0x17, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x41, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x17, 0x77, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x72, 0x75, 0x6e, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x77, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x75, 0x6e, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x49,
	0x64, 0x12, 0x3c, 0x0a, 0x1b, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x6a, 0x6f,
	0x62, 0x5f, 0x72, 0x75, 0x6e, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x4a, 0x6f, 0x62, 0x52, 0x75, 0x6e, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x4b, 0x0a, 0x18, 0x46, 0x69, 0x6e,
	0x61, 0x6c, 0x69, 0x7a, 0x65, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x61, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x49, 0x64, 0x22, 0x84, 0x02, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x41,
	0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x35, 0x0a, 0x17, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x72, 0x75, 0x6e, 0x5f,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x14, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x75, 0x6e, 0x42, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x3c, 0x0a, 0x1b, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x72, 0x75, 0x6e, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x77, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x4a, 0x6f, 0x62, 0x52, 0x75, 0x6e, 0x42, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x09, 0x69, 0x64, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x08, 0x69, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x7c, 0x0a,
	0x15, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x63, 0x0a, 0x09, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x45, 0x2e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f,
	0x4d, 0x6f, 0x6e, 0x6f, 0x6c, 0x69, 0x74, 0x68, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x52, 0x09, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x22, 0xa1, 0x02, 0x0a, 0x26,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x4d, 0x6f, 0x6e, 0x6f, 0x6c, 0x69, 0x74, 0x68, 0x41, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x12, 0x35, 0x0a, 0x17, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x5f, 0x72, 0x75, 0x6e, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x52, 0x75, 0x6e, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x3c, 0x0a,
	0x1b, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x72, 0x75,
	0x6e, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x17, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x4a, 0x6f, 0x62, 0x52,
	0x75, 0x6e, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0xa6, 0x01, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x41, 0x72, 0x74,
	0x69, 0x66, 0x61, 0x63, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x35, 0x0a, 0x17, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x72, 0x75, 0x6e, 0x5f,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x14, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x75, 0x6e, 0x42, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x3c, 0x0a, 0x1b, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x72, 0x75, 0x6e, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x77, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x4a, 0x6f, 0x62, 0x52, 0x75, 0x6e, 0x42, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3d, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x53,
	0x69, 0x67, 0x6e, 0x65, 0x64, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x55, 0x52, 0x4c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x69, 0x67, 0x6e,
	0x65, 0x64, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x65, 0x64, 0x55, 0x72, 0x6c, 0x22, 0xa0, 0x01, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x35, 0x0a, 0x17, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x72, 0x75,
	0x6e, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x14, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x75, 0x6e, 0x42,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x3c, 0x0a, 0x1b, 0x77, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x72, 0x75, 0x6e, 0x5f, 0x62, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x77,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x4a, 0x6f, 0x62, 0x52, 0x75, 0x6e, 0x42, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x49, 0x0a, 0x16, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x02, 0x6f, 0x6b, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x61, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x49, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_artifact_proto_rawDescOnce sync.Once
	file_artifact_proto_rawDescData = file_artifact_proto_rawDesc
)

func file_artifact_proto_rawDescGZIP() []byte {
	file_artifact_proto_rawDescOnce.Do(func() {
		file_artifact_proto_rawDescData = protoimpl.X.CompressGZIP(file_artifact_proto_rawDescData)
	})
	return file_artifact_proto_rawDescData
}

var (
	file_artifact_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
	file_artifact_proto_goTypes  = []interface{}{
		(*CreateArtifactRequest)(nil),                  // 0: github.actions.results.api.v1.CreateArtifactRequest
		(*CreateArtifactResponse)(nil),                 // 1: github.actions.results.api.v1.CreateArtifactResponse
		(*FinalizeArtifactRequest)(nil),                // 2: github.actions.results.api.v1.FinalizeArtifactRequest
		(*FinalizeArtifactResponse)(nil),               // 3: github.actions.results.api.v1.FinalizeArtifactResponse
		(*ListArtifactsRequest)(nil),                   // 4: github.actions.results.api.v1.ListArtifactsRequest
		(*ListArtifactsResponse)(nil),                  // 5: github.actions.results.api.v1.ListArtifactsResponse
		(*ListArtifactsResponse_MonolithArtifact)(nil), // 6: github.actions.results.api.v1.ListArtifactsResponse_MonolithArtifact
		(*GetSignedArtifactURLRequest)(nil),            // 7: github.actions.results.api.v1.GetSignedArtifactURLRequest
		(*GetSignedArtifactURLResponse)(nil),           // 8: github.actions.results.api.v1.GetSignedArtifactURLResponse
		(*DeleteArtifactRequest)(nil),                  // 9: github.actions.results.api.v1.DeleteArtifactRequest
		(*DeleteArtifactResponse)(nil),                 // 10: github.actions.results.api.v1.DeleteArtifactResponse
		(*timestamppb.Timestamp)(nil),                  // 11: google.protobuf.Timestamp
		(*wrapperspb.StringValue)(nil),                 // 12: google.protobuf.StringValue
		(*wrapperspb.Int64Value)(nil),                  // 13: google.protobuf.Int64Value
	}
)

var file_artifact_proto_depIdxs = []int32{
	11, // 0: github.actions.results.api.v1.CreateArtifactRequest.expires_at:type_name -> google.protobuf.Timestamp
	12, // 1: github.actions.results.api.v1.FinalizeArtifactRequest.hash:type_name -> google.protobuf.StringValue
	12, // 2: github.actions.results.api.v1.ListArtifactsRequest.name_filter:type_name -> google.protobuf.StringValue
	13, // 3: github.actions.results.api.v1.ListArtifactsRequest.id_filter:type_name -> google.protobuf.Int64Value
	6,  // 4: github.actions.results.api.v1.ListArtifactsResponse.artifacts:type_name -> github.actions.results.api.v1.ListArtifactsResponse_MonolithArtifact
	11, // 5: github.actions.results.api.v1.ListArtifactsResponse_MonolithArtifact.created_at:type_name -> google.protobuf.Timestamp
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_artifact_proto_init() }
func file_artifact_proto_init() {
	if File_artifact_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_artifact_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateArtifactRequest); i {
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
		file_artifact_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateArtifactResponse); i {
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
		file_artifact_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinalizeArtifactRequest); i {
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
		file_artifact_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinalizeArtifactResponse); i {
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
		file_artifact_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListArtifactsRequest); i {
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
		file_artifact_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListArtifactsResponse); i {
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
		file_artifact_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListArtifactsResponse_MonolithArtifact); i {
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
		file_artifact_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSignedArtifactURLRequest); i {
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
		file_artifact_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSignedArtifactURLResponse); i {
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
		file_artifact_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteArtifactRequest); i {
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
		file_artifact_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteArtifactResponse); i {
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
			RawDescriptor: file_artifact_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_artifact_proto_goTypes,
		DependencyIndexes: file_artifact_proto_depIdxs,
		MessageInfos:      file_artifact_proto_msgTypes,
	}.Build()
	File_artifact_proto = out.File
	file_artifact_proto_rawDesc = nil
	file_artifact_proto_goTypes = nil
	file_artifact_proto_depIdxs = nil
}
