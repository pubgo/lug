// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: proto/hello/example.proto

// GoGo gRPC Example
//
// This example is used to show how to use gRPC and
// gRPC-Gateway with GoGo Protobuf.

package hello

import (
	_ "github.com/gogo/protobuf/gogoproto"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Role int32

const (
	Role_GUEST  Role = 0
	Role_MEMBER Role = 1
	Role_ADMIN  Role = 2
)

// Enum value maps for Role.
var (
	Role_name = map[int32]string{
		0: "GUEST",
		1: "MEMBER",
		2: "ADMIN",
	}
	Role_value = map[string]int32{
		"GUEST":  0,
		"MEMBER": 1,
		"ADMIN":  2,
	}
)

func (x Role) Enum() *Role {
	p := new(Role)
	*p = x
	return p
}

func (x Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Role) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_hello_example_proto_enumTypes[0].Descriptor()
}

func (Role) Type() protoreflect.EnumType {
	return &file_proto_hello_example_proto_enumTypes[0]
}

func (x Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Role.Descriptor instead.
func (Role) EnumDescriptor() ([]byte, []int) {
	return file_proto_hello_example_proto_rawDescGZIP(), []int{0}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Role       Role                   `protobuf:"varint,2,opt,name=role,proto3,enum=hello.Role" json:"role,omitempty"`
	CreateDate *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=create_date,json=createDate,proto3" json:"create_date,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hello_example_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hello_example_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_proto_hello_example_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetRole() Role {
	if x != nil {
		return x.Role
	}
	return Role_GUEST
}

func (x *User) GetCreateDate() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateDate
	}
	return nil
}

type UserRole struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role    Role                           `protobuf:"varint,1,opt,name=role,proto3,enum=hello.Role" json:"role,omitempty"`
	Lists   *structpb.ListValue            `protobuf:"bytes,2,opt,name=lists,proto3" json:"lists,omitempty"`
	Headers map[string]*structpb.ListValue `protobuf:"bytes,3,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *UserRole) Reset() {
	*x = UserRole{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hello_example_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRole) ProtoMessage() {}

func (x *UserRole) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hello_example_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRole.ProtoReflect.Descriptor instead.
func (*UserRole) Descriptor() ([]byte, []int) {
	return file_proto_hello_example_proto_rawDescGZIP(), []int{1}
}

func (x *UserRole) GetRole() Role {
	if x != nil {
		return x.Role
	}
	return Role_GUEST
}

func (x *UserRole) GetLists() *structpb.ListValue {
	if x != nil {
		return x.Lists
	}
	return nil
}

func (x *UserRole) GetHeaders() map[string]*structpb.ListValue {
	if x != nil {
		return x.Headers
	}
	return nil
}

type UpdateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The user resource which replaces the resource on the server.
	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// The update mask applies to the resource. For the `FieldMask` definition,
	// see https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#fieldmask
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateUserRequest) Reset() {
	*x = UpdateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hello_example_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserRequest) ProtoMessage() {}

func (x *UpdateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hello_example_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return file_proto_hello_example_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateUserRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UpdateUserRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

type ListUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Only list users created after this timestamp
	CreatedSince *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=created_since,json=createdSince,proto3" json:"created_since,omitempty"`
	// Only list users older than this Duration
	OlderThan *durationpb.Duration `protobuf:"bytes,2,opt,name=older_than,json=olderThan,proto3" json:"older_than,omitempty"`
}

func (x *ListUsersRequest) Reset() {
	*x = ListUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hello_example_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUsersRequest) ProtoMessage() {}

func (x *ListUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hello_example_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUsersRequest.ProtoReflect.Descriptor instead.
func (*ListUsersRequest) Descriptor() ([]byte, []int) {
	return file_proto_hello_example_proto_rawDescGZIP(), []int{3}
}

func (x *ListUsersRequest) GetCreatedSince() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedSince
	}
	return nil
}

func (x *ListUsersRequest) GetOlderThan() *durationpb.Duration {
	if x != nil {
		return x.OlderThan
	}
	return nil
}

var File_proto_hello_example_proto protoreflect.FileDescriptor

