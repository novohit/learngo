package main

import (
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"

	"learngo/rpc/8-grpc-metadata/pb"
)

const PORT = ":8088"

type server struct {
	pb.UnimplementedUserServiceServer
}

// 传统模式
func (s *server) GetData(ctx context.Context, req *pb.RequestData) (resp *pb.ResponseData, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		for k, v := range md {
			log.Println(k, v)
		}
	}
	return &pb.ResponseData{
		Data:       req.Data + "!",
		Code:       200,
		CreateTime: timestamppb.New(time.Now()),
	}, nil
}

func main() {
	// 创建监听
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// 实例化server并注册服务service
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	log.Printf("server listening at %v", listener.Addr())
	// 启动服务 将listener传进去
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
