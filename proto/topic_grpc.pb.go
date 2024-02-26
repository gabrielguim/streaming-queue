// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: proto/topic.proto

package proto

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

// TopicClient is the client API for Topic service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TopicClient interface {
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (Topic_SubscribeClient, error)
	Publish(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Void, error)
	Create(ctx context.Context, in *TopicRequest, opts ...grpc.CallOption) (*Void, error)
	Delete(ctx context.Context, in *TopicRequest, opts ...grpc.CallOption) (*Void, error)
}

type topicClient struct {
	cc grpc.ClientConnInterface
}

func NewTopicClient(cc grpc.ClientConnInterface) TopicClient {
	return &topicClient{cc}
}

func (c *topicClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (Topic_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &Topic_ServiceDesc.Streams[0], "/topic.Topic/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &topicSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Topic_SubscribeClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type topicSubscribeClient struct {
	grpc.ClientStream
}

func (x *topicSubscribeClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *topicClient) Publish(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/topic.Topic/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicClient) Create(ctx context.Context, in *TopicRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/topic.Topic/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicClient) Delete(ctx context.Context, in *TopicRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/topic.Topic/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TopicServer is the server API for Topic service.
// All implementations must embed UnimplementedTopicServer
// for forward compatibility
type TopicServer interface {
	Subscribe(*SubscribeRequest, Topic_SubscribeServer) error
	Publish(context.Context, *Message) (*Void, error)
	Create(context.Context, *TopicRequest) (*Void, error)
	Delete(context.Context, *TopicRequest) (*Void, error)
	mustEmbedUnimplementedTopicServer()
}

// UnimplementedTopicServer must be embedded to have forward compatible implementations.
type UnimplementedTopicServer struct {
}

func (UnimplementedTopicServer) Subscribe(*SubscribeRequest, Topic_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedTopicServer) Publish(context.Context, *Message) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedTopicServer) Create(context.Context, *TopicRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedTopicServer) Delete(context.Context, *TopicRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedTopicServer) mustEmbedUnimplementedTopicServer() {}

// UnsafeTopicServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TopicServer will
// result in compilation errors.
type UnsafeTopicServer interface {
	mustEmbedUnimplementedTopicServer()
}

func RegisterTopicServer(s grpc.ServiceRegistrar, srv TopicServer) {
	s.RegisterService(&Topic_ServiceDesc, srv)
}

func _Topic_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TopicServer).Subscribe(m, &topicSubscribeServer{stream})
}

type Topic_SubscribeServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type topicSubscribeServer struct {
	grpc.ServerStream
}

func (x *topicSubscribeServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func _Topic_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/topic.Topic/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServer).Publish(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Topic_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/topic.Topic/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServer).Create(ctx, req.(*TopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Topic_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/topic.Topic/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServer).Delete(ctx, req.(*TopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Topic_ServiceDesc is the grpc.ServiceDesc for Topic service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Topic_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "topic.Topic",
	HandlerType: (*TopicServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _Topic_Publish_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Topic_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Topic_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _Topic_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/topic.proto",
}
