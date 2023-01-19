package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc/status"

	"google.golang.org/grpc"

	"learngo/rpc/13-grpc-RTO/pb"
)

type CustomCredentials struct {
}

func (c CustomCredentials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"token": "123456"}, nil
}
func (c CustomCredentials) RequireTransportSecurity() bool {
	return false
}

func main() {
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
	// 设置超时时间
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	// 3 拿到客户端后就可以进行调用
	responseData, err := c.GetData(ctx, &pb.RequestData{
		Data:  "hello",
		Id:    10,
		Email: "inti@tt.com",
	})
	if err != nil {
		status, ok := status.FromError(err)
		if !ok {
			log.Fatalf("error解析为status发生错误")
		}
		log.Println(status.Code())
		log.Println(status.Message())
		log.Println(status.Details())
		return
	}
	log.Println(responseData)

	//data := pb.RequestData{
	//	Id:    -100,
	//	Email: "inti@tt.com",
	//}
	//err = data.Validate()
	//if err != nil {
	//	log.Println(err)
	//}
	//req := new(pb.RequestData)
	//err = req.Validate()
}
