package main

import (
	"bufio"
	"log"
	"net"
)

func main() {
	// 监听本地1080端口
	server, err := net.Listen("tcp", "127.0.0.1:1080")
	if err != nil {
		panic(err)
	}
	for {
		// 不断地接收请求
		// 命令 nc 127.0.0.1 1080 可以建立一个tcp连接
		client, err := server.Accept()
		if err != nil {
			log.Printf("Accept failed %v", err)
			continue
		}
		// 成功接收到请求后 开子线程处理
		// 类似新建一个子线程 但是开销比其他语言建立子线程小
		go process(client)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	// 简单的处理 发送什么打印什么
	reader := bufio.NewReader(conn)
	for {
		b, err := reader.ReadByte()
		if err != nil {
			break
		}
		_, err = conn.Write([]byte{b})
		if err != nil {
			break
		}
	}
}
