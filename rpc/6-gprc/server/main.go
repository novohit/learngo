package main

import (
	"context"
	"learngo/rpc/6-gprc/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

// 1 取出Server
type Server struct {
	pb.UnimplementedHelloServiceServer
}

// 2 挂载方法
func (s *Server) SayHello(ctx context.Context, req *pb.UserRequest) (resp *pb.UserResponse, err error) {
	return &pb.UserResponse{Code: 200, Data: req.GetName() + " ,hello grpc"}, nil
}

func main() {
	// 创建监听
	listener, err := net.Listen("tcp", ":8088")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// 实例化server并注册服务service
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &Server{})
	log.Printf("server listening at %v", listener.Addr())
	// 启动服务 将listener传进去
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
