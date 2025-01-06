// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: devicecallback/devicecallback.proto

package devicecallback

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	DeviceCallBackService_CreateDeviceCallback_FullMethodName = "/devicecallback.DeviceCallBackService/CreateDeviceCallback"
	DeviceCallBackService_UpdateDeviceCallback_FullMethodName = "/devicecallback.DeviceCallBackService/UpdateDeviceCallback"
	DeviceCallBackService_DeleteDeviceCallback_FullMethodName = "/devicecallback.DeviceCallBackService/DeleteDeviceCallback"
)

// DeviceCallBackServiceClient is the client API for DeviceCallBackService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeviceCallBackServiceClient interface {
	// 创建设备回调 edge = c driver = s
	CreateDeviceCallback(ctx context.Context, in *CreateDeviceCallbackRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 更新设备回调 edge = c driver = s
	UpdateDeviceCallback(ctx context.Context, in *UpdateDeviceCallbackRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 删除设备回调 edge = c driver = s
	DeleteDeviceCallback(ctx context.Context, in *DeleteDeviceCallbackRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type deviceCallBackServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeviceCallBackServiceClient(cc grpc.ClientConnInterface) DeviceCallBackServiceClient {
	return &deviceCallBackServiceClient{cc}
}

func (c *deviceCallBackServiceClient) CreateDeviceCallback(ctx context.Context, in *CreateDeviceCallbackRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DeviceCallBackService_CreateDeviceCallback_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceCallBackServiceClient) UpdateDeviceCallback(ctx context.Context, in *UpdateDeviceCallbackRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DeviceCallBackService_UpdateDeviceCallback_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceCallBackServiceClient) DeleteDeviceCallback(ctx context.Context, in *DeleteDeviceCallbackRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DeviceCallBackService_DeleteDeviceCallback_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeviceCallBackServiceServer is the server API for DeviceCallBackService service.
// All implementations must embed UnimplementedDeviceCallBackServiceServer
// for forward compatibility.
type DeviceCallBackServiceServer interface {
	// 创建设备回调 edge = c driver = s
	CreateDeviceCallback(context.Context, *CreateDeviceCallbackRequest) (*emptypb.Empty, error)
	// 更新设备回调 edge = c driver = s
	UpdateDeviceCallback(context.Context, *UpdateDeviceCallbackRequest) (*emptypb.Empty, error)
	// 删除设备回调 edge = c driver = s
	DeleteDeviceCallback(context.Context, *DeleteDeviceCallbackRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedDeviceCallBackServiceServer()
}

// UnimplementedDeviceCallBackServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDeviceCallBackServiceServer struct{}

func (UnimplementedDeviceCallBackServiceServer) CreateDeviceCallback(context.Context, *CreateDeviceCallbackRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDeviceCallback not implemented")
}
func (UnimplementedDeviceCallBackServiceServer) UpdateDeviceCallback(context.Context, *UpdateDeviceCallbackRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDeviceCallback not implemented")
}
func (UnimplementedDeviceCallBackServiceServer) DeleteDeviceCallback(context.Context, *DeleteDeviceCallbackRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDeviceCallback not implemented")
}
func (UnimplementedDeviceCallBackServiceServer) mustEmbedUnimplementedDeviceCallBackServiceServer() {}
func (UnimplementedDeviceCallBackServiceServer) testEmbeddedByValue()                               {}

// UnsafeDeviceCallBackServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeviceCallBackServiceServer will
// result in compilation errors.
type UnsafeDeviceCallBackServiceServer interface {
	mustEmbedUnimplementedDeviceCallBackServiceServer()
}

func RegisterDeviceCallBackServiceServer(s grpc.ServiceRegistrar, srv DeviceCallBackServiceServer) {
	// If the following call pancis, it indicates UnimplementedDeviceCallBackServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DeviceCallBackService_ServiceDesc, srv)
}

func _DeviceCallBackService_CreateDeviceCallback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDeviceCallbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceCallBackServiceServer).CreateDeviceCallback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeviceCallBackService_CreateDeviceCallback_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceCallBackServiceServer).CreateDeviceCallback(ctx, req.(*CreateDeviceCallbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceCallBackService_UpdateDeviceCallback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDeviceCallbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceCallBackServiceServer).UpdateDeviceCallback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeviceCallBackService_UpdateDeviceCallback_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceCallBackServiceServer).UpdateDeviceCallback(ctx, req.(*UpdateDeviceCallbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceCallBackService_DeleteDeviceCallback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDeviceCallbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceCallBackServiceServer).DeleteDeviceCallback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeviceCallBackService_DeleteDeviceCallback_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceCallBackServiceServer).DeleteDeviceCallback(ctx, req.(*DeleteDeviceCallbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DeviceCallBackService_ServiceDesc is the grpc.ServiceDesc for DeviceCallBackService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeviceCallBackService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "devicecallback.DeviceCallBackService",
	HandlerType: (*DeviceCallBackServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDeviceCallback",
			Handler:    _DeviceCallBackService_CreateDeviceCallback_Handler,
		},
		{
			MethodName: "UpdateDeviceCallback",
			Handler:    _DeviceCallBackService_UpdateDeviceCallback_Handler,
		},
		{
			MethodName: "DeleteDeviceCallback",
			Handler:    _DeviceCallBackService_DeleteDeviceCallback_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "devicecallback/devicecallback.proto",
}
