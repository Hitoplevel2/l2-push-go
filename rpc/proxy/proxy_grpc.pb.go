// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.5
// source: proxy.proto

package proxy

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	entity "l2-push-go/rpc/entity"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ProxyClient is the client API for Proxy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProxyClient interface {
	//查询订阅
	GetSubscription(ctx context.Context, in *entity.Void, opts ...grpc.CallOption) (*SubscriptionRsp, error)
	//新增订阅
	AddSubscription(ctx context.Context, in *entity.String, opts ...grpc.CallOption) (*Rsp, error)
	//取消订阅
	DelSubscription(ctx context.Context, in *entity.String, opts ...grpc.CallOption) (*Rsp, error)
	//推送逐笔成交行情数据
	NewTickRecordStream(ctx context.Context, in *entity.Void, opts ...grpc.CallOption) (Proxy_NewTickRecordStreamClient, error)
	//推送逐笔委托行情数据
	NewOrderRecordStream(ctx context.Context, in *entity.Void, opts ...grpc.CallOption) (Proxy_NewOrderRecordStreamClient, error)
	//推送委托队列行情数据
	NewOrderQueueRecordStream(ctx context.Context, in *entity.Void, opts ...grpc.CallOption) (Proxy_NewOrderQueueRecordStreamClient, error)
	//推送股票十档行情行情数据
	NewStockQuoteRecordStream(ctx context.Context, in *entity.Void, opts ...grpc.CallOption) (Proxy_NewStockQuoteRecordStreamClient, error)
}

type proxyClient struct {
	cc grpc.ClientConnInterface
}

func NewProxyClient(cc grpc.ClientConnInterface) ProxyClient {
	return &proxyClient{cc}
}

