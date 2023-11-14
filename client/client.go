package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "github.com/AiLiaa/grpc-demo/proto"
)

const address = "localhost:50051"

func main() {
	//连接到这个地址，并创建一个 gRPC 连接实例
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	//创建一个 GreeterClient 的客户端实例 c，该实例用于发送 gRPC 请求
	c := pb.NewGreeterClient(conn)

	//请求参数
	name := "world"
	//创建了一个具有超时上下文的 ctx 1秒
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//发送 SayHello 请求 接收响应结果 r
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
