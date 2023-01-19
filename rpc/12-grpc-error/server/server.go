package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"learngo/rpc/12-grpc-error/pb"

	"google.golang.org/grpc"
)

const PORT = ":8088"

type server struct {
	pb.UnimplementedUserServiceServer
}

// 传统模式
func (s *server) GetData(ctx context.Context, req *pb.RequestData) (resp *pb.ResponseData, err error) {
	// status.New(c, msg)是Status类型 有Message、Code和Details字段
	return nil, status.Errorf(codes.InvalidArgument, "参数错误 %s", req.Email)
}

type Validator interface {
	Validate() error
}

func main() {
	// 定义一个拦截器 只需要实现拦截器的方法
	myInterceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		res, err := handler(ctx, req)
		return res, err
	}

	// 创建监听
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// 配置拦截器
	option := grpc.UnaryInterceptor(myInterceptor)
	// 实例化server并注册服务service
	s := grpc.NewServer(option)
	pb.RegisterUserServiceServer(s, &server{})
	log.Printf("server listening at %v", listener.Addr())
	// 启动服务 将listener传进去
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
