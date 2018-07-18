package main

import (
	"log"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "./proto"
)

const (
	address     = "localhost:50051" // 服务端地址
)

func main() {
	// 启动grpc客户端，连接grpc服务端
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// 使用连接，创建HelloService实例
	helloService := pb.NewHelloServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 调用SayHello，向服务端发送消息
	response, err := helloService.SayHello(ctx, &pb.HelloRequest{
		Code: "0",
		Message: "来自GO客户端的OK",
	})
	if err != nil {
		log.Fatalf("could not sayHello: %v", err)
	}

	// 打印服务端回应内容
	log.Printf("SayHello: Code: %s,Message: %s", response.Code, response.Message)
}