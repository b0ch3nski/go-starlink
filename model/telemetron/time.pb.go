// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        v4.25.3
// source: spacex/api/telemetron/public/common/time.proto

package telemetron

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

type Epoch int32

const (
	Epoch_UNIX Epoch = 0
	Epoch_GPS  Epoch = 1
)

// Enum value maps for Epoch.
var (
	Epoch_name = map[int32]string{
		0: "UNIX",
		1: "GPS",
	}
	Epoch_value = map[string]int32{
		"UNIX": 0,
		"GPS":  1,
	}
)

func (x Epoch) Enum() *Epoch {
	p := new(Epoch)
	*p = x
	return p
}

func (x Epoch) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Epoch) Descriptor() protoreflect.EnumDescriptor {
	return file_spacex_api_telemetron_public_common_time_proto_enumTypes[0].Descriptor()
}

func (Epoch) Type() protoreflect.EnumType {
	return &file_spacex_api_telemetron_public_common_time_proto_enumTypes[0]
}

func (x Epoch) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Epoch.Descriptor instead.
func (Epoch) EnumDescriptor() ([]byte, []int) {
	return file_spacex_api_telemetron_public_common_time_proto_rawDescGZIP(), []int{0}
}

type TimestampInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Epoch       Epoch `protobuf:"varint,1,opt,name=epoch,proto3,enum=SpaceX.API.Telemetron.Public.Common.Epoch" json:"epoch,omitempty"`
	Nanoseconds int64 `protobuf:"varint,2,opt,name=nanoseconds,proto3" json:"nanoseconds,omitempty"`
}

func (x *TimestampInfo) Reset() {
	*x = TimestampInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacex_api_telemetron_public_common_time_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimestampInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimestampInfo) ProtoMessage() {}

func (x *TimestampInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spacex_api_telemetron_public_common_time_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimestampInfo.ProtoReflect.Descriptor instead.
func (*TimestampInfo) Descriptor() ([]byte, []int) {
	return file_spacex_api_telemetron_public_common_time_proto_rawDescGZIP(), []int{0}
}

func (x *TimestampInfo) GetEpoch() Epoch {
	if x != nil {
		return x.Epoch
	}
	return Epoch_UNIX
}

func (x *TimestampInfo) GetNanoseconds() int64 {
	if x != nil {
		return x.Nanoseconds
	}
	return 0
}

var File_spacex_api_telemetron_public_common_time_proto protoreflect.FileDescriptor

var file_spacex_api_telemetron_public_common_time_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x78, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65, 0x6c,
	0x65, 0x6d, 0x65, 0x74, 0x72, 0x6f, 0x6e, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x23, 0x53, 0x70, 0x61, 0x63, 0x65, 0x58, 0x2e, 0x41, 0x50, 0x49, 0x2e, 0x54, 0x65, 0x6c,
	0x65, 0x6d, 0x65, 0x74, 0x72, 0x6f, 0x6e, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0x73, 0x0a, 0x0d, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x40, 0x0a, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x53, 0x70, 0x61, 0x63, 0x65, 0x58, 0x2e, 0x41,
	0x50, 0x49, 0x2e, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x6f, 0x6e, 0x2e, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x70, 0x6f, 0x63,
	0x68, 0x52, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x61, 0x6e, 0x6f,
	0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6e,
	0x61, 0x6e, 0x6f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x2a, 0x1a, 0x0a, 0x05, 0x45, 0x70,
	0x6f, 0x63, 0x68, 0x12, 0x08, 0x0a, 0x04, 0x55, 0x4e, 0x49, 0x58, 0x10, 0x00, 0x12, 0x07, 0x0a,
	0x03, 0x47, 0x50, 0x53, 0x10, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spacex_api_telemetron_public_common_time_proto_rawDescOnce sync.Once
	file_spacex_api_telemetron_public_common_time_proto_rawDescData = file_spacex_api_telemetron_public_common_time_proto_rawDesc
)

func file_spacex_api_telemetron_public_common_time_proto_rawDescGZIP() []byte {
	file_spacex_api_telemetron_public_common_time_proto_rawDescOnce.Do(func() {
		file_spacex_api_telemetron_public_common_time_proto_rawDescData = protoimpl.X.CompressGZIP(file_spacex_api_telemetron_public_common_time_proto_rawDescData)
	})
	return file_spacex_api_telemetron_public_common_time_proto_rawDescData
}

var file_spacex_api_telemetron_public_common_time_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_spacex_api_telemetron_public_common_time_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_spacex_api_telemetron_public_common_time_proto_goTypes = []interface{}{
	(Epoch)(0),            // 0: SpaceX.API.Telemetron.Public.Common.Epoch
	(*TimestampInfo)(nil), // 1: SpaceX.API.Telemetron.Public.Common.TimestampInfo
}
var file_spacex_api_telemetron_public_common_time_proto_depIdxs = []int32{
	0, // 0: SpaceX.API.Telemetron.Public.Common.TimestampInfo.epoch:type_name -> SpaceX.API.Telemetron.Public.Common.Epoch
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_spacex_api_telemetron_public_common_time_proto_init() }
func file_spacex_api_telemetron_public_common_time_proto_init() {
	if File_spacex_api_telemetron_public_common_time_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spacex_api_telemetron_public_common_time_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimestampInfo); i {
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
			RawDescriptor: file_spacex_api_telemetron_public_common_time_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_spacex_api_telemetron_public_common_time_proto_goTypes,
		DependencyIndexes: file_spacex_api_telemetron_public_common_time_proto_depIdxs,
		EnumInfos:         file_spacex_api_telemetron_public_common_time_proto_enumTypes,
		MessageInfos:      file_spacex_api_telemetron_public_common_time_proto_msgTypes,
	}.Build()
	File_spacex_api_telemetron_public_common_time_proto = out.File
	file_spacex_api_telemetron_public_common_time_proto_rawDesc = nil
	file_spacex_api_telemetron_public_common_time_proto_goTypes = nil
	file_spacex_api_telemetron_public_common_time_proto_depIdxs = nil
}
