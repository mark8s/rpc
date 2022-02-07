package main

import (
	"log"
	"net"
	"net/rpc"
	"time"
)

type HelloService struct {
}

// 方法是导出的
// 方法有两个参数，都是导出类型或内建类型
// 方法的第二个参数是指针
// 方法只有一个error接口类型的返回值

func (h *HelloService) Say(request string, response *string) error {
	format := time.Now().Format("2006-01-02 15:04:05")
	*response = request + " -- " + format
	return nil
}

func main() {

	// 1. 注册服务
	// 参数name为注册的服务名，参数rcvr是注册的服务
	rpc.RegisterName("HelloService", new(HelloService))

	// 2.监听端口
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		return
	}

	for {
		// 3.监听请求
		accept, err := listen.Accept()
		log.Println("服务端接收请求")
		if err != nil {
			log.Fatalf("Accept Error: %s", err)
		}
		go rpc.ServeConn(accept)
	}

}
