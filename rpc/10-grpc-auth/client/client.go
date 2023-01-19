package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	"learngo/rpc/10-grpc-auth/pb"
)

type CustomCredentials struct {
}

func (c CustomCredentials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"token": "12345"}, nil
}
func (c CustomCredentials) RequireTransportSecurity() bool {
	return false
}

func main() {
	// 客户端也可以定义拦截器
	//myInterceptor := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	//	md := metadata.Pairs("key1", "value1", "key1", "value2", "token", "fadfdaf")
	//	ctx = metadata.NewOutgoingContext(context.Background(), md)
	//	startTime := time.Now()
	//	err := invoker(ctx, method, req, reply, cc, opts...)
	//	log.Printf("本次调用耗时: %s\n", time.Since(startTime))
	//	return err
	//}
	//
	//option := grpc.WithUnaryInterceptor(myInterceptor)

	// grpc内置了一个封装了设置metadata的拦截器
	option := grpc.WithPerRPCCredentials(CustomCredentials{})

	// 1 拨号创建连接 Set up a connection to the server.
	conn, err := grpc.Dial("localhost:8088", grpc.WithInsecure(), option)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// 2 实例化client
	c := pb.NewUserServiceClient(conn)
	ctx := context.Background()
	// 3 拿到客户端后就可以进行调用
	responseData, err := c.GetData(ctx, &pb.RequestData{
		Data: "hello",
	})
	if err != nil {
		log.Println(err)
	}
	log.Println(responseData)

}
