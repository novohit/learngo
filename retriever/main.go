package main

import (
	"fmt"
	"learngo/retriever/mock"
	"learngo/retriever/real"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("https://www.imooc.com")
}

// go语言中接口是由使用者定义的
// 不同于传统面向对象语言，一个类实现接口后会告诉其它人，该类实现这个接口
// go语言接口的实现是隐式的，只要实现接口里的方法
func main() {
	var r Retriever
	r = mock.Retriever{Contents: "this is a fake message"}
	fmt.Println(download(r))
	r = real.Retriever{}
	fmt.Println(download(r))
}
