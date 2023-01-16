package main

import (
	"encoding/json"
	"fmt"

	"github.com/kirinlabs/HttpRequest"
)

// rpc 远程过程调用，怎么做到像本地调用

type ResponseData struct {
	Data int `json:"data"`
}

func Add(a, b int) int {
	request := HttpRequest.NewRequest()
	response, _ := request.Get(fmt.Sprintf("http://localhost:8000/%s?a=%d&b=%d", "add", a, b))
	body, _ := response.Body()
	responseData := ResponseData{}
	_ = json.Unmarshal(body, &responseData)
	fmt.Println(string(body))
	return responseData.Data
}
func main() {
	println(Add(1, 5))
}
