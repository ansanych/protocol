// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: wb.proto

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

// WBClient is the client API for WB service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WBClient interface {
	SetWBAuth(ctx context.Context, in *SetWBAuthRequest, opts ...grpc.CallOption) (*BoolReply, error)
	GetWBAuth(ctx context.Context, in *Access, opts ...grpc.CallOption) (*WBAuth, error)
	ErrorWBAuth(ctx context.Context, in *Access, opts ...grpc.CallOption) (*BoolReply, error)
	GetWidgetDataWB(ctx context.Context, in *Access, opts ...grpc.CallOption) (*WidgetReply, error)
	GetProductListWB(ctx context.Context, in *SelectRequest, opts ...grpc.CallOption) (*ShopProductList, error)
	GetProductWB(ctx context.Context, in *ShopProductRequest, opts ...grpc.CallOption) (*ShopProduct, error)
	SearchShopProductWB(ctx context.Context, in *SearchShopProductRequest, opts ...grpc.CallOption) (*ShopProductList, error)
	GetUnsortedCountWB(ctx context.Context, in *Access, opts ...grpc.CallOption) (*ProductCount, error)
	GetUnsortedListWB(ctx context.Context, in *SelectRequest, opts ...grpc.CallOption) (*ShopProductList, error)
}

type wBClient struct {
	cc grpc.ClientConnInterface
}

func NewWBClient(cc grpc.ClientConnInterface) WBClient {
	return &wBClient{cc}
}

func (c *wBClient) SetWBAuth(ctx context.Context, in *SetWBAuthRequest, opts ...grpc.CallOption) (*BoolReply, error) {
	out := new(BoolReply)
	err := c.cc.Invoke(ctx, "/protocol.WB/SetWBAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wBClient) GetWBAuth(ctx context.Context, in *Access, opts ...grpc.CallOption) (*WBAuth, error) {
	out := new(WBAuth)
	err := c.cc.Invoke(ctx, "/protocol.WB/GetWBAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wBClient) ErrorWBAuth(ctx context.Context, in *Access, opts ...grpc.CallOption) (*BoolReply, error) {
	out := new(BoolReply)
	err := c.cc.Invoke(ctx, "/protocol.WB/ErrorWBAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wBClient) GetWidgetDataWB(ctx context.Context, in *Access, opts ...grpc.CallOption) (*WidgetReply, error) {
	out := new(WidgetReply)
	err := c.cc.Invoke(ctx, "/protocol.WB/GetWidgetDataWB", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wBClient) GetProductListWB(ctx context.Context, in *SelectRequest, opts ...grpc.CallOption) (*ShopProductList, error) {
	out := new(ShopProductList)
	err := c.cc.Invoke(ctx, "/protocol.WB/GetProductListWB", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wBClient) GetProductWB(ctx context.Context, in *ShopProductRequest, opts ...grpc.CallOption) (*ShopProduct, error) {
	out := new(ShopProduct)
	err := c.cc.Invoke(ctx, "/protocol.WB/GetProductWB", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wBClient) SearchShopProductWB(ctx context.Context, in *SearchShopProductRequest, opts ...grpc.CallOption) (*ShopProductList, error) {
	out := new(ShopProductList)
	err := c.cc.Invoke(ctx, "/protocol.WB/SearchShopProductWB", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wBClient) GetUnsortedCountWB(ctx context.Context, in *Access, opts ...grpc.CallOption) (*ProductCount, error) {
	out := new(ProductCount)
	err := c.cc.Invoke(ctx, "/protocol.WB/GetUnsortedCountWB", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wBClient) GetUnsortedListWB(ctx context.Context, in *SelectRequest, opts ...grpc.CallOption) (*ShopProductList, error) {
	out := new(ShopProductList)
	err := c.cc.Invoke(ctx, "/protocol.WB/GetUnsortedListWB", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WBServer is the server API for WB service.
// All implementations must embed UnimplementedWBServer
// for forward compatibility
type WBServer interface {
	SetWBAuth(context.Context, *SetWBAuthRequest) (*BoolReply, error)
	GetWBAuth(context.Context, *Access) (*WBAuth, error)
	ErrorWBAuth(context.Context, *Access) (*BoolReply, error)
	GetWidgetDataWB(context.Context, *Access) (*WidgetReply, error)
	GetProductListWB(context.Context, *SelectRequest) (*ShopProductList, error)
	GetProductWB(context.Context, *ShopProductRequest) (*ShopProduct, error)
	SearchShopProductWB(context.Context, *SearchShopProductRequest) (*ShopProductList, error)
	GetUnsortedCountWB(context.Context, *Access) (*ProductCount, error)
	GetUnsortedListWB(context.Context, *SelectRequest) (*ShopProductList, error)
	mustEmbedUnimplementedWBServer()
}

// UnimplementedWBServer must be embedded to have forward compatible implementations.
type UnimplementedWBServer struct {
}

func (UnimplementedWBServer) SetWBAuth(context.Context, *SetWBAuthRequest) (*BoolReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetWBAuth not implemented")
}
func (UnimplementedWBServer) GetWBAuth(context.Context, *Access) (*WBAuth, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWBAuth not implemented")
}
func (UnimplementedWBServer) ErrorWBAuth(context.Context, *Access) (*BoolReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ErrorWBAuth not implemented")
}
func (UnimplementedWBServer) GetWidgetDataWB(context.Context, *Access) (*WidgetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWidgetDataWB not implemented")
}
func (UnimplementedWBServer) GetProductListWB(context.Context, *SelectRequest) (*ShopProductList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductListWB not implemented")
}
func (UnimplementedWBServer) GetProductWB(context.Context, *ShopProductRequest) (*ShopProduct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductWB not implemented")
}
func (UnimplementedWBServer) SearchShopProductWB(context.Context, *SearchShopProductRequest) (*ShopProductList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchShopProductWB not implemented")
}
func (UnimplementedWBServer) GetUnsortedCountWB(context.Context, *Access) (*ProductCount, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUnsortedCountWB not implemented")
}
func (UnimplementedWBServer) GetUnsortedListWB(context.Context, *SelectRequest) (*ShopProductList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUnsortedListWB not implemented")
}
func (UnimplementedWBServer) mustEmbedUnimplementedWBServer() {}

// UnsafeWBServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WBServer will
// result in compilation errors.
type UnsafeWBServer interface {
	mustEmbedUnimplementedWBServer()
}

func RegisterWBServer(s grpc.ServiceRegistrar, srv WBServer) {
	s.RegisterService(&WB_ServiceDesc, srv)
}

func _WB_SetWBAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetWBAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WBServer).SetWBAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.WB/SetWBAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WBServer).SetWBAuth(ctx, req.(*SetWBAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WB_GetWBAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Access)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WBServer).GetWBAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.WB/GetWBAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WBServer).GetWBAuth(ctx, req.(*Access))
	}
	return interceptor(ctx, in, info, handler)
}

func _WB_ErrorWBAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Access)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WBServer).ErrorWBAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.WB/ErrorWBAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WBServer).ErrorWBAuth(ctx, req.(*Access))
	}
	return interceptor(ctx, in, info, handler)
}

