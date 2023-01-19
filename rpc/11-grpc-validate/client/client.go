package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	"learngo/rpc/11-grpc-validate/pb"
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
	ctx := context.Background()
	// 3 拿到客户端后就可以进行调用
	responseData, err := c.GetData(ctx, &pb.RequestData{
		Data:  "hello",
		Id:    10,
		Email: "inti@tt.com",
	})
	if err != nil {
		log.Println(err)
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
