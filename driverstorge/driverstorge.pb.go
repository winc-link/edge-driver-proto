// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: driverstorge/driverstorge.proto

package driverstorage

import (
	drivercommon "github.com/winc-link/edge-driver-proto/drivercommon"
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

type PutReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DriverServiceId string `protobuf:"bytes,1,opt,name=driver_service_id,json=driverServiceId,proto3" json:"driver_service_id,omitempty"`
	Data            []*KV  `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *PutReq) Reset() {
	*x = PutReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_driverstorge_driverstorge_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutReq) ProtoMessage() {}

func (x *PutReq) ProtoReflect() protoreflect.Message {
	mi := &file_driverstorge_driverstorge_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutReq.ProtoReflect.Descriptor instead.
func (*PutReq) Descriptor() ([]byte, []int) {
	return file_driverstorge_driverstorge_proto_rawDescGZIP(), []int{0}
}

func (x *PutReq) GetDriverServiceId() string {
	if x != nil {
		return x.DriverServiceId
	}
	return ""
}

func (x *PutReq) GetData() []*KV {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DriverServiceId string   `protobuf:"bytes,1,opt,name=driver_service_id,json=driverServiceId,proto3" json:"driver_service_id,omitempty"`
	Keys            []string `protobuf:"bytes,2,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *GetReq) Reset() {
	*x = GetReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_driverstorge_driverstorge_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReq) ProtoMessage() {}

func (x *GetReq) ProtoReflect() protoreflect.Message {
	mi := &file_driverstorge_driverstorge_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReq.ProtoReflect.Descriptor instead.
func (*GetReq) Descriptor() ([]byte, []int) {
	return file_driverstorge_driverstorge_proto_rawDescGZIP(), []int{1}
}

func (x *GetReq) GetDriverServiceId() string {
	if x != nil {
		return x.DriverServiceId
	}
	return ""
}

func (x *GetReq) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

type AllReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DriverServiceId string `protobuf:"bytes,1,opt,name=driver_service_id,json=driverServiceId,proto3" json:"driver_service_id,omitempty"`
}

func (x *AllReq) Reset() {
	*x = AllReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_driverstorge_driverstorge_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllReq) ProtoMessage() {}

