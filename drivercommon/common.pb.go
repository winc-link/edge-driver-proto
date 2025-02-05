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

type MetaBasesType int32

const (
	MetaBasesType_Mysql  MetaBasesType = 0
	MetaBasesType_Sqlite MetaBasesType = 1
)

var MetaBasesType_name = map[int32]string{
	0: "Mysql",
	1: "Sqlite",
}

var MetaBasesType_value = map[string]int32{
	"Mysql":  0,
	"Sqlite": 1,
}

func (x MetaBasesType) String() string {
	return proto.EnumName(MetaBasesType_name, int32(x))
}

func (MetaBasesType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c20f641e39ad35ce, []int{1}
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

type ConfigResponse struct {
	MetaBases            *MetaBases    `protobuf:"bytes,1,opt,name=metaBases,proto3" json:"metaBases,omitempty"`
	DataBases            *DataBases    `protobuf:"bytes,2,opt,name=dataBases,proto3" json:"dataBases,omitempty"`
	MessageQueue         *MessageQueue `protobuf:"bytes,3,opt,name=messageQueue,proto3" json:"messageQueue,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ConfigResponse) Reset()         { *m = ConfigResponse{} }
func (m *ConfigResponse) String() string { return proto.CompactTextString(m) }
func (*ConfigResponse) ProtoMessage()    {}
func (*ConfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c20f641e39ad35ce, []int{5}
}

func (m *ConfigResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigResponse.Unmarshal(m, b)
}
func (m *ConfigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigResponse.Marshal(b, m, deterministic)
}
func (m *ConfigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigResponse.Merge(m, src)
}
func (m *ConfigResponse) XXX_Size() int {
	return xxx_messageInfo_ConfigResponse.Size(m)
}
func (m *ConfigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigResponse proto.InternalMessageInfo

func (m *ConfigResponse) GetMetaBases() *MetaBases {
	if m != nil {
		return m.MetaBases
	}
	return nil
}

func (m *ConfigResponse) GetDataBases() *DataBases {
	if m != nil {
		return m.DataBases
	}
	return nil
}

func (m *ConfigResponse) GetMessageQueue() *MessageQueue {
	if m != nil {
		return m.MessageQueue
	}
	return nil
}

type MetaBases struct {
	Type                 MetaBasesType `protobuf:"varint,1,opt,name=Type,proto3,enum=drivercommon.MetaBasesType" json:"Type,omitempty"`
	Source               string        `protobuf:"bytes,2,opt,name=Source,proto3" json:"Source,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *MetaBases) Reset()         { *m = MetaBases{} }
func (m *MetaBases) String() string { return proto.CompactTextString(m) }
func (*MetaBases) ProtoMessage()    {}
func (*MetaBases) Descriptor() ([]byte, []int) {
	return fileDescriptor_c20f641e39ad35ce, []int{6}
}

func (m *MetaBases) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetaBases.Unmarshal(m, b)
}
func (m *MetaBases) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetaBases.Marshal(b, m, deterministic)
}
func (m *MetaBases) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetaBases.Merge(m, src)
}
func (m *MetaBases) XXX_Size() int {
	return xxx_messageInfo_MetaBases.Size(m)
}
func (m *MetaBases) XXX_DiscardUnknown() {
	xxx_messageInfo_MetaBases.DiscardUnknown(m)
}

var xxx_messageInfo_MetaBases proto.InternalMessageInfo

func (m *MetaBases) GetType() MetaBasesType {
	if m != nil {
		return m.Type
	}
	return MetaBasesType_Mysql
}

func (m *MetaBases) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

