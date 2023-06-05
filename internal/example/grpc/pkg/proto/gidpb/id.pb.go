// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: gid/id.proto

package gidpb

import (
	_ "github.com/pubgo/funk/proto/errorpb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type SrvCode int32

const (
	SrvCode_OK SrvCode = 0
	// id generate error
	SrvCode_IDGenerateFailed SrvCode = 1
)

// Enum value maps for SrvCode.
var (
	SrvCode_name = map[int32]string{
		0: "OK",
		1: "IDGenerateFailed",
	}
	SrvCode_value = map[string]int32{
		"OK":               0,
		"IDGenerateFailed": 1,
	}
)

func (x SrvCode) Enum() *SrvCode {
	p := new(SrvCode)
	*p = x
	return p
}

func (x SrvCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SrvCode) Descriptor() protoreflect.EnumDescriptor {
	return file_gid_id_proto_enumTypes[0].Descriptor()
}

func (SrvCode) Type() protoreflect.EnumType {
	return &file_gid_id_proto_enumTypes[0]
}

func (x SrvCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SrvCode.Descriptor instead.
func (SrvCode) EnumDescriptor() ([]byte, []int) {
	return file_gid_id_proto_rawDescGZIP(), []int{0}
}

type Tag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Tag) Reset() {
	*x = Tag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gid_id_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tag) ProtoMessage() {}

func (x *Tag) ProtoReflect() protoreflect.Message {
	mi := &file_gid_id_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tag.ProtoReflect.Descriptor instead.
func (*Tag) Descriptor() ([]byte, []int) {
	return file_gid_id_proto_rawDescGZIP(), []int{0}
}

func (x *Tag) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Tag) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type GenerateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the unique id generated
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// the type of id generated
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *GenerateResponse) Reset() {
	*x = GenerateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gid_id_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateResponse) ProtoMessage() {}

