// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.14.0
// source: productcallback/productcallback.proto

package productcallback

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ProductCallBackServiceClient is the client API for ProductCallBackService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductCallBackServiceClient interface {
	// 创建设备回调
	CreateProductCallback(ctx context.Context, in *CreateProductCallbackRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 更新设备回调
	UpdateProductCallback(ctx context.Context, in *UpdateProductCallbackRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 删除设备回调
	DeleteProductCallback(ctx context.Context, in *DeleteProductCallbackRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type productCallBackServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductCallBackServiceClient(cc grpc.ClientConnInterface) ProductCallBackServiceClient {
	return &productCallBackServiceClient{cc}
}

func (c *productCallBackServiceClient) CreateProductCallback(ctx context.Context, in *CreateProductCallbackRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/productcallback.ProductCallBackService/CreateProductCallback", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCallBackServiceClient) UpdateProductCallback(ctx context.Context, in *UpdateProductCallbackRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/productcallback.ProductCallBackService/UpdateProductCallback", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCallBackServiceClient) DeleteProductCallback(ctx context.Context, in *DeleteProductCallbackRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/productcallback.ProductCallBackService/DeleteProductCallback", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductCallBackServiceServer is the server API for ProductCallBackService service.
// All implementations must embed UnimplementedProductCallBackServiceServer
// for forward compatibility
type ProductCallBackServiceServer interface {
	// 创建设备回调
	CreateProductCallback(context.Context, *CreateProductCallbackRequest) (*emptypb.Empty, error)
	// 更新设备回调
	UpdateProductCallback(context.Context, *UpdateProductCallbackRequest) (*emptypb.Empty, error)
	// 删除设备回调
	DeleteProductCallback(context.Context, *DeleteProductCallbackRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedProductCallBackServiceServer()
}

// UnimplementedProductCallBackServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProductCallBackServiceServer struct {
}

func (UnimplementedProductCallBackServiceServer) CreateProductCallback(context.Context, *CreateProductCallbackRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProductCallback not implemented")
}
func (UnimplementedProductCallBackServiceServer) UpdateProductCallback(context.Context, *UpdateProductCallbackRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProductCallback not implemented")
}
func (UnimplementedProductCallBackServiceServer) DeleteProductCallback(context.Context, *DeleteProductCallbackRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProductCallback not implemented")
}
func (UnimplementedProductCallBackServiceServer) mustEmbedUnimplementedProductCallBackServiceServer() {
}

// UnsafeProductCallBackServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductCallBackServiceServer will
// result in compilation errors.
type UnsafeProductCallBackServiceServer interface {
	mustEmbedUnimplementedProductCallBackServiceServer()
}

func RegisterProductCallBackServiceServer(s grpc.ServiceRegistrar, srv ProductCallBackServiceServer) {
	s.RegisterService(&ProductCallBackService_ServiceDesc, srv)
}

func _ProductCallBackService_CreateProductCallback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductCallbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCallBackServiceServer).CreateProductCallback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/productcallback.ProductCallBackService/CreateProductCallback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCallBackServiceServer).CreateProductCallback(ctx, req.(*CreateProductCallbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductCallBackService_UpdateProductCallback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProductCallbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCallBackServiceServer).UpdateProductCallback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/productcallback.ProductCallBackService/UpdateProductCallback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCallBackServiceServer).UpdateProductCallback(ctx, req.(*UpdateProductCallbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductCallBackService_DeleteProductCallback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProductCallbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCallBackServiceServer).DeleteProductCallback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/productcallback.ProductCallBackService/DeleteProductCallback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCallBackServiceServer).DeleteProductCallback(ctx, req.(*DeleteProductCallbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductCallBackService_ServiceDesc is the grpc.ServiceDesc for ProductCallBackService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductCallBackService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "productcallback.ProductCallBackService",
	HandlerType: (*ProductCallBackServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProductCallback",
			Handler:    _ProductCallBackService_CreateProductCallback_Handler,
		},
		{
			MethodName: "UpdateProductCallback",
			Handler:    _ProductCallBackService_UpdateProductCallback_Handler,
		},
		{
			MethodName: "DeleteProductCallback",
			Handler:    _ProductCallBackService_DeleteProductCallback_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "productcallback/productcallback.proto",
}
