package main

import "fmt"

// adder返回值是一样函数func(int) int
func adder() func(int) int {
	sum := 0 // sum这个变量在func函数外，是一个环境，叫自由变量
	return func(val int) int {
		sum += val
		return sum
	} // 返回的是一个闭包，不是返回代码，返回sum的引用
}

// 正统的函数式编程，没有状态
type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(val int) (int, iAdder) {
		return base + val, adder2(base + val)
	}
}
func main() {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("0 + 1 + ... + %d = %d\n", i, a(i))
	}
	fmt.Println(a(10))
	fmt.Println(a(10))
	fmt.Println(a(10))
	fmt.Println("=========")
	a2 := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, a2 = a2(i)
		fmt.Printf("0 + 1 + ... + %d = %d\n", i, s)
	}
}
