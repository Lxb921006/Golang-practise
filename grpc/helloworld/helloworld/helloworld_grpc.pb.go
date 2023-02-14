// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: helloworld.proto

package __

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

// TestGrpcHelloWorldClient is the client API for TestGrpcHelloWorld service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestGrpcHelloWorldClient interface {
	SayHelloWorld(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type testGrpcHelloWorldClient struct {
	cc grpc.ClientConnInterface
}

func NewTestGrpcHelloWorldClient(cc grpc.ClientConnInterface) TestGrpcHelloWorldClient {
	return &testGrpcHelloWorldClient{cc}
}

func (c *testGrpcHelloWorldClient) SayHelloWorld(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/helloworld.TestGrpcHelloWorld/SayHelloWorld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestGrpcHelloWorldServer is the server API for TestGrpcHelloWorld service.
// All implementations must embed UnimplementedTestGrpcHelloWorldServer
// for forward compatibility
type TestGrpcHelloWorldServer interface {
	SayHelloWorld(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedTestGrpcHelloWorldServer()
}

// UnimplementedTestGrpcHelloWorldServer must be embedded to have forward compatible implementations.
type UnimplementedTestGrpcHelloWorldServer struct {
}

func (UnimplementedTestGrpcHelloWorldServer) SayHelloWorld(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHelloWorld not implemented")
}
func (UnimplementedTestGrpcHelloWorldServer) mustEmbedUnimplementedTestGrpcHelloWorldServer() {}

// UnsafeTestGrpcHelloWorldServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestGrpcHelloWorldServer will
// result in compilation errors.
type UnsafeTestGrpcHelloWorldServer interface {
	mustEmbedUnimplementedTestGrpcHelloWorldServer()
}

func RegisterTestGrpcHelloWorldServer(s grpc.ServiceRegistrar, srv TestGrpcHelloWorldServer) {
	s.RegisterService(&TestGrpcHelloWorld_ServiceDesc, srv)
}

func _TestGrpcHelloWorld_SayHelloWorld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestGrpcHelloWorldServer).SayHelloWorld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.TestGrpcHelloWorld/SayHelloWorld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestGrpcHelloWorldServer).SayHelloWorld(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TestGrpcHelloWorld_ServiceDesc is the grpc.ServiceDesc for TestGrpcHelloWorld service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TestGrpcHelloWorld_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.TestGrpcHelloWorld",
	HandlerType: (*TestGrpcHelloWorldServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHelloWorld",
			Handler:    _TestGrpcHelloWorld_SayHelloWorld_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld.proto",
}
