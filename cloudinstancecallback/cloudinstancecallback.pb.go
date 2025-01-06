// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v3.21.12
// source: cloudinstancecallback/cloudinstancecallback.proto

package cloudinstancecallback

import (
	cloudinstance "github.com/winc-link/edge-driver-proto/cloudinstance"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CloudInstanceStatueCallbackRequest struct {
	state             protoimpl.MessageState            `protogen:"open.v1"`
	CloudInstanceName string                            `protobuf:"bytes,1,opt,name=cloudInstanceName,proto3" json:"cloudInstanceName,omitempty"`
	Status            cloudinstance.CloudInstanceStatus `protobuf:"varint,2,opt,name=status,proto3,enum=cloudinstance.CloudInstanceStatus" json:"status,omitempty"`
	HappenTime        uint64                            `protobuf:"varint,3,opt,name=happenTime,proto3" json:"happenTime,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *CloudInstanceStatueCallbackRequest) Reset() {
	*x = CloudInstanceStatueCallbackRequest{}
	mi := &file_cloudinstancecallback_cloudinstancecallback_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CloudInstanceStatueCallbackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudInstanceStatueCallbackRequest) ProtoMessage() {}

func (x *CloudInstanceStatueCallbackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cloudinstancecallback_cloudinstancecallback_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudInstanceStatueCallbackRequest.ProtoReflect.Descriptor instead.
func (*CloudInstanceStatueCallbackRequest) Descriptor() ([]byte, []int) {
	return file_cloudinstancecallback_cloudinstancecallback_proto_rawDescGZIP(), []int{0}
}

func (x *CloudInstanceStatueCallbackRequest) GetCloudInstanceName() string {
	if x != nil {
		return x.CloudInstanceName
	}
	return ""
}

func (x *CloudInstanceStatueCallbackRequest) GetStatus() cloudinstance.CloudInstanceStatus {
	if x != nil {
		return x.Status
	}
	return cloudinstance.CloudInstanceStatus(0)
}

func (x *CloudInstanceStatueCallbackRequest) GetHappenTime() uint64 {
	if x != nil {
		return x.HappenTime
	}
	return 0
}

var File_cloudinstancecallback_cloudinstancecallback_proto protoreflect.FileDescriptor

var file_cloudinstancecallback_cloudinstancecallback_proto_rawDesc = []byte{
	0x0a, 0x31, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x63,
	0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x15, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xae, 0x01, 0x0a, 0x22, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2c, 0x0a, 0x11, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x3a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x22, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x68,
	0x61, 0x70, 0x70, 0x65, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0a, 0x68, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x32, 0x92, 0x01, 0x0a, 0x1c,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x61, 0x6c,
	0x6c, 0x42, 0x61, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x72, 0x0a, 0x1b,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x39, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x63, 0x61, 0x6c, 0x6c, 0x62,
	0x61, 0x63, 0x6b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77,
	0x69, 0x6e, 0x63, 0x2d, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x2d, 0x64, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloudinstancecallback_cloudinstancecallback_proto_rawDescOnce sync.Once
	file_cloudinstancecallback_cloudinstancecallback_proto_rawDescData = file_cloudinstancecallback_cloudinstancecallback_proto_rawDesc
)

func file_cloudinstancecallback_cloudinstancecallback_proto_rawDescGZIP() []byte {
	file_cloudinstancecallback_cloudinstancecallback_proto_rawDescOnce.Do(func() {
		file_cloudinstancecallback_cloudinstancecallback_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloudinstancecallback_cloudinstancecallback_proto_rawDescData)
	})
	return file_cloudinstancecallback_cloudinstancecallback_proto_rawDescData
}

var file_cloudinstancecallback_cloudinstancecallback_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cloudinstancecallback_cloudinstancecallback_proto_goTypes = []any{
	(*CloudInstanceStatueCallbackRequest)(nil), // 0: cloudinstancecallback.cloudInstanceStatueCallbackRequest
	(cloudinstance.CloudInstanceStatus)(0),     // 1: cloudinstance.CloudInstanceStatus
	(*emptypb.Empty)(nil),                      // 2: google.protobuf.Empty
}
var file_cloudinstancecallback_cloudinstancecallback_proto_depIdxs = []int32{
	1, // 0: cloudinstancecallback.cloudInstanceStatueCallbackRequest.status:type_name -> cloudinstance.CloudInstanceStatus
	0, // 1: cloudinstancecallback.CloudInstanceCallBackService.CloudInstanceStatueCallback:input_type -> cloudinstancecallback.cloudInstanceStatueCallbackRequest
	2, // 2: cloudinstancecallback.CloudInstanceCallBackService.CloudInstanceStatueCallback:output_type -> google.protobuf.Empty
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cloudinstancecallback_cloudinstancecallback_proto_init() }
func file_cloudinstancecallback_cloudinstancecallback_proto_init() {
	if File_cloudinstancecallback_cloudinstancecallback_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloudinstancecallback_cloudinstancecallback_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cloudinstancecallback_cloudinstancecallback_proto_goTypes,
		DependencyIndexes: file_cloudinstancecallback_cloudinstancecallback_proto_depIdxs,
		MessageInfos:      file_cloudinstancecallback_cloudinstancecallback_proto_msgTypes,
	}.Build()
	File_cloudinstancecallback_cloudinstancecallback_proto = out.File
	file_cloudinstancecallback_cloudinstancecallback_proto_rawDesc = nil
	file_cloudinstancecallback_cloudinstancecallback_proto_goTypes = nil
	file_cloudinstancecallback_cloudinstancecallback_proto_depIdxs = nil
}
