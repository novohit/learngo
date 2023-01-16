package main

import (
	"io"
	"net/http"
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
	// 1 实例化一个server
	http.HandleFunc("/jsonrpc", func(writer http.ResponseWriter, request *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: request.Body,
			Writer:     writer,
		}
		// 使用http连接的rpc
		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})
	// 2 注册处理逻辑handler
	_ = rpc.RegisterName("hello-service", &HelloService{})
	// 3 启动服务
	http.ListenAndServe(":1234", nil)
}
