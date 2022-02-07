package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {

	// 建立连接
	dial, err := rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		return
	}

	var result string
	// 发起请求
	// 参数serviceMethod 注册服务名+“.”+方法名
	// reply是返回值，这里需要传指针
	err = dial.Call("HelloService.Say", "golang", &result)
	if err != nil {
		log.Fatal("Dial error ", err)
	}

	fmt.Println("result:", result)
}
