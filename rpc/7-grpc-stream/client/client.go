package main

import (
	"context"
	"fmt"
	"learngo/rpc/7-grpc-stream/pb/person"
	"log"
	"sync"
	"time"

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
	c := person.NewModelTestServiceClient(conn)

	ctx := context.Background()
	// 3 拿到客户端后就可以进行调用
	log.Printf("服务端数据流模式")
	sstreamClient, err := c.ServerStreamModel(ctx, &person.StreamReq{Data: "hello"})
	for {
		// socket模式
		resp, err := sstreamClient.Recv()
		if err != nil {
			log.Printf("could not greet: %v\n", err)
			break
		}
		log.Println(resp)
	}

	// 客户端流模式
	log.Printf("客户端数据流模式")
	cstreamClient, err := c.ClientStreamModel(ctx)
	i := 1
	for {
		err := cstreamClient.Send(&person.StreamReq{
			Data: fmt.Sprintf("hello %d", i),
		})
		if err != nil {
			log.Printf("server receive fail: %v\n", err)
		}
		i++
		time.Sleep(time.Second)
		if i > 10 {
			recv, _ := cstreamClient.CloseAndRecv()
			log.Println(recv)
			break
		}
	}

	// 双向流模式
	log.Printf("双向数据流模式")
	doubleClient, err := c.DoubleStreamModel(ctx)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			err := doubleClient.Send(&person.StreamReq{
				Data: "a message by client",
			})
			if err != nil {
				log.Printf("client send error")
				break
			}
			time.Sleep(time.Second)
		}
	}()
	go func() {
		defer wg.Done()
		for {
			data, err := doubleClient.Recv()
			if err != nil {
				log.Printf("client receive error")
				break
			}
			log.Println("receive by server, ", data)
		}
	}()
	wg.Wait()

	// 入参为流的模式 client有中断的方法 CloseAndRecv() CloseSend()
}
