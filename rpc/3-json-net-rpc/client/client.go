package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	// 1 建立连接
	// 这里用net拨号 之前用rpc拨号是因为直接用的gob编码
	// net拨号返回的是一个连接 我们再用rpc设置json编码包装成client
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic("connection fail")
	}
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	var reply string
	// {"method":"hello-service.Hello", "params":["hello"], "id":0}
	err = client.Call("hello-service.Hello", "novo", &reply)
	if err != nil {
		panic("call fail")
	}
	fmt.Println(reply)
}
