// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: products.proto

package protocol

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

// ProductsClient is the client API for Products service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductsClient interface {
	GetProducts(ctx context.Context, in *SelectRequest, opts ...grpc.CallOption) (*ProductsReply, error)
	GetProduct(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*Product, error)
	CreateFromShop(ctx context.Context, in *ProductLink, opts ...grpc.CallOption) (*BoolReply, error)
	LinkToShop(ctx context.Context, in *ProductLink, opts ...grpc.CallOption) (*BoolReply, error)
	UnlinkToShop(ctx context.Context, in *ProductLink, opts ...grpc.CallOption) (*BoolReply, error)
}

type productsClient struct {
	cc grpc.ClientConnInterface
}

func NewProductsClient(cc grpc.ClientConnInterface) ProductsClient {
	return &productsClient{cc}
}

func (c *productsClient) GetProducts(ctx context.Context, in *SelectRequest, opts ...grpc.CallOption) (*ProductsReply, error) {
	out := new(ProductsReply)
	err := c.cc.Invoke(ctx, "/protocol.Products/GetProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsClient) GetProduct(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, "/protocol.Products/GetProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsClient) CreateFromShop(ctx context.Context, in *ProductLink, opts ...grpc.CallOption) (*BoolReply, error) {
	out := new(BoolReply)
	err := c.cc.Invoke(ctx, "/protocol.Products/CreateFromShop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsClient) LinkToShop(ctx context.Context, in *ProductLink, opts ...grpc.CallOption) (*BoolReply, error) {
	out := new(BoolReply)
	err := c.cc.Invoke(ctx, "/protocol.Products/LinkToShop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsClient) UnlinkToShop(ctx context.Context, in *ProductLink, opts ...grpc.CallOption) (*BoolReply, error) {
	out := new(BoolReply)
	err := c.cc.Invoke(ctx, "/protocol.Products/UnlinkToShop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductsServer is the server API for Products service.
// All implementations must embed UnimplementedProductsServer
// for forward compatibility
type ProductsServer interface {
	GetProducts(context.Context, *SelectRequest) (*ProductsReply, error)
	GetProduct(context.Context, *ProductRequest) (*Product, error)
	CreateFromShop(context.Context, *ProductLink) (*BoolReply, error)
	LinkToShop(context.Context, *ProductLink) (*BoolReply, error)
	UnlinkToShop(context.Context, *ProductLink) (*BoolReply, error)
	mustEmbedUnimplementedProductsServer()
}

// UnimplementedProductsServer must be embedded to have forward compatible implementations.
type UnimplementedProductsServer struct {
}

func (UnimplementedProductsServer) GetProducts(context.Context, *SelectRequest) (*ProductsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProducts not implemented")
}
func (UnimplementedProductsServer) GetProduct(context.Context, *ProductRequest) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedProductsServer) CreateFromShop(context.Context, *ProductLink) (*BoolReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFromShop not implemented")
}
func (UnimplementedProductsServer) LinkToShop(context.Context, *ProductLink) (*BoolReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LinkToShop not implemented")
}
func (UnimplementedProductsServer) UnlinkToShop(context.Context, *ProductLink) (*BoolReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnlinkToShop not implemented")
}
func (UnimplementedProductsServer) mustEmbedUnimplementedProductsServer() {}

// UnsafeProductsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductsServer will
// result in compilation errors.
type UnsafeProductsServer interface {
	mustEmbedUnimplementedProductsServer()
}

func RegisterProductsServer(s grpc.ServiceRegistrar, srv ProductsServer) {
	s.RegisterService(&Products_ServiceDesc, srv)
}

func _Products_GetProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServer).GetProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Products/GetProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServer).GetProducts(ctx, req.(*SelectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Products_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Products/GetProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServer).GetProduct(ctx, req.(*ProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Products_CreateFromShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductLink)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServer).CreateFromShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Products/CreateFromShop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServer).CreateFromShop(ctx, req.(*ProductLink))
	}
	return interceptor(ctx, in, info, handler)
}

func _Products_LinkToShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductLink)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServer).LinkToShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Products/LinkToShop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServer).LinkToShop(ctx, req.(*ProductLink))
	}
	return interceptor(ctx, in, info, handler)
}

func _Products_UnlinkToShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductLink)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServer).UnlinkToShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Products/UnlinkToShop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServer).UnlinkToShop(ctx, req.(*ProductLink))
	}
	return interceptor(ctx, in, info, handler)
}

// Products_ServiceDesc is the grpc.ServiceDesc for Products service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Products_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.Products",
	HandlerType: (*ProductsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProducts",
			Handler:    _Products_GetProducts_Handler,
		},
		{
			MethodName: "GetProduct",
			Handler:    _Products_GetProduct_Handler,
		},
		{
			MethodName: "CreateFromShop",
			Handler:    _Products_CreateFromShop_Handler,
		},
		{
			MethodName: "LinkToShop",
			Handler:    _Products_LinkToShop_Handler,
		},
		{
			MethodName: "UnlinkToShop",
			Handler:    _Products_UnlinkToShop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "products.proto",
}
