// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// ExchangeClient is the client API for Exchange service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExchangeClient interface {
	Connection(ctx context.Context, opts ...grpc.CallOption) (Exchange_ConnectionClient, error)
}

type exchangeClient struct {
	cc grpc.ClientConnInterface
}

func NewExchangeClient(cc grpc.ClientConnInterface) ExchangeClient {
	return &exchangeClient{cc}
}

func (c *exchangeClient) Connection(ctx context.Context, opts ...grpc.CallOption) (Exchange_ConnectionClient, error) {
	stream, err := c.cc.NewStream(ctx, &Exchange_ServiceDesc.Streams[0], "/protocol.Exchange/Connection", opts...)
	if err != nil {
		return nil, err
	}
	x := &exchangeConnectionClient{stream}
	return x, nil
}

type Exchange_ConnectionClient interface {
	Send(*InMessage) error
	Recv() (*OutMessage, error)
	grpc.ClientStream
}

type exchangeConnectionClient struct {
	grpc.ClientStream
}

func (x *exchangeConnectionClient) Send(m *InMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *exchangeConnectionClient) Recv() (*OutMessage, error) {
	m := new(OutMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ExchangeServer is the server API for Exchange service.
// All implementations must embed UnimplementedExchangeServer
// for forward compatibility
type ExchangeServer interface {
	Connection(Exchange_ConnectionServer) error
}

// UnimplementedExchangeServer must be embedded to have forward compatible implementations.
type UnimplementedExchangeServer struct {
}

func (UnimplementedExchangeServer) Connection(Exchange_ConnectionServer) error {
	return status.Errorf(codes.Unimplemented, "method Connection not implemented")
}
func (UnimplementedExchangeServer) mustEmbedUnimplementedExchangeServer() {}

// UnsafeExchangeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExchangeServer will
// result in compilation errors.
type UnsafeExchangeServer interface {
	mustEmbedUnimplementedExchangeServer()
}

func RegisterExchangeServer(s grpc.ServiceRegistrar, srv ExchangeServer) {
	s.RegisterService(&Exchange_ServiceDesc, srv)
}

func _Exchange_Connection_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ExchangeServer).Connection(&exchangeConnectionServer{stream})
}

type Exchange_ConnectionServer interface {
	Send(*OutMessage) error
	Recv() (*InMessage, error)
	grpc.ServerStream
}

type exchangeConnectionServer struct {
	grpc.ServerStream
}

func (x *exchangeConnectionServer) Send(m *OutMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *exchangeConnectionServer) Recv() (*InMessage, error) {
	m := new(InMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Exchange_ServiceDesc is the grpc.ServiceDesc for Exchange service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Exchange_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.Exchange",
	HandlerType: (*ExchangeServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connection",
			Handler:       _Exchange_Connection_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "exchange.proto",
}