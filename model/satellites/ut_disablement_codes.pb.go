// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.25.3
// source: spacex/api/satellites/network/ut_disablement_codes.proto

package satellites

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

type UtDisablementCode int32

const (
	UtDisablementCode_UNKNOWN_STATE                UtDisablementCode = 0
	UtDisablementCode_OKAY                         UtDisablementCode = 1
	UtDisablementCode_NO_ACTIVE_ACCOUNT            UtDisablementCode = 2
	UtDisablementCode_TOO_FAR_FROM_SERVICE_ADDRESS UtDisablementCode = 3
	UtDisablementCode_IN_OCEAN                     UtDisablementCode = 4
	UtDisablementCode_INVALID_COUNTRY              UtDisablementCode = 5
	UtDisablementCode_BLOCKED_COUNTRY              UtDisablementCode = 6
	UtDisablementCode_DATA_OVERAGE_SANDBOX_POLICY  UtDisablementCode = 7
)

// Enum value maps for UtDisablementCode.
var (
	UtDisablementCode_name = map[int32]string{
		0: "UNKNOWN_STATE",
		1: "OKAY",
		2: "NO_ACTIVE_ACCOUNT",
		3: "TOO_FAR_FROM_SERVICE_ADDRESS",
		4: "IN_OCEAN",
		5: "INVALID_COUNTRY",
		6: "BLOCKED_COUNTRY",
		7: "DATA_OVERAGE_SANDBOX_POLICY",
	}
	UtDisablementCode_value = map[string]int32{
		"UNKNOWN_STATE":                0,
		"OKAY":                         1,
		"NO_ACTIVE_ACCOUNT":            2,
		"TOO_FAR_FROM_SERVICE_ADDRESS": 3,
		"IN_OCEAN":                     4,
		"INVALID_COUNTRY":              5,
		"BLOCKED_COUNTRY":              6,
		"DATA_OVERAGE_SANDBOX_POLICY":  7,
	}
)

func (x UtDisablementCode) Enum() *UtDisablementCode {
	p := new(UtDisablementCode)
	*p = x
	return p
}

func (x UtDisablementCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UtDisablementCode) Descriptor() protoreflect.EnumDescriptor {
	return file_spacex_api_satellites_network_ut_disablement_codes_proto_enumTypes[0].Descriptor()
}

func (UtDisablementCode) Type() protoreflect.EnumType {
	return &file_spacex_api_satellites_network_ut_disablement_codes_proto_enumTypes[0]
}

func (x UtDisablementCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UtDisablementCode.Descriptor instead.
func (UtDisablementCode) EnumDescriptor() ([]byte, []int) {
	return file_spacex_api_satellites_network_ut_disablement_codes_proto_rawDescGZIP(), []int{0}
}

var File_spacex_api_satellites_network_ut_disablement_codes_proto protoreflect.FileDescriptor

var file_spacex_api_satellites_network_ut_disablement_codes_proto_rawDesc = []byte{
	0x0a, 0x38, 0x73, 0x70, 0x61, 0x63, 0x65, 0x78, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x61, 0x74,
	0x65, 0x6c, 0x6c, 0x69, 0x74, 0x65, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f,
	0x75, 0x74, 0x5f, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x53, 0x70, 0x61, 0x63,
	0x65, 0x58, 0x2e, 0x41, 0x50, 0x49, 0x2e, 0x53, 0x61, 0x74, 0x65, 0x6c, 0x6c, 0x69, 0x74, 0x65,
	0x73, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2a, 0xc2, 0x01, 0x0a, 0x11, 0x55, 0x74,
	0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45,
	0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4f, 0x4b, 0x41, 0x59, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11,
	0x4e, 0x4f, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e,
	0x54, 0x10, 0x02, 0x12, 0x20, 0x0a, 0x1c, 0x54, 0x4f, 0x4f, 0x5f, 0x46, 0x41, 0x52, 0x5f, 0x46,
	0x52, 0x4f, 0x4d, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x41, 0x44, 0x44, 0x52,
	0x45, 0x53, 0x53, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x4e, 0x5f, 0x4f, 0x43, 0x45, 0x41,
	0x4e, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x43,
	0x4f, 0x55, 0x4e, 0x54, 0x52, 0x59, 0x10, 0x05, 0x12, 0x13, 0x0a, 0x0f, 0x42, 0x4c, 0x4f, 0x43,
	0x4b, 0x45, 0x44, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x52, 0x59, 0x10, 0x06, 0x12, 0x1f, 0x0a,
	0x1b, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x4f, 0x56, 0x45, 0x52, 0x41, 0x47, 0x45, 0x5f, 0x53, 0x41,
	0x4e, 0x44, 0x42, 0x4f, 0x58, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x10, 0x07, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spacex_api_satellites_network_ut_disablement_codes_proto_rawDescOnce sync.Once
	file_spacex_api_satellites_network_ut_disablement_codes_proto_rawDescData = file_spacex_api_satellites_network_ut_disablement_codes_proto_rawDesc
)

func file_spacex_api_satellites_network_ut_disablement_codes_proto_rawDescGZIP() []byte {
	file_spacex_api_satellites_network_ut_disablement_codes_proto_rawDescOnce.Do(func() {
		file_spacex_api_satellites_network_ut_disablement_codes_proto_rawDescData = protoimpl.X.CompressGZIP(file_spacex_api_satellites_network_ut_disablement_codes_proto_rawDescData)
	})
	return file_spacex_api_satellites_network_ut_disablement_codes_proto_rawDescData
}

var file_spacex_api_satellites_network_ut_disablement_codes_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_spacex_api_satellites_network_ut_disablement_codes_proto_goTypes = []interface{}{
	(UtDisablementCode)(0), // 0: SpaceX.API.Satellites.Network.UtDisablementCode
}
var file_spacex_api_satellites_network_ut_disablement_codes_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_spacex_api_satellites_network_ut_disablement_codes_proto_init() }
func file_spacex_api_satellites_network_ut_disablement_codes_proto_init() {
	if File_spacex_api_satellites_network_ut_disablement_codes_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spacex_api_satellites_network_ut_disablement_codes_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_spacex_api_satellites_network_ut_disablement_codes_proto_goTypes,
		DependencyIndexes: file_spacex_api_satellites_network_ut_disablement_codes_proto_depIdxs,
		EnumInfos:         file_spacex_api_satellites_network_ut_disablement_codes_proto_enumTypes,
	}.Build()
	File_spacex_api_satellites_network_ut_disablement_codes_proto = out.File
	file_spacex_api_satellites_network_ut_disablement_codes_proto_rawDesc = nil
	file_spacex_api_satellites_network_ut_disablement_codes_proto_goTypes = nil
	file_spacex_api_satellites_network_ut_disablement_codes_proto_depIdxs = nil
}
