// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: proto/test/test.proto

package test

import (
	_ "github.com/pubgo/lava/proto/lava"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_	= protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_	= protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GenerateResponse struct {
	state		protoimpl.MessageState
	sizeCache	protoimpl.SizeCache
	unknownFields	protoimpl.UnknownFields

	Id	string	`protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are assignable to Image:
	//	*GenerateResponse_MediaFile
	//	*GenerateResponse_Data
	//	*GenerateResponse_AdIdToCopyImageFrom
	Image	isGenerateResponse_Image	`protobuf_oneof:"image"`
	// Types that are assignable to Image1:
	//	*GenerateResponse_MediaFile1
	//	*GenerateResponse_Data1
	//	*GenerateResponse_AdIdToCopyImageFrom1
	Image1	isGenerateResponse_Image1	`protobuf_oneof:"image1"`
	Type	string				`protobuf:"bytes,2,opt,name=type,proto3" json:"types,omitempty" yaml1:"types,omitempty"`
}

func (x *GenerateResponse) Reset() {
	*x = GenerateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_test_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateResponse) ProtoMessage()	{}

func (x *GenerateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_test_test_proto_msgTypes[0]
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
	return file_proto_test_test_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (m *GenerateResponse) GetImage() isGenerateResponse_Image {
	if m != nil {
		return m.Image
	}
	return nil
}

func (x *GenerateResponse) GetMediaFile() string {
	if x, ok := x.GetImage().(*GenerateResponse_MediaFile); ok {
		return x.MediaFile
	}
	return ""
}

func (x *GenerateResponse) GetData() []byte {
	if x, ok := x.GetImage().(*GenerateResponse_Data); ok {
		return x.Data
	}
	return nil
}

func (x *GenerateResponse) GetAdIdToCopyImageFrom() int64 {
	if x, ok := x.GetImage().(*GenerateResponse_AdIdToCopyImageFrom); ok {
		return x.AdIdToCopyImageFrom
	}
	return 0
}

func (m *GenerateResponse) GetImage1() isGenerateResponse_Image1 {
	if m != nil {
		return m.Image1
	}
	return nil
}

func (x *GenerateResponse) GetMediaFile1() string {
	if x, ok := x.GetImage1().(*GenerateResponse_MediaFile1); ok {
		return x.MediaFile1
	}
	return ""
}

func (x *GenerateResponse) GetData1() []byte {
	if x, ok := x.GetImage1().(*GenerateResponse_Data1); ok {
		return x.Data1
	}
	return nil
}

func (x *GenerateResponse) GetAdIdToCopyImageFrom1() int64 {
	if x, ok := x.GetImage1().(*GenerateResponse_AdIdToCopyImageFrom1); ok {
		return x.AdIdToCopyImageFrom1
	}
	return 0
}

func (x *GenerateResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type isGenerateResponse_Image interface {
	isGenerateResponse_Image()
}

type GenerateResponse_MediaFile struct {
	MediaFile string `protobuf:"bytes,12,opt,name=media_file,json=mediaFile,proto3,oneof"`
}

type GenerateResponse_Data struct {
	Data []byte `protobuf:"bytes,13,opt,name=data,proto3,oneof"`
}

type GenerateResponse_AdIdToCopyImageFrom struct {
	AdIdToCopyImageFrom int64 `protobuf:"varint,14,opt,name=ad_id_to_copy_image_from,json=adIdToCopyImageFrom,proto3,oneof"`
}

func (*GenerateResponse_MediaFile) isGenerateResponse_Image()	{}

func (*GenerateResponse_Data) isGenerateResponse_Image()	{}

func (*GenerateResponse_AdIdToCopyImageFrom) isGenerateResponse_Image()	{}

type isGenerateResponse_Image1 interface {
	isGenerateResponse_Image1()
}

type GenerateResponse_MediaFile1 struct {
	MediaFile1 string `protobuf:"bytes,15,opt,name=media_file1,json=mediaFile1,proto3,oneof"`
}

type GenerateResponse_Data1 struct {
	Data1 []byte `protobuf:"bytes,16,opt,name=data1,proto3,oneof"`
}

type GenerateResponse_AdIdToCopyImageFrom1 struct {
	AdIdToCopyImageFrom1 int64 `protobuf:"varint,17,opt,name=ad_id_to_copy_image_from1,json=adIdToCopyImageFrom1,proto3,oneof"`
}

func (*GenerateResponse_MediaFile1) isGenerateResponse_Image1()	{}

func (*GenerateResponse_Data1) isGenerateResponse_Image1()	{}

func (*GenerateResponse_AdIdToCopyImageFrom1) isGenerateResponse_Image1()	{}

var File_proto_test_test_proto protoreflect.FileDescriptor

var file_proto_test_test_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x61, 0x76, 0x61, 0x2f, 0x74, 0x61, 0x67, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81, 0x03, 0x0a, 0x10, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x38, 0x0a, 0x0a, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x42, 0x17, 0x82,
	0xea, 0x30, 0x13, 0x0a, 0x05, 0x6a, 0x73, 0x6f, 0x6e, 0x31, 0x12, 0x0a, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x48, 0x00, 0x52, 0x09, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x46,
	0x69, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0c, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x37, 0x0a, 0x18, 0x61, 0x64, 0x5f,
	0x69, 0x64, 0x5f, 0x74, 0x6f, 0x5f, 0x63, 0x6f, 0x70, 0x79, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x13, 0x61,
	0x64, 0x49, 0x64, 0x54, 0x6f, 0x43, 0x6f, 0x70, 0x79, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x46, 0x72,
	0x6f, 0x6d, 0x12, 0x21, 0x0a, 0x0b, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x66, 0x69, 0x6c, 0x65,
	0x31, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0a, 0x6d, 0x65, 0x64, 0x69, 0x61,
	0x46, 0x69, 0x6c, 0x65, 0x31, 0x12, 0x16, 0x0a, 0x05, 0x64, 0x61, 0x74, 0x61, 0x31, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x0c, 0x48, 0x01, 0x52, 0x05, 0x64, 0x61, 0x74, 0x61, 0x31, 0x12, 0x39, 0x0a,
	0x19, 0x61, 0x64, 0x5f, 0x69, 0x64, 0x5f, 0x74, 0x6f, 0x5f, 0x63, 0x6f, 0x70, 0x79, 0x5f, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x31, 0x18, 0x11, 0x20, 0x01, 0x28, 0x03,
	0x48, 0x01, 0x52, 0x14, 0x61, 0x64, 0x49, 0x64, 0x54, 0x6f, 0x43, 0x6f, 0x70, 0x79, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x31, 0x12, 0x4b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x37, 0x82, 0xea, 0x30, 0x17, 0x0a, 0x04, 0x6a, 0x73,
	0x6f, 0x6e, 0x12, 0x0f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x82, 0xea, 0x30, 0x18, 0x0a, 0x05, 0x79, 0x61, 0x6d, 0x6c, 0x31, 0x12, 0x0f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x08,
	0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x31, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x3b, 0x74,
	0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_test_test_proto_rawDescOnce	sync.Once
	file_proto_test_test_proto_rawDescData	= file_proto_test_test_proto_rawDesc
)

func file_proto_test_test_proto_rawDescGZIP() []byte {
	file_proto_test_test_proto_rawDescOnce.Do(func() {
		file_proto_test_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_test_test_proto_rawDescData)
	})
	return file_proto_test_test_proto_rawDescData
}

var file_proto_test_test_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_test_test_proto_goTypes = []interface{}{
	(*GenerateResponse)(nil),	// 0: test.GenerateResponse
}
var file_proto_test_test_proto_depIdxs = []int32{
	0,	// [0:0] is the sub-list for method output_type
	0,	// [0:0] is the sub-list for method input_type
	0,	// [0:0] is the sub-list for extension type_name
	0,	// [0:0] is the sub-list for extension extendee
	0,	// [0:0] is the sub-list for field type_name
}

func init()	{ file_proto_test_test_proto_init() }
func file_proto_test_test_proto_init() {
	if File_proto_test_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_test_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
	}
	file_proto_test_test_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*GenerateResponse_MediaFile)(nil),
		(*GenerateResponse_Data)(nil),
		(*GenerateResponse_AdIdToCopyImageFrom)(nil),
		(*GenerateResponse_MediaFile1)(nil),
		(*GenerateResponse_Data1)(nil),
		(*GenerateResponse_AdIdToCopyImageFrom1)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath:	reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor:	file_proto_test_test_proto_rawDesc,
			NumEnums:	0,
			NumMessages:	1,
			NumExtensions:	0,
			NumServices:	0,
		},
		GoTypes:		file_proto_test_test_proto_goTypes,
		DependencyIndexes:	file_proto_test_test_proto_depIdxs,
		MessageInfos:		file_proto_test_test_proto_msgTypes,
	}.Build()
	File_proto_test_test_proto = out.File
	file_proto_test_test_proto_rawDesc = nil
	file_proto_test_test_proto_goTypes = nil
	file_proto_test_test_proto_depIdxs = nil
}
