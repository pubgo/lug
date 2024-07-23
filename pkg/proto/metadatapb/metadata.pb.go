// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: proto/lava/services/metadata/metadata.proto

package metadatapb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ListServicesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListServicesRequest) Reset() {
	*x = ListServicesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_lava_services_metadata_metadata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListServicesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListServicesRequest) ProtoMessage() {}

func (x *ListServicesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_lava_services_metadata_metadata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListServicesRequest.ProtoReflect.Descriptor instead.
func (*ListServicesRequest) Descriptor() ([]byte, []int) {
	return file_proto_lava_services_metadata_metadata_proto_rawDescGZIP(), []int{0}
}

type ListServicesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Services []string `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
	Methods  []string `protobuf:"bytes,2,rep,name=methods,proto3" json:"methods,omitempty"`
}

func (x *ListServicesReply) Reset() {
	*x = ListServicesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_lava_services_metadata_metadata_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListServicesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListServicesReply) ProtoMessage() {}

func (x *ListServicesReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_lava_services_metadata_metadata_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListServicesReply.ProtoReflect.Descriptor instead.
func (*ListServicesReply) Descriptor() ([]byte, []int) {
	return file_proto_lava_services_metadata_metadata_proto_rawDescGZIP(), []int{1}
}

func (x *ListServicesReply) GetServices() []string {
	if x != nil {
		return x.Services
	}
	return nil
}

func (x *ListServicesReply) GetMethods() []string {
	if x != nil {
		return x.Methods
	}
	return nil
}

type GetServiceDescRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetServiceDescRequest) Reset() {
	*x = GetServiceDescRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_lava_services_metadata_metadata_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServiceDescRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServiceDescRequest) ProtoMessage() {}

func (x *GetServiceDescRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_lava_services_metadata_metadata_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServiceDescRequest.ProtoReflect.Descriptor instead.
func (*GetServiceDescRequest) Descriptor() ([]byte, []int) {
	return file_proto_lava_services_metadata_metadata_proto_rawDescGZIP(), []int{2}
}

func (x *GetServiceDescRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetServiceDescReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileDescSet *descriptorpb.FileDescriptorSet `protobuf:"bytes,1,opt,name=file_desc_set,json=fileDescSet,proto3" json:"file_desc_set,omitempty"`
}

func (x *GetServiceDescReply) Reset() {
	*x = GetServiceDescReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_lava_services_metadata_metadata_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServiceDescReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServiceDescReply) ProtoMessage() {}

func (x *GetServiceDescReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_lava_services_metadata_metadata_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServiceDescReply.ProtoReflect.Descriptor instead.
func (*GetServiceDescReply) Descriptor() ([]byte, []int) {
	return file_proto_lava_services_metadata_metadata_proto_rawDescGZIP(), []int{3}
}

func (x *GetServiceDescReply) GetFileDescSet() *descriptorpb.FileDescriptorSet {
	if x != nil {
		return x.FileDescSet
	}
	return nil
}

var File_proto_lava_services_metadata_metadata_proto protoreflect.FileDescriptor

var file_proto_lava_services_metadata_metadata_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x61, 0x76, 0x61, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6c,
	0x61, 0x76, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x15, 0x0a, 0x13, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x49, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x22, 0x2b, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x5d, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x46, 0x0a, 0x0d, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x5f, 0x73,
	0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x52, 0x0b, 0x66, 0x69,
	0x6c, 0x65, 0x44, 0x65, 0x73, 0x63, 0x53, 0x65, 0x74, 0x32, 0x82, 0x02, 0x0a, 0x08, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x73, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x21, 0x2e, 0x6c, 0x61, 0x76, 0x61, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6c, 0x61, 0x76, 0x61,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x19, 0x12, 0x17, 0x2f, 0x6c, 0x61, 0x76, 0x61, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x80, 0x01, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x12, 0x23,
	0x2e, 0x6c, 0x61, 0x76, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6c, 0x61, 0x76, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x73,
	0x63, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x12, 0x1e,
	0x2f, 0x6c, 0x61, 0x76, 0x61, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x42, 0x37,
	0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x75, 0x62,
	0x67, 0x6f, 0x2f, 0x6c, 0x61, 0x76, 0x61, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x70, 0x62, 0x3b, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_lava_services_metadata_metadata_proto_rawDescOnce sync.Once
	file_proto_lava_services_metadata_metadata_proto_rawDescData = file_proto_lava_services_metadata_metadata_proto_rawDesc
)

func file_proto_lava_services_metadata_metadata_proto_rawDescGZIP() []byte {
	file_proto_lava_services_metadata_metadata_proto_rawDescOnce.Do(func() {
		file_proto_lava_services_metadata_metadata_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_lava_services_metadata_metadata_proto_rawDescData)
	})
	return file_proto_lava_services_metadata_metadata_proto_rawDescData
}

var file_proto_lava_services_metadata_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_lava_services_metadata_metadata_proto_goTypes = []any{
	(*ListServicesRequest)(nil),            // 0: lava.service.ListServicesRequest
	(*ListServicesReply)(nil),              // 1: lava.service.ListServicesReply
	(*GetServiceDescRequest)(nil),          // 2: lava.service.GetServiceDescRequest
	(*GetServiceDescReply)(nil),            // 3: lava.service.GetServiceDescReply
	(*descriptorpb.FileDescriptorSet)(nil), // 4: google.protobuf.FileDescriptorSet
}
var file_proto_lava_services_metadata_metadata_proto_depIdxs = []int32{
	4, // 0: lava.service.GetServiceDescReply.file_desc_set:type_name -> google.protobuf.FileDescriptorSet
	0, // 1: lava.service.Metadata.ListServices:input_type -> lava.service.ListServicesRequest
	2, // 2: lava.service.Metadata.GetServiceDesc:input_type -> lava.service.GetServiceDescRequest
	1, // 3: lava.service.Metadata.ListServices:output_type -> lava.service.ListServicesReply
	3, // 4: lava.service.Metadata.GetServiceDesc:output_type -> lava.service.GetServiceDescReply
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_lava_services_metadata_metadata_proto_init() }
func file_proto_lava_services_metadata_metadata_proto_init() {
	if File_proto_lava_services_metadata_metadata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_lava_services_metadata_metadata_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ListServicesRequest); i {
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
		file_proto_lava_services_metadata_metadata_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ListServicesReply); i {
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
		file_proto_lava_services_metadata_metadata_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetServiceDescRequest); i {
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
		file_proto_lava_services_metadata_metadata_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetServiceDescReply); i {
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
			RawDescriptor: file_proto_lava_services_metadata_metadata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_lava_services_metadata_metadata_proto_goTypes,
		DependencyIndexes: file_proto_lava_services_metadata_metadata_proto_depIdxs,
		MessageInfos:      file_proto_lava_services_metadata_metadata_proto_msgTypes,
	}.Build()
	File_proto_lava_services_metadata_metadata_proto = out.File
	file_proto_lava_services_metadata_metadata_proto_rawDesc = nil
	file_proto_lava_services_metadata_metadata_proto_goTypes = nil
	file_proto_lava_services_metadata_metadata_proto_depIdxs = nil
}
