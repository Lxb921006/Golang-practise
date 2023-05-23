// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: streamrpc.proto

package streamrpc

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

// StreamRpcServiceClient is the client API for StreamRpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamRpcServiceClient interface {
	SayHelloWorld(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (StreamRpcService_SayHelloWorldClient, error)
}

type streamRpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamRpcServiceClient(cc grpc.ClientConnInterface) StreamRpcServiceClient {
	return &streamRpcServiceClient{cc}
}

func (c *streamRpcServiceClient) SayHelloWorld(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (StreamRpcService_SayHelloWorldClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamRpcService_ServiceDesc.Streams[0], "/streamrpc.StreamRpcService/SayHelloWorld", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamRpcServiceSayHelloWorldClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StreamRpcService_SayHelloWorldClient interface {
	Recv() (*StreamReply, error)
	grpc.ClientStream
}

type streamRpcServiceSayHelloWorldClient struct {
	grpc.ClientStream
}

func (x *streamRpcServiceSayHelloWorldClient) Recv() (*StreamReply, error) {
	m := new(StreamReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamRpcServiceServer is the server API for StreamRpcService service.
// All implementations must embed UnimplementedStreamRpcServiceServer
// for forward compatibility
type StreamRpcServiceServer interface {
	SayHelloWorld(*StreamRequest, StreamRpcService_SayHelloWorldServer) error
	mustEmbedUnimplementedStreamRpcServiceServer()
}

// UnimplementedStreamRpcServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStreamRpcServiceServer struct {
}

func (UnimplementedStreamRpcServiceServer) SayHelloWorld(*StreamRequest, StreamRpcService_SayHelloWorldServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloWorld not implemented")
}
func (UnimplementedStreamRpcServiceServer) mustEmbedUnimplementedStreamRpcServiceServer() {}

// UnsafeStreamRpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamRpcServiceServer will
// result in compilation errors.
type UnsafeStreamRpcServiceServer interface {
	mustEmbedUnimplementedStreamRpcServiceServer()
}

func RegisterStreamRpcServiceServer(s grpc.ServiceRegistrar, srv StreamRpcServiceServer) {
	s.RegisterService(&StreamRpcService_ServiceDesc, srv)
}

func _StreamRpcService_SayHelloWorld_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamRpcServiceServer).SayHelloWorld(m, &streamRpcServiceSayHelloWorldServer{stream})
}

type StreamRpcService_SayHelloWorldServer interface {
	Send(*StreamReply) error
	grpc.ServerStream
}

type streamRpcServiceSayHelloWorldServer struct {
	grpc.ServerStream
}

func (x *streamRpcServiceSayHelloWorldServer) Send(m *StreamReply) error {
	return x.ServerStream.SendMsg(m)
}

// StreamRpcService_ServiceDesc is the grpc.ServiceDesc for StreamRpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamRpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "streamrpc.StreamRpcService",
	HandlerType: (*StreamRpcServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayHelloWorld",
			Handler:       _StreamRpcService_SayHelloWorld_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "streamrpc.proto",
}