func (c *proxyClient) GetSubscription(ctx context.Context, in *entity.Void, opts ...grpc.CallOption) (*SubscriptionRsp, error) {
	out := new(SubscriptionRsp)
	err := c.cc.Invoke(ctx, "/sa.rpc.cli.proxy.Proxy/GetSubscription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyClient) AddSubscription(ctx context.Context, in *entity.String, opts ...grpc.CallOption) (*Rsp, error) {
	out := new(Rsp)
	err := c.cc.Invoke(ctx, "/sa.rpc.cli.proxy.Proxy/AddSubscription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyClient) DelSubscription(ctx context.Context, in *entity.String, opts ...grpc.CallOption) (*Rsp, error) {
	out := new(Rsp)
	err := c.cc.Invoke(ctx, "/sa.rpc.cli.proxy.Proxy/DelSubscription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyClient) NewTickRecordStream(ctx context.Context, in *entity.Void, opts ...grpc.CallOption) (Proxy_NewTickRecordStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Proxy_ServiceDesc.Streams[0], "/sa.rpc.cli.proxy.Proxy/NewTickRecordStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &proxyNewTickRecordStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Proxy_NewTickRecordStreamClient interface {
	Recv() (*entity.TickRecord, error)
	grpc.ClientStream
}

type proxyNewTickRecordStreamClient struct {
	grpc.ClientStream
}

func (x *proxyNewTickRecordStreamClient) Recv() (*entity.TickRecord, error) {
	m := new(entity.TickRecord)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *proxyClient) NewOrderRecordStream(ctx context.Context, in *entity.Void, opts ...grpc.CallOption) (Proxy_NewOrderRecordStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Proxy_ServiceDesc.Streams[1], "/sa.rpc.cli.proxy.Proxy/NewOrderRecordStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &proxyNewOrderRecordStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Proxy_NewOrderRecordStreamClient interface {
	Recv() (*entity.OrderRecord, error)
	grpc.ClientStream
}

type proxyNewOrderRecordStreamClient struct {
	grpc.ClientStream
}

func (x *proxyNewOrderRecordStreamClient) Recv() (*entity.OrderRecord, error) {
	m := new(entity.OrderRecord)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *proxyClient) NewOrderQueueRecordStream(ctx context.Context, in *entity.Void, opts ...grpc.CallOption) (Proxy_NewOrderQueueRecordStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Proxy_ServiceDesc.Streams[2], "/sa.rpc.cli.proxy.Proxy/NewOrderQueueRecordStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &proxyNewOrderQueueRecordStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Proxy_NewOrderQueueRecordStreamClient interface {
	Recv() (*entity.OrderQueueRecord, error)
	grpc.ClientStream
}

type proxyNewOrderQueueRecordStreamClient struct {
	grpc.ClientStream
}

func (x *proxyNewOrderQueueRecordStreamClient) Recv() (*entity.OrderQueueRecord, error) {
	m := new(entity.OrderQueueRecord)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *proxyClient) NewStockQuoteRecordStream(ctx context.Context, in *entity.Void, opts ...grpc.CallOption) (Proxy_NewStockQuoteRecordStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Proxy_ServiceDesc.Streams[3], "/sa.rpc.cli.proxy.Proxy/NewStockQuoteRecordStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &proxyNewStockQuoteRecordStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Proxy_NewStockQuoteRecordStreamClient interface {
	Recv() (*entity.StockQuoteRecord, error)
	grpc.ClientStream
}

type proxyNewStockQuoteRecordStreamClient struct {
	grpc.ClientStream
}

func (x *proxyNewStockQuoteRecordStreamClient) Recv() (*entity.StockQuoteRecord, error) {
	m := new(entity.StockQuoteRecord)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProxyServer is the server API for Proxy service.
// All implementations must embed UnimplementedProxyServer
// for forward compatibility
type ProxyServer interface {
	//查询订阅
	GetSubscription(context.Context, *entity.Void) (*SubscriptionRsp, error)
	//新增订阅
	AddSubscription(context.Context, *entity.String) (*Rsp, error)
	//取消订阅
	DelSubscription(context.Context, *entity.String) (*Rsp, error)
	//推送逐笔成交行情数据
	NewTickRecordStream(*entity.Void, Proxy_NewTickRecordStreamServer) error
	//推送逐笔委托行情数据
	NewOrderRecordStream(*entity.Void, Proxy_NewOrderRecordStreamServer) error
	//推送委托队列行情数据
	NewOrderQueueRecordStream(*entity.Void, Proxy_NewOrderQueueRecordStreamServer) error
	//推送股票十档行情行情数据
	NewStockQuoteRecordStream(*entity.Void, Proxy_NewStockQuoteRecordStreamServer) error
	mustEmbedUnimplementedProxyServer()
}

// UnimplementedProxyServer must be embedded to have forward compatible implementations.
type UnimplementedProxyServer struct {
}

func (UnimplementedProxyServer) GetSubscription(context.Context, *entity.Void) (*SubscriptionRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubscription not implemented")
}
func (UnimplementedProxyServer) AddSubscription(context.Context, *entity.String) (*Rsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSubscription not implemented")
}
func (UnimplementedProxyServer) DelSubscription(context.Context, *entity.String) (*Rsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelSubscription not implemented")
}
func (UnimplementedProxyServer) NewTickRecordStream(*entity.Void, Proxy_NewTickRecordStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method NewTickRecordStream not implemented")
}
func (UnimplementedProxyServer) NewOrderRecordStream(*entity.Void, Proxy_NewOrderRecordStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method NewOrderRecordStream not implemented")
}
func (UnimplementedProxyServer) NewOrderQueueRecordStream(*entity.Void, Proxy_NewOrderQueueRecordStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method NewOrderQueueRecordStream not implemented")
}
func (UnimplementedProxyServer) NewStockQuoteRecordStream(*entity.Void, Proxy_NewStockQuoteRecordStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method NewStockQuoteRecordStream not implemented")
}
func (UnimplementedProxyServer) mustEmbedUnimplementedProxyServer() {}

// UnsafeProxyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProxyServer will
// result in compilation errors.
type UnsafeProxyServer interface {
	mustEmbedUnimplementedProxyServer()
}

func RegisterProxyServer(s grpc.ServiceRegistrar, srv ProxyServer) {
	s.RegisterService(&Proxy_ServiceDesc, srv)
}

func _Proxy_GetSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(entity.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServer).GetSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sa.rpc.cli.proxy.Proxy/GetSubscription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServer).GetSubscription(ctx, req.(*entity.Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _Proxy_AddSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(entity.String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServer).AddSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sa.rpc.cli.proxy.Proxy/AddSubscription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServer).AddSubscription(ctx, req.(*entity.String))
	}
	return interceptor(ctx, in, info, handler)
}

func _Proxy_DelSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(entity.String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServer).DelSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sa.rpc.cli.proxy.Proxy/DelSubscription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServer).DelSubscription(ctx, req.(*entity.String))
	}
	return interceptor(ctx, in, info, handler)
}

