package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc"

	"learngo/rpc/9-grpc-interceptor/pb"
)

func main() {

	// 客户端也可以定义拦截器
	myInterceptor := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		startTime := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)
		log.Printf("本次调用耗时: %s\n", time.Since(startTime))
		return err
	}

	option := grpc.WithUnaryInterceptor(myInterceptor)
	// 1 拨号创建连接 Set up a connection to the server.
	conn, err := grpc.Dial("localhost:8088", grpc.WithInsecure(), option)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// 2 实例化client
	c := pb.NewUserServiceClient(conn)
	md := metadata.Pairs("key1", "value1", "key1", "value2", "token", "fadfdaf")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	// 3 拿到客户端后就可以进行调用
	responseData, err := c.GetData(ctx, &pb.RequestData{
		Data: "hello",
	})
	log.Println(responseData)

}
