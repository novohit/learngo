package main

import (
	"fmt"
	"learngo/interface/infra"
)

func getRetriever() retriever {
	// 直到return才知道返回的是什么类型
	return infra.Retriever{}
}

// Something that can "Get"
// infra和testing并没有地方指定实现了这个interface，为什么就可以用了？
type retriever interface {
	Get(string) string
}

func main() {

	r := getRetriever()
	fmt.Println(r.Get("https://www.imooc.com"))
}
