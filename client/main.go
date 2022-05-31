package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	hello_grpc "grpc-test/pb"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
	if err != nil {
		log.Println("client: ", err)
	}
	defer conn.Close()
	client := hello_grpc.NewHelloGRPCClient(conn)
	req, _ := client.SayHi(context.Background(), &hello_grpc.Req{Message: "客户端请求"})
	fmt.Println(req.GetMessage())
}
