package main

import (
	"fmt"
	"io/ioutil"
)

// go switch会自动break，除非使用fallthrough
func eval(a, b int, op string) int {
	var res int
	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "/":
		res = a / b
	case "*":
		res = a * b
	}
	return res
}

// 函数名(参数名 参数类型) 返回值类型
func grade(score int) string {
	res := ""
	// go的case可以使用表达式，前提是switch后不写表达式， 其他语言底层case都是用的整型比较
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score)) // 会中断程序
	case score < 60:
		res = "F"
	case score < 80:
		res = "C"
	case score < 90:
		res = "B"
	case score <= 100:
		res = "A"
	}
	return res
}
func main() {
	const filename = "basic/abc.txt"
	// 获取文件的全部内容 content是一个字节数组
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	// if 的条件里可以赋值且赋值的变量作用域只在这个if语句里
	if contents, err := ioutil.ReadFile(filename); err == nil {
		// content是一个字节数组
		fmt.Println(contents)
		fmt.Println(string(contents))
	} else {
		fmt.Println(err)
	}

	fmt.Println(
		grade(9),
		grade(44),
		grade(66),
		grade(77),
		grade(88),
		grade(99),
	)

	fmt.Println(eval(3, 4, "*"))
}
