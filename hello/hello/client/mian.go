package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	pb "grpc-test/hello/proto/hello"
)

const (
	address = "127.0.0.1:8084"

	// OpenTLS 是否开启TLS认证
	OpenTLS = true
)

// CustomCredential 自定义认证
type CustomCredential struct{}

// GetRequestMateData 实现自定义认证接口
func (c CustomCredential) GetRequestMateData(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "101010",
		"appkey": "i am key",
	}, nil
}
func main() {

	//

	// 创建连接
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalln(err)
	}

	// 初始化客户端
	client := pb.NewHelloClient(conn)
	// 调用方法
	req := &pb.HelloRequest{Name: "grpc"}

	resp, err := client.SayHello(context.Background(), req)
	if err != nil {
		grpclog.Fatalln(err)
	}
	fmt.Println(resp.Message)
}