func _Proxy_NewTickRecordStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(entity.Void)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProxyServer).NewTickRecordStream(m, &proxyNewTickRecordStreamServer{stream})
}

type Proxy_NewTickRecordStreamServer interface {
	Send(*entity.TickRecord) error
	grpc.ServerStream
}

type proxyNewTickRecordStreamServer struct {
	grpc.ServerStream
}

func (x *proxyNewTickRecordStreamServer) Send(m *entity.TickRecord) error {
	return x.ServerStream.SendMsg(m)
}

func _Proxy_NewOrderRecordStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(entity.Void)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProxyServer).NewOrderRecordStream(m, &proxyNewOrderRecordStreamServer{stream})
}

type Proxy_NewOrderRecordStreamServer interface {
	Send(*entity.OrderRecord) error
	grpc.ServerStream
}

type proxyNewOrderRecordStreamServer struct {
	grpc.ServerStream
}

func (x *proxyNewOrderRecordStreamServer) Send(m *entity.OrderRecord) error {
	return x.ServerStream.SendMsg(m)
}

func _Proxy_NewOrderQueueRecordStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(entity.Void)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProxyServer).NewOrderQueueRecordStream(m, &proxyNewOrderQueueRecordStreamServer{stream})
}

type Proxy_NewOrderQueueRecordStreamServer interface {
	Send(*entity.OrderQueueRecord) error
	grpc.ServerStream
}

type proxyNewOrderQueueRecordStreamServer struct {
	grpc.ServerStream
}

func (x *proxyNewOrderQueueRecordStreamServer) Send(m *entity.OrderQueueRecord) error {
	return x.ServerStream.SendMsg(m)
}

func _Proxy_NewStockQuoteRecordStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(entity.Void)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProxyServer).NewStockQuoteRecordStream(m, &proxyNewStockQuoteRecordStreamServer{stream})
}

type Proxy_NewStockQuoteRecordStreamServer interface {
	Send(*entity.StockQuoteRecord) error
	grpc.ServerStream
}

type proxyNewStockQuoteRecordStreamServer struct {
	grpc.ServerStream
}

func (x *proxyNewStockQuoteRecordStreamServer) Send(m *entity.StockQuoteRecord) error {
	return x.ServerStream.SendMsg(m)
}

// Proxy_ServiceDesc is the grpc.ServiceDesc for Proxy service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Proxy_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sa.rpc.cli.proxy.Proxy",
	HandlerType: (*ProxyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSubscription",
			Handler:    _Proxy_GetSubscription_Handler,
		},
		{
			MethodName: "AddSubscription",
			Handler:    _Proxy_AddSubscription_Handler,
		},
		{
			MethodName: "DelSubscription",
			Handler:    _Proxy_DelSubscription_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "NewTickRecordStream",
			Handler:       _Proxy_NewTickRecordStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "NewOrderRecordStream",
			Handler:       _Proxy_NewOrderRecordStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "NewOrderQueueRecordStream",
			Handler:       _Proxy_NewOrderQueueRecordStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "NewStockQuoteRecordStream",
			Handler:       _Proxy_NewStockQuoteRecordStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proxy.proto",
}