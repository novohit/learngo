package main

import (
	"context"
	"log"

	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc"

	"learngo/rpc/8-grpc-metadata/pb"
)

func main() {
	// 1 拨号创建连接 Set up a connection to the server.
	conn, err := grpc.Dial("localhost:8088", grpc.WithInsecure())
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
