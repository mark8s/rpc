package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"rpc/v4-grpc/proto/hello"
)

func main() {
	// 建立连接
	conn, err := grpc.Dial("127.0.0.1:1234", grpc.WithInsecure())
	if err != nil {
		return
	}
	// 关闭连接
	defer conn.Close()
	// 实例化客户端
	client := hello.NewUserServiceClient(conn)
	// 发起请求
	reply, err := client.Say(context.TODO(), &hello.Request{
		Name: "mark",
	})
	if err != nil {
		return
	}
	fmt.Println("返回：", reply.Result)

}
