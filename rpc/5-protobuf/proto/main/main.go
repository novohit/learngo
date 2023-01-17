package main

import (
	"encoding/json"
	"fmt"
	user "learngo/rpc/5-protobuf/proto"

	"github.com/golang/protobuf/proto"
)

type User struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Courses []string `json:"courses"`
}

func main() {
	req := user.UserRequest{Name: "novo", Age: 11, Courses: []string{"微服务", "go"}}
	protoResp, _ := proto.Marshal(&req)
	// TODO 具体的编码是怎么实现的 protobuf的原理
	fmt.Println(string(protoResp))
	fmt.Println("切片长度", len(protoResp))
	fmt.Println("================")
	jsonStruct := User{Name: "novo", Age: 11, Courses: []string{"微服务", "go"}}
	jsonResp, _ := json.Marshal(jsonStruct)
	fmt.Println(string(jsonResp))
	fmt.Println("切片长度", len(jsonResp))

	// protobuf反序列化
	newReq := user.UserRequest{}
	_ = proto.Unmarshal(protoResp, &newReq)
	fmt.Println(newReq.Name, newReq.Age, newReq.Courses)
}
