// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.2
// source: ai/v1/agent.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AIAgent_Chat_FullMethodName = "/api.ai.v1.AIAgent/Chat"
)

// AIAgentClient is the client API for AIAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// AI Agent 服务
type AIAgentClient interface {
	// 与 AI 聊天
	Chat(ctx context.Context, in *ChatRequest, opts ...grpc.CallOption) (*ChatReply, error)
}

type aIAgentClient struct {
	cc grpc.ClientConnInterface
}

func NewAIAgentClient(cc grpc.ClientConnInterface) AIAgentClient {
	return &aIAgentClient{cc}
}

func (c *aIAgentClient) Chat(ctx context.Context, in *ChatRequest, opts ...grpc.CallOption) (*ChatReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChatReply)
	err := c.cc.Invoke(ctx, AIAgent_Chat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AIAgentServer is the server API for AIAgent service.
// All implementations must embed UnimplementedAIAgentServer
// for forward compatibility.
//
// AI Agent 服务
type AIAgentServer interface {
	// 与 AI 聊天
	Chat(context.Context, *ChatRequest) (*ChatReply, error)
	mustEmbedUnimplementedAIAgentServer()
}

// UnimplementedAIAgentServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAIAgentServer struct{}

func (UnimplementedAIAgentServer) Chat(context.Context, *ChatRequest) (*ChatReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Chat not implemented")
}
func (UnimplementedAIAgentServer) mustEmbedUnimplementedAIAgentServer() {}
func (UnimplementedAIAgentServer) testEmbeddedByValue()                 {}

// UnsafeAIAgentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AIAgentServer will
// result in compilation errors.
type UnsafeAIAgentServer interface {
	mustEmbedUnimplementedAIAgentServer()
}

func RegisterAIAgentServer(s grpc.ServiceRegistrar, srv AIAgentServer) {
	// If the following call pancis, it indicates UnimplementedAIAgentServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AIAgent_ServiceDesc, srv)
}

func _AIAgent_Chat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AIAgentServer).Chat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AIAgent_Chat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AIAgentServer).Chat(ctx, req.(*ChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AIAgent_ServiceDesc is the grpc.ServiceDesc for AIAgent service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AIAgent_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.ai.v1.AIAgent",
	HandlerType: (*AIAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Chat",
			Handler:    _AIAgent_Chat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ai/v1/agent.proto",
}
