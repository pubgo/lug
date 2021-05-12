// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.11.4
// source: runmode.proto

package config

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

type RunMode int32

const (
	RunMode_dev     RunMode = 0
	RunMode_test    RunMode = 1
	RunMode_stag    RunMode = 2
	RunMode_prod    RunMode = 3
	RunMode_release RunMode = 4
	RunMode_unknown RunMode = 5
)

// Enum value maps for RunMode.
var (
	RunMode_name = map[int32]string{
		0: "dev",
		1: "test",
		2: "stag",
		3: "prod",
		4: "release",
		5: "unknown",
	}
	RunMode_value = map[string]int32{
		"dev":     0,
		"test":    1,
		"stag":    2,
		"prod":    3,
		"release": 4,
		"unknown": 5,
	}
)

func (x RunMode) Enum() *RunMode {
	p := new(RunMode)
	*p = x
	return p
}

func (x RunMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RunMode) Descriptor() protoreflect.EnumDescriptor {
	return file_runmode_proto_enumTypes[0].Descriptor()
}

func (RunMode) Type() protoreflect.EnumType {
	return &file_runmode_proto_enumTypes[0]
}

func (x RunMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RunMode.Descriptor instead.
func (RunMode) EnumDescriptor() ([]byte, []int) {
	return file_runmode_proto_rawDescGZIP(), []int{0}
}

var File_runmode_proto protoreflect.FileDescriptor

var file_runmode_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x75, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2a, 0x4a, 0x0a, 0x07, 0x52, 0x75, 0x6e, 0x4d, 0x6f,
	0x64, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x64, 0x65, 0x76, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x74,
	0x65, 0x73, 0x74, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x73, 0x74, 0x61, 0x67, 0x10, 0x02, 0x12,
	0x08, 0x0a, 0x04, 0x70, 0x72, 0x6f, 0x64, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x72, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77,
	0x6e, 0x10, 0x05, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_runmode_proto_rawDescOnce sync.Once
	file_runmode_proto_rawDescData = file_runmode_proto_rawDesc
)

func file_runmode_proto_rawDescGZIP() []byte {
	file_runmode_proto_rawDescOnce.Do(func() {
		file_runmode_proto_rawDescData = protoimpl.X.CompressGZIP(file_runmode_proto_rawDescData)
	})
	return file_runmode_proto_rawDescData
}

var file_runmode_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_runmode_proto_goTypes = []interface{}{
	(RunMode)(0), // 0: config.RunMode
}
var file_runmode_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_runmode_proto_init() }
func file_runmode_proto_init() {
	if File_runmode_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_runmode_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_runmode_proto_goTypes,
		DependencyIndexes: file_runmode_proto_depIdxs,
		EnumInfos:         file_runmode_proto_enumTypes,
	}.Build()
	File_runmode_proto = out.File
	file_runmode_proto_rawDesc = nil
	file_runmode_proto_goTypes = nil
	file_runmode_proto_depIdxs = nil
}
