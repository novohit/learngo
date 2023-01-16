package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	// http://127.0.0.1:8000/add?a=1&b=2
	// 返回格式: json{"data":3}
	// http严格上来讲不算是一种数据传输协议，数据传输协议是底层的TCP协议
	// 1.callID r.URL.Path 2.数据的传输协议(url的参数传递协议) 3.网络传输协议
	http.HandleFunc("/add", func(writer http.ResponseWriter, request *http.Request) {
		_ = request.ParseForm() // 解析参数
		fmt.Println("path: ", request.URL.Path)
		a, err := strconv.Atoi(request.Form["a"][0])
		if err != nil {
			log.Printf("Atoi failed %v", err)
		}
		b, err := strconv.Atoi(request.Form["b"][0])
		if err != nil {
			log.Printf("Atoi failed %v", err)
		}
		writer.Header().Set("Content-Type", "application/json")
		data, err := json.Marshal(map[string]int{"data": a + b})

		_, _ = writer.Write(data)
	})

	_ = http.ListenAndServe(":8000", nil)
}
