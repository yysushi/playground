// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: example/v1/example.proto

package examplev1

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

type HogeEnum int32

const (
	HogeEnum_HOGE_ENUM_UNSPECIFIED HogeEnum = 0
	// Deprecated: Marked as deprecated in example/v1/example.proto.
	HogeEnum_HOGE_ENUM_SEARCH  HogeEnum = 1
	HogeEnum_HOGE_ENUM_DISPLAY HogeEnum = 2
)

// Enum value maps for HogeEnum.
var (
	HogeEnum_name = map[int32]string{
		0: "HOGE_ENUM_UNSPECIFIED",
		1: "HOGE_ENUM_SEARCH",
		2: "HOGE_ENUM_DISPLAY",
	}
	HogeEnum_value = map[string]int32{
		"HOGE_ENUM_UNSPECIFIED": 0,
		"HOGE_ENUM_SEARCH":      1,
		"HOGE_ENUM_DISPLAY":     2,
	}
)

func (x HogeEnum) Enum() *HogeEnum {
	p := new(HogeEnum)
	*p = x
	return p
}

func (x HogeEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HogeEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_example_v1_example_proto_enumTypes[0].Descriptor()
}

func (HogeEnum) Type() protoreflect.EnumType {
	return &file_example_v1_example_proto_enumTypes[0]
}

func (x HogeEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HogeEnum.Descriptor instead.
func (HogeEnum) EnumDescriptor() ([]byte, []int) {
	return file_example_v1_example_proto_rawDescGZIP(), []int{0}
}

type HogeMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HogeField string `protobuf:"bytes,1,opt,name=hoge_field,json=hogeField,proto3" json:"hoge_field,omitempty"`
}

func (x *HogeMessage) Reset() {
	*x = HogeMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_v1_example_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HogeMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HogeMessage) ProtoMessage() {}

func (x *HogeMessage) ProtoReflect() protoreflect.Message {
	mi := &file_example_v1_example_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HogeMessage.ProtoReflect.Descriptor instead.
func (*HogeMessage) Descriptor() ([]byte, []int) {
	return file_example_v1_example_proto_rawDescGZIP(), []int{0}
}

func (x *HogeMessage) GetHogeField() string {
	if x != nil {
		return x.HogeField
	}
	return ""
}

var file_example_v1_example_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         123456789,
		Name:          "example.v1.string_name",
		Tag:           "bytes,123456789,opt,name=string_name",
		Filename:      "example/v1/example.proto",
	},
}

// Extension fields to descriptorpb.EnumValueOptions.
var (
	// optional string string_name = 123456789;
	E_StringName = &file_example_v1_example_proto_extTypes[0]
)

var File_example_v1_example_proto protoreflect.FileDescriptor

var file_example_v1_example_proto_rawDesc = []byte{
	0x0a, 0x18, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x0b, 0x48, 0x6f, 0x67, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x6f, 0x67, 0x65, 0x5f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x68, 0x6f, 0x67,
	0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x2a, 0x6b, 0x0a, 0x08, 0x48, 0x6f, 0x67, 0x65, 0x45, 0x6e,
	0x75, 0x6d, 0x12, 0x19, 0x0a, 0x15, 0x48, 0x4f, 0x47, 0x45, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x18, 0x0a,
	0x10, 0x48, 0x4f, 0x47, 0x45, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x53, 0x45, 0x41, 0x52, 0x43,
	0x48, 0x10, 0x01, 0x1a, 0x02, 0x08, 0x01, 0x12, 0x2a, 0x0a, 0x11, 0x48, 0x4f, 0x47, 0x45, 0x5f,
	0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x44, 0x49, 0x53, 0x50, 0x4c, 0x41, 0x59, 0x10, 0x02, 0x1a, 0x13,
	0xaa, 0xd1, 0xf9, 0xd6, 0x03, 0x0d, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x48, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x95, 0x9a, 0xef, 0x3a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0xc1, 0x01,
	0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x42, 0x0c, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x58, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x79, 0x73,
	0x75, 0x73, 0x68, 0x69, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x75,
	0x66, 0x66, 0x65, 0x72, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x76, 0x31,
	0x3b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x45, 0x58, 0x58,
	0xaa, 0x02, 0x0a, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0a,
	0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x16, 0x45, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x0b, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_example_v1_example_proto_rawDescOnce sync.Once
	file_example_v1_example_proto_rawDescData = file_example_v1_example_proto_rawDesc
)

func file_example_v1_example_proto_rawDescGZIP() []byte {
	file_example_v1_example_proto_rawDescOnce.Do(func() {
		file_example_v1_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_v1_example_proto_rawDescData)
	})
	return file_example_v1_example_proto_rawDescData
}

var file_example_v1_example_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_example_v1_example_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_example_v1_example_proto_goTypes = []interface{}{
	(HogeEnum)(0),                         // 0: example.v1.HogeEnum
	(*HogeMessage)(nil),                   // 1: example.v1.HogeMessage
	(*descriptorpb.EnumValueOptions)(nil), // 2: google.protobuf.EnumValueOptions
}
var file_example_v1_example_proto_depIdxs = []int32{
	2, // 0: example.v1.string_name:extendee -> google.protobuf.EnumValueOptions
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_example_v1_example_proto_init() }
func file_example_v1_example_proto_init() {
	if File_example_v1_example_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_example_v1_example_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HogeMessage); i {
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
			RawDescriptor: file_example_v1_example_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_example_v1_example_proto_goTypes,
		DependencyIndexes: file_example_v1_example_proto_depIdxs,
		EnumInfos:         file_example_v1_example_proto_enumTypes,
		MessageInfos:      file_example_v1_example_proto_msgTypes,
		ExtensionInfos:    file_example_v1_example_proto_extTypes,
	}.Build()
	File_example_v1_example_proto = out.File
	file_example_v1_example_proto_rawDesc = nil
	file_example_v1_example_proto_goTypes = nil
	file_example_v1_example_proto_depIdxs = nil
}