func _WB_GetWidgetDataWB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Access)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WBServer).GetWidgetDataWB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.WB/GetWidgetDataWB",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WBServer).GetWidgetDataWB(ctx, req.(*Access))
	}
	return interceptor(ctx, in, info, handler)
}

func _WB_GetProductListWB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WBServer).GetProductListWB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.WB/GetProductListWB",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WBServer).GetProductListWB(ctx, req.(*SelectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WB_GetProductWB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShopProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WBServer).GetProductWB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.WB/GetProductWB",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WBServer).GetProductWB(ctx, req.(*ShopProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WB_SearchShopProductWB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchShopProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WBServer).SearchShopProductWB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.WB/SearchShopProductWB",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WBServer).SearchShopProductWB(ctx, req.(*SearchShopProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WB_GetUnsortedCountWB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Access)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WBServer).GetUnsortedCountWB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.WB/GetUnsortedCountWB",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WBServer).GetUnsortedCountWB(ctx, req.(*Access))
	}
	return interceptor(ctx, in, info, handler)
}

func _WB_GetUnsortedListWB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WBServer).GetUnsortedListWB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.WB/GetUnsortedListWB",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WBServer).GetUnsortedListWB(ctx, req.(*SelectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WB_ServiceDesc is the grpc.ServiceDesc for WB service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WB_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.WB",
	HandlerType: (*WBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetWBAuth",
			Handler:    _WB_SetWBAuth_Handler,
		},
		{
			MethodName: "GetWBAuth",
			Handler:    _WB_GetWBAuth_Handler,
		},
		{
			MethodName: "ErrorWBAuth",
			Handler:    _WB_ErrorWBAuth_Handler,
		},
		{
			MethodName: "GetWidgetDataWB",
			Handler:    _WB_GetWidgetDataWB_Handler,
		},
		{
			MethodName: "GetProductListWB",
			Handler:    _WB_GetProductListWB_Handler,
		},
		{
			MethodName: "GetProductWB",
			Handler:    _WB_GetProductWB_Handler,
		},
		{
			MethodName: "SearchShopProductWB",
			Handler:    _WB_SearchShopProductWB_Handler,
		},
		{
			MethodName: "GetUnsortedCountWB",
			Handler:    _WB_GetUnsortedCountWB_Handler,
		},
		{
			MethodName: "GetUnsortedListWB",
			Handler:    _WB_GetUnsortedListWB_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wb.proto",
}