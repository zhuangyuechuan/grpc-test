package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	hello_grpc "grpc-test/pb"
	"log"
	"net"
)

type server struct {
	hello_grpc.UnimplementedHelloGRPCServer
}

func (s *server) SayHi(cxt context.Context, req *hello_grpc.Req) (res *hello_grpc.Res, err error) {
	fmt.Println(req.GetMessage())
	return &hello_grpc.Res{Message: "从服务端放回grpc内容"}, nil
}

func main() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Println(err)
	}
	s := grpc.NewServer()
	hello_grpc.RegisterHelloGRPCServer(s, &server{})
	s.Serve(l)
}
