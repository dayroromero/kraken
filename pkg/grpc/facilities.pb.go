// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: pkg/grpc/facilities.proto

package grpc

import (
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type CreateFacilityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status             bool            `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`
	DisableNote        string          `protobuf:"bytes,2,opt,name=DisableNote,proto3" json:"DisableNote,omitempty"`
	FacilityName       string          `protobuf:"bytes,3,opt,name=FacilityName,proto3" json:"FacilityName,omitempty"`
	IsPort             bool            `protobuf:"varint,4,opt,name=IsPort,proto3" json:"IsPort,omitempty"`
	PortId             string          `protobuf:"bytes,5,opt,name=PortId,proto3" json:"PortId,omitempty"`
	TypeOfTerminal     string          `protobuf:"bytes,6,opt,name=TypeOfTerminal,proto3" json:"TypeOfTerminal,omitempty"`
	ThirdPartyServices bool            `protobuf:"varint,7,opt,name=ThirdPartyServices,proto3" json:"ThirdPartyServices,omitempty"`
	FacilityAddress    *_struct.Struct `protobuf:"bytes,8,opt,name=FacilityAddress,proto3" json:"FacilityAddress,omitempty"`
	FacilityCoordinate *_struct.Struct `protobuf:"bytes,9,opt,name=FacilityCoordinate,proto3" json:"FacilityCoordinate,omitempty"`
	EntityId           string          `protobuf:"bytes,10,opt,name=EntityId,proto3" json:"EntityId,omitempty"`
	FacilityManager    *_struct.Struct `protobuf:"bytes,11,opt,name=FacilityManager,proto3" json:"FacilityManager,omitempty"`
}

func (x *CreateFacilityRequest) Reset() {
	*x = CreateFacilityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_facilities_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFacilityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFacilityRequest) ProtoMessage() {}

func (x *CreateFacilityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_facilities_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFacilityRequest.ProtoReflect.Descriptor instead.
func (*CreateFacilityRequest) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_facilities_proto_rawDescGZIP(), []int{0}
}

func (x *CreateFacilityRequest) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *CreateFacilityRequest) GetDisableNote() string {
	if x != nil {
		return x.DisableNote
	}
	return ""
}

func (x *CreateFacilityRequest) GetFacilityName() string {
	if x != nil {
		return x.FacilityName
	}
	return ""
}

func (x *CreateFacilityRequest) GetIsPort() bool {
	if x != nil {
		return x.IsPort
	}
	return false
}

func (x *CreateFacilityRequest) GetPortId() string {
	if x != nil {
		return x.PortId
	}
	return ""
}

func (x *CreateFacilityRequest) GetTypeOfTerminal() string {
	if x != nil {
		return x.TypeOfTerminal
	}
	return ""
}

func (x *CreateFacilityRequest) GetThirdPartyServices() bool {
	if x != nil {
		return x.ThirdPartyServices
	}
	return false
}

func (x *CreateFacilityRequest) GetFacilityAddress() *_struct.Struct {
	if x != nil {
		return x.FacilityAddress
	}
	return nil
}

func (x *CreateFacilityRequest) GetFacilityCoordinate() *_struct.Struct {
	if x != nil {
		return x.FacilityCoordinate
	}
	return nil
}

func (x *CreateFacilityRequest) GetEntityId() string {
	if x != nil {
		return x.EntityId
	}
	return ""
}

func (x *CreateFacilityRequest) GetFacilityManager() *_struct.Struct {
	if x != nil {
		return x.FacilityManager
	}
	return nil
}

type CreateFacilityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *CreateFacilityResponse) Reset() {
	*x = CreateFacilityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_facilities_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFacilityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFacilityResponse) ProtoMessage() {}

func (x *CreateFacilityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_facilities_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFacilityResponse.ProtoReflect.Descriptor instead.
func (*CreateFacilityResponse) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_facilities_proto_rawDescGZIP(), []int{1}
}

