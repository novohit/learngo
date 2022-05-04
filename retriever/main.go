package main

import (
	"fmt"
	"learngo/retriever/mock"
	"learngo/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("https://www.imooc.com")
}

// 查看r的类型 Type Switch
func inspect(r Retriever) {
	fmt.Printf("%T %v", r, r)
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}

// go语言中接口是由使用者定义的
// 不同于传统面向对象语言，一个类实现接口后会告诉其它人，该类实现这个接口
// go语言接口的实现是隐式的，只要实现接口里的方法
// 接口变量相当于一个结构，含有实现者的类型和实现者的指针，实现者的指针指向实现者
// 接口变量自带指针，接口变量同样采用值传递，几乎不需要使用接口的指针，指针接收者实现只能以指针方式使用，值接收者都可
// 表示任何类型: interface{}
// Type Assertion Type Switch
func main() {
	var r Retriever
	// mock是指接收者实现，也可以传指针
	r = mock.Retriever{Contents: "this is a fake message"}
	r = &mock.Retriever{Contents: "this is a fake message"}

	fmt.Printf("%T %v\n", r, r) // %T取类型，%v取值
	fmt.Println(download(r))
	inspect(r)
	// real是指针接收者实现，只能以指针方式使用
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	fmt.Printf("%T %v\n", r, r)
	inspect(r)
	//fmt.Println(download(r))

	// Type Assertion 类型的断言
	realRetriever := r.(*real.Retriever) // 类似于向下转型
	fmt.Println(realRetriever.TimeOut)
	if mockRetriever, ok := r.(mock.Retriever); !ok {
		fmt.Println("main.Retriever is not mock.Retrieve")
	} else {
		fmt.Println(mockRetriever.Contents) // 如果断言失败：interface conversion: main.Retriever is real.Retriever, not mock.Retrieve
	}

}