var file_proto_hello_example_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2f, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14,
	0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x01, 0x0a, 0x04, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x35, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x25,
	0xe2, 0xdf, 0x1f, 0x21, 0x10, 0x00, 0x2a, 0x1d, 0x49, 0x44, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20,
	0x62, 0x65, 0x20, 0x61, 0x20, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x20, 0x69, 0x6e,
	0x74, 0x65, 0x67, 0x65, 0x72, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x04, 0x72, 0x6f, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e,
	0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x41, 0x0a, 0x0b, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0x90, 0xdf, 0x1f,
	0x01, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x65, 0x22, 0xed, 0x01,
	0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x6c,
	0x69, 0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x12, 0x36, 0x0a,
	0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x2e,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x73, 0x1a, 0x56, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x30, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x71, 0x0a,
	0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1f, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61,
	0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b,
	0x22, 0x99, 0x01, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0x90, 0xdf, 0x1f, 0x01, 0x52, 0x0c,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x53, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x0a,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x68, 0x61, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x04, 0x98, 0xdf, 0x1f,
	0x01, 0x52, 0x09, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x54, 0x68, 0x61, 0x6e, 0x2a, 0x28, 0x0a, 0x04,
	0x52, 0x6f, 0x6c, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x55, 0x45, 0x53, 0x54, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x41,
	0x44, 0x4d, 0x49, 0x4e, 0x10, 0x02, 0x32, 0x94, 0x03, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x0b, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x3a, 0x01, 0x2a,
	0x12, 0x45, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0b, 0x2e, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x4a, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x12, 0x17, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x30, 0x01, 0x12, 0x4f, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x42, 0x79, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0f, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x1a, 0x0b, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x72, 0x6f, 0x6c, 0x65,
	0x28, 0x01, 0x30, 0x01, 0x12, 0x57, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x18, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1c, 0x32, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x42, 0x6c, 0x5a,
	0x08, 0x2e, 0x2f, 0x3b, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x92, 0x41, 0x4b, 0x12, 0x05, 0x32, 0x03,
	0x31, 0x2e, 0x30, 0x2a, 0x01, 0x02, 0x72, 0x3f, 0x0a, 0x17, 0x67, 0x52, 0x50, 0x43, 0x20, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x20, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72,
	0x79, 0x12, 0x24, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0xc8, 0xe2, 0x1e, 0x01, 0xd0, 0xe2, 0x1e, 0x01, 0xe0,
	0xe2, 0x1e, 0x01, 0xc0, 0xe3, 0x1e, 0x01, 0xc8, 0xe3, 0x1e, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_hello_example_proto_rawDescOnce sync.Once
	file_proto_hello_example_proto_rawDescData = file_proto_hello_example_proto_rawDesc
)

func file_proto_hello_example_proto_rawDescGZIP() []byte {
	file_proto_hello_example_proto_rawDescOnce.Do(func() {
		file_proto_hello_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_hello_example_proto_rawDescData)
	})
	return file_proto_hello_example_proto_rawDescData
}

var file_proto_hello_example_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_hello_example_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_hello_example_proto_goTypes = []interface{}{
	(Role)(0),                     // 0: hello.Role
	(*User)(nil),                  // 1: hello.User
	(*UserRole)(nil),              // 2: hello.UserRole
	(*UpdateUserRequest)(nil),     // 3: hello.UpdateUserRequest
	(*ListUsersRequest)(nil),      // 4: hello.ListUsersRequest
	nil,                           // 5: hello.UserRole.HeadersEntry
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(*structpb.ListValue)(nil),    // 7: google.protobuf.ListValue
	(*fieldmaskpb.FieldMask)(nil), // 8: google.protobuf.FieldMask
	(*durationpb.Duration)(nil),   // 9: google.protobuf.Duration
	(*emptypb.Empty)(nil),         // 10: google.protobuf.Empty
}
var file_proto_hello_example_proto_depIdxs = []int32{
	0,  // 0: hello.User.role:type_name -> hello.Role
	6,  // 1: hello.User.create_date:type_name -> google.protobuf.Timestamp
	0,  // 2: hello.UserRole.role:type_name -> hello.Role
	7,  // 3: hello.UserRole.lists:type_name -> google.protobuf.ListValue
	5,  // 4: hello.UserRole.headers:type_name -> hello.UserRole.HeadersEntry
	1,  // 5: hello.UpdateUserRequest.user:type_name -> hello.User
	8,  // 6: hello.UpdateUserRequest.update_mask:type_name -> google.protobuf.FieldMask
	6,  // 7: hello.ListUsersRequest.created_since:type_name -> google.protobuf.Timestamp
	9,  // 8: hello.ListUsersRequest.older_than:type_name -> google.protobuf.Duration
	7,  // 9: hello.UserRole.HeadersEntry.value:type_name -> google.protobuf.ListValue
	1,  // 10: hello.UserService.AddUser:input_type -> hello.User
	1,  // 11: hello.UserService.GetUser:input_type -> hello.User
	4,  // 12: hello.UserService.ListUsers:input_type -> hello.ListUsersRequest
	2,  // 13: hello.UserService.ListUsersByRole:input_type -> hello.UserRole
	3,  // 14: hello.UserService.UpdateUser:input_type -> hello.UpdateUserRequest
	10, // 15: hello.UserService.AddUser:output_type -> google.protobuf.Empty
	10, // 16: hello.UserService.GetUser:output_type -> google.protobuf.Empty
	1,  // 17: hello.UserService.ListUsers:output_type -> hello.User
	1,  // 18: hello.UserService.ListUsersByRole:output_type -> hello.User
	1,  // 19: hello.UserService.UpdateUser:output_type -> hello.User
	15, // [15:20] is the sub-list for method output_type
	10, // [10:15] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_proto_hello_example_proto_init() }
func file_proto_hello_example_proto_init() {
	if File_proto_hello_example_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_hello_example_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_proto_hello_example_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRole); i {
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
		file_proto_hello_example_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserRequest); i {
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
		file_proto_hello_example_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUsersRequest); i {
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
			RawDescriptor: file_proto_hello_example_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_hello_example_proto_goTypes,
		DependencyIndexes: file_proto_hello_example_proto_depIdxs,
		EnumInfos:         file_proto_hello_example_proto_enumTypes,
		MessageInfos:      file_proto_hello_example_proto_msgTypes,
	}.Build()
	File_proto_hello_example_proto = out.File
	file_proto_hello_example_proto_rawDesc = nil
	file_proto_hello_example_proto_goTypes = nil
	file_proto_hello_example_proto_depIdxs = nil
}
