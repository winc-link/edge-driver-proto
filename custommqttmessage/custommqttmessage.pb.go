// Code generated by protoc-gen-go. DO NOT EDIT.
// source: custommqttmessage/custommqttmessage.proto

package custommqttmessage

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	drivercommon "github.com/winc-link/edge-driver-proto/drivercommon"
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

type PlatformCustomPublishRequest struct {
	BaseRequest          *drivercommon.BaseRequestMessage `protobuf:"bytes,1,opt,name=baseRequest,proto3" json:"baseRequest,omitempty"`
	DeviceId             string                           `protobuf:"bytes,2,opt,name=deviceId,proto3" json:"deviceId,omitempty"`
	Topic                string                           `protobuf:"bytes,3,opt,name=topic,proto3" json:"topic,omitempty"`
	Qos                  int32                            `protobuf:"varint,4,opt,name=qos,proto3" json:"qos,omitempty"`
	Retained             bool                             `protobuf:"varint,5,opt,name=retained,proto3" json:"retained,omitempty"`
	Payload              string                           `protobuf:"bytes,6,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *PlatformCustomPublishRequest) Reset()         { *m = PlatformCustomPublishRequest{} }
func (m *PlatformCustomPublishRequest) String() string { return proto.CompactTextString(m) }
func (*PlatformCustomPublishRequest) ProtoMessage()    {}
func (*PlatformCustomPublishRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_738b77d40ff46d3e, []int{0}
}

func (m *PlatformCustomPublishRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlatformCustomPublishRequest.Unmarshal(m, b)
}
func (m *PlatformCustomPublishRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlatformCustomPublishRequest.Marshal(b, m, deterministic)
}
func (m *PlatformCustomPublishRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlatformCustomPublishRequest.Merge(m, src)
}
func (m *PlatformCustomPublishRequest) XXX_Size() int {
	return xxx_messageInfo_PlatformCustomPublishRequest.Size(m)
}
func (m *PlatformCustomPublishRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PlatformCustomPublishRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PlatformCustomPublishRequest proto.InternalMessageInfo

func (m *PlatformCustomPublishRequest) GetBaseRequest() *drivercommon.BaseRequestMessage {
	if m != nil {
		return m.BaseRequest
	}
	return nil
}

func (m *PlatformCustomPublishRequest) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *PlatformCustomPublishRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *PlatformCustomPublishRequest) GetQos() int32 {
	if m != nil {
		return m.Qos
	}
	return 0
}

func (m *PlatformCustomPublishRequest) GetRetained() bool {
	if m != nil {
		return m.Retained
	}
	return false
}

func (m *PlatformCustomPublishRequest) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

type PlatformCustomSubscribeRequest struct {
	BaseRequest          *drivercommon.BaseRequestMessage `protobuf:"bytes,1,opt,name=baseRequest,proto3" json:"baseRequest,omitempty"`
	DeviceId             string                           `protobuf:"bytes,2,opt,name=deviceId,proto3" json:"deviceId,omitempty"`
	Topic                string                           `protobuf:"bytes,3,opt,name=topic,proto3" json:"topic,omitempty"`
	Qos                  int32                            `protobuf:"varint,4,opt,name=qos,proto3" json:"qos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *PlatformCustomSubscribeRequest) Reset()         { *m = PlatformCustomSubscribeRequest{} }
func (m *PlatformCustomSubscribeRequest) String() string { return proto.CompactTextString(m) }
func (*PlatformCustomSubscribeRequest) ProtoMessage()    {}
func (*PlatformCustomSubscribeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_738b77d40ff46d3e, []int{1}
}

func (m *PlatformCustomSubscribeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlatformCustomSubscribeRequest.Unmarshal(m, b)
}
func (m *PlatformCustomSubscribeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlatformCustomSubscribeRequest.Marshal(b, m, deterministic)
}
func (m *PlatformCustomSubscribeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlatformCustomSubscribeRequest.Merge(m, src)
}
func (m *PlatformCustomSubscribeRequest) XXX_Size() int {
	return xxx_messageInfo_PlatformCustomSubscribeRequest.Size(m)
}
func (m *PlatformCustomSubscribeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PlatformCustomSubscribeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PlatformCustomSubscribeRequest proto.InternalMessageInfo

func (m *PlatformCustomSubscribeRequest) GetBaseRequest() *drivercommon.BaseRequestMessage {
	if m != nil {
		return m.BaseRequest
	}
	return nil
}

func (m *PlatformCustomSubscribeRequest) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *PlatformCustomSubscribeRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *PlatformCustomSubscribeRequest) GetQos() int32 {
	if m != nil {
		return m.Qos
	}
	return 0
}

type PlatformCustomUnSubscribeRequest struct {
	BaseRequest          *drivercommon.BaseRequestMessage `protobuf:"bytes,1,opt,name=baseRequest,proto3" json:"baseRequest,omitempty"`
	DeviceId             string                           `protobuf:"bytes,2,opt,name=deviceId,proto3" json:"deviceId,omitempty"`
	Topics               []string                         `protobuf:"bytes,3,rep,name=topics,proto3" json:"topics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *PlatformCustomUnSubscribeRequest) Reset()         { *m = PlatformCustomUnSubscribeRequest{} }
func (m *PlatformCustomUnSubscribeRequest) String() string { return proto.CompactTextString(m) }
func (*PlatformCustomUnSubscribeRequest) ProtoMessage()    {}
func (*PlatformCustomUnSubscribeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_738b77d40ff46d3e, []int{2}
}

func (m *PlatformCustomUnSubscribeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlatformCustomUnSubscribeRequest.Unmarshal(m, b)
}
func (m *PlatformCustomUnSubscribeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlatformCustomUnSubscribeRequest.Marshal(b, m, deterministic)
}
func (m *PlatformCustomUnSubscribeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlatformCustomUnSubscribeRequest.Merge(m, src)
}
func (m *PlatformCustomUnSubscribeRequest) XXX_Size() int {
	return xxx_messageInfo_PlatformCustomUnSubscribeRequest.Size(m)
}
func (m *PlatformCustomUnSubscribeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PlatformCustomUnSubscribeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PlatformCustomUnSubscribeRequest proto.InternalMessageInfo

func (m *PlatformCustomUnSubscribeRequest) GetBaseRequest() *drivercommon.BaseRequestMessage {
	if m != nil {
		return m.BaseRequest
	}
	return nil
}

func (m *PlatformCustomUnSubscribeRequest) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *PlatformCustomUnSubscribeRequest) GetTopics() []string {
	if m != nil {
		return m.Topics
	}
	return nil
}

func init() {
	proto.RegisterType((*PlatformCustomPublishRequest)(nil), "custommqttmessage.PlatformCustomPublishRequest")
	proto.RegisterType((*PlatformCustomSubscribeRequest)(nil), "custommqttmessage.PlatformCustomSubscribeRequest")
	proto.RegisterType((*PlatformCustomUnSubscribeRequest)(nil), "custommqttmessage.PlatformCustomUnSubscribeRequest")
}

func init() {
	proto.RegisterFile("custommqttmessage/custommqttmessage.proto", fileDescriptor_738b77d40ff46d3e)
}

var fileDescriptor_738b77d40ff46d3e = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x53, 0xc1, 0x6e, 0xe2, 0x30,
	0x10, 0x5d, 0x93, 0x85, 0x05, 0x73, 0xd9, 0xb5, 0xd8, 0x5d, 0x83, 0xd0, 0x2a, 0xca, 0x29, 0x7b,
	0x20, 0xd1, 0xc2, 0x65, 0xd5, 0x23, 0x9c, 0x7a, 0x40, 0x42, 0xae, 0x7a, 0xe9, 0x2d, 0x71, 0xa6,
	0x60, 0x35, 0x89, 0x93, 0xd8, 0xa1, 0xea, 0xaf, 0xf4, 0x0b, 0xfa, 0x51, 0x3d, 0xf4, 0x53, 0x2a,
	0x12, 0xa0, 0xd0, 0x94, 0xe6, 0x56, 0xf5, 0xe4, 0xbc, 0x8c, 0xe7, 0xbd, 0x37, 0x9e, 0x19, 0xfc,
	0x97, 0xe7, 0x4a, 0xcb, 0x28, 0x4a, 0xb5, 0x8e, 0x40, 0x29, 0x6f, 0x09, 0x6e, 0xe5, 0x8f, 0x93,
	0x64, 0x52, 0x4b, 0xf2, 0xa3, 0x12, 0x18, 0xf4, 0x83, 0x4c, 0xac, 0x21, 0xe3, 0x32, 0x8a, 0x64,
	0xec, 0x96, 0x47, 0x79, 0xdb, 0x7a, 0x44, 0x78, 0xb8, 0x08, 0x3d, 0x7d, 0x2d, 0xb3, 0x68, 0x56,
	0x24, 0x2e, 0x72, 0x3f, 0x14, 0x6a, 0xc5, 0x20, 0xcd, 0x41, 0x69, 0x32, 0xc5, 0x5d, 0xdf, 0x53,
	0xb0, 0x85, 0x14, 0x99, 0xc8, 0xee, 0x8e, 0x4d, 0xe7, 0x90, 0xd1, 0x99, 0xbe, 0x5c, 0x98, 0x97,
	0x92, 0xec, 0x30, 0x89, 0x0c, 0x70, 0x3b, 0x80, 0xb5, 0xe0, 0x70, 0x1e, 0xd0, 0x86, 0x89, 0xec,
	0x0e, 0xdb, 0x63, 0xd2, 0xc3, 0x4d, 0x2d, 0x13, 0xc1, 0xa9, 0x51, 0x04, 0x4a, 0x40, 0xbe, 0x63,
	0x23, 0x95, 0x8a, 0x7e, 0x35, 0x91, 0xdd, 0x64, 0x9b, 0xcf, 0x0d, 0x47, 0x06, 0xda, 0x13, 0x31,
	0x04, 0xb4, 0x69, 0x22, 0xbb, 0xcd, 0xf6, 0x98, 0x50, 0xfc, 0x2d, 0xf1, 0xee, 0x42, 0xe9, 0x05,
	0xb4, 0x55, 0xb0, 0xec, 0xa0, 0xf5, 0x80, 0xf0, 0x9f, 0xe3, 0xf2, 0x2e, 0x72, 0x5f, 0xf1, 0x4c,
	0xf8, 0xf0, 0xc9, 0x0a, 0xb4, 0xee, 0x11, 0x36, 0x8f, 0xad, 0x5e, 0xc6, 0x1f, 0x6e, 0xf6, 0x17,
	0x6e, 0x15, 0xfe, 0x14, 0x35, 0x4c, 0xc3, 0xee, 0xb0, 0x2d, 0x1a, 0x3f, 0x35, 0x70, 0x8f, 0x25,
	0xbc, 0xf4, 0x35, 0x4f, 0xf5, 0x8e, 0x99, 0xac, 0xf0, 0xcf, 0x37, 0xc7, 0x87, 0xb8, 0x4e, 0x75,
	0x40, 0xdf, 0x1b, 0xb4, 0xc1, 0xf0, 0xb8, 0x8a, 0x59, 0x71, 0x30, 0x50, 0x89, 0x8c, 0x15, 0x58,
	0x5f, 0x48, 0x88, 0x7f, 0x9f, 0xe8, 0x24, 0xf9, 0x57, 0xab, 0xf5, 0xfa, 0x21, 0x6b, 0xd5, 0x12,
	0xdc, 0x3f, 0xd9, 0x0c, 0x32, 0xa9, 0xd5, 0xab, 0xb6, 0xae, 0x4e, 0x71, 0x7a, 0x76, 0xf5, 0x7f,
	0x29, 0xf4, 0x2a, 0xf7, 0x1d, 0x2e, 0x23, 0xf7, 0x56, 0xc4, 0x7c, 0x14, 0x8a, 0xf8, 0xc6, 0x85,
	0x60, 0x09, 0xa3, 0x32, 0x75, 0x54, 0xec, 0x6c, 0x75, 0xf3, 0xfd, 0x56, 0x11, 0x98, 0x3c, 0x07,
	0x00, 0x00, 0xff, 0xff, 0xac, 0x92, 0xef, 0x54, 0x27, 0x04, 0x00, 0x00,
}
