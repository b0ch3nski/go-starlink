// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.2
// source: spacex/api/device/dish_config.proto

package device

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

type DishConfig_SnowMeltMode int32

const (
	DishConfig_AUTO       DishConfig_SnowMeltMode = 0
	DishConfig_ALWAYS_ON  DishConfig_SnowMeltMode = 1
	DishConfig_ALWAYS_OFF DishConfig_SnowMeltMode = 2
)

// Enum value maps for DishConfig_SnowMeltMode.
var (
	DishConfig_SnowMeltMode_name = map[int32]string{
		0: "AUTO",
		1: "ALWAYS_ON",
		2: "ALWAYS_OFF",
	}
	DishConfig_SnowMeltMode_value = map[string]int32{
		"AUTO":       0,
		"ALWAYS_ON":  1,
		"ALWAYS_OFF": 2,
	}
)

func (x DishConfig_SnowMeltMode) Enum() *DishConfig_SnowMeltMode {
	p := new(DishConfig_SnowMeltMode)
	*p = x
	return p
}

func (x DishConfig_SnowMeltMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DishConfig_SnowMeltMode) Descriptor() protoreflect.EnumDescriptor {
	return file_spacex_api_device_dish_config_proto_enumTypes[0].Descriptor()
}

func (DishConfig_SnowMeltMode) Type() protoreflect.EnumType {
	return &file_spacex_api_device_dish_config_proto_enumTypes[0]
}

