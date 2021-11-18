// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: errors/test/test.proto

package test

import (
	_ "github.com/pubgo/lava/errors"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Test int32

const (
	// NotFound 找不到
	Test_NotFound Test = 0
	// Unknown 未知
	Test_Unknown Test = 1
)

// Enum value maps for Test.
var (
	Test_name = map[int32]string{
		0: "NotFound",
		1: "Unknown",
	}
	Test_value = map[string]int32{
		"NotFound": 0,
		"Unknown":  1,
	}
)

func (x Test) Enum() *Test {
	p := new(Test)
	*p = x
	return p
}

func (x Test) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Test) Descriptor() protoreflect.EnumDescriptor {
	return file_errors_test_test_proto_enumTypes[0].Descriptor()
}

func (Test) Type() protoreflect.EnumType {
	return &file_errors_test_test_proto_enumTypes[0]
}

func (x Test) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Test.Descriptor instead.
func (Test) EnumDescriptor() ([]byte, []int) {
	return file_errors_test_test_proto_rawDescGZIP(), []int{0}
}

var File_errors_test_test_proto protoreflect.FileDescriptor

var file_errors_test_test_proto_rawDesc = []byte{
	0x0a, 0x16, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x65, 0x73, 0x74, 0x1a, 0x20,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x31, 0x0a, 0x04, 0x54, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x08, 0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x10, 0x00, 0x1a, 0x06, 0x80, 0xea, 0x30,
	0xa0, 0x8d, 0x06, 0x12, 0x13, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x01,
	0x1a, 0x06, 0x80, 0xea, 0x30, 0xa1, 0x8d, 0x06, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x3b, 0x74,
	0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_errors_test_test_proto_rawDescOnce sync.Once
	file_errors_test_test_proto_rawDescData = file_errors_test_test_proto_rawDesc
)

func file_errors_test_test_proto_rawDescGZIP() []byte {
	file_errors_test_test_proto_rawDescOnce.Do(func() {
		file_errors_test_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_errors_test_test_proto_rawDescData)
	})
	return file_errors_test_test_proto_rawDescData
}

var file_errors_test_test_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_errors_test_test_proto_goTypes = []interface{}{
	(Test)(0), // 0: test.Test
}
var file_errors_test_test_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_errors_test_test_proto_init() }
func file_errors_test_test_proto_init() {
	if File_errors_test_test_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_errors_test_test_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_errors_test_test_proto_goTypes,
		DependencyIndexes: file_errors_test_test_proto_depIdxs,
		EnumInfos:         file_errors_test_test_proto_enumTypes,
	}.Build()
	File_errors_test_test_proto = out.File
	file_errors_test_test_proto_rawDesc = nil
	file_errors_test_test_proto_goTypes = nil
	file_errors_test_test_proto_depIdxs = nil
}
