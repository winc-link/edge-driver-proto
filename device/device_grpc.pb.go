// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.14.0
// source: device/device.proto

package device

import (
	context "context"
	drivercommon "github.com/winc-link/edgex-driver-proto/drivercommon"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RpcDeviceClient is the client API for RpcDevice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RpcDeviceClient interface {
	//设备连接云服务
	ConnectIotCloud(ctx context.Context, in *ConnectIotCloudRequest, opts ...grpc.CallOption) (*drivercommon.CommonResponse, error)
}

type rpcDeviceClient struct {
	cc grpc.ClientConnInterface
}

func NewRpcDeviceClient(cc grpc.ClientConnInterface) RpcDeviceClient {
	return &rpcDeviceClient{cc}
}

func (c *rpcDeviceClient) ConnectIotCloud(ctx context.Context, in *ConnectIotCloudRequest, opts ...grpc.CallOption) (*drivercommon.CommonResponse, error) {
	out := new(drivercommon.CommonResponse)
	err := c.cc.Invoke(ctx, "/device.RpcDevice/ConnectIotCloud", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RpcDeviceServer is the server API for RpcDevice service.
// All implementations must embed UnimplementedRpcDeviceServer
// for forward compatibility
type RpcDeviceServer interface {
	//设备连接云服务
	ConnectIotCloud(context.Context, *ConnectIotCloudRequest) (*drivercommon.CommonResponse, error)
	mustEmbedUnimplementedRpcDeviceServer()
}

// UnimplementedRpcDeviceServer must be embedded to have forward compatible implementations.
type UnimplementedRpcDeviceServer struct {
}

func (UnimplementedRpcDeviceServer) ConnectIotCloud(context.Context, *ConnectIotCloudRequest) (*drivercommon.CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConnectIotCloud not implemented")
}
func (UnimplementedRpcDeviceServer) mustEmbedUnimplementedRpcDeviceServer() {}

// UnsafeRpcDeviceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RpcDeviceServer will
// result in compilation errors.
type UnsafeRpcDeviceServer interface {
	mustEmbedUnimplementedRpcDeviceServer()
}

func RegisterRpcDeviceServer(s grpc.ServiceRegistrar, srv RpcDeviceServer) {
	s.RegisterService(&RpcDevice_ServiceDesc, srv)
}

func _RpcDevice_ConnectIotCloud_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectIotCloudRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcDeviceServer).ConnectIotCloud(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/device.RpcDevice/ConnectIotCloud",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcDeviceServer).ConnectIotCloud(ctx, req.(*ConnectIotCloudRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RpcDevice_ServiceDesc is the grpc.ServiceDesc for RpcDevice service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RpcDevice_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "device.RpcDevice",
	HandlerType: (*RpcDeviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConnectIotCloud",
			Handler:    _RpcDevice_ConnectIotCloud_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "device/device.proto",
}
