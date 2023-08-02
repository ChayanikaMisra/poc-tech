// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: hello_world.proto

package pb

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

const (
	HelloWorldService_Greeting_FullMethodName = "/HelloWorldService/Greeting"
)

// HelloWorldServiceClient is the client API for HelloWorldService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelloWorldServiceClient interface {
	Greeting(ctx context.Context, in *HelloWorldServiceRequest, opts ...grpc.CallOption) (*HelloWorldServiceReply, error)
}

type helloWorldServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloWorldServiceClient(cc grpc.ClientConnInterface) HelloWorldServiceClient {
	return &helloWorldServiceClient{cc}
}

func (c *helloWorldServiceClient) Greeting(ctx context.Context, in *HelloWorldServiceRequest, opts ...grpc.CallOption) (*HelloWorldServiceReply, error) {
	out := new(HelloWorldServiceReply)
	err := c.cc.Invoke(ctx, HelloWorldService_Greeting_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloWorldServiceServer is the server API for HelloWorldService service.
// All implementations must embed UnimplementedHelloWorldServiceServer
// for forward compatibility
type HelloWorldServiceServer interface {
	Greeting(context.Context, *HelloWorldServiceRequest) (*HelloWorldServiceReply, error)
	mustEmbedUnimplementedHelloWorldServiceServer()
}

// UnimplementedHelloWorldServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHelloWorldServiceServer struct {
}

func (UnimplementedHelloWorldServiceServer) Greeting(context.Context, *HelloWorldServiceRequest) (*HelloWorldServiceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greeting not implemented")
}
func (UnimplementedHelloWorldServiceServer) mustEmbedUnimplementedHelloWorldServiceServer() {}

// UnsafeHelloWorldServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelloWorldServiceServer will
// result in compilation errors.
type UnsafeHelloWorldServiceServer interface {
	mustEmbedUnimplementedHelloWorldServiceServer()
}

func RegisterHelloWorldServiceServer(s grpc.ServiceRegistrar, srv HelloWorldServiceServer) {
	s.RegisterService(&HelloWorldService_ServiceDesc, srv)
}

func _HelloWorldService_Greeting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloWorldServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloWorldServiceServer).Greeting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HelloWorldService_Greeting_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloWorldServiceServer).Greeting(ctx, req.(*HelloWorldServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HelloWorldService_ServiceDesc is the grpc.ServiceDesc for HelloWorldService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HelloWorldService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "HelloWorldService",
	HandlerType: (*HelloWorldServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greeting",
			Handler:    _HelloWorldService_Greeting_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello_world.proto",
}