func (x *CreateFacilityResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CreateFacilityResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type UpdateFacilityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FacilityId         string          `protobuf:"bytes,1,opt,name=FacilityId,proto3" json:"FacilityId,omitempty"`
	Status             bool            `protobuf:"varint,2,opt,name=Status,proto3" json:"Status,omitempty"`
	DisableNote        string          `protobuf:"bytes,3,opt,name=DisableNote,proto3" json:"DisableNote,omitempty"`
	FacilityName       string          `protobuf:"bytes,4,opt,name=FacilityName,proto3" json:"FacilityName,omitempty"`
	IsPort             bool            `protobuf:"varint,5,opt,name=IsPort,proto3" json:"IsPort,omitempty"`
	PortId             string          `protobuf:"bytes,6,opt,name=PortId,proto3" json:"PortId,omitempty"`
	TypeOfTerminal     string          `protobuf:"bytes,7,opt,name=TypeOfTerminal,proto3" json:"TypeOfTerminal,omitempty"`
	ThirdPartyServices bool            `protobuf:"varint,8,opt,name=ThirdPartyServices,proto3" json:"ThirdPartyServices,omitempty"`
	FacilityAddress    *_struct.Struct `protobuf:"bytes,9,opt,name=FacilityAddress,proto3" json:"FacilityAddress,omitempty"`
	FacilityCoordinate *_struct.Struct `protobuf:"bytes,10,opt,name=FacilityCoordinate,proto3" json:"FacilityCoordinate,omitempty"`
	EntityId           string          `protobuf:"bytes,11,opt,name=EntityId,proto3" json:"EntityId,omitempty"`
	FacilityManager    *_struct.Struct `protobuf:"bytes,12,opt,name=FacilityManager,proto3" json:"FacilityManager,omitempty"`
}

func (x *UpdateFacilityRequest) Reset() {
	*x = UpdateFacilityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_facilities_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFacilityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFacilityRequest) ProtoMessage() {}

func (x *UpdateFacilityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_facilities_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFacilityRequest.ProtoReflect.Descriptor instead.
func (*UpdateFacilityRequest) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_facilities_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateFacilityRequest) GetFacilityId() string {
	if x != nil {
		return x.FacilityId
	}
	return ""
}

func (x *UpdateFacilityRequest) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *UpdateFacilityRequest) GetDisableNote() string {
	if x != nil {
		return x.DisableNote
	}
	return ""
}

func (x *UpdateFacilityRequest) GetFacilityName() string {
	if x != nil {
		return x.FacilityName
	}
	return ""
}

func (x *UpdateFacilityRequest) GetIsPort() bool {
	if x != nil {
		return x.IsPort
	}
	return false
}

func (x *UpdateFacilityRequest) GetPortId() string {
	if x != nil {
		return x.PortId
	}
	return ""
}

func (x *UpdateFacilityRequest) GetTypeOfTerminal() string {
	if x != nil {
		return x.TypeOfTerminal
	}
	return ""
}

func (x *UpdateFacilityRequest) GetThirdPartyServices() bool {
	if x != nil {
		return x.ThirdPartyServices
	}
	return false
}

func (x *UpdateFacilityRequest) GetFacilityAddress() *_struct.Struct {
	if x != nil {
		return x.FacilityAddress
	}
	return nil
}

func (x *UpdateFacilityRequest) GetFacilityCoordinate() *_struct.Struct {
	if x != nil {
		return x.FacilityCoordinate
	}
	return nil
}

func (x *UpdateFacilityRequest) GetEntityId() string {
	if x != nil {
		return x.EntityId
	}
	return ""
}

func (x *UpdateFacilityRequest) GetFacilityManager() *_struct.Struct {
	if x != nil {
		return x.FacilityManager
	}
	return nil
}

var File_pkg_grpc_facilities_proto protoreflect.FileDescriptor

