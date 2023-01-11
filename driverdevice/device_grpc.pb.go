// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.14.0
// source: driverdevice/device.proto

package driverdevice

import (
	context "context"
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
	// 设备连接云服务
	ConnectIotPlatform(ctx context.Context, in *ConnectIotPlatformRequest, opts ...grpc.CallOption) (*ConnectIotPlatformResponse, error)
	// 设备断开连接云服务
	DisconnectIotPlatform(ctx context.Context, in *DisconnectIotPlatformRequest, opts ...grpc.CallOption) (*DisconnectIotPlatformResponse, error)
	// 设备连接状态
	GetDeviceConnectStatus(ctx context.Context, in *GetDeviceConnectStatusRequest, opts ...grpc.CallOption) (*GetDeviceConnectStatusResponse, error)
	// 获取所有设备
	QueryDeviceList(ctx context.Context, in *QueryDeviceListRequest, opts ...grpc.CallOption) (*QueryDeviceListResponse, error)
	// 获取设备
	QueryDeviceById(ctx context.Context, in *QueryDeviceByIdRequest, opts ...grpc.CallOption) (*QueryDeviceByIdRequestResponse, error)
	// 创建设备
	CreateDevice(ctx context.Context, in *CreateDeviceRequest, opts ...grpc.CallOption) (*CreateDeviceRequestResponse, error)
	// 创建设备并且建立连接
	CreateDeviceAndConnect(ctx context.Context, in *CreateDeviceAndConnectRequest, opts ...grpc.CallOption) (*CreateDeviceAndConnectRequestResponse, error)
	// 删除设备
	DeleteDevice(ctx context.Context, in *DeleteDeviceRequest, opts ...grpc.CallOption) (*DeleteDeviceResponse, error)
}

type rpcDeviceClient struct {
	cc grpc.ClientConnInterface
}

func NewRpcDeviceClient(cc grpc.ClientConnInterface) RpcDeviceClient {
	return &rpcDeviceClient{cc}
}

