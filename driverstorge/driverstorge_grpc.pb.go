// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: driverstorge/driverstorge.proto

package driverstorage

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
	DriverStorage_All_FullMethodName    = "/driverstorage.DriverStorage/All"
	DriverStorage_Get_FullMethodName    = "/driverstorage.DriverStorage/Get"
	DriverStorage_Put_FullMethodName    = "/driverstorage.DriverStorage/Put"
	DriverStorage_Delete_FullMethodName = "/driverstorage.DriverStorage/Delete"
)

// DriverStorageClient is the client API for DriverStorage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DriverStorageClient interface {
	// all
	All(ctx context.Context, in *AllReq, opts ...grpc.CallOption) (*KVs, error)
	// get
	Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*KVs, error)
	// put
	Put(ctx context.Context, in *PutReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// delete
	Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type driverStorageClient struct {
	cc grpc.ClientConnInterface
}

func NewDriverStorageClient(cc grpc.ClientConnInterface) DriverStorageClient {
	return &driverStorageClient{cc}
}

func (c *driverStorageClient) All(ctx context.Context, in *AllReq, opts ...grpc.CallOption) (*KVs, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(KVs)
	err := c.cc.Invoke(ctx, DriverStorage_All_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverStorageClient) Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*KVs, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(KVs)
	err := c.cc.Invoke(ctx, DriverStorage_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverStorageClient) Put(ctx context.Context, in *PutReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DriverStorage_Put_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverStorageClient) Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DriverStorage_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DriverStorageServer is the server API for DriverStorage service.
// All implementations must embed UnimplementedDriverStorageServer
// for forward compatibility.
type DriverStorageServer interface {
	// all
	All(context.Context, *AllReq) (*KVs, error)
	// get
	Get(context.Context, *GetReq) (*KVs, error)
	// put
	Put(context.Context, *PutReq) (*emptypb.Empty, error)
	// delete
	Delete(context.Context, *DeleteReq) (*emptypb.Empty, error)
	mustEmbedUnimplementedDriverStorageServer()
}

// UnimplementedDriverStorageServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDriverStorageServer struct{}

func (UnimplementedDriverStorageServer) All(context.Context, *AllReq) (*KVs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method All not implemented")
}
func (UnimplementedDriverStorageServer) Get(context.Context, *GetReq) (*KVs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedDriverStorageServer) Put(context.Context, *PutReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (UnimplementedDriverStorageServer) Delete(context.Context, *DeleteReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedDriverStorageServer) mustEmbedUnimplementedDriverStorageServer() {}
func (UnimplementedDriverStorageServer) testEmbeddedByValue()                       {}

// UnsafeDriverStorageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DriverStorageServer will
// result in compilation errors.
type UnsafeDriverStorageServer interface {
	mustEmbedUnimplementedDriverStorageServer()
}

func RegisterDriverStorageServer(s grpc.ServiceRegistrar, srv DriverStorageServer) {
	// If the following call pancis, it indicates UnimplementedDriverStorageServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DriverStorage_ServiceDesc, srv)
}

func _DriverStorage_All_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverStorageServer).All(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DriverStorage_All_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverStorageServer).All(ctx, req.(*AllReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverStorage_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverStorageServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DriverStorage_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverStorageServer).Get(ctx, req.(*GetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverStorage_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverStorageServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DriverStorage_Put_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverStorageServer).Put(ctx, req.(*PutReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverStorage_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverStorageServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DriverStorage_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverStorageServer).Delete(ctx, req.(*DeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// DriverStorage_ServiceDesc is the grpc.ServiceDesc for DriverStorage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DriverStorage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "driverstorage.DriverStorage",
	HandlerType: (*DriverStorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "All",
			Handler:    _DriverStorage_All_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _DriverStorage_Get_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _DriverStorage_Put_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _DriverStorage_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "driverstorge/driverstorge.proto",
}
