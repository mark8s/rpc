package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"rpc/v3-json/dto"
	"time"
)

type MathService struct {
}

func (m *MathService) Sum(in dto.SumParam, out *dto.SumRes) error {
	fmt.Printf("in arg: %#v\n", in)
	*out = dto.SumRes{
		Sum:  in.A + in.B,
		Time: time.Now(),
	}
	return nil
}

func main() {
	// 注册服务
	err := rpc.RegisterName("MathService", new(MathService))
	if err != nil {
		return
	}

	// 监听端口
	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		return
	}

	// 监听
	for {
		conn, err := listen.Accept()
		if err != nil {
			return
		}
		// 使用json编码
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