type DataBases struct {
	Type                 string          `protobuf:"bytes,1,opt,name=Type,proto3" json:"Type,omitempty"`
	Tdengine             *TDengineSource `protobuf:"bytes,2,opt,name=tdengine,proto3" json:"tdengine,omitempty"`
	InfluxDB             *InfluxDB       `protobuf:"bytes,3,opt,name=influxDB,proto3" json:"influxDB,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *DataBases) Reset()         { *m = DataBases{} }
func (m *DataBases) String() string { return proto.CompactTextString(m) }
func (*DataBases) ProtoMessage()    {}
func (*DataBases) Descriptor() ([]byte, []int) {
	return fileDescriptor_c20f641e39ad35ce, []int{7}
}

func (m *DataBases) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataBases.Unmarshal(m, b)
}
func (m *DataBases) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataBases.Marshal(b, m, deterministic)
}
func (m *DataBases) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataBases.Merge(m, src)
}
func (m *DataBases) XXX_Size() int {
	return xxx_messageInfo_DataBases.Size(m)
}
func (m *DataBases) XXX_DiscardUnknown() {
	xxx_messageInfo_DataBases.DiscardUnknown(m)
}

var xxx_messageInfo_DataBases proto.InternalMessageInfo

func (m *DataBases) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *DataBases) GetTdengine() *TDengineSource {
	if m != nil {
		return m.Tdengine
	}
	return nil
}

func (m *DataBases) GetInfluxDB() *InfluxDB {
	if m != nil {
		return m.InfluxDB
	}
	return nil
}

type TDengineSource struct {
	Dsn                  string   `protobuf:"bytes,1,opt,name=Dsn,proto3" json:"Dsn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TDengineSource) Reset()         { *m = TDengineSource{} }
func (m *TDengineSource) String() string { return proto.CompactTextString(m) }
func (*TDengineSource) ProtoMessage()    {}
func (*TDengineSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_c20f641e39ad35ce, []int{8}
}

func (m *TDengineSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TDengineSource.Unmarshal(m, b)
}
func (m *TDengineSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TDengineSource.Marshal(b, m, deterministic)
}
func (m *TDengineSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TDengineSource.Merge(m, src)
}
func (m *TDengineSource) XXX_Size() int {
	return xxx_messageInfo_TDengineSource.Size(m)
}
func (m *TDengineSource) XXX_DiscardUnknown() {
	xxx_messageInfo_TDengineSource.DiscardUnknown(m)
}

var xxx_messageInfo_TDengineSource proto.InternalMessageInfo

func (m *TDengineSource) GetDsn() string {
	if m != nil {
		return m.Dsn
	}
	return ""
}

type InfluxDB struct {
	Org                  string   `protobuf:"bytes,1,opt,name=org,proto3" json:"org,omitempty"`
	Bucket               string   `protobuf:"bytes,2,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Url                  string   `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	Token                string   `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InfluxDB) Reset()         { *m = InfluxDB{} }
func (m *InfluxDB) String() string { return proto.CompactTextString(m) }
func (*InfluxDB) ProtoMessage()    {}
func (*InfluxDB) Descriptor() ([]byte, []int) {
	return fileDescriptor_c20f641e39ad35ce, []int{9}
}

func (m *InfluxDB) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InfluxDB.Unmarshal(m, b)
}
func (m *InfluxDB) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InfluxDB.Marshal(b, m, deterministic)
}
func (m *InfluxDB) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InfluxDB.Merge(m, src)
}
func (m *InfluxDB) XXX_Size() int {
	return xxx_messageInfo_InfluxDB.Size(m)
}
func (m *InfluxDB) XXX_DiscardUnknown() {
	xxx_messageInfo_InfluxDB.DiscardUnknown(m)
}

var xxx_messageInfo_InfluxDB proto.InternalMessageInfo

func (m *InfluxDB) GetOrg() string {
	if m != nil {
		return m.Org
	}
	return ""
}

func (m *InfluxDB) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *InfluxDB) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *InfluxDB) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type MessageQueue struct {
	Protocol             string   `protobuf:"bytes,1,opt,name=Protocol,proto3" json:"Protocol,omitempty"`
	Host                 string   `protobuf:"bytes,2,opt,name=Host,proto3" json:"Host,omitempty"`
	Port                 uint32   `protobuf:"varint,3,opt,name=Port,proto3" json:"Port,omitempty"`
	Type                 string   `protobuf:"bytes,4,opt,name=Type,proto3" json:"Type,omitempty"`
	PublishTopicPrefix   string   `protobuf:"bytes,5,opt,name=PublishTopicPrefix,proto3" json:"PublishTopicPrefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageQueue) Reset()         { *m = MessageQueue{} }
func (m *MessageQueue) String() string { return proto.CompactTextString(m) }
func (*MessageQueue) ProtoMessage()    {}
func (*MessageQueue) Descriptor() ([]byte, []int) {
	return fileDescriptor_c20f641e39ad35ce, []int{10}
}

func (m *MessageQueue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageQueue.Unmarshal(m, b)
}
func (m *MessageQueue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageQueue.Marshal(b, m, deterministic)
}
func (m *MessageQueue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageQueue.Merge(m, src)
}
func (m *MessageQueue) XXX_Size() int {
	return xxx_messageInfo_MessageQueue.Size(m)
}
func (m *MessageQueue) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageQueue.DiscardUnknown(m)
}

var xxx_messageInfo_MessageQueue proto.InternalMessageInfo

func (m *MessageQueue) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *MessageQueue) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *MessageQueue) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *MessageQueue) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *MessageQueue) GetPublishTopicPrefix() string {
	if m != nil {
		return m.PublishTopicPrefix
	}
	return ""
}

