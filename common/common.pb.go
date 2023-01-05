// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: common/common.proto

package common

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

type IotPlatform int32

const (
	IotPlatform_CustomerIot IotPlatform = 0 //用户自定义
	IotPlatform_WinCLinkIot IotPlatform = 1 //赢创万联
	IotPlatform_AliIot      IotPlatform = 2 //阿里
	IotPlatform_HuaweiIot   IotPlatform = 3 //华为
	IotPlatform_TencentIot  IotPlatform = 4 //腾讯
	IotPlatform_TuyaIot     IotPlatform = 5 //涂鸦
	IotPlatform_OneNetIot   IotPlatform = 6 //中国移动
)

// Enum value maps for IotPlatform.
var (
	IotPlatform_name = map[int32]string{
		0: "CustomerIot",
		1: "WinCLinkIot",
		2: "AliIot",
		3: "HuaweiIot",
		4: "TencentIot",
		5: "TuyaIot",
		6: "OneNetIot",
	}
	IotPlatform_value = map[string]int32{
		"CustomerIot": 0,
		"WinCLinkIot": 1,
		"AliIot":      2,
		"HuaweiIot":   3,
		"TencentIot":  4,
		"TuyaIot":     5,
		"OneNetIot":   6,
	}
)

func (x IotPlatform) Enum() *IotPlatform {
	p := new(IotPlatform)
	*p = x
	return p
}

func (x IotPlatform) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IotPlatform) Descriptor() protoreflect.EnumDescriptor {
	return file_common_common_proto_enumTypes[0].Descriptor()
}

func (IotPlatform) Type() protoreflect.EnumType {
	return &file_common_common_proto_enumTypes[0]
}

func (x IotPlatform) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IotPlatform.Descriptor instead.
func (IotPlatform) EnumDescriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{0}
}

type CloudInstanceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CloudInstanceId string `protobuf:"bytes,2,opt,name=cloudInstanceId,proto3" json:"cloudInstanceId,omitempty"`
	BaseAddress     string `protobuf:"bytes,3,opt,name=baseAddress,proto3" json:"baseAddress,omitempty"`
}

func (x *CloudInstanceInfo) Reset() {
	*x = CloudInstanceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudInstanceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudInstanceInfo) ProtoMessage() {}

func (x *CloudInstanceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudInstanceInfo.ProtoReflect.Descriptor instead.
func (*CloudInstanceInfo) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{0}
}

func (x *CloudInstanceInfo) GetCloudInstanceId() string {
	if x != nil {
		return x.CloudInstanceId
	}
	return ""
}

func (x *CloudInstanceInfo) GetBaseAddress() string {
	if x != nil {
		return x.BaseAddress
	}
	return ""
}

type BaseIotPlatformMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UseCloudPlatform  bool               `protobuf:"varint,1,opt,name=useCloudPlatform,proto3" json:"useCloudPlatform,omitempty"`
	CloudInstanceInfo *CloudInstanceInfo `protobuf:"bytes,2,opt,name=cloudInstanceInfo,proto3" json:"cloudInstanceInfo,omitempty"`
}

func (x *BaseIotPlatformMessage) Reset() {
	*x = BaseIotPlatformMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseIotPlatformMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseIotPlatformMessage) ProtoMessage() {}

func (x *BaseIotPlatformMessage) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseIotPlatformMessage.ProtoReflect.Descriptor instead.
func (*BaseIotPlatformMessage) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{1}
}

func (x *BaseIotPlatformMessage) GetUseCloudPlatform() bool {
	if x != nil {
		return x.UseCloudPlatform
	}
	return false
}

func (x *BaseIotPlatformMessage) GetCloudInstanceInfo() *CloudInstanceInfo {
	if x != nil {
		return x.CloudInstanceInfo
	}
	return nil
}

type CommonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId    string `protobuf:"bytes,1,opt,name=RequestId,proto3" json:"RequestId,omitempty"`
	ErrorMessage string `protobuf:"bytes,2,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	Code         string `protobuf:"bytes,3,opt,name=Code,proto3" json:"Code,omitempty"`
	Success      bool   `protobuf:"varint,4,opt,name=Success,proto3" json:"Success,omitempty"`
}

func (x *CommonResponse) Reset() {
	*x = CommonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonResponse) ProtoMessage() {}

func (x *CommonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonResponse.ProtoReflect.Descriptor instead.
func (*CommonResponse) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{2}
}

func (x *CommonResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *CommonResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *CommonResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CommonResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_common_common_proto protoreflect.FileDescriptor

var file_common_common_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0x5f, 0x0a,
	0x11, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b,
	0x62, 0x61, 0x73, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x8d,
	0x01, 0x0a, 0x16, 0x62, 0x61, 0x73, 0x65, 0x49, 0x6f, 0x74, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x75, 0x73, 0x65,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x10, 0x75, 0x73, 0x65, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x47, 0x0a, 0x11, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x11, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x80,
	0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12,
	0x22, 0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x2a, 0x76, 0x0a, 0x0b, 0x49, 0x6f, 0x74, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x6f, 0x74, 0x10,
	0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x57, 0x69, 0x6e, 0x43, 0x4c, 0x69, 0x6e, 0x6b, 0x49, 0x6f, 0x74,
	0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x6c, 0x69, 0x49, 0x6f, 0x74, 0x10, 0x02, 0x12, 0x0d,
	0x0a, 0x09, 0x48, 0x75, 0x61, 0x77, 0x65, 0x69, 0x49, 0x6f, 0x74, 0x10, 0x03, 0x12, 0x0e, 0x0a,
	0x0a, 0x54, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x49, 0x6f, 0x74, 0x10, 0x04, 0x12, 0x0b, 0x0a,
	0x07, 0x54, 0x75, 0x79, 0x61, 0x49, 0x6f, 0x74, 0x10, 0x05, 0x12, 0x0d, 0x0a, 0x09, 0x4f, 0x6e,
	0x65, 0x4e, 0x65, 0x74, 0x49, 0x6f, 0x74, 0x10, 0x06, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x69, 0x6e, 0x63, 0x2d, 0x6c, 0x69, 0x6e,
	0x6b, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x78, 0x2d, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_common_common_proto_rawDescOnce sync.Once
	file_common_common_proto_rawDescData = file_common_common_proto_rawDesc
)

func file_common_common_proto_rawDescGZIP() []byte {
	file_common_common_proto_rawDescOnce.Do(func() {
		file_common_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_common_proto_rawDescData)
	})
	return file_common_common_proto_rawDescData
}

var file_common_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_common_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_common_common_proto_goTypes = []interface{}{
	(IotPlatform)(0),               // 0: common.IotPlatform
	(*CloudInstanceInfo)(nil),      // 1: common.CloudInstanceInfo
	(*BaseIotPlatformMessage)(nil), // 2: common.baseIotPlatformMessage
	(*CommonResponse)(nil),         // 3: common.CommonResponse
}
var file_common_common_proto_depIdxs = []int32{
	1, // 0: common.baseIotPlatformMessage.cloudInstanceInfo:type_name -> common.CloudInstanceInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_common_common_proto_init() }
func file_common_common_proto_init() {
	if File_common_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloudInstanceInfo); i {
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
		file_common_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseIotPlatformMessage); i {
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
		file_common_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonResponse); i {
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
			RawDescriptor: file_common_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_common_proto_goTypes,
		DependencyIndexes: file_common_common_proto_depIdxs,
		EnumInfos:         file_common_common_proto_enumTypes,
		MessageInfos:      file_common_common_proto_msgTypes,
	}.Build()
	File_common_common_proto = out.File
	file_common_common_proto_rawDesc = nil
	file_common_common_proto_goTypes = nil
	file_common_common_proto_depIdxs = nil
}