// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: proto/lava/rpc.proto

package lavapbv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RpcMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version *string           `protobuf:"bytes,2,opt,name=version,proto3,oneof" json:"version,omitempty"`
	Tags    map[string]string `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *RpcMeta) Reset() {
	*x = RpcMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_lava_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcMeta) ProtoMessage() {}

func (x *RpcMeta) ProtoReflect() protoreflect.Message {
	mi := &file_proto_lava_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpcMeta.ProtoReflect.Descriptor instead.
func (*RpcMeta) Descriptor() ([]byte, []int) {
	return file_proto_lava_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *RpcMeta) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RpcMeta) GetVersion() string {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return ""
}

func (x *RpcMeta) GetTags() map[string]string {
	if x != nil {
		return x.Tags
	}
	return nil
}

var file_proto_lava_rpc_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*RpcMeta)(nil),
		Field:         100004,
		Name:          "lava.rpc.options",
		Tag:           "bytes,100004,opt,name=options",
		Filename:      "proto/lava/rpc.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional lava.rpc.RpcMeta options = 100004;
	E_Options = &file_proto_lava_rpc_proto_extTypes[0]
)

var File_proto_lava_rpc_proto protoreflect.FileDescriptor

var file_proto_lava_rpc_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x61, 0x76, 0x61, 0x2f, 0x72, 0x70, 0x63,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6c, 0x61, 0x76, 0x61, 0x2e, 0x72, 0x70, 0x63,
	0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xb2, 0x01, 0x0a, 0x07, 0x52, 0x70, 0x63, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x88, 0x01,
	0x01, 0x12, 0x2f, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x6c, 0x61, 0x76, 0x61, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x70, 0x63, 0x4d, 0x65,
	0x74, 0x61, 0x2e, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x1a, 0x37, 0x0a, 0x09, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x3a, 0x4d, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0xa4, 0x8d, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6c, 0x61, 0x76,
	0x61, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x70, 0x63, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x07, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x75, 0x62, 0x67, 0x6f, 0x2f, 0x6c, 0x61, 0x76, 0x61, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x61, 0x76, 0x61, 0x70, 0x62,
	0x76, 0x31, 0x3b, 0x6c, 0x61, 0x76, 0x61, 0x70, 0x62, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_lava_rpc_proto_rawDescOnce sync.Once
	file_proto_lava_rpc_proto_rawDescData = file_proto_lava_rpc_proto_rawDesc
)

func file_proto_lava_rpc_proto_rawDescGZIP() []byte {
	file_proto_lava_rpc_proto_rawDescOnce.Do(func() {
		file_proto_lava_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_lava_rpc_proto_rawDescData)
	})
	return file_proto_lava_rpc_proto_rawDescData
}

var file_proto_lava_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_lava_rpc_proto_goTypes = []any{
	(*RpcMeta)(nil),                    // 0: lava.rpc.RpcMeta
	nil,                                // 1: lava.rpc.RpcMeta.TagsEntry
	(*descriptorpb.MethodOptions)(nil), // 2: google.protobuf.MethodOptions
}
var file_proto_lava_rpc_proto_depIdxs = []int32{
	1, // 0: lava.rpc.RpcMeta.tags:type_name -> lava.rpc.RpcMeta.TagsEntry
	2, // 1: lava.rpc.options:extendee -> google.protobuf.MethodOptions
	0, // 2: lava.rpc.options:type_name -> lava.rpc.RpcMeta
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	2, // [2:3] is the sub-list for extension type_name
	1, // [1:2] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_lava_rpc_proto_init() }
func file_proto_lava_rpc_proto_init() {
	if File_proto_lava_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_lava_rpc_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*RpcMeta); i {
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
	file_proto_lava_rpc_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_lava_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_proto_lava_rpc_proto_goTypes,
		DependencyIndexes: file_proto_lava_rpc_proto_depIdxs,
		MessageInfos:      file_proto_lava_rpc_proto_msgTypes,
		ExtensionInfos:    file_proto_lava_rpc_proto_extTypes,
	}.Build()
	File_proto_lava_rpc_proto = out.File
	file_proto_lava_rpc_proto_rawDesc = nil
	file_proto_lava_rpc_proto_goTypes = nil
	file_proto_lava_rpc_proto_depIdxs = nil
}