func init() {
	proto.RegisterEnum("drivercommon.IotPlatform", IotPlatform_name, IotPlatform_value)
	proto.RegisterEnum("drivercommon.MetaBasesType", MetaBasesType_name, MetaBasesType_value)
	proto.RegisterType((*CloudInstanceInfo)(nil), "drivercommon.CloudInstanceInfo")
	proto.RegisterType((*BaseRequestMessage)(nil), "drivercommon.BaseRequestMessage")
	proto.RegisterType((*CommonResponse)(nil), "drivercommon.CommonResponse")
	proto.RegisterType((*Pong)(nil), "drivercommon.Pong")
	proto.RegisterType((*VersionResponse)(nil), "drivercommon.VersionResponse")
	proto.RegisterType((*ConfigResponse)(nil), "drivercommon.ConfigResponse")
	proto.RegisterType((*MetaBases)(nil), "drivercommon.MetaBases")
	proto.RegisterType((*DataBases)(nil), "drivercommon.DataBases")
	proto.RegisterType((*TDengineSource)(nil), "drivercommon.TDengineSource")
	proto.RegisterType((*InfluxDB)(nil), "drivercommon.InfluxDB")
	proto.RegisterType((*MessageQueue)(nil), "drivercommon.MessageQueue")
}

func init() { proto.RegisterFile("drivercommon/common.proto", fileDescriptor_c20f641e39ad35ce) }

