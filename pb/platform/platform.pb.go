// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: platform.proto

package platform

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

type NewPlatform struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UnitName        string  `protobuf:"bytes,1,opt,name=unitName,proto3" json:"unitName,omitempty"`
	UnitClass       string  `protobuf:"bytes,2,opt,name=unitClass,proto3" json:"unitClass,omitempty"`
	OperationField  string  `protobuf:"bytes,3,opt,name=operation_field,json=operationField,proto3" json:"operation_field,omitempty"`
	GeneralCategory string  `protobuf:"bytes,4,opt,name=generalCategory,proto3" json:"generalCategory,omitempty"`
	GeneralType     string  `protobuf:"bytes,5,opt,name=generalType,proto3" json:"generalType,omitempty"`
	LethalityLevel  string  `protobuf:"bytes,6,opt,name=lethalityLevel,proto3" json:"lethalityLevel,omitempty"`
	CruisingSpeed   float32 `protobuf:"fixed32,7,opt,name=cruisingSpeed,proto3" json:"cruisingSpeed,omitempty"`
	MaxSpeed        float32 `protobuf:"fixed32,8,opt,name=maxSpeed,proto3" json:"maxSpeed,omitempty"`
	CombatRange     float32 `protobuf:"fixed32,9,opt,name=combatRange,proto3" json:"combatRange,omitempty"`
	FuelLoad        float32 `protobuf:"fixed32,10,opt,name=fuelLoad,proto3" json:"fuelLoad,omitempty"`
	CountryOrigin   string  `protobuf:"bytes,11,opt,name=countryOrigin,proto3" json:"countryOrigin,omitempty"`
}

func (x *NewPlatform) Reset() {
	*x = NewPlatform{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewPlatform) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewPlatform) ProtoMessage() {}