func (x *AllReq) ProtoReflect() protoreflect.Message {
	mi := &file_driverstorge_driverstorge_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllReq.ProtoReflect.Descriptor instead.
func (*AllReq) Descriptor() ([]byte, []int) {
	return file_driverstorge_driverstorge_proto_rawDescGZIP(), []int{2}
}

func (x *AllReq) GetDriverServiceId() string {
	if x != nil {
		return x.DriverServiceId
	}
	return ""
}

type DeleteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DriverServiceId string   `protobuf:"bytes,1,opt,name=driver_service_id,json=driverServiceId,proto3" json:"driver_service_id,omitempty"`
	Keys            []string `protobuf:"bytes,2,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *DeleteReq) Reset() {
	*x = DeleteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_driverstorge_driverstorge_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteReq) ProtoMessage() {}

func (x *DeleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_driverstorge_driverstorge_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteReq.ProtoReflect.Descriptor instead.
func (*DeleteReq) Descriptor() ([]byte, []int) {
	return file_driverstorge_driverstorge_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteReq) GetDriverServiceId() string {
	if x != nil {
		return x.DriverServiceId
	}
	return ""
}

func (x *DeleteReq) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

type KVs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResponse *drivercommon.CommonResponse `protobuf:"bytes,1,opt,name=baseResponse,proto3" json:"baseResponse,omitempty"`
	Kvs          []*KV                        `protobuf:"bytes,2,rep,name=kvs,proto3" json:"kvs,omitempty"`
}

func (x *KVs) Reset() {
	*x = KVs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_driverstorge_driverstorge_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KVs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KVs) ProtoMessage() {}

func (x *KVs) ProtoReflect() protoreflect.Message {
	mi := &file_driverstorge_driverstorge_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KVs.ProtoReflect.Descriptor instead.
func (*KVs) Descriptor() ([]byte, []int) {
	return file_driverstorge_driverstorge_proto_rawDescGZIP(), []int{4}
}

func (x *KVs) GetBaseResponse() *drivercommon.CommonResponse {
	if x != nil {
		return x.BaseResponse
	}
	return nil
}

func (x *KVs) GetKvs() []*KV {
	if x != nil {
		return x.Kvs
	}
	return nil
}

type KV struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResponse *drivercommon.CommonResponse `protobuf:"bytes,1,opt,name=baseResponse,proto3" json:"baseResponse,omitempty"`
	Key          string                       `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value        []byte                       `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *KV) Reset() {
	*x = KV{}
	if protoimpl.UnsafeEnabled {
		mi := &file_driverstorge_driverstorge_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KV) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KV) ProtoMessage() {}

func (x *KV) ProtoReflect() protoreflect.Message {
	mi := &file_driverstorge_driverstorge_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KV.ProtoReflect.Descriptor instead.
func (*KV) Descriptor() ([]byte, []int) {
	return file_driverstorge_driverstorge_proto_rawDescGZIP(), []int{5}
}

func (x *KV) GetBaseResponse() *drivercommon.CommonResponse {
	if x != nil {
		return x.BaseResponse
	}
	return nil
}

func (x *KV) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KV) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_driverstorge_driverstorge_proto protoreflect.FileDescriptor

var file_driverstorge_driverstorge_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x74, 0x6f, 0x72, 0x67, 0x65, 0x2f, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x74, 0x6f, 0x72, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0d, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x1a, 0x19, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b, 0x0a, 0x06, 0x50,
	0x75, 0x74, 0x52, 0x65, 0x71, 0x12, 0x2a, 0x0a, 0x11, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49,
	0x64, 0x12, 0x25, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e,
	0x4b, 0x56, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x48, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x2a, 0x0a, 0x11, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x65,
	0x79, 0x73, 0x22, 0x34, 0x0a, 0x06, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x2a, 0x0a, 0x11,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x22, 0x4b, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x2a, 0x0a, 0x11, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x04, 0x6b, 0x65, 0x79, 0x73, 0x22, 0x6c, 0x0a, 0x03, 0x4b, 0x56, 0x73, 0x12, 0x40, 0x0a, 0x0c,
	0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23,
	0x0a, 0x03, 0x6b, 0x76, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x64, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x4b, 0x56, 0x52, 0x03,
	0x6b, 0x76, 0x73, 0x22, 0x6e, 0x0a, 0x02, 0x4b, 0x56, 0x12, 0x40, 0x0a, 0x0c, 0x62, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0c, 0x62,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x32, 0xf1, 0x01, 0x0a, 0x0d, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x30, 0x0a, 0x03, 0x41, 0x6c, 0x6c, 0x12, 0x15, 0x2e, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x6c, 0x6c,
	0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2e, 0x4b, 0x56, 0x73, 0x12, 0x30, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x15,
	0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x4b, 0x56, 0x73, 0x12, 0x3a, 0x0a, 0x03, 0x50, 0x75, 0x74,
	0x12, 0x15, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x50, 0x75, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x18, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x64, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x69, 0x6e, 0x63, 0x2d, 0x6c, 0x69, 0x6e, 0x6b, 0x2f,
	0x65, 0x64, 0x67, 0x65, 0x2d, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_driverstorge_driverstorge_proto_rawDescOnce sync.Once
	file_driverstorge_driverstorge_proto_rawDescData = file_driverstorge_driverstorge_proto_rawDesc
)

func file_driverstorge_driverstorge_proto_rawDescGZIP() []byte {
	file_driverstorge_driverstorge_proto_rawDescOnce.Do(func() {
		file_driverstorge_driverstorge_proto_rawDescData = protoimpl.X.CompressGZIP(file_driverstorge_driverstorge_proto_rawDescData)
	})
	return file_driverstorge_driverstorge_proto_rawDescData
}

var file_driverstorge_driverstorge_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_driverstorge_driverstorge_proto_goTypes = []interface{}{
	(*PutReq)(nil),                      // 0: driverstorage.PutReq
	(*GetReq)(nil),                      // 1: driverstorage.GetReq
	(*AllReq)(nil),                      // 2: driverstorage.AllReq
	(*DeleteReq)(nil),                   // 3: driverstorage.DeleteReq
	(*KVs)(nil),                         // 4: driverstorage.KVs
	(*KV)(nil),                          // 5: driverstorage.KV
	(*drivercommon.CommonResponse)(nil), // 6: drivercommon.CommonResponse
}
var file_driverstorge_driverstorge_proto_depIdxs = []int32{
	5, // 0: driverstorage.PutReq.data:type_name -> driverstorage.KV
	6, // 1: driverstorage.KVs.baseResponse:type_name -> drivercommon.CommonResponse
	5, // 2: driverstorage.KVs.kvs:type_name -> driverstorage.KV
	6, // 3: driverstorage.KV.baseResponse:type_name -> drivercommon.CommonResponse
	2, // 4: driverstorage.DriverStorage.All:input_type -> driverstorage.AllReq
	1, // 5: driverstorage.DriverStorage.Get:input_type -> driverstorage.GetReq
	0, // 6: driverstorage.DriverStorage.Put:input_type -> driverstorage.PutReq
	3, // 7: driverstorage.DriverStorage.Delete:input_type -> driverstorage.DeleteReq
	4, // 8: driverstorage.DriverStorage.All:output_type -> driverstorage.KVs
	4, // 9: driverstorage.DriverStorage.Get:output_type -> driverstorage.KVs
	6, // 10: driverstorage.DriverStorage.Put:output_type -> drivercommon.CommonResponse
	6, // 11: driverstorage.DriverStorage.Delete:output_type -> drivercommon.CommonResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_driverstorge_driverstorge_proto_init() }
func file_driverstorge_driverstorge_proto_init() {
	if File_driverstorge_driverstorge_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_driverstorge_driverstorge_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutReq); i {
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
		file_driverstorge_driverstorge_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReq); i {
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
		file_driverstorge_driverstorge_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllReq); i {
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
		file_driverstorge_driverstorge_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteReq); i {
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
		file_driverstorge_driverstorge_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KVs); i {
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
		file_driverstorge_driverstorge_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KV); i {
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
			RawDescriptor: file_driverstorge_driverstorge_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_driverstorge_driverstorge_proto_goTypes,
		DependencyIndexes: file_driverstorge_driverstorge_proto_depIdxs,
		MessageInfos:      file_driverstorge_driverstorge_proto_msgTypes,
	}.Build()
	File_driverstorge_driverstorge_proto = out.File
	file_driverstorge_driverstorge_proto_rawDesc = nil
	file_driverstorge_driverstorge_proto_goTypes = nil
	file_driverstorge_driverstorge_proto_depIdxs = nil
}
