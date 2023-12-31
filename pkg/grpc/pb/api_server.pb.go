// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.6
// source: proto/api_server.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateApplicationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string                     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Desc        string                     `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	CreatedBy   string                     `protobuf:"bytes,3,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	Permissions []*CreatePermissionRequest `protobuf:"bytes,4,rep,name=permissions,proto3" json:"permissions,omitempty"`
}

func (x *CreateApplicationRequest) Reset() {
	*x = CreateApplicationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateApplicationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateApplicationRequest) ProtoMessage() {}

func (x *CreateApplicationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateApplicationRequest.ProtoReflect.Descriptor instead.
func (*CreateApplicationRequest) Descriptor() ([]byte, []int) {
	return file_proto_api_server_proto_rawDescGZIP(), []int{0}
}

func (x *CreateApplicationRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateApplicationRequest) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *CreateApplicationRequest) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *CreateApplicationRequest) GetPermissions() []*CreatePermissionRequest {
	if x != nil {
		return x.Permissions
	}
	return nil
}

type CreateApplicationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          string                `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name        string                `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description *string               `protobuf:"bytes,3,opt,name=description,proto3,oneof" json:"description,omitempty"`
	CreatedBy   string                `protobuf:"bytes,4,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	Slug        string                `protobuf:"bytes,5,opt,name=slug,proto3" json:"slug,omitempty"`
	TeamID      string                `protobuf:"bytes,6,opt,name=teamID,proto3" json:"teamID,omitempty"`
	Permissions []*PermissionResponse `protobuf:"bytes,7,rep,name=permissions,proto3" json:"permissions,omitempty"`
}

func (x *CreateApplicationResponse) Reset() {
	*x = CreateApplicationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateApplicationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateApplicationResponse) ProtoMessage() {}

func (x *CreateApplicationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateApplicationResponse.ProtoReflect.Descriptor instead.
func (*CreateApplicationResponse) Descriptor() ([]byte, []int) {
	return file_proto_api_server_proto_rawDescGZIP(), []int{1}
}

func (x *CreateApplicationResponse) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *CreateApplicationResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateApplicationResponse) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *CreateApplicationResponse) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *CreateApplicationResponse) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *CreateApplicationResponse) GetTeamID() string {
	if x != nil {
		return x.TeamID
	}
	return ""
}

func (x *CreateApplicationResponse) GetPermissions() []*PermissionResponse {
	if x != nil {
		return x.Permissions
	}
	return nil
}

type CreatePermissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApplicationID       string `protobuf:"bytes,1,opt,name=applicationID,proto3" json:"applicationID,omitempty"`
	ApplicationResource string `protobuf:"bytes,2,opt,name=applicationResource,proto3" json:"applicationResource,omitempty"`
	Action              string `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
}

func (x *CreatePermissionRequest) Reset() {
	*x = CreatePermissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePermissionRequest) ProtoMessage() {}

func (x *CreatePermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePermissionRequest.ProtoReflect.Descriptor instead.
func (*CreatePermissionRequest) Descriptor() ([]byte, []int) {
	return file_proto_api_server_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePermissionRequest) GetApplicationID() string {
	if x != nil {
		return x.ApplicationID
	}
	return ""
}

func (x *CreatePermissionRequest) GetApplicationResource() string {
	if x != nil {
		return x.ApplicationResource
	}
	return ""
}

func (x *CreatePermissionRequest) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

type PermissionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID                  string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	ApplicationID       string `protobuf:"bytes,2,opt,name=applicationID,proto3" json:"applicationID,omitempty"`
	ApplicationResource string `protobuf:"bytes,3,opt,name=applicationResource,proto3" json:"applicationResource,omitempty"`
	Action              string `protobuf:"bytes,4,opt,name=action,proto3" json:"action,omitempty"`
	CreatedBy           string `protobuf:"bytes,5,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
}

func (x *PermissionResponse) Reset() {
	*x = PermissionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionResponse) ProtoMessage() {}

func (x *PermissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionResponse.ProtoReflect.Descriptor instead.
func (*PermissionResponse) Descriptor() ([]byte, []int) {
	return file_proto_api_server_proto_rawDescGZIP(), []int{3}
}

func (x *PermissionResponse) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *PermissionResponse) GetApplicationID() string {
	if x != nil {
		return x.ApplicationID
	}
	return ""
}

func (x *PermissionResponse) GetApplicationResource() string {
	if x != nil {
		return x.ApplicationResource
	}
	return ""
}

func (x *PermissionResponse) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *PermissionResponse) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

var File_proto_api_server_proto protoreflect.FileDescriptor

var file_proto_api_server_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0xa1,
	0x01, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64,
	0x65, 0x73, 0x63, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42,
	0x79, 0x12, 0x3f, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x22, 0xfc, 0x01, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75,
	0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x12, 0x16, 0x0a,
	0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74,
	0x65, 0x61, 0x6d, 0x49, 0x44, 0x12, 0x3a, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x89, 0x01, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a,
	0x0d, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x44, 0x12, 0x30, 0x0a, 0x13, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x13, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xb2, 0x01,
	0x0a, 0x12, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x30, 0x0a, 0x13, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42,
	0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x42, 0x79, 0x32, 0x58, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x49, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0d, 0x5a, 0x0b,
	0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_api_server_proto_rawDescOnce sync.Once
	file_proto_api_server_proto_rawDescData = file_proto_api_server_proto_rawDesc
)

func file_proto_api_server_proto_rawDescGZIP() []byte {
	file_proto_api_server_proto_rawDescOnce.Do(func() {
		file_proto_api_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_api_server_proto_rawDescData)
	})
	return file_proto_api_server_proto_rawDescData
}

var file_proto_api_server_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_api_server_proto_goTypes = []interface{}{
	(*CreateApplicationRequest)(nil),  // 0: grpc.CreateApplicationRequest
	(*CreateApplicationResponse)(nil), // 1: grpc.CreateApplicationResponse
	(*CreatePermissionRequest)(nil),   // 2: grpc.CreatePermissionRequest
	(*PermissionResponse)(nil),        // 3: grpc.PermissionResponse
}
var file_proto_api_server_proto_depIdxs = []int32{
	2, // 0: grpc.CreateApplicationRequest.permissions:type_name -> grpc.CreatePermissionRequest
	3, // 1: grpc.CreateApplicationResponse.permissions:type_name -> grpc.PermissionResponse
	0, // 2: grpc.Application.Create:input_type -> grpc.CreateApplicationRequest
	1, // 3: grpc.Application.Create:output_type -> grpc.CreateApplicationResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_api_server_proto_init() }
func file_proto_api_server_proto_init() {
	if File_proto_api_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_api_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateApplicationRequest); i {
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
		file_proto_api_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateApplicationResponse); i {
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
		file_proto_api_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePermissionRequest); i {
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
		file_proto_api_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionResponse); i {
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
	file_proto_api_server_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_api_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_api_server_proto_goTypes,
		DependencyIndexes: file_proto_api_server_proto_depIdxs,
		MessageInfos:      file_proto_api_server_proto_msgTypes,
	}.Build()
	File_proto_api_server_proto = out.File
	file_proto_api_server_proto_rawDesc = nil
	file_proto_api_server_proto_goTypes = nil
	file_proto_api_server_proto_depIdxs = nil
}