func (x *NewPlatform) ProtoReflect() protoreflect.Message {
	mi := &file_platform_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewPlatform.ProtoReflect.Descriptor instead.
func (*NewPlatform) Descriptor() ([]byte, []int) {
	return file_platform_proto_rawDescGZIP(), []int{0}
}

func (x *NewPlatform) GetUnitName() string {
	if x != nil {
		return x.UnitName
	}
	return ""
}

func (x *NewPlatform) GetUnitClass() string {
	if x != nil {
		return x.UnitClass
	}
	return ""
}

func (x *NewPlatform) GetOperationField() string {
	if x != nil {
		return x.OperationField
	}
	return ""
}

func (x *NewPlatform) GetGeneralCategory() string {
	if x != nil {
		return x.GeneralCategory
	}
	return ""
}

func (x *NewPlatform) GetGeneralType() string {
	if x != nil {
		return x.GeneralType
	}
	return ""
}

func (x *NewPlatform) GetLethalityLevel() string {
	if x != nil {
		return x.LethalityLevel
	}
	return ""
}

func (x *NewPlatform) GetCruisingSpeed() float32 {
	if x != nil {
		return x.CruisingSpeed
	}
	return 0
}

func (x *NewPlatform) GetMaxSpeed() float32 {
	if x != nil {
		return x.MaxSpeed
	}
	return 0
}

func (x *NewPlatform) GetCombatRange() float32 {
	if x != nil {
		return x.CombatRange
	}
	return 0
}

func (x *NewPlatform) GetFuelLoad() float32 {
	if x != nil {
		return x.FuelLoad
	}
	return 0
}

func (x *NewPlatform) GetCountryOrigin() string {
	if x != nil {
		return x.CountryOrigin
	}
	return ""
}

type Platform struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UnitName        string  `protobuf:"bytes,1,opt,name=unitName,proto3" json:"unitName,omitempty"`
	UnitClass       string  `protobuf:"bytes,2,opt,name=unitClass,proto3" json:"unitClass,omitempty"`
	OperationField  string  `protobuf:"bytes,3,opt,name=operation_field,json=operationField,proto3" json:"operation_field,omitempty"`
	GeneralCategory string  `protobuf:"bytes,4,opt,name=generalCategory,proto3" json:"generalCategory,omitempty"`
	GeneralType     string  `protobuf:"bytes,5,opt,name=generalType,proto3" json:"generalType,omitempty"`
	LethalityLevel  string  `protobuf:"bytes,6,opt,name=lethalityLevel,proto3" json:"lethalityLevel,omitempty"`
	CruisingSpeed   float32 `protobuf:"fixed32,7,opt,name=cruisingSpeed,proto3" json:"cruisingSpeed,omitempty"`
	MaxSpeed        float32 `protobuf:"fixed32,8,opt,name=maxSpeed,proto3" json:"maxSpeed,omitempty"`
	CombatRange     float32 `protobuf:"fixed32,9,opt,name=combatRange,proto3" json:"combatRange,omitempty"`
	FuelLoad        float32 `protobuf:"fixed32,10,opt,name=fuelLoad,proto3" json:"fuelLoad,omitempty"`
	CountryOrigin   string  `protobuf:"bytes,11,opt,name=countryOrigin,proto3" json:"countryOrigin,omitempty"`
	Id              string  `protobuf:"bytes,12,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Platform) Reset() {
	*x = Platform{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Platform) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Platform) ProtoMessage() {}

func (x *Platform) ProtoReflect() protoreflect.Message {
	mi := &file_platform_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Platform.ProtoReflect.Descriptor instead.
func (*Platform) Descriptor() ([]byte, []int) {
	return file_platform_proto_rawDescGZIP(), []int{1}
}

func (x *Platform) GetUnitName() string {
	if x != nil {
		return x.UnitName
	}
	return ""
}

func (x *Platform) GetUnitClass() string {
	if x != nil {
		return x.UnitClass
	}
	return ""
}

func (x *Platform) GetOperationField() string {
	if x != nil {
		return x.OperationField
	}
	return ""
}

func (x *Platform) GetGeneralCategory() string {
	if x != nil {
		return x.GeneralCategory
	}
	return ""
}

func (x *Platform) GetGeneralType() string {
	if x != nil {
		return x.GeneralType
	}
	return ""
}

func (x *Platform) GetLethalityLevel() string {
	if x != nil {
		return x.LethalityLevel
	}
	return ""
}

func (x *Platform) GetCruisingSpeed() float32 {
	if x != nil {
		return x.CruisingSpeed
	}
	return 0
}

func (x *Platform) GetMaxSpeed() float32 {
	if x != nil {
		return x.MaxSpeed
	}
	return 0
}

func (x *Platform) GetCombatRange() float32 {
	if x != nil {
		return x.CombatRange
	}
	return 0
}

func (x *Platform) GetFuelLoad() float32 {
	if x != nil {
		return x.FuelLoad
	}
	return 0
}

func (x *Platform) GetCountryOrigin() string {
	if x != nil {
		return x.CountryOrigin
	}
	return ""
}

func (x *Platform) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type PlatformListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PlatformListRequest) Reset() {
	*x = PlatformListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlatformListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlatformListRequest) ProtoMessage() {}

func (x *PlatformListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_platform_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlatformListRequest.ProtoReflect.Descriptor instead.
func (*PlatformListRequest) Descriptor() ([]byte, []int) {
	return file_platform_proto_rawDescGZIP(), []int{2}
}

type PlatformListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Platforms []*Platform `protobuf:"bytes,1,rep,name=platforms,proto3" json:"platforms,omitempty"`
}

func (x *PlatformListResponse) Reset() {
	*x = PlatformListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlatformListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlatformListResponse) ProtoMessage() {}

func (x *PlatformListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_platform_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlatformListResponse.ProtoReflect.Descriptor instead.
func (*PlatformListResponse) Descriptor() ([]byte, []int) {
	return file_platform_proto_rawDescGZIP(), []int{3}
}

func (x *PlatformListResponse) GetPlatforms() []*Platform {
	if x != nil {
		return x.Platforms
	}
	return nil
}

var File_platform_proto protoreflect.FileDescriptor

var file_platform_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x03, 0x0a, 0x0b, 0x4e, 0x65, 0x77, 0x50,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x6e, 0x69, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x6e, 0x69, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x6e, 0x69, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x6e, 0x69, 0x74, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x6c, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x6c, 0x65, 0x74, 0x68, 0x61, 0x6c,
	0x69, 0x74, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x6c, 0x65, 0x74, 0x68, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x24,
	0x0a, 0x0d, 0x63, 0x72, 0x75, 0x69, 0x73, 0x69, 0x6e, 0x67, 0x53, 0x70, 0x65, 0x65, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x63, 0x72, 0x75, 0x69, 0x73, 0x69, 0x6e, 0x67, 0x53,
	0x70, 0x65, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x61, 0x78, 0x53, 0x70, 0x65, 0x65, 0x64,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x53, 0x70, 0x65, 0x65, 0x64,
	0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x65, 0x6c, 0x4c, 0x6f, 0x61, 0x64, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x66, 0x75, 0x65, 0x6c, 0x4c, 0x6f, 0x61, 0x64, 0x12, 0x24,
	0x0a, 0x0d, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x4f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x22, 0x97, 0x03, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x6e, 0x69, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x6e, 0x69, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x75, 0x6e, 0x69, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x75, 0x6e, 0x69, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x20,
	0x0a, 0x0b, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x26, 0x0a, 0x0e, 0x6c, 0x65, 0x74, 0x68, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x65, 0x74, 0x68, 0x61, 0x6c,
	0x69, 0x74, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x72, 0x75, 0x69,
	0x73, 0x69, 0x6e, 0x67, 0x53, 0x70, 0x65, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x0d, 0x63, 0x72, 0x75, 0x69, 0x73, 0x69, 0x6e, 0x67, 0x53, 0x70, 0x65, 0x65, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x6d, 0x61, 0x78, 0x53, 0x70, 0x65, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x08, 0x6d, 0x61, 0x78, 0x53, 0x70, 0x65, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f,
	0x6d, 0x62, 0x61, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x0b, 0x63, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x66, 0x75, 0x65, 0x6c, 0x4c, 0x6f, 0x61, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08,
	0x66, 0x75, 0x65, 0x6c, 0x4c, 0x6f, 0x61, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x15,
	0x0a, 0x13, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x45, 0x0a, 0x14, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a,
	0x09, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x32, 0x4d, 0x0a, 0x0f,
	0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x3a, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x65, 0x77,
	0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x70,
	0x62, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_platform_proto_rawDescOnce sync.Once
	file_platform_proto_rawDescData = file_platform_proto_rawDesc
)

func file_platform_proto_rawDescGZIP() []byte {
	file_platform_proto_rawDescOnce.Do(func() {
		file_platform_proto_rawDescData = protoimpl.X.CompressGZIP(file_platform_proto_rawDescData)
	})
	return file_platform_proto_rawDescData
}

var file_platform_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_platform_proto_goTypes = []interface{}{
	(*NewPlatform)(nil),          // 0: proto.NewPlatform
	(*Platform)(nil),             // 1: proto.Platform
	(*PlatformListRequest)(nil),  // 2: proto.PlatformListRequest
	(*PlatformListResponse)(nil), // 3: proto.PlatformListResponse
}
var file_platform_proto_depIdxs = []int32{
	1, // 0: proto.PlatformListResponse.platforms:type_name -> proto.Platform
	0, // 1: proto.PlatformService.CreateNewPlatform:input_type -> proto.NewPlatform
	1, // 2: proto.PlatformService.CreateNewPlatform:output_type -> proto.Platform
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_platform_proto_init() }
func file_platform_proto_init() {
	if File_platform_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_platform_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewPlatform); i {
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
		file_platform_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Platform); i {
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
		file_platform_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlatformListRequest); i {
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
		file_platform_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlatformListResponse); i {
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
			RawDescriptor: file_platform_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_platform_proto_goTypes,
		DependencyIndexes: file_platform_proto_depIdxs,
		MessageInfos:      file_platform_proto_msgTypes,
	}.Build()
	File_platform_proto = out.File
	file_platform_proto_rawDesc = nil
	file_platform_proto_goTypes = nil
	file_platform_proto_depIdxs = nil
}
