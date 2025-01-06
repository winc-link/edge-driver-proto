// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: drivercommon/common.proto

package drivercommon

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
	Common_Ping_FullMethodName    = "/drivercommon.Common/Ping"
	Common_Version_FullMethodName = "/drivercommon.Common/Version"
)

// CommonClient is the client API for Common service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommonClient interface {
	// Ping tests whether the service is working
	Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Pong, error)
	// Version obtains version information from the target service.
	Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VersionResponse, error)
}

type commonClient struct {
	cc grpc.ClientConnInterface
}

func NewCommonClient(cc grpc.ClientConnInterface) CommonClient {
	return &commonClient{cc}
}

func (c *commonClient) Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Pong, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Pong)
	err := c.cc.Invoke(ctx, Common_Ping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonClient) Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VersionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, Common_Version_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommonServer is the server API for Common service.
// All implementations must embed UnimplementedCommonServer
// for forward compatibility.
type CommonServer interface {
	// Ping tests whether the service is working
	Ping(context.Context, *emptypb.Empty) (*Pong, error)
	// Version obtains version information from the target service.
	Version(context.Context, *emptypb.Empty) (*VersionResponse, error)
	mustEmbedUnimplementedCommonServer()
}

// UnimplementedCommonServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCommonServer struct{}

func (UnimplementedCommonServer) Ping(context.Context, *emptypb.Empty) (*Pong, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedCommonServer) Version(context.Context, *emptypb.Empty) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedCommonServer) mustEmbedUnimplementedCommonServer() {}
func (UnimplementedCommonServer) testEmbeddedByValue()                {}

// UnsafeCommonServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommonServer will
// result in compilation errors.
type UnsafeCommonServer interface {
	mustEmbedUnimplementedCommonServer()
}

func RegisterCommonServer(s grpc.ServiceRegistrar, srv CommonServer) {
	// If the following call pancis, it indicates UnimplementedCommonServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Common_ServiceDesc, srv)
}

func _Common_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Common_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServer).Ping(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Common_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Common_Version_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServer).Version(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Common_ServiceDesc is the grpc.ServiceDesc for Common service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Common_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "drivercommon.Common",
	HandlerType: (*CommonServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Common_Ping_Handler,
		},
		{
			MethodName: "Version",
			Handler:    _Common_Version_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "drivercommon/common.proto",
}
