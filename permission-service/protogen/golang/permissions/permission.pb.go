// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: permissions/permission.proto

package permissions

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Enum for permission levels
type PermissionLevel int32

const (
	PermissionLevel_OWNER  PermissionLevel = 0
	PermissionLevel_EDITOR PermissionLevel = 1
	PermissionLevel_VIEWER PermissionLevel = 2
)

// Enum value maps for PermissionLevel.
var (
	PermissionLevel_name = map[int32]string{
		0: "OWNER",
		1: "EDITOR",
		2: "VIEWER",
	}
	PermissionLevel_value = map[string]int32{
		"OWNER":  0,
		"EDITOR": 1,
		"VIEWER": 2,
	}
)

func (x PermissionLevel) Enum() *PermissionLevel {
	p := new(PermissionLevel)
	*p = x
	return p
}

func (x PermissionLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PermissionLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_permissions_permission_proto_enumTypes[0].Descriptor()
}

func (PermissionLevel) Type() protoreflect.EnumType {
	return &file_permissions_permission_proto_enumTypes[0]
}

func (x PermissionLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PermissionLevel.Descriptor instead.
func (PermissionLevel) EnumDescriptor() ([]byte, []int) {
	return file_permissions_permission_proto_rawDescGZIP(), []int{0}
}

// Message representing a permission entry
type Permission struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileId        string                 `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Level         PermissionLevel        `protobuf:"varint,3,opt,name=level,proto3,enum=PermissionLevel" json:"level,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Permission) Reset() {
	*x = Permission{}
	mi := &file_permissions_permission_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Permission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Permission) ProtoMessage() {}

func (x *Permission) ProtoReflect() protoreflect.Message {
	mi := &file_permissions_permission_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Permission.ProtoReflect.Descriptor instead.
func (*Permission) Descriptor() ([]byte, []int) {
	return file_permissions_permission_proto_rawDescGZIP(), []int{0}
}

func (x *Permission) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *Permission) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Permission) GetLevel() PermissionLevel {
	if x != nil {
		return x.Level
	}
	return PermissionLevel_OWNER
}

// Create
type CreatePermissionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Permission    *Permission            `protobuf:"bytes,1,opt,name=permission,proto3" json:"permission,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreatePermissionRequest) Reset() {
	*x = CreatePermissionRequest{}
	mi := &file_permissions_permission_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePermissionRequest) ProtoMessage() {}

func (x *CreatePermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_permissions_permission_proto_msgTypes[1]
	if x != nil {
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
	return file_permissions_permission_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePermissionRequest) GetPermission() *Permission {
	if x != nil {
		return x.Permission
	}
	return nil
}

type CreatePermissionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Permission    *Permission            `protobuf:"bytes,1,opt,name=permission,proto3" json:"permission,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreatePermissionResponse) Reset() {
	*x = CreatePermissionResponse{}
	mi := &file_permissions_permission_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePermissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePermissionResponse) ProtoMessage() {}

func (x *CreatePermissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_permissions_permission_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePermissionResponse.ProtoReflect.Descriptor instead.
func (*CreatePermissionResponse) Descriptor() ([]byte, []int) {
	return file_permissions_permission_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePermissionResponse) GetPermission() *Permission {
	if x != nil {
		return x.Permission
	}
	return nil
}

// Read (Get single permission)
type GetPermissionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileId        string                 `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPermissionRequest) Reset() {
	*x = GetPermissionRequest{}
	mi := &file_permissions_permission_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPermissionRequest) ProtoMessage() {}

func (x *GetPermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_permissions_permission_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPermissionRequest.ProtoReflect.Descriptor instead.
func (*GetPermissionRequest) Descriptor() ([]byte, []int) {
	return file_permissions_permission_proto_rawDescGZIP(), []int{3}
}

func (x *GetPermissionRequest) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *GetPermissionRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetPermissionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Permission    *Permission            `protobuf:"bytes,1,opt,name=permission,proto3" json:"permission,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPermissionResponse) Reset() {
	*x = GetPermissionResponse{}
	mi := &file_permissions_permission_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPermissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPermissionResponse) ProtoMessage() {}

