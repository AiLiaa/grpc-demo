package main

import (
	"context"
	pb "github.com/AiLiaa/grpc-demo/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	//大坑,实现GreeterServer接口,嵌入结构体UnimplementedGreeterServer
	//UnimplementedGreeterServer must be embedded to have forwarded compatible implementations.
	pb.UnimplementedGreeterServer
}

// SayHello
// 实现了 SayHello 方法，该方法是 GreeterServer 接口中定义的方法
// 处理客户端发来的 HelloRequest 请求，并返回一个 HelloReply 响应
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	//创建了一个 TCP 监听器，并将其绑定到端口 :50051 上
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//创建了一个 gRPC 服务端实例
	s := grpc.NewServer()
	//将我们的服务端注册到 gRPC 服务器上
	pb.RegisterGreeterServer(s, &server{})

	//启动 gRPC 服务器并开始监听来自客户端的请求
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