var file_pkg_grpc_facilities_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x66, 0x61, 0x63, 0x69, 0x6c,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70,
	0x63, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xe8, 0x03, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x61, 0x63, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x4e,
	0x6f, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x46, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x46, 0x61, 0x63, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x73, 0x50, 0x6f, 0x72,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x49, 0x73, 0x50, 0x6f, 0x72, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x50, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x50, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x54, 0x79, 0x70, 0x65, 0x4f,
	0x66, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x54, 0x79, 0x70, 0x65, 0x4f, 0x66, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x12,
	0x2e, 0x0a, 0x12, 0x54, 0x68, 0x69, 0x72, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x54, 0x68, 0x69,
	0x72, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12,
	0x41, 0x0a, 0x0f, 0x46, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x52, 0x0f, 0x46, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x47, 0x0a, 0x12, 0x46, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x43, 0x6f,
	0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x12, 0x46, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x41, 0x0a, 0x0f, 0x46, 0x61, 0x63, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0f, 0x46, 0x61, 0x63, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x22, 0x46, 0x0a, 0x16, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x46, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x22, 0x88, 0x04, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x61, 0x63,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a,
	0x46, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x46, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x4e,
	0x6f, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x69, 0x73, 0x61, 0x62,
	0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x46, 0x61, 0x63, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x46, 0x61,
	0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x73,
	0x50, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x49, 0x73, 0x50, 0x6f,
	0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x50, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x54, 0x79,
	0x70, 0x65, 0x4f, 0x66, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x54, 0x79, 0x70, 0x65, 0x4f, 0x66, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e,
	0x61, 0x6c, 0x12, 0x2e, 0x0a, 0x12, 0x54, 0x68, 0x69, 0x72, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12,
	0x54, 0x68, 0x69, 0x72, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x12, 0x41, 0x0a, 0x0f, 0x46, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x52, 0x0f, 0x46, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x47, 0x0a, 0x12, 0x46, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x12, 0x46, 0x61, 0x63, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x41, 0x0a, 0x0f, 0x46, 0x61,
	0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0f, 0x46, 0x61,
	0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x32, 0xb1, 0x01,
	0x0a, 0x11, 0x46, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x61, 0x63,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x1b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x46, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x46, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x61, 0x63, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x12, 0x1b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x46, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46,
	0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_grpc_facilities_proto_rawDescOnce sync.Once
	file_pkg_grpc_facilities_proto_rawDescData = file_pkg_grpc_facilities_proto_rawDesc
)

func file_pkg_grpc_facilities_proto_rawDescGZIP() []byte {
	file_pkg_grpc_facilities_proto_rawDescOnce.Do(func() {
		file_pkg_grpc_facilities_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_grpc_facilities_proto_rawDescData)
	})
	return file_pkg_grpc_facilities_proto_rawDescData
}

var file_pkg_grpc_facilities_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pkg_grpc_facilities_proto_goTypes = []interface{}{
	(*CreateFacilityRequest)(nil),  // 0: grpc.CreateFacilityRequest
	(*CreateFacilityResponse)(nil), // 1: grpc.CreateFacilityResponse
	(*UpdateFacilityRequest)(nil),  // 2: grpc.UpdateFacilityRequest
	(*_struct.Struct)(nil),         // 3: google.protobuf.Struct
}
var file_pkg_grpc_facilities_proto_depIdxs = []int32{
	3, // 0: grpc.CreateFacilityRequest.FacilityAddress:type_name -> google.protobuf.Struct
	3, // 1: grpc.CreateFacilityRequest.FacilityCoordinate:type_name -> google.protobuf.Struct
	3, // 2: grpc.CreateFacilityRequest.FacilityManager:type_name -> google.protobuf.Struct
	3, // 3: grpc.UpdateFacilityRequest.FacilityAddress:type_name -> google.protobuf.Struct
	3, // 4: grpc.UpdateFacilityRequest.FacilityCoordinate:type_name -> google.protobuf.Struct
	3, // 5: grpc.UpdateFacilityRequest.FacilityManager:type_name -> google.protobuf.Struct
	0, // 6: grpc.FacilitiesService.CreateFacility:input_type -> grpc.CreateFacilityRequest
	2, // 7: grpc.FacilitiesService.UpdateFacility:input_type -> grpc.UpdateFacilityRequest
	1, // 8: grpc.FacilitiesService.CreateFacility:output_type -> grpc.CreateFacilityResponse
	1, // 9: grpc.FacilitiesService.UpdateFacility:output_type -> grpc.CreateFacilityResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_pkg_grpc_facilities_proto_init() }
func file_pkg_grpc_facilities_proto_init() {
	if File_pkg_grpc_facilities_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_grpc_facilities_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFacilityRequest); i {
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
		file_pkg_grpc_facilities_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFacilityResponse); i {
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
		file_pkg_grpc_facilities_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFacilityRequest); i {
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
			RawDescriptor: file_pkg_grpc_facilities_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_grpc_facilities_proto_goTypes,
		DependencyIndexes: file_pkg_grpc_facilities_proto_depIdxs,
		MessageInfos:      file_pkg_grpc_facilities_proto_msgTypes,
	}.Build()
	File_pkg_grpc_facilities_proto = out.File
	file_pkg_grpc_facilities_proto_rawDesc = nil
	file_pkg_grpc_facilities_proto_goTypes = nil
	file_pkg_grpc_facilities_proto_depIdxs = nil
}