var fileDescriptor_c20f641e39ad35ce = []byte{
	// 854 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x55, 0xdd, 0x6a, 0xe3, 0x46,
	0x14, 0x8e, 0x12, 0xc7, 0xb6, 0x8e, 0x1d, 0xc7, 0x3b, 0x2c, 0xa9, 0x37, 0xbb, 0xa5, 0x41, 0x94,
	0x12, 0xd2, 0xc6, 0x06, 0x6f, 0x17, 0x0a, 0x85, 0xc2, 0xc6, 0x29, 0x8d, 0x61, 0x93, 0xba, 0x8a,
	0x69, 0xa1, 0xf4, 0x46, 0x96, 0x8e, 0xb5, 0x43, 0xa4, 0x19, 0x67, 0x7e, 0x76, 0x37, 0x77, 0xbd,
	0xe8, 0x6d, 0x6f, 0xfb, 0x32, 0x85, 0x3e, 0x40, 0x2f, 0xfb, 0x44, 0x65, 0x46, 0x3f, 0x96, 0xec,
	0xcd, 0x95, 0xe6, 0x7c, 0xe7, 0x3b, 0xff, 0xa3, 0x33, 0xf0, 0x2c, 0x12, 0xf4, 0x1d, 0x8a, 0x90,
	0xa7, 0x29, 0x67, 0xa3, 0xec, 0x33, 0x5c, 0x09, 0xae, 0x38, 0xe9, 0x56, 0x55, 0xc7, 0xcf, 0x63,
	0xce, 0xe3, 0x04, 0x47, 0x56, 0xb7, 0xd0, 0xcb, 0x11, 0xa6, 0x2b, 0xf5, 0x90, 0x51, 0xbd, 0x7f,
	0x1d, 0x78, 0x32, 0x49, 0xb8, 0x8e, 0xa6, 0x4c, 0xaa, 0x80, 0x85, 0x38, 0x65, 0x4b, 0x4e, 0x4e,
	0xe1, 0x30, 0xac, 0x81, 0xd1, 0x60, 0xf7, 0xc4, 0x39, 0x75, 0xfd, 0x4d, 0x98, 0x9c, 0x40, 0x67,
	0x11, 0x48, 0x7c, 0x1d, 0x45, 0x02, 0xa5, 0x1c, 0xec, 0x59, 0x56, 0x15, 0x22, 0x5f, 0xc1, 0x93,
	0x9a, 0xd1, 0x4d, 0x90, 0xe2, 0xa0, 0x61, 0x79, 0xdb, 0x0a, 0xf2, 0x2d, 0x74, 0x28, 0x57, 0xb3,
	0x24, 0x50, 0x4b, 0x2e, 0xd2, 0xc1, 0xfe, 0x89, 0x73, 0xda, 0x1b, 0x3f, 0x1b, 0x56, 0x0b, 0x1a,
	0x4e, 0xd7, 0x04, 0xbf, 0xca, 0xf6, 0xfe, 0x76, 0x80, 0x5c, 0x04, 0x12, 0x7d, 0xbc, 0xd7, 0x28,
	0xd5, 0x35, 0x4a, 0x19, 0xc4, 0x48, 0xce, 0xa0, 0xaf, 0x25, 0xda, 0x2a, 0x4b, 0xc7, 0xce, 0x89,
	0x73, 0xda, 0xf6, 0xb7, 0x70, 0x72, 0xbd, 0x91, 0xad, 0x69, 0x87, 0xad, 0xbd, 0x33, 0xfe, 0xac,
	0x9e, 0xc5, 0x56, 0xd7, 0xfc, 0x6d, 0x4b, 0x13, 0x3a, 0x33, 0xaa, 0x74, 0x32, 0xeb, 0xd1, 0x16,
	0xee, 0xfd, 0xee, 0x40, 0x6f, 0x62, 0x7d, 0xfb, 0x28, 0x57, 0x9c, 0x49, 0x24, 0x2f, 0xc0, 0xcd,
	0x6b, 0x99, 0x46, 0x36, 0x65, 0xd7, 0x5f, 0x03, 0xc4, 0x83, 0x2e, 0x0a, 0xc1, 0x45, 0x5e, 0x67,
	0x3e, 0xa2, 0x1a, 0x46, 0x08, 0x34, 0x26, 0x3c, 0xc2, 0x3c, 0xa8, 0x3d, 0x93, 0x01, 0xb4, 0x6e,
	0x75, 0x18, 0x9a, 0x79, 0x35, 0x6c, 0x1b, 0x0a, 0xd1, 0xfb, 0x1c, 0x1a, 0x33, 0xce, 0x62, 0x13,
	0x57, 0xd1, 0x14, 0xa5, 0x0a, 0xd2, 0x55, 0xee, 0x76, 0x0d, 0x78, 0x5f, 0xc2, 0xe1, 0xcf, 0x28,
	0x24, 0xad, 0x24, 0x3a, 0x80, 0xd6, 0xbb, 0x0c, 0xca, 0xe9, 0x85, 0xe8, 0xfd, 0x63, 0xab, 0x62,
	0x4b, 0x1a, 0x97, 0xe4, 0x57, 0xe0, 0xa6, 0xa8, 0x02, 0x33, 0x29, 0x69, 0xab, 0xea, 0x8c, 0x3f,
	0xa9, 0xf7, 0xf6, 0xba, 0x50, 0xfb, 0x6b, 0xa6, 0x31, 0x8b, 0x82, 0xc2, 0x6c, 0xf7, 0x63, 0x66,
	0x97, 0x41, 0x69, 0x56, 0x32, 0xc9, 0x77, 0xd0, 0x4d, 0xb3, 0x66, 0xfc, 0xa4, 0x51, 0x67, 0x9d,
	0xe8, 0x8c, 0x8f, 0x37, 0x03, 0xae, 0x19, 0x7e, 0x8d, 0xef, 0xcd, 0xc1, 0x2d, 0xd3, 0x21, 0x23,
	0x68, 0xcc, 0x1f, 0x56, 0x68, 0xb3, 0xee, 0x8d, 0x9f, 0x3f, 0x92, 0xb5, 0xa1, 0xf8, 0x96, 0x48,
	0x8e, 0xa0, 0x79, 0xcb, 0xb5, 0x08, 0x8b, 0xe9, 0xe4, 0x92, 0xf7, 0xa7, 0x03, 0x6e, 0x99, 0xae,
	0x99, 0x52, 0xe9, 0xd6, 0xcd, 0x2d, 0xbf, 0x81, 0xb6, 0x8a, 0x90, 0xc5, 0x94, 0x61, 0x5e, 0xed,
	0x8b, 0x7a, 0xb8, 0xf9, 0x65, 0xa6, 0xcd, 0x3c, 0xfa, 0x25, 0x9b, 0x8c, 0xa1, 0x4d, 0xd9, 0x32,
	0xd1, 0x1f, 0x2e, 0x2f, 0xf2, 0x6a, 0x8f, 0x36, 0x7e, 0xa0, 0x5c, 0xeb, 0x97, 0x3c, 0xcf, 0x83,
	0x5e, 0xdd, 0x1f, 0xe9, 0xc3, 0xde, 0xa5, 0x64, 0x79, 0x4a, 0xe6, 0xe8, 0xfd, 0x06, 0xed, 0xc2,
	0xd2, 0x68, 0xb9, 0x88, 0x0b, 0x2d, 0x17, 0xb1, 0xa9, 0x74, 0xa1, 0xc3, 0x3b, 0x54, 0x45, 0xa5,
	0x99, 0x64, 0x98, 0x5a, 0x24, 0xf9, 0x05, 0x34, 0x47, 0xf2, 0x14, 0xf6, 0x15, 0xbf, 0x43, 0x96,
	0x6f, 0x81, 0x4c, 0xf0, 0xfe, 0x72, 0xa0, 0x5b, 0x1d, 0x03, 0x39, 0x86, 0xf6, 0xcc, 0xec, 0xa8,
	0x90, 0x27, 0x79, 0x9c, 0x52, 0x36, 0x0d, 0xbb, 0xe2, 0xb2, 0x08, 0x65, 0xcf, 0x06, 0x9b, 0x71,
	0xa1, 0x6c, 0xa4, 0x03, 0xdf, 0x9e, 0xcb, 0xc6, 0x36, 0x2a, 0x8d, 0x1d, 0x02, 0x99, 0xe9, 0x45,
	0x42, 0xe5, 0xdb, 0x39, 0x5f, 0xd1, 0x70, 0x26, 0x70, 0x49, 0x3f, 0xd8, 0x4d, 0xe3, 0xfa, 0x1f,
	0xd1, 0x9c, 0xfd, 0xe1, 0x40, 0xa7, 0xb2, 0x72, 0x48, 0x17, 0xda, 0x6f, 0x78, 0x18, 0x24, 0x53,
	0xae, 0xfa, 0x3b, 0xe4, 0x10, 0x3a, 0x13, 0x2d, 0x15, 0x4f, 0x51, 0x18, 0xc0, 0x31, 0xc0, 0x2f,
	0x94, 0x4d, 0xde, 0x50, 0x76, 0x67, 0x80, 0x5d, 0x02, 0xd0, 0x7c, 0x9d, 0x50, 0x73, 0xde, 0x23,
	0x07, 0xe0, 0x5e, 0xe9, 0xe0, 0x3d, 0x5a, 0xb1, 0x41, 0x7a, 0x00, 0x73, 0x64, 0x21, 0x32, 0x65,
	0xe4, 0x7d, 0xd2, 0x81, 0xd6, 0x5c, 0x3f, 0x04, 0x46, 0x68, 0x1a, 0xee, 0x8f, 0x0c, 0x6f, 0xd0,
	0xea, 0x5a, 0x67, 0x5f, 0xc0, 0x41, 0xed, 0x82, 0x11, 0x17, 0xf6, 0xaf, 0x1f, 0xe4, 0x7d, 0xd2,
	0xdf, 0x31, 0x21, 0x6e, 0xef, 0x13, 0xaa, 0xb0, 0xef, 0x8c, 0xff, 0x73, 0xa0, 0x99, 0xad, 0x11,
	0xf2, 0x35, 0x34, 0x66, 0x94, 0xc5, 0xe4, 0x68, 0x98, 0x3d, 0x01, 0xc3, 0xe2, 0x09, 0x18, 0x7e,
	0x6f, 0x9e, 0x80, 0x63, 0x52, 0xbf, 0x16, 0xe6, 0xd7, 0xf7, 0x76, 0xc8, 0x05, 0xb4, 0xf2, 0xdf,
	0xfb, 0x51, 0xc3, 0x4f, 0xeb, 0x86, 0x1b, 0xdb, 0xc0, 0xdb, 0x21, 0x37, 0xf0, 0xf4, 0x07, 0x54,
	0x57, 0x3a, 0x4d, 0x29, 0x8b, 0x17, 0x54, 0x44, 0xd9, 0x0a, 0x78, 0xd4, 0xe1, 0xc6, 0xd5, 0xae,
	0x2f, 0x0c, 0x6f, 0xe7, 0xe2, 0xd5, 0xaf, 0x2f, 0x63, 0xaa, 0xde, 0xea, 0xc5, 0x30, 0xe4, 0xe9,
	0xe8, 0x3d, 0x65, 0xe1, 0x79, 0x42, 0xd9, 0xdd, 0x08, 0xa3, 0x18, 0xcf, 0x33, 0xd3, 0x73, 0xeb,
	0x72, 0x54, 0xf5, 0xb3, 0x68, 0x5a, 0xec, 0xe5, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x71, 0x58,
	0x6f, 0x3e, 0x2c, 0x07, 0x00, 0x00,
}
