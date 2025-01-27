// Code generated by protoc-gen-go. DO NOT EDIT.
// source: drivercommon/common.proto

package drivercommon

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/emptypb"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type IotPlatform int32

const (
	IotPlatform_LocalIot    IotPlatform = 0
	IotPlatform_CustomerIot IotPlatform = 1
	IotPlatform_WinCLinkIot IotPlatform = 2
	IotPlatform_AliIot      IotPlatform = 3
	IotPlatform_HuaweiIot   IotPlatform = 4
	IotPlatform_TencentIot  IotPlatform = 5
	IotPlatform_TuyaIot     IotPlatform = 6
	IotPlatform_OneNetIot   IotPlatform = 7
)

var IotPlatform_name = map[int32]string{
	0: "LocalIot",
	1: "CustomerIot",
	2: "WinCLinkIot",
	3: "AliIot",
	4: "HuaweiIot",
	5: "TencentIot",
	6: "TuyaIot",
	7: "OneNetIot",
}

var IotPlatform_value = map[string]int32{
	"LocalIot":    0,
	"CustomerIot": 1,
	"WinCLinkIot": 2,
	"AliIot":      3,
	"HuaweiIot":   4,
	"TencentIot":  5,
	"TuyaIot":     6,
	"OneNetIot":   7,
}

func (x IotPlatform) String() string {
	return proto.EnumName(IotPlatform_name, int32(x))
}

func (IotPlatform) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c20f641e39ad35ce, []int{0}
}

