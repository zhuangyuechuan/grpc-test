package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	pb "grpc-test/hello/proto/hello"
	"net"
)

const (
	address = "127.0.0.1:8084"
)

type helloServer struct{}

var HelloServer = helloServer{}

func (h *helloServer) SayHello(ctx context.Context, req *pb.HelloRequest) (resp *pb.HelloResponse, err error) {
	resp = new(pb.HelloResponse)
	resp.Message = fmt.Sprintf("Hello %s", req.Name)
	return resp, nil
}

func main() {
	// 设置监听
	l, err := net.Listen("tcp", address)
	if err != nil {
		grpclog.Fatalln("Faild to Listen: %v", err)
	}

	// TLS认证
	creds, err := credentials.NewServerTLSFromFile("hello/keys/server.pem", "hello/keys/server.key")
	if err != nil {
		grpclog.Fatalf("Failed to generate credentials %v", err)
	}

	// 实例化grpc server
	s := grpc.NewServer(grpc.Creds(creds))
	// 注册HelloService
	pb.RegisterHelloServer(s, &HelloServer)

	grpclog.Println("Listen on " + address)
	// 开启服务
	s.Serve(l)
}