func (c *rpcDeviceClient) ConnectIotPlatform(ctx context.Context, in *ConnectIotPlatformRequest, opts ...grpc.CallOption) (*ConnectIotPlatformResponse, error) {
	out := new(ConnectIotPlatformResponse)
	err := c.cc.Invoke(ctx, "/driverdevice.RpcDevice/ConnectIotPlatform", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcDeviceClient) DisconnectIotPlatform(ctx context.Context, in *DisconnectIotPlatformRequest, opts ...grpc.CallOption) (*DisconnectIotPlatformResponse, error) {
	out := new(DisconnectIotPlatformResponse)
	err := c.cc.Invoke(ctx, "/driverdevice.RpcDevice/DisconnectIotPlatform", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcDeviceClient) GetDeviceConnectStatus(ctx context.Context, in *GetDeviceConnectStatusRequest, opts ...grpc.CallOption) (*GetDeviceConnectStatusResponse, error) {
	out := new(GetDeviceConnectStatusResponse)
	err := c.cc.Invoke(ctx, "/driverdevice.RpcDevice/GetDeviceConnectStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcDeviceClient) QueryDeviceList(ctx context.Context, in *QueryDeviceListRequest, opts ...grpc.CallOption) (*QueryDeviceListResponse, error) {
	out := new(QueryDeviceListResponse)
	err := c.cc.Invoke(ctx, "/driverdevice.RpcDevice/QueryDeviceList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcDeviceClient) QueryDeviceById(ctx context.Context, in *QueryDeviceByIdRequest, opts ...grpc.CallOption) (*QueryDeviceByIdRequestResponse, error) {
	out := new(QueryDeviceByIdRequestResponse)
	err := c.cc.Invoke(ctx, "/driverdevice.RpcDevice/QueryDeviceById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcDeviceClient) CreateDevice(ctx context.Context, in *CreateDeviceRequest, opts ...grpc.CallOption) (*CreateDeviceRequestResponse, error) {
	out := new(CreateDeviceRequestResponse)
	err := c.cc.Invoke(ctx, "/driverdevice.RpcDevice/CreateDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcDeviceClient) CreateDeviceAndConnect(ctx context.Context, in *CreateDeviceAndConnectRequest, opts ...grpc.CallOption) (*CreateDeviceAndConnectRequestResponse, error) {
	out := new(CreateDeviceAndConnectRequestResponse)
	err := c.cc.Invoke(ctx, "/driverdevice.RpcDevice/CreateDeviceAndConnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcDeviceClient) DeleteDevice(ctx context.Context, in *DeleteDeviceRequest, opts ...grpc.CallOption) (*DeleteDeviceResponse, error) {
	out := new(DeleteDeviceResponse)
	err := c.cc.Invoke(ctx, "/driverdevice.RpcDevice/DeleteDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RpcDeviceServer is the server API for RpcDevice service.
// All implementations must embed UnimplementedRpcDeviceServer
// for forward compatibility
type RpcDeviceServer interface {
	// 设备连接云服务
	ConnectIotPlatform(context.Context, *ConnectIotPlatformRequest) (*ConnectIotPlatformResponse, error)
	// 设备断开连接云服务
	DisconnectIotPlatform(context.Context, *DisconnectIotPlatformRequest) (*DisconnectIotPlatformResponse, error)
	// 设备连接状态
	GetDeviceConnectStatus(context.Context, *GetDeviceConnectStatusRequest) (*GetDeviceConnectStatusResponse, error)
	// 获取所有设备
	QueryDeviceList(context.Context, *QueryDeviceListRequest) (*QueryDeviceListResponse, error)
	// 获取设备
	QueryDeviceById(context.Context, *QueryDeviceByIdRequest) (*QueryDeviceByIdRequestResponse, error)
	// 创建设备
	CreateDevice(context.Context, *CreateDeviceRequest) (*CreateDeviceRequestResponse, error)
	// 创建设备并且建立连接
	CreateDeviceAndConnect(context.Context, *CreateDeviceAndConnectRequest) (*CreateDeviceAndConnectRequestResponse, error)
	// 删除设备
	DeleteDevice(context.Context, *DeleteDeviceRequest) (*DeleteDeviceResponse, error)
	mustEmbedUnimplementedRpcDeviceServer()
}

// UnimplementedRpcDeviceServer must be embedded to have forward compatible implementations.
type UnimplementedRpcDeviceServer struct {
}

func (UnimplementedRpcDeviceServer) ConnectIotPlatform(context.Context, *ConnectIotPlatformRequest) (*ConnectIotPlatformResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConnectIotPlatform not implemented")
}
func (UnimplementedRpcDeviceServer) DisconnectIotPlatform(context.Context, *DisconnectIotPlatformRequest) (*DisconnectIotPlatformResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisconnectIotPlatform not implemented")
}
func (UnimplementedRpcDeviceServer) GetDeviceConnectStatus(context.Context, *GetDeviceConnectStatusRequest) (*GetDeviceConnectStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceConnectStatus not implemented")
}
func (UnimplementedRpcDeviceServer) QueryDeviceList(context.Context, *QueryDeviceListRequest) (*QueryDeviceListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryDeviceList not implemented")
}
func (UnimplementedRpcDeviceServer) QueryDeviceById(context.Context, *QueryDeviceByIdRequest) (*QueryDeviceByIdRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryDeviceById not implemented")
}
func (UnimplementedRpcDeviceServer) CreateDevice(context.Context, *CreateDeviceRequest) (*CreateDeviceRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDevice not implemented")
}
func (UnimplementedRpcDeviceServer) CreateDeviceAndConnect(context.Context, *CreateDeviceAndConnectRequest) (*CreateDeviceAndConnectRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDeviceAndConnect not implemented")
}
func (UnimplementedRpcDeviceServer) DeleteDevice(context.Context, *DeleteDeviceRequest) (*DeleteDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDevice not implemented")
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

func _RpcDevice_ConnectIotPlatform_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectIotPlatformRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcDeviceServer).ConnectIotPlatform(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/driverdevice.RpcDevice/ConnectIotPlatform",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcDeviceServer).ConnectIotPlatform(ctx, req.(*ConnectIotPlatformRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcDevice_DisconnectIotPlatform_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisconnectIotPlatformRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcDeviceServer).DisconnectIotPlatform(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/driverdevice.RpcDevice/DisconnectIotPlatform",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcDeviceServer).DisconnectIotPlatform(ctx, req.(*DisconnectIotPlatformRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcDevice_GetDeviceConnectStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceConnectStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcDeviceServer).GetDeviceConnectStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/driverdevice.RpcDevice/GetDeviceConnectStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcDeviceServer).GetDeviceConnectStatus(ctx, req.(*GetDeviceConnectStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcDevice_QueryDeviceList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDeviceListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcDeviceServer).QueryDeviceList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/driverdevice.RpcDevice/QueryDeviceList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcDeviceServer).QueryDeviceList(ctx, req.(*QueryDeviceListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcDevice_QueryDeviceById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDeviceByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcDeviceServer).QueryDeviceById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/driverdevice.RpcDevice/QueryDeviceById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcDeviceServer).QueryDeviceById(ctx, req.(*QueryDeviceByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcDevice_CreateDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcDeviceServer).CreateDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/driverdevice.RpcDevice/CreateDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcDeviceServer).CreateDevice(ctx, req.(*CreateDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcDevice_CreateDeviceAndConnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDeviceAndConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcDeviceServer).CreateDeviceAndConnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/driverdevice.RpcDevice/CreateDeviceAndConnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcDeviceServer).CreateDeviceAndConnect(ctx, req.(*CreateDeviceAndConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcDevice_DeleteDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcDeviceServer).DeleteDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/driverdevice.RpcDevice/DeleteDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcDeviceServer).DeleteDevice(ctx, req.(*DeleteDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RpcDevice_ServiceDesc is the grpc.ServiceDesc for RpcDevice service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RpcDevice_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "driverdevice.RpcDevice",
	HandlerType: (*RpcDeviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConnectIotPlatform",
			Handler:    _RpcDevice_ConnectIotPlatform_Handler,
		},
		{
			MethodName: "DisconnectIotPlatform",
			Handler:    _RpcDevice_DisconnectIotPlatform_Handler,
		},
		{
			MethodName: "GetDeviceConnectStatus",
			Handler:    _RpcDevice_GetDeviceConnectStatus_Handler,
		},
		{
			MethodName: "QueryDeviceList",
			Handler:    _RpcDevice_QueryDeviceList_Handler,
		},
		{
			MethodName: "QueryDeviceById",
			Handler:    _RpcDevice_QueryDeviceById_Handler,
		},
		{
			MethodName: "CreateDevice",
			Handler:    _RpcDevice_CreateDevice_Handler,
		},
		{
			MethodName: "CreateDeviceAndConnect",
			Handler:    _RpcDevice_CreateDeviceAndConnect_Handler,
		},
		{
			MethodName: "DeleteDevice",
			Handler:    _RpcDevice_DeleteDevice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "driverdevice/device.proto",
}
