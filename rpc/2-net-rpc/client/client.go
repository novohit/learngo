package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// 1 建立连接
	client, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		panic("connection fail")
	}
	var reply string
	err = client.Call("hello-service.Hello", "novo", &reply)
	if err != nil {
		panic("call fail")
	}
	fmt.Println(reply)
}