type CloudInstanceInfo struct {
	CloudInstanceId      string      `protobuf:"bytes,2,opt,name=cloudInstanceId,proto3" json:"cloudInstanceId,omitempty"`
	BaseAddress          string      `protobuf:"bytes,3,opt,name=baseAddress,proto3" json:"baseAddress,omitempty"`
	CloudInstanceName    string      `protobuf:"bytes,4,opt,name=cloudInstanceName,proto3" json:"cloudInstanceName,omitempty"`
	IotPlatform          IotPlatform `protobuf:"varint,5,opt,name=iotPlatform,proto3,enum=drivercommon.IotPlatform" json:"iotPlatform,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CloudInstanceInfo) Reset()         { *m = CloudInstanceInfo{} }
func (m *CloudInstanceInfo) String() string { return proto.CompactTextString(m) }
func (*CloudInstanceInfo) ProtoMessage()    {}
func (*CloudInstanceInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_c20f641e39ad35ce, []int{0}
}

func (m *CloudInstanceInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloudInstanceInfo.Unmarshal(m, b)
}
func (m *CloudInstanceInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloudInstanceInfo.Marshal(b, m, deterministic)
}
func (m *CloudInstanceInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloudInstanceInfo.Merge(m, src)
}
func (m *CloudInstanceInfo) XXX_Size() int {
	return xxx_messageInfo_CloudInstanceInfo.Size(m)
}
func (m *CloudInstanceInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CloudInstanceInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CloudInstanceInfo proto.InternalMessageInfo

func (m *CloudInstanceInfo) GetCloudInstanceId() string {
	if m != nil {
		return m.CloudInstanceId
	}
	return ""
}

func (m *CloudInstanceInfo) GetBaseAddress() string {
	if m != nil {
		return m.BaseAddress
	}
	return ""
}

func (m *CloudInstanceInfo) GetCloudInstanceName() string {
	if m != nil {
		return m.CloudInstanceName
	}
	return ""
}

func (m *CloudInstanceInfo) GetIotPlatform() IotPlatform {
	if m != nil {
		return m.IotPlatform
	}
	return IotPlatform_LocalIot
}

type BaseRequestMessage struct {
	UseCloudPlatform     bool               `protobuf:"varint,1,opt,name=useCloudPlatform,proto3" json:"useCloudPlatform,omitempty"`
	CloudInstanceInfo    *CloudInstanceInfo `protobuf:"bytes,2,opt,name=cloudInstanceInfo,proto3" json:"cloudInstanceInfo,omitempty"`
	DriverInstanceId     string             `protobuf:"bytes,3,opt,name=driverInstanceId,proto3" json:"driverInstanceId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *BaseRequestMessage) Reset()         { *m = BaseRequestMessage{} }
func (m *BaseRequestMessage) String() string { return proto.CompactTextString(m) }
func (*BaseRequestMessage) ProtoMessage()    {}
func (*BaseRequestMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_c20f641e39ad35ce, []int{1}
}

func (m *BaseRequestMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseRequestMessage.Unmarshal(m, b)
}
func (m *BaseRequestMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseRequestMessage.Marshal(b, m, deterministic)
}
func (m *BaseRequestMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseRequestMessage.Merge(m, src)
}
func (m *BaseRequestMessage) XXX_Size() int {
	return xxx_messageInfo_BaseRequestMessage.Size(m)
}
func (m *BaseRequestMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseRequestMessage.DiscardUnknown(m)
}

var xxx_messageInfo_BaseRequestMessage proto.InternalMessageInfo

func (m *BaseRequestMessage) GetUseCloudPlatform() bool {
	if m != nil {
		return m.UseCloudPlatform
	}
	return false
}

func (m *BaseRequestMessage) GetCloudInstanceInfo() *CloudInstanceInfo {
	if m != nil {
		return m.CloudInstanceInfo
	}
	return nil
}

func (m *BaseRequestMessage) GetDriverInstanceId() string {
	if m != nil {
		return m.DriverInstanceId
	}
	return ""
}

type CommonResponse struct {
	RequestId            string   `protobuf:"bytes,1,opt,name=RequestId,proto3" json:"RequestId,omitempty"`
	ErrorMessage         string   `protobuf:"bytes,2,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	Code                 string   `protobuf:"bytes,3,opt,name=Code,proto3" json:"Code,omitempty"`
	Success              bool     `protobuf:"varint,4,opt,name=Success,proto3" json:"Success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommonResponse) Reset()         { *m = CommonResponse{} }
func (m *CommonResponse) String() string { return proto.CompactTextString(m) }
func (*CommonResponse) ProtoMessage()    {}
func (*CommonResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c20f641e39ad35ce, []int{2}
}

func (m *CommonResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonResponse.Unmarshal(m, b)
}
func (m *CommonResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonResponse.Marshal(b, m, deterministic)
}
func (m *CommonResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonResponse.Merge(m, src)
}
func (m *CommonResponse) XXX_Size() int {
	return xxx_messageInfo_CommonResponse.Size(m)
}
func (m *CommonResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CommonResponse proto.InternalMessageInfo

func (m *CommonResponse) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *CommonResponse) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func (m *CommonResponse) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *CommonResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type Pong struct {
	Timestamp            string   `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pong) Reset()         { *m = Pong{} }
func (m *Pong) String() string { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()    {}
func (*Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_c20f641e39ad35ce, []int{3}
}

func (m *Pong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pong.Unmarshal(m, b)
}
func (m *Pong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pong.Marshal(b, m, deterministic)
}
func (m *Pong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pong.Merge(m, src)
}
func (m *Pong) XXX_Size() int {
	return xxx_messageInfo_Pong.Size(m)
}
func (m *Pong) XXX_DiscardUnknown() {
	xxx_messageInfo_Pong.DiscardUnknown(m)
}

var xxx_messageInfo_Pong proto.InternalMessageInfo

func (m *Pong) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

type VersionResponse struct {
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionResponse) Reset()         { *m = VersionResponse{} }
func (m *VersionResponse) String() string { return proto.CompactTextString(m) }
func (*VersionResponse) ProtoMessage()    {}
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c20f641e39ad35ce, []int{4}
}

func (m *VersionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionResponse.Unmarshal(m, b)
}
func (m *VersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionResponse.Marshal(b, m, deterministic)
}
func (m *VersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionResponse.Merge(m, src)
}
func (m *VersionResponse) XXX_Size() int {
	return xxx_messageInfo_VersionResponse.Size(m)
}
func (m *VersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VersionResponse proto.InternalMessageInfo

func (m *VersionResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func init() {
	proto.RegisterEnum("drivercommon.IotPlatform", IotPlatform_name, IotPlatform_value)
	proto.RegisterType((*CloudInstanceInfo)(nil), "drivercommon.CloudInstanceInfo")
	proto.RegisterType((*BaseRequestMessage)(nil), "drivercommon.BaseRequestMessage")
	proto.RegisterType((*CommonResponse)(nil), "drivercommon.CommonResponse")
	proto.RegisterType((*Pong)(nil), "drivercommon.Pong")
	proto.RegisterType((*VersionResponse)(nil), "drivercommon.VersionResponse")
}

func init() { proto.RegisterFile("drivercommon/common.proto", fileDescriptor_c20f641e39ad35ce) }

var fileDescriptor_c20f641e39ad35ce = []byte{
	// 526 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xc1, 0x6e, 0xd3, 0x4c,
	0x10, 0xc7, 0xeb, 0xd6, 0x4d, 0x9a, 0x71, 0xbe, 0xa4, 0x9d, 0xc3, 0x27, 0xb7, 0x80, 0x88, 0x2c,
	0x0e, 0x51, 0x21, 0xb6, 0x94, 0xc2, 0x89, 0x53, 0x13, 0x21, 0x11, 0xa9, 0x2d, 0x91, 0xa9, 0x40,
	0xe2, 0xe6, 0xd8, 0x13, 0x63, 0xd5, 0xde, 0x0d, 0xde, 0x75, 0xab, 0xde, 0x10, 0xe2, 0xad, 0x78,
	0x03, 0x9e, 0x0a, 0xed, 0xda, 0x21, 0x76, 0x22, 0x4e, 0xf6, 0xfc, 0xf6, 0x3f, 0xb3, 0x33, 0xff,
	0xd1, 0xc2, 0x69, 0x94, 0x27, 0xf7, 0x94, 0x87, 0x3c, 0xcb, 0x38, 0xf3, 0xca, 0x8f, 0xbb, 0xca,
	0xb9, 0xe4, 0xd8, 0xad, 0x1f, 0x9d, 0x3d, 0x89, 0x39, 0x8f, 0x53, 0xf2, 0xf4, 0xd9, 0xa2, 0x58,
	0x7a, 0x94, 0xad, 0xe4, 0x63, 0x29, 0x75, 0x7e, 0x1b, 0x70, 0x32, 0x4d, 0x79, 0x11, 0xcd, 0x98,
	0x90, 0x01, 0x0b, 0x69, 0xc6, 0x96, 0x1c, 0x87, 0xd0, 0x0f, 0x1b, 0x30, 0xb2, 0xf7, 0x07, 0xc6,
	0xb0, 0xe3, 0x6f, 0x63, 0x1c, 0x80, 0xb5, 0x08, 0x04, 0x5d, 0x46, 0x51, 0x4e, 0x42, 0xd8, 0x07,
	0x5a, 0x55, 0x47, 0xf8, 0x0a, 0x4e, 0x1a, 0x49, 0x37, 0x41, 0x46, 0xb6, 0xa9, 0x75, 0xbb, 0x07,
	0xf8, 0x16, 0xac, 0x84, 0xcb, 0x79, 0x1a, 0xc8, 0x25, 0xcf, 0x33, 0xfb, 0x70, 0x60, 0x0c, 0x7b,
	0xe3, 0x53, 0xb7, 0x3e, 0x90, 0x3b, 0xdb, 0x08, 0xfc, 0xba, 0xda, 0xf9, 0x65, 0x00, 0x4e, 0x02,
	0x41, 0x3e, 0x7d, 0x2b, 0x48, 0xc8, 0x6b, 0x12, 0x22, 0x88, 0x09, 0xcf, 0xe1, 0xb8, 0x10, 0xa4,
	0xa7, 0xfc, 0x5b, 0xd8, 0x18, 0x18, 0xc3, 0x23, 0x7f, 0x87, 0xe3, 0xf5, 0x56, 0xb7, 0xca, 0x0e,
	0x3d, 0xbb, 0x35, 0x7e, 0xde, 0xec, 0x62, 0xc7, 0x35, 0x7f, 0x37, 0x53, 0x5d, 0x5d, 0x26, 0xd5,
	0x9c, 0x2c, 0x3d, 0xda, 0xe1, 0xce, 0x77, 0x03, 0x7a, 0x53, 0x5d, 0xdb, 0x27, 0xb1, 0xe2, 0x4c,
	0x10, 0x3e, 0x85, 0x4e, 0x35, 0xcb, 0x2c, 0xd2, 0x2d, 0x77, 0xfc, 0x0d, 0x40, 0x07, 0xba, 0x94,
	0xe7, 0x3c, 0xaf, 0xe6, 0xac, 0x56, 0xd4, 0x60, 0x88, 0x60, 0x4e, 0x79, 0x44, 0xd5, 0xa5, 0xfa,
	0x1f, 0x6d, 0x68, 0x7f, 0x2c, 0xc2, 0x50, 0xed, 0xcb, 0xd4, 0x36, 0xac, 0x43, 0xe7, 0x05, 0x98,
	0x73, 0xce, 0x62, 0x75, 0xaf, 0x4c, 0x32, 0x12, 0x32, 0xc8, 0x56, 0x55, 0xd9, 0x0d, 0x70, 0x5e,
	0x42, 0xff, 0x13, 0xe5, 0x22, 0xa9, 0x35, 0x6a, 0x43, 0xfb, 0xbe, 0x44, 0x95, 0x7c, 0x1d, 0x9e,
	0xff, 0x34, 0xc0, 0xaa, 0x2d, 0x0c, 0xbb, 0x70, 0x74, 0xc5, 0xc3, 0x20, 0x9d, 0x71, 0x79, 0xbc,
	0x87, 0x7d, 0xb0, 0xa6, 0x85, 0x90, 0x3c, 0xa3, 0x5c, 0x01, 0x43, 0x81, 0xcf, 0x09, 0x9b, 0x5e,
	0x25, 0xec, 0x4e, 0x81, 0x7d, 0x04, 0x68, 0x5d, 0xa6, 0x89, 0xfa, 0x3f, 0xc0, 0xff, 0xa0, 0xf3,
	0xbe, 0x08, 0x1e, 0x48, 0x87, 0x26, 0xf6, 0x00, 0x6e, 0x89, 0x85, 0xc4, 0xa4, 0x8a, 0x0f, 0xd1,
	0x82, 0xf6, 0x6d, 0xf1, 0x18, 0xa8, 0xa0, 0xa5, 0xb4, 0x1f, 0x18, 0xdd, 0x90, 0x3e, 0x6b, 0x8f,
	0x7f, 0x18, 0xd0, 0x2a, 0xcd, 0xc5, 0xd7, 0x60, 0xce, 0x13, 0x16, 0xe3, 0xff, 0x6e, 0xf9, 0x30,
	0xdc, 0xf5, 0xc3, 0x70, 0xdf, 0xa9, 0x87, 0x71, 0x86, 0xcd, 0x3d, 0x2b, 0x43, 0x9c, 0x3d, 0x9c,
	0x40, 0xbb, 0x1a, 0xfa, 0x9f, 0x89, 0xcf, 0x9a, 0x89, 0x5b, 0x1e, 0x39, 0x7b, 0x93, 0x37, 0x5f,
	0x2e, 0xe2, 0x44, 0x7e, 0x2d, 0x16, 0x6e, 0xc8, 0x33, 0xef, 0x21, 0x61, 0xe1, 0x28, 0x4d, 0xd8,
	0x9d, 0x47, 0x51, 0x4c, 0xa3, 0x32, 0x77, 0xa4, 0x6b, 0x7a, 0xf5, 0x42, 0x8b, 0x96, 0x66, 0x17,
	0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x65, 0x01, 0x17, 0x74, 0xf2, 0x03, 0x00, 0x00,
}
