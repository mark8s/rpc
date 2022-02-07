package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// 建立连接
	client, err := rpc.DialHTTP("tcp", ":8080")
	if err != nil {
		return
	}

	var result int
	err = client.Call("MathService.Multi", 2, &result)
	if err != nil {
		return
	}
	fmt.Printf("result:%v \n", result)
}