func (x *GetPermissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_permissions_permission_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPermissionResponse.ProtoReflect.Descriptor instead.
func (*GetPermissionResponse) Descriptor() ([]byte, []int) {
	return file_permissions_permission_proto_rawDescGZIP(), []int{4}
}

func (x *GetPermissionResponse) GetPermission() *Permission {
	if x != nil {
		return x.Permission
	}
	return nil
}

// Read (List permissions for a file)
type ListPermissionsByFileRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileId        string                 `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListPermissionsByFileRequest) Reset() {
	*x = ListPermissionsByFileRequest{}
	mi := &file_permissions_permission_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPermissionsByFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPermissionsByFileRequest) ProtoMessage() {}

func (x *ListPermissionsByFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_permissions_permission_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPermissionsByFileRequest.ProtoReflect.Descriptor instead.
func (*ListPermissionsByFileRequest) Descriptor() ([]byte, []int) {
	return file_permissions_permission_proto_rawDescGZIP(), []int{5}
}

func (x *ListPermissionsByFileRequest) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

type ListPermissionsByFileResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Permissions   []*Permission          `protobuf:"bytes,1,rep,name=permissions,proto3" json:"permissions,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListPermissionsByFileResponse) Reset() {
	*x = ListPermissionsByFileResponse{}
	mi := &file_permissions_permission_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPermissionsByFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPermissionsByFileResponse) ProtoMessage() {}

func (x *ListPermissionsByFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_permissions_permission_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPermissionsByFileResponse.ProtoReflect.Descriptor instead.
func (*ListPermissionsByFileResponse) Descriptor() ([]byte, []int) {
	return file_permissions_permission_proto_rawDescGZIP(), []int{6}
}

func (x *ListPermissionsByFileResponse) GetPermissions() []*Permission {
	if x != nil {
		return x.Permissions
	}
	return nil
}

// Read (List permissions for a user)
type ListPermissionsByUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListPermissionsByUserRequest) Reset() {
	*x = ListPermissionsByUserRequest{}
	mi := &file_permissions_permission_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPermissionsByUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPermissionsByUserRequest) ProtoMessage() {}

