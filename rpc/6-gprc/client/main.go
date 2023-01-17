package main

import (
	"context"
	"learngo/rpc/6-gprc/pb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	// 1 拨号创建连接 Set up a connection to the server.
	conn, err := grpc.Dial("localhost:8088", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// 2 实例化client
	c := pb.NewHelloServiceClient(conn)

	ctx := context.Background()
	// 3 拿到客户端后就可以进行调用
	resp, err := c.SayHello(ctx, &pb.UserRequest{Name: "novo", Age: 14})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("code: %d, data: %s", resp.GetCode(), resp.GetData())
}
