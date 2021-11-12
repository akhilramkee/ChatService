// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package chatserver

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

// ChatServicesClient is the client API for ChatServices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServicesClient interface {
	MessageChannel(ctx context.Context, opts ...grpc.CallOption) (ChatServices_MessageChannelClient, error)
	SendStatus(ctx context.Context, in *MessageStatus, opts ...grpc.CallOption) (*MessageStatus, error)
}

type chatServicesClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServicesClient(cc grpc.ClientConnInterface) ChatServicesClient {
	return &chatServicesClient{cc}
}

func (c *chatServicesClient) MessageChannel(ctx context.Context, opts ...grpc.CallOption) (ChatServices_MessageChannelClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatServices_ServiceDesc.Streams[0], "/chatserver.ChatServices/MessageChannel", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServicesMessageChannelClient{stream}
	return x, nil
}

type ChatServices_MessageChannelClient interface {
	Send(*MessageComponent) error
	Recv() (*MessageComponent, error)
	grpc.ClientStream
}

type chatServicesMessageChannelClient struct {
	grpc.ClientStream
}

func (x *chatServicesMessageChannelClient) Send(m *MessageComponent) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatServicesMessageChannelClient) Recv() (*MessageComponent, error) {
	m := new(MessageComponent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatServicesClient) SendStatus(ctx context.Context, in *MessageStatus, opts ...grpc.CallOption) (*MessageStatus, error) {
	out := new(MessageStatus)
	err := c.cc.Invoke(ctx, "/chatserver.ChatServices/SendStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServicesServer is the server API for ChatServices service.
// All implementations must embed UnimplementedChatServicesServer
// for forward compatibility
type ChatServicesServer interface {
	MessageChannel(ChatServices_MessageChannelServer) error
	SendStatus(context.Context, *MessageStatus) (*MessageStatus, error)
	mustEmbedUnimplementedChatServicesServer()
}

// UnimplementedChatServicesServer must be embedded to have forward compatible implementations.
type UnimplementedChatServicesServer struct {
}

func (UnimplementedChatServicesServer) MessageChannel(ChatServices_MessageChannelServer) error {
	return status.Errorf(codes.Unimplemented, "method MessageChannel not implemented")
}
func (UnimplementedChatServicesServer) SendStatus(context.Context, *MessageStatus) (*MessageStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendStatus not implemented")
}
func (UnimplementedChatServicesServer) mustEmbedUnimplementedChatServicesServer() {}

// UnsafeChatServicesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServicesServer will
// result in compilation errors.
type UnsafeChatServicesServer interface {
	mustEmbedUnimplementedChatServicesServer()
}

func RegisterChatServicesServer(s grpc.ServiceRegistrar, srv ChatServicesServer) {
	s.RegisterService(&ChatServices_ServiceDesc, srv)
}

func _ChatServices_MessageChannel_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServicesServer).MessageChannel(&chatServicesMessageChannelServer{stream})
}

type ChatServices_MessageChannelServer interface {
	Send(*MessageComponent) error
	Recv() (*MessageComponent, error)
	grpc.ServerStream
}

type chatServicesMessageChannelServer struct {
	grpc.ServerStream
}

func (x *chatServicesMessageChannelServer) Send(m *MessageComponent) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatServicesMessageChannelServer) Recv() (*MessageComponent, error) {
	m := new(MessageComponent)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ChatServices_SendStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageStatus)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServicesServer).SendStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatserver.ChatServices/SendStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServicesServer).SendStatus(ctx, req.(*MessageStatus))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatServices_ServiceDesc is the grpc.ServiceDesc for ChatServices service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatServices_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chatserver.ChatServices",
	HandlerType: (*ChatServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendStatus",
			Handler:    _ChatServices_SendStatus_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MessageChannel",
			Handler:       _ChatServices_MessageChannel_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/chat.proto",
}
