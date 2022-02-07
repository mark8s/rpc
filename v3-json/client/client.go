package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"rpc/v3-json/dto"
)

func main() {
	// 建立连接
	conn, err := net.Dial("tcp", ":9090")
	if err != nil {
		return
	}
	// 创建客户端
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	param := dto.SumParam{A: 2, B: 8}
	var sum dto.SumRes
	err = client.Call("MathService.Sum", param, &sum)
	if err != nil {
		return
	}
	fmt.Printf("res:%+v \n", sum)
}
