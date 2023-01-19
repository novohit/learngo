package main

import (
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"learngo/rpc/10-grpc-auth/pb"
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
	// 定义一个拦截器 只需要实现拦截器的方法
	myInterceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		log.Println("请求前")
		// handler就是执行原来的逻辑
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return resp, status.Error(codes.Unauthenticated, "无metadata信息")
		}

		token, ok := md["token"]
		if !ok {
			return resp, status.Error(codes.Unauthenticated, "用户未登录")
		}
		if token[0] != "123456" {
			return resp, status.Error(codes.Unauthenticated, "用户未登录")
		}
		log.Println("token:", token)
		res, err := handler(ctx, req)
		log.Printf("请求结束")
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