func (x *GenerateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gid_id_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateResponse.ProtoReflect.Descriptor instead.
func (*GenerateResponse) Descriptor() ([]byte, []int) {
	return file_gid_id_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GenerateResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

// Generate a unique ID. Defaults to uuid.
type GenerateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// type of id e.g uuid, shortid, snowflake (64 bit), bigflake (128 bit)
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *GenerateRequest) Reset() {
	*x = GenerateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gid_id_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateRequest) ProtoMessage() {}

func (x *GenerateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gid_id_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateRequest.ProtoReflect.Descriptor instead.
func (*GenerateRequest) Descriptor() ([]byte, []int) {
	return file_gid_id_proto_rawDescGZIP(), []int{2}
}

func (x *GenerateRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

// List the types of IDs available. No query params needed.
type TypesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TypesRequest) Reset() {
	*x = TypesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gid_id_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TypesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TypesRequest) ProtoMessage() {}

func (x *TypesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gid_id_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TypesRequest.ProtoReflect.Descriptor instead.
func (*TypesRequest) Descriptor() ([]byte, []int) {
	return file_gid_id_proto_rawDescGZIP(), []int{3}
}

// TypesResponse 返回值类型
type TypesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Types []string `protobuf:"bytes,1,rep,name=types,proto3" json:"types,omitempty"`
}

func (x *TypesResponse) Reset() {
	*x = TypesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gid_id_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TypesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TypesResponse) ProtoMessage() {}

func (x *TypesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gid_id_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TypesResponse.ProtoReflect.Descriptor instead.
func (*TypesResponse) Descriptor() ([]byte, []int) {
	return file_gid_id_proto_rawDescGZIP(), []int{4}
}

func (x *TypesResponse) GetTypes() []string {
	if x != nil {
		return x.Types
	}
	return nil
}

var File_gid_id_proto protoreflect.FileDescriptor

var file_gid_id_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x67, 0x69, 0x64, 0x2f, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03,
	0x67, 0x69, 0x64, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x70, 0x62, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x2d, 0x0a, 0x03, 0x54, 0x61, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x36, 0x0a, 0x10, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x25, 0x0a, 0x0f, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x0e,
	0x0a, 0x0c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x25,
	0x0a, 0x0d, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2a, 0x37, 0x0a, 0x07, 0x53, 0x72, 0x76, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x10, 0x49, 0x44, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x01, 0x1a, 0x06,
	0x92, 0xea, 0x30, 0x02, 0x10, 0x0d, 0x1a, 0x06, 0x8a, 0xea, 0x30, 0x02, 0x08, 0x01, 0x32, 0x80,
	0x02, 0x0a, 0x02, 0x49, 0x64, 0x12, 0x53, 0x0a, 0x08, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x12, 0x14, 0x2e, 0x67, 0x69, 0x64, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x67, 0x69, 0x64, 0x2e, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x64, 0x2f, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x4c, 0x0a, 0x0a, 0x54, 0x79,
	0x70, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x11, 0x2e, 0x67, 0x69, 0x64, 0x2e, 0x54,
	0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x67, 0x69,
	0x64, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x64, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x31, 0x30, 0x01, 0x12, 0x44, 0x0a, 0x05, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x12, 0x11, 0x2e, 0x67, 0x69, 0x64, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x67, 0x69, 0x64, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e,
	0x12, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x64, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x11,
	0xca, 0x41, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68, 0x6f, 0x73, 0x74, 0x3a, 0x38, 0x30, 0x38,
	0x30, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x67, 0x69, 0x64, 0x70, 0x62, 0x3b, 0x67, 0x69, 0x64,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gid_id_proto_rawDescOnce sync.Once
	file_gid_id_proto_rawDescData = file_gid_id_proto_rawDesc
)

func file_gid_id_proto_rawDescGZIP() []byte {
	file_gid_id_proto_rawDescOnce.Do(func() {
		file_gid_id_proto_rawDescData = protoimpl.X.CompressGZIP(file_gid_id_proto_rawDescData)
	})
	return file_gid_id_proto_rawDescData
}

var file_gid_id_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_gid_id_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_gid_id_proto_goTypes = []interface{}{
	(SrvCode)(0),             // 0: gid.SrvCode
	(*Tag)(nil),              // 1: gid.Tag
	(*GenerateResponse)(nil), // 2: gid.GenerateResponse
	(*GenerateRequest)(nil),  // 3: gid.GenerateRequest
	(*TypesRequest)(nil),     // 4: gid.TypesRequest
	(*TypesResponse)(nil),    // 5: gid.TypesResponse
}
var file_gid_id_proto_depIdxs = []int32{
	3, // 0: gid.Id.Generate:input_type -> gid.GenerateRequest
	4, // 1: gid.Id.TypeStream:input_type -> gid.TypesRequest
	4, // 2: gid.Id.Types:input_type -> gid.TypesRequest
	2, // 3: gid.Id.Generate:output_type -> gid.GenerateResponse
	5, // 4: gid.Id.TypeStream:output_type -> gid.TypesResponse
	5, // 5: gid.Id.Types:output_type -> gid.TypesResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gid_id_proto_init() }
func file_gid_id_proto_init() {
	if File_gid_id_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gid_id_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tag); i {
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
		file_gid_id_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateResponse); i {
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
		file_gid_id_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateRequest); i {
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
		file_gid_id_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TypesRequest); i {
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
		file_gid_id_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TypesResponse); i {
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
			RawDescriptor: file_gid_id_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gid_id_proto_goTypes,
		DependencyIndexes: file_gid_id_proto_depIdxs,
		EnumInfos:         file_gid_id_proto_enumTypes,
		MessageInfos:      file_gid_id_proto_msgTypes,
	}.Build()
	File_gid_id_proto = out.File
	file_gid_id_proto_rawDesc = nil
	file_gid_id_proto_goTypes = nil
	file_gid_id_proto_depIdxs = nil
}
