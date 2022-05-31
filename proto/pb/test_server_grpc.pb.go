// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.9.0
// source: test_server.proto

package test_go

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

// TestServerClient is the client API for TestServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestServerClient interface {
	// Test 测试方法
	Test(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type testServerClient struct {
	cc grpc.ClientConnInterface
}

func NewTestServerClient(cc grpc.ClientConnInterface) TestServerClient {
	return &testServerClient{cc}
}

func (c *testServerClient) Test(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/test_go.TestServer/Test", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestServerServer is the server API for TestServer service.
// All implementations must embed UnimplementedTestServerServer
// for forward compatibility
type TestServerServer interface {
	// Test 测试方法
	Test(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedTestServerServer()
}

// UnimplementedTestServerServer must be embedded to have forward compatible implementations.
type UnimplementedTestServerServer struct {
}

func (UnimplementedTestServerServer) Test(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Test not implemented")
}
func (UnimplementedTestServerServer) mustEmbedUnimplementedTestServerServer() {}

// UnsafeTestServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestServerServer will
// result in compilation errors.
type UnsafeTestServerServer interface {
	mustEmbedUnimplementedTestServerServer()
}

func RegisterTestServerServer(s grpc.ServiceRegistrar, srv TestServerServer) {
	s.RegisterService(&TestServer_ServiceDesc, srv)
}

func _TestServer_Test_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServerServer).Test(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test_go.TestServer/Test",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServerServer).Test(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// TestServer_ServiceDesc is the grpc.ServiceDesc for TestServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TestServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "test_go.TestServer",
	HandlerType: (*TestServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Test",
			Handler:    _TestServer_Test_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test_server.proto",
}
