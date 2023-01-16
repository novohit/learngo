package main

import (
	"fmt"

	"github.com/kirinlabs/HttpRequest"
)

func main() {
	request := HttpRequest.NewRequest()
	response, _ := request.Post("http://localhost:1234/jsonrpc", "{\"method\":\"hello-service.Hello\", \"params\":[\"zwx\"], \"id\":0}")
	body, _ := response.Body()
	fmt.Println(string(body))
}
