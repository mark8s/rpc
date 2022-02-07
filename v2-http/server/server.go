package main

import (
	"fmt"
	"net/http"
	"net/rpc"
)

type MathService struct {
}

func (m *MathService) Multi(a int, sum *int) error {
	*sum = a * a
	return nil
}

func main() {

	// 1.注册服务
	err := rpc.RegisterName("MathService", new(MathService))
	if err != nil {
		return
	}

	rpc.HandleHTTP()
	// 监听端口
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
