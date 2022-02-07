package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"rpc/v4-grpc/proto/hello"
)

func main() {

	// 创建grpc服务
	grpcServer := grpc.NewServer()
	// 注册服务
	hello.RegisterUserServiceServer(grpcServer, new(hello.HelloService))
	// 监听端口
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		return
	}
	log.Print("服务启动.... \n")
	grpcServer.Serve(listen)
}