func (x *ListPermissionsByUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_permissions_permission_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPermissionsByUserRequest.ProtoReflect.Descriptor instead.
func (*ListPermissionsByUserRequest) Descriptor() ([]byte, []int) {
	return file_permissions_permission_proto_rawDescGZIP(), []int{7}
}

func (x *ListPermissionsByUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type ListPermissionsByUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Permissions   []*Permission          `protobuf:"bytes,1,rep,name=permissions,proto3" json:"permissions,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListPermissionsByUserResponse) Reset() {
	*x = ListPermissionsByUserResponse{}
	mi := &file_permissions_permission_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPermissionsByUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPermissionsByUserResponse) ProtoMessage() {}

func (x *ListPermissionsByUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_permissions_permission_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPermissionsByUserResponse.ProtoReflect.Descriptor instead.
func (*ListPermissionsByUserResponse) Descriptor() ([]byte, []int) {
	return file_permissions_permission_proto_rawDescGZIP(), []int{8}
}

func (x *ListPermissionsByUserResponse) GetPermissions() []*Permission {
	if x != nil {
		return x.Permissions
	}
	return nil
}

// Update
type UpdatePermissionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Permission    *Permission            `protobuf:"bytes,1,opt,name=permission,proto3" json:"permission,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdatePermissionRequest) Reset() {
	*x = UpdatePermissionRequest{}
	mi := &file_permissions_permission_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdatePermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePermissionRequest) ProtoMessage() {}

func (x *UpdatePermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_permissions_permission_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePermissionRequest.ProtoReflect.Descriptor instead.
func (*UpdatePermissionRequest) Descriptor() ([]byte, []int) {
	return file_permissions_permission_proto_rawDescGZIP(), []int{9}
}

func (x *UpdatePermissionRequest) GetPermission() *Permission {
	if x != nil {
		return x.Permission
	}
	return nil
}

type UpdatePermissionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Permission    *Permission            `protobuf:"bytes,1,opt,name=permission,proto3" json:"permission,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdatePermissionResponse) Reset() {
	*x = UpdatePermissionResponse{}
	mi := &file_permissions_permission_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdatePermissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePermissionResponse) ProtoMessage() {}

func (x *UpdatePermissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_permissions_permission_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePermissionResponse.ProtoReflect.Descriptor instead.
func (*UpdatePermissionResponse) Descriptor() ([]byte, []int) {
	return file_permissions_permission_proto_rawDescGZIP(), []int{10}
}

func (x *UpdatePermissionResponse) GetPermission() *Permission {
	if x != nil {
		return x.Permission
	}
	return nil
}

// Delete
type DeletePermissionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileId        string                 `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeletePermissionRequest) Reset() {
	*x = DeletePermissionRequest{}
	mi := &file_permissions_permission_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeletePermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePermissionRequest) ProtoMessage() {}

func (x *DeletePermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_permissions_permission_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePermissionRequest.ProtoReflect.Descriptor instead.
func (*DeletePermissionRequest) Descriptor() ([]byte, []int) {
	return file_permissions_permission_proto_rawDescGZIP(), []int{11}
}

func (x *DeletePermissionRequest) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *DeletePermissionRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type DeletePermissionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeletePermissionResponse) Reset() {
	*x = DeletePermissionResponse{}
	mi := &file_permissions_permission_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeletePermissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePermissionResponse) ProtoMessage() {}

func (x *DeletePermissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_permissions_permission_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePermissionResponse.ProtoReflect.Descriptor instead.
func (*DeletePermissionResponse) Descriptor() ([]byte, []int) {
	return file_permissions_permission_proto_rawDescGZIP(), []int{12}
}

func (x *DeletePermissionResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_permissions_permission_proto protoreflect.FileDescriptor

var file_permissions_permission_proto_rawDesc = string([]byte{
	0x0a, 0x1c, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a, 0x0a,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c,
	0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x05,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x05, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x22, 0x46, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2b, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x47, 0x0a, 0x18,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x48, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x44, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x37, 0x0a, 0x1c, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x4e,
	0x0a, 0x1d, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x42, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2d, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x37,
	0x0a, 0x1c, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x4e, 0x0a, 0x1d, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x46, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2b, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22,
	0x47, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x0a, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x4b, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x34, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2a, 0x34, 0x0a, 0x0f, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x09,
	0x0a, 0x05, 0x4f, 0x57, 0x4e, 0x45, 0x52, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x44, 0x49,
	0x54, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x56, 0x49, 0x45, 0x57, 0x45, 0x52, 0x10,
	0x02, 0x32, 0xe1, 0x05, 0x0a, 0x11, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x69, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x22, 0x0c, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x68, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x15, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x12, 0x20, 0x2f, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69,
	0x64, 0x7d, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x76, 0x0a, 0x15,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x42,
	0x79, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x1d, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x69, 0x64, 0x7d, 0x12, 0x76, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x42,
	0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x93, 0x01, 0x0a,
	0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x18, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x44, 0x3a, 0x0a,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x36, 0x2f, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x7b, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x7d, 0x12, 0x71, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x22, 0x2a, 0x20, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x7b, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x7b, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x1e, 0x5a, 0x1c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67,
	0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_permissions_permission_proto_rawDescOnce sync.Once
	file_permissions_permission_proto_rawDescData []byte
)

func file_permissions_permission_proto_rawDescGZIP() []byte {
	file_permissions_permission_proto_rawDescOnce.Do(func() {
		file_permissions_permission_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_permissions_permission_proto_rawDesc), len(file_permissions_permission_proto_rawDesc)))
	})
	return file_permissions_permission_proto_rawDescData
}

var file_permissions_permission_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_permissions_permission_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_permissions_permission_proto_goTypes = []any{
	(PermissionLevel)(0),                  // 0: PermissionLevel
	(*Permission)(nil),                    // 1: Permission
	(*CreatePermissionRequest)(nil),       // 2: CreatePermissionRequest
	(*CreatePermissionResponse)(nil),      // 3: CreatePermissionResponse
	(*GetPermissionRequest)(nil),          // 4: GetPermissionRequest
	(*GetPermissionResponse)(nil),         // 5: GetPermissionResponse
	(*ListPermissionsByFileRequest)(nil),  // 6: ListPermissionsByFileRequest
	(*ListPermissionsByFileResponse)(nil), // 7: ListPermissionsByFileResponse
	(*ListPermissionsByUserRequest)(nil),  // 8: ListPermissionsByUserRequest
	(*ListPermissionsByUserResponse)(nil), // 9: ListPermissionsByUserResponse
	(*UpdatePermissionRequest)(nil),       // 10: UpdatePermissionRequest
	(*UpdatePermissionResponse)(nil),      // 11: UpdatePermissionResponse
	(*DeletePermissionRequest)(nil),       // 12: DeletePermissionRequest
	(*DeletePermissionResponse)(nil),      // 13: DeletePermissionResponse
}
var file_permissions_permission_proto_depIdxs = []int32{
	0,  // 0: Permission.level:type_name -> PermissionLevel
	1,  // 1: CreatePermissionRequest.permission:type_name -> Permission
	1,  // 2: CreatePermissionResponse.permission:type_name -> Permission
	1,  // 3: GetPermissionResponse.permission:type_name -> Permission
	1,  // 4: ListPermissionsByFileResponse.permissions:type_name -> Permission
	1,  // 5: ListPermissionsByUserResponse.permissions:type_name -> Permission
	1,  // 6: UpdatePermissionRequest.permission:type_name -> Permission
	1,  // 7: UpdatePermissionResponse.permission:type_name -> Permission
	2,  // 8: PermissionService.CreatePermission:input_type -> CreatePermissionRequest
	4,  // 9: PermissionService.GetPermission:input_type -> GetPermissionRequest
	6,  // 10: PermissionService.ListPermissionsByFile:input_type -> ListPermissionsByFileRequest
	8,  // 11: PermissionService.ListPermissionsByUser:input_type -> ListPermissionsByUserRequest
	10, // 12: PermissionService.UpdatePermission:input_type -> UpdatePermissionRequest
	12, // 13: PermissionService.DeletePermission:input_type -> DeletePermissionRequest
	3,  // 14: PermissionService.CreatePermission:output_type -> CreatePermissionResponse
	5,  // 15: PermissionService.GetPermission:output_type -> GetPermissionResponse
	7,  // 16: PermissionService.ListPermissionsByFile:output_type -> ListPermissionsByFileResponse
	9,  // 17: PermissionService.ListPermissionsByUser:output_type -> ListPermissionsByUserResponse
	11, // 18: PermissionService.UpdatePermission:output_type -> UpdatePermissionResponse
	13, // 19: PermissionService.DeletePermission:output_type -> DeletePermissionResponse
	14, // [14:20] is the sub-list for method output_type
	8,  // [8:14] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_permissions_permission_proto_init() }
func file_permissions_permission_proto_init() {
	if File_permissions_permission_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_permissions_permission_proto_rawDesc), len(file_permissions_permission_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_permissions_permission_proto_goTypes,
		DependencyIndexes: file_permissions_permission_proto_depIdxs,
		EnumInfos:         file_permissions_permission_proto_enumTypes,
		MessageInfos:      file_permissions_permission_proto_msgTypes,
	}.Build()
	File_permissions_permission_proto = out.File
	file_permissions_permission_proto_goTypes = nil
	file_permissions_permission_proto_depIdxs = nil
}
