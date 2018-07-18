package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "./proto"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051" // RPC服务端端口
)

type HelloService struct{}

// 实现hello的HelloServiceServer中的SayHello方法
// 表示实现了HelloServiceServer接口
func (s *HelloService) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {

	// 打印客户端请求内容
	log.Printf("SayHello: Code: %s,Message: %s", request.Code, request.Message)

	// 响应客户端
	return &pb.HelloResponse{
		Code: "0",
		Message: "来自GO服务端的OK",
	}, nil
}

func main() {
	// 启动服务监听
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 新建一个grpc服务器
	s := grpc.NewServer()

	// 将实现HelloService注册到grpc服务器中
	pb.RegisterHelloServiceServer(s, &HelloService{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}