func (x DishConfig_SnowMeltMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DishConfig_SnowMeltMode.Descriptor instead.
func (DishConfig_SnowMeltMode) EnumDescriptor() ([]byte, []int) {
	return file_spacex_api_device_dish_config_proto_rawDescGZIP(), []int{0, 0}
}

type DishConfig_LocationRequestMode int32

const (
	DishConfig_NONE  DishConfig_LocationRequestMode = 0
	DishConfig_LOCAL DishConfig_LocationRequestMode = 1
)

// Enum value maps for DishConfig_LocationRequestMode.
var (
	DishConfig_LocationRequestMode_name = map[int32]string{
		0: "NONE",
		1: "LOCAL",
	}
	DishConfig_LocationRequestMode_value = map[string]int32{
		"NONE":  0,
		"LOCAL": 1,
	}
)

func (x DishConfig_LocationRequestMode) Enum() *DishConfig_LocationRequestMode {
	p := new(DishConfig_LocationRequestMode)
	*p = x
	return p
}

func (x DishConfig_LocationRequestMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DishConfig_LocationRequestMode) Descriptor() protoreflect.EnumDescriptor {
	return file_spacex_api_device_dish_config_proto_enumTypes[1].Descriptor()
}

func (DishConfig_LocationRequestMode) Type() protoreflect.EnumType {
	return &file_spacex_api_device_dish_config_proto_enumTypes[1]
}

func (x DishConfig_LocationRequestMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DishConfig_LocationRequestMode.Descriptor instead.
func (DishConfig_LocationRequestMode) EnumDescriptor() ([]byte, []int) {
	return file_spacex_api_device_dish_config_proto_rawDescGZIP(), []int{0, 1}
}

type DishConfig_LevelDishMode int32

const (
	DishConfig_TILT_LIKE_NORMAL DishConfig_LevelDishMode = 0
	DishConfig_FORCE_LEVEL      DishConfig_LevelDishMode = 1
)

// Enum value maps for DishConfig_LevelDishMode.
var (
	DishConfig_LevelDishMode_name = map[int32]string{
		0: "TILT_LIKE_NORMAL",
		1: "FORCE_LEVEL",
	}
	DishConfig_LevelDishMode_value = map[string]int32{
		"TILT_LIKE_NORMAL": 0,
		"FORCE_LEVEL":      1,
	}
)

func (x DishConfig_LevelDishMode) Enum() *DishConfig_LevelDishMode {
	p := new(DishConfig_LevelDishMode)
	*p = x
	return p
}

func (x DishConfig_LevelDishMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DishConfig_LevelDishMode) Descriptor() protoreflect.EnumDescriptor {
	return file_spacex_api_device_dish_config_proto_enumTypes[2].Descriptor()
}

func (DishConfig_LevelDishMode) Type() protoreflect.EnumType {
	return &file_spacex_api_device_dish_config_proto_enumTypes[2]
}

func (x DishConfig_LevelDishMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DishConfig_LevelDishMode.Descriptor instead.
func (DishConfig_LevelDishMode) EnumDescriptor() ([]byte, []int) {
	return file_spacex_api_device_dish_config_proto_rawDescGZIP(), []int{0, 2}
}

type DishConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SnowMeltMode                         DishConfig_SnowMeltMode        `protobuf:"varint,1,opt,name=snow_melt_mode,json=snowMeltMode,proto3,enum=SpaceX.API.Device.DishConfig_SnowMeltMode" json:"snow_melt_mode,omitempty"`
	ApplySnowMeltMode                    bool                           `protobuf:"varint,1001,opt,name=apply_snow_melt_mode,json=applySnowMeltMode,proto3" json:"apply_snow_melt_mode,omitempty"`
	LocationRequestMode                  DishConfig_LocationRequestMode `protobuf:"varint,2,opt,name=location_request_mode,json=locationRequestMode,proto3,enum=SpaceX.API.Device.DishConfig_LocationRequestMode" json:"location_request_mode,omitempty"`
	ApplyLocationRequestMode             bool                           `protobuf:"varint,2001,opt,name=apply_location_request_mode,json=applyLocationRequestMode,proto3" json:"apply_location_request_mode,omitempty"`
	LevelDishMode                        DishConfig_LevelDishMode       `protobuf:"varint,3,opt,name=level_dish_mode,json=levelDishMode,proto3,enum=SpaceX.API.Device.DishConfig_LevelDishMode" json:"level_dish_mode,omitempty"`
	ApplyLevelDishMode                   bool                           `protobuf:"varint,3001,opt,name=apply_level_dish_mode,json=applyLevelDishMode,proto3" json:"apply_level_dish_mode,omitempty"`
	PowerSaveStartMinutes                uint32                         `protobuf:"varint,4,opt,name=power_save_start_minutes,json=powerSaveStartMinutes,proto3" json:"power_save_start_minutes,omitempty"`
	ApplyPowerSaveStartMinutes           bool                           `protobuf:"varint,4001,opt,name=apply_power_save_start_minutes,json=applyPowerSaveStartMinutes,proto3" json:"apply_power_save_start_minutes,omitempty"`
	PowerSaveDurationMinutes             uint32                         `protobuf:"varint,5,opt,name=power_save_duration_minutes,json=powerSaveDurationMinutes,proto3" json:"power_save_duration_minutes,omitempty"`
	ApplyPowerSaveDurationMinutes        bool                           `protobuf:"varint,5001,opt,name=apply_power_save_duration_minutes,json=applyPowerSaveDurationMinutes,proto3" json:"apply_power_save_duration_minutes,omitempty"`
	PowerSaveMode                        bool                           `protobuf:"varint,6,opt,name=power_save_mode,json=powerSaveMode,proto3" json:"power_save_mode,omitempty"`
	ApplyPowerSaveMode                   bool                           `protobuf:"varint,6001,opt,name=apply_power_save_mode,json=applyPowerSaveMode,proto3" json:"apply_power_save_mode,omitempty"`
	SwupdateThreeDayDeferralEnabled      bool                           `protobuf:"varint,7,opt,name=swupdate_three_day_deferral_enabled,json=swupdateThreeDayDeferralEnabled,proto3" json:"swupdate_three_day_deferral_enabled,omitempty"`
	ApplySwupdateThreeDayDeferralEnabled bool                           `protobuf:"varint,7001,opt,name=apply_swupdate_three_day_deferral_enabled,json=applySwupdateThreeDayDeferralEnabled,proto3" json:"apply_swupdate_three_day_deferral_enabled,omitempty"`
	AssetClass                           uint32                         `protobuf:"varint,8,opt,name=asset_class,json=assetClass,proto3" json:"asset_class,omitempty"`
	ApplyAssetClass                      bool                           `protobuf:"varint,8001,opt,name=apply_asset_class,json=applyAssetClass,proto3" json:"apply_asset_class,omitempty"`
}

func (x *DishConfig) Reset() {
	*x = DishConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spacex_api_device_dish_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DishConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DishConfig) ProtoMessage() {}

func (x *DishConfig) ProtoReflect() protoreflect.Message {
	mi := &file_spacex_api_device_dish_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DishConfig.ProtoReflect.Descriptor instead.
func (*DishConfig) Descriptor() ([]byte, []int) {
	return file_spacex_api_device_dish_config_proto_rawDescGZIP(), []int{0}
}

func (x *DishConfig) GetSnowMeltMode() DishConfig_SnowMeltMode {
	if x != nil {
		return x.SnowMeltMode
	}
	return DishConfig_AUTO
}

func (x *DishConfig) GetApplySnowMeltMode() bool {
	if x != nil {
		return x.ApplySnowMeltMode
	}
	return false
}

func (x *DishConfig) GetLocationRequestMode() DishConfig_LocationRequestMode {
	if x != nil {
		return x.LocationRequestMode
	}
	return DishConfig_NONE
}

func (x *DishConfig) GetApplyLocationRequestMode() bool {
	if x != nil {
		return x.ApplyLocationRequestMode
	}
	return false
}

func (x *DishConfig) GetLevelDishMode() DishConfig_LevelDishMode {
	if x != nil {
		return x.LevelDishMode
	}
	return DishConfig_TILT_LIKE_NORMAL
}

func (x *DishConfig) GetApplyLevelDishMode() bool {
	if x != nil {
		return x.ApplyLevelDishMode
	}
	return false
}

func (x *DishConfig) GetPowerSaveStartMinutes() uint32 {
	if x != nil {
		return x.PowerSaveStartMinutes
	}
	return 0
}

func (x *DishConfig) GetApplyPowerSaveStartMinutes() bool {
	if x != nil {
		return x.ApplyPowerSaveStartMinutes
	}
	return false
}

func (x *DishConfig) GetPowerSaveDurationMinutes() uint32 {
	if x != nil {
		return x.PowerSaveDurationMinutes
	}
	return 0
}

func (x *DishConfig) GetApplyPowerSaveDurationMinutes() bool {
	if x != nil {
		return x.ApplyPowerSaveDurationMinutes
	}
	return false
}

func (x *DishConfig) GetPowerSaveMode() bool {
	if x != nil {
		return x.PowerSaveMode
	}
	return false
}

func (x *DishConfig) GetApplyPowerSaveMode() bool {
	if x != nil {
		return x.ApplyPowerSaveMode
	}
	return false
}

func (x *DishConfig) GetSwupdateThreeDayDeferralEnabled() bool {
	if x != nil {
		return x.SwupdateThreeDayDeferralEnabled
	}
	return false
}

func (x *DishConfig) GetApplySwupdateThreeDayDeferralEnabled() bool {
	if x != nil {
		return x.ApplySwupdateThreeDayDeferralEnabled
	}
	return false
}

func (x *DishConfig) GetAssetClass() uint32 {
	if x != nil {
		return x.AssetClass
	}
	return 0
}

func (x *DishConfig) GetApplyAssetClass() bool {
	if x != nil {
		return x.ApplyAssetClass
	}
	return false
}

var File_spacex_api_device_dish_config_proto protoreflect.FileDescriptor

var file_spacex_api_device_dish_config_proto_rawDesc = []byte{
	0x0a, 0x23, 0x73, 0x70, 0x61, 0x63, 0x65, 0x78, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x64, 0x69, 0x73, 0x68, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x53, 0x70, 0x61, 0x63, 0x65, 0x58, 0x2e, 0x41, 0x50,
	0x49, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x22, 0xb7, 0x09, 0x0a, 0x0a, 0x44, 0x69, 0x73,
	0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x50, 0x0a, 0x0e, 0x73, 0x6e, 0x6f, 0x77, 0x5f,
	0x6d, 0x65, 0x6c, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x2a, 0x2e, 0x53, 0x70, 0x61, 0x63, 0x65, 0x58, 0x2e, 0x41, 0x50, 0x49, 0x2e, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x44, 0x69, 0x73, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53,
	0x6e, 0x6f, 0x77, 0x4d, 0x65, 0x6c, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0c, 0x73, 0x6e, 0x6f,
	0x77, 0x4d, 0x65, 0x6c, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x61, 0x70, 0x70,
	0x6c, 0x79, 0x5f, 0x73, 0x6e, 0x6f, 0x77, 0x5f, 0x6d, 0x65, 0x6c, 0x74, 0x5f, 0x6d, 0x6f, 0x64,
	0x65, 0x18, 0xe9, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x53,
	0x6e, 0x6f, 0x77, 0x4d, 0x65, 0x6c, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x65, 0x0a, 0x15, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x31, 0x2e, 0x53, 0x70, 0x61,
	0x63, 0x65, 0x58, 0x2e, 0x41, 0x50, 0x49, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44,
	0x69, 0x73, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x13, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x6f,
	0x64, 0x65, 0x12, 0x3e, 0x0a, 0x1b, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x6d, 0x6f, 0x64,
	0x65, 0x18, 0xd1, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x18, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x6f,
	0x64, 0x65, 0x12, 0x53, 0x0a, 0x0f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x64, 0x69, 0x73, 0x68,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x53, 0x70,
	0x61, 0x63, 0x65, 0x58, 0x2e, 0x41, 0x50, 0x49, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x44, 0x69, 0x73, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x44, 0x69, 0x73, 0x68, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0d, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x44,
	0x69, 0x73, 0x68, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x32, 0x0a, 0x15, 0x61, 0x70, 0x70, 0x6c, 0x79,
	0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x64, 0x69, 0x73, 0x68, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x18, 0xb9, 0x17, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x44, 0x69, 0x73, 0x68, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x37, 0x0a, 0x18, 0x70,
	0x6f, 0x77, 0x65, 0x72, 0x5f, 0x73, 0x61, 0x76, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x15, 0x70,
	0x6f, 0x77, 0x65, 0x72, 0x53, 0x61, 0x76, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4d, 0x69, 0x6e,
	0x75, 0x74, 0x65, 0x73, 0x12, 0x43, 0x0a, 0x1e, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x70, 0x6f,
	0x77, 0x65, 0x72, 0x5f, 0x73, 0x61, 0x76, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x6d,
	0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x18, 0xa1, 0x1f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1a, 0x61,
	0x70, 0x70, 0x6c, 0x79, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x53, 0x61, 0x76, 0x65, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x1b, 0x70, 0x6f, 0x77,
	0x65, 0x72, 0x5f, 0x73, 0x61, 0x76, 0x65, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x18,
	0x70, 0x6f, 0x77, 0x65, 0x72, 0x53, 0x61, 0x76, 0x65, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x12, 0x49, 0x0a, 0x21, 0x61, 0x70, 0x70, 0x6c,
	0x79, 0x5f, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x73, 0x61, 0x76, 0x65, 0x5f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x18, 0x89, 0x27,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x1d, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x50, 0x6f, 0x77, 0x65, 0x72,
	0x53, 0x61, 0x76, 0x65, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x69, 0x6e, 0x75,
	0x74, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x73, 0x61, 0x76,
	0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x70, 0x6f,
	0x77, 0x65, 0x72, 0x53, 0x61, 0x76, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x32, 0x0a, 0x15, 0x61,
	0x70, 0x70, 0x6c, 0x79, 0x5f, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x73, 0x61, 0x76, 0x65, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x18, 0xf1, 0x2e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x61, 0x70, 0x70,
	0x6c, 0x79, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x53, 0x61, 0x76, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12,
	0x4c, 0x0a, 0x23, 0x73, 0x77, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x68, 0x72, 0x65,
	0x65, 0x5f, 0x64, 0x61, 0x79, 0x5f, 0x64, 0x65, 0x66, 0x65, 0x72, 0x72, 0x61, 0x6c, 0x5f, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1f, 0x73, 0x77,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x68, 0x72, 0x65, 0x65, 0x44, 0x61, 0x79, 0x44, 0x65,
	0x66, 0x65, 0x72, 0x72, 0x61, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x58, 0x0a,
	0x29, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x73, 0x77, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x74, 0x68, 0x72, 0x65, 0x65, 0x5f, 0x64, 0x61, 0x79, 0x5f, 0x64, 0x65, 0x66, 0x65, 0x72, 0x72,
	0x61, 0x6c, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0xd9, 0x36, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x24, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x53, 0x77, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x68, 0x72, 0x65, 0x65, 0x44, 0x61, 0x79, 0x44, 0x65, 0x66, 0x65, 0x72, 0x72, 0x61, 0x6c,
	0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x61, 0x70, 0x70, 0x6c,
	0x79, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0xc1, 0x3e,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x41, 0x73, 0x73, 0x65, 0x74,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x22, 0x37, 0x0a, 0x0c, 0x53, 0x6e, 0x6f, 0x77, 0x4d, 0x65, 0x6c,
	0x74, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x41, 0x55, 0x54, 0x4f, 0x10, 0x00, 0x12,
	0x0d, 0x0a, 0x09, 0x41, 0x4c, 0x57, 0x41, 0x59, 0x53, 0x5f, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x0e,
	0x0a, 0x0a, 0x41, 0x4c, 0x57, 0x41, 0x59, 0x53, 0x5f, 0x4f, 0x46, 0x46, 0x10, 0x02, 0x22, 0x2a,
	0x0a, 0x13, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12,
	0x09, 0x0a, 0x05, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x10, 0x01, 0x22, 0x36, 0x0a, 0x0d, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x44, 0x69, 0x73, 0x68, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54,
	0x49, 0x4c, 0x54, 0x5f, 0x4c, 0x49, 0x4b, 0x45, 0x5f, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10,
	0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x46, 0x4f, 0x52, 0x43, 0x45, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c,
	0x10, 0x01, 0x42, 0x17, 0x5a, 0x15, 0x73, 0x70, 0x61, 0x63, 0x65, 0x78, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_spacex_api_device_dish_config_proto_rawDescOnce sync.Once
	file_spacex_api_device_dish_config_proto_rawDescData = file_spacex_api_device_dish_config_proto_rawDesc
)

func file_spacex_api_device_dish_config_proto_rawDescGZIP() []byte {
	file_spacex_api_device_dish_config_proto_rawDescOnce.Do(func() {
		file_spacex_api_device_dish_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_spacex_api_device_dish_config_proto_rawDescData)
	})
	return file_spacex_api_device_dish_config_proto_rawDescData
}

var file_spacex_api_device_dish_config_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_spacex_api_device_dish_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_spacex_api_device_dish_config_proto_goTypes = []interface{}{
	(DishConfig_SnowMeltMode)(0),        // 0: SpaceX.API.Device.DishConfig.SnowMeltMode
	(DishConfig_LocationRequestMode)(0), // 1: SpaceX.API.Device.DishConfig.LocationRequestMode
	(DishConfig_LevelDishMode)(0),       // 2: SpaceX.API.Device.DishConfig.LevelDishMode
	(*DishConfig)(nil),                  // 3: SpaceX.API.Device.DishConfig
}
var file_spacex_api_device_dish_config_proto_depIdxs = []int32{
	0, // 0: SpaceX.API.Device.DishConfig.snow_melt_mode:type_name -> SpaceX.API.Device.DishConfig.SnowMeltMode
	1, // 1: SpaceX.API.Device.DishConfig.location_request_mode:type_name -> SpaceX.API.Device.DishConfig.LocationRequestMode
	2, // 2: SpaceX.API.Device.DishConfig.level_dish_mode:type_name -> SpaceX.API.Device.DishConfig.LevelDishMode
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_spacex_api_device_dish_config_proto_init() }
func file_spacex_api_device_dish_config_proto_init() {
	if File_spacex_api_device_dish_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spacex_api_device_dish_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DishConfig); i {
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
			RawDescriptor: file_spacex_api_device_dish_config_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_spacex_api_device_dish_config_proto_goTypes,
		DependencyIndexes: file_spacex_api_device_dish_config_proto_depIdxs,
		EnumInfos:         file_spacex_api_device_dish_config_proto_enumTypes,
		MessageInfos:      file_spacex_api_device_dish_config_proto_msgTypes,
	}.Build()
	File_spacex_api_device_dish_config_proto = out.File
	file_spacex_api_device_dish_config_proto_rawDesc = nil
	file_spacex_api_device_dish_config_proto_goTypes = nil
	file_spacex_api_device_dish_config_proto_depIdxs = nil
}
