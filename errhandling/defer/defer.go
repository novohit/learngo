package main

import (
	"bufio"
	"fmt"
	"learngo/errhandling"
	"os"
)

func tryDefer() {
	// defer 相当于有一个栈 所以2先打印
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	panic("err")
	fmt.Println(4)
}

func tryDefer2() {
	for i := 0; i < 100; i++ {
		// 参数在defer语句时计算
		defer fmt.Println(i) // 这里的i已经保留了当时的i 而不是打印出30行个30
		if i == 30 {
			panic("printed too many")
		}
	}
}

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush() // 将缓存的内容刷新出来

	f := errhandling.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}

}

// defer调用：确保调用在函数结束时发生
// 参数在defer语句时计算
// defer列表为先进后出
func main() {
	//tryDefer()
	writeFile("errhandling/fib.txt")
	tryDefer2()
}
