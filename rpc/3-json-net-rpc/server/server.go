package main

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct {
}

func (s *HelloService) Hello(request string, reply *string) error {
	// 返回值是通过指针修改reply的值
	*reply = "hello " + request
	return nil
}

func main() {
	// 1 实例化一个server 用net包监听tcp连接
	listener, _ := net.Listen("tcp", ":1234")
	// 2 注册处理逻辑handler
	_ = rpc.RegisterName("hello-service", &HelloService{})
	// 3 启动服务
	for {
		conn, _ := listener.Accept() // 当有监听到请求连接时 传给rpc
		// 将rpc内置的数据编码协议改为json
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
