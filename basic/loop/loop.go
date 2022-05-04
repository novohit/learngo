package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBin(n int) string {
	res := ""
	// 省略初始条件 相当于while
	for ; n > 0; n /= 2 {
		low := n % 2
		res = strconv.Itoa(low) + res
	}
	return res
}

func getSum(n int) int {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	return sum
}

func printFile(filename string) {
	file, err := os.Open(filename)
	// file 是一个文件类型
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	// 初始条件和递增条件都省略 go语言没有while 相当于while 一行一行打印
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	// 死循环很好写，因为go语言经常要需要用死循环
	for {
		fmt.Println("hello")
	}
}
func main() {
	fmt.Println(
		convertToBin(5),
		convertToBin(13),
		convertToBin(123213123123),
	)
	fmt.Println(getSum(100))
	printFile("basic/branch/abc.txt")
	forever()
}
