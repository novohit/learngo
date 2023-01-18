package main

import (
	"context"
	"fmt"
	"learngo/rpc/7-grpc-stream/pb/person"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
)

const PORT = ":8088"

type server struct {
	person.UnimplementedModelTestServiceServer
}

// 传统模式
func (s *server) SimpleModel(cnt context.Context, req *person.StreamReq) (resp *person.StreamReq, err error) {
	return nil, nil
}

// 客户端数据流模式
func (s *server) ClientStreamModel(server person.ModelTestService_ClientStreamModelServer) error {
	log.Printf("客户端数据流模式")
	i := 1
	for {
		req, err := server.Recv()
		if err != nil {
			_ = server.SendAndClose(&person.StreamReq{Data: "server receive finish"})
			break
		}
		log.Printf("server receive NO.%d, data: %s\n", i, req.Data)
		i++
	}
	return nil
}

// 服务端数据流模式
func (s *server) ServerStreamModel(req *person.StreamReq, server person.ModelTestService_ServerStreamModelServer) error {
	log.Printf("服务端数据流模式")
	i := 1
	for {
		_ = server.Send(&person.StreamResp{
			Data: fmt.Sprintf("server send NO.%d", i),
		})
		i++
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}
	return nil
}

// 双向数据流模式
func (s *server) DoubleStreamModel(server person.ModelTestService_DoubleStreamModelServer) error {
	log.Printf("双向数据流模式")
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			data, err := server.Recv()
			if err != nil {
				log.Printf("server receive error")
				break
			}
			log.Println("receive by client, ", data)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			err := server.Send(&person.StreamReq{
				Data: "a message by server",
			})
			if err != nil {
				log.Printf("server send error")
				break
			}
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
	return nil
}

func main() {
	// 创建监听
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// 实例化server并注册服务service
	s := grpc.NewServer()
	person.RegisterModelTestServiceServer(s, &server{})
	log.Printf("server listening at %v", listener.Addr())
	// 启动服务 将listener传进去
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
