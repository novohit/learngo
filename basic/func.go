package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

// 多个返回值一般用于返回error
func eval2(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		// 如果不想用第二个返回值，要用下划线定义第二个返回值，否则定义了r但是不使用 是编译错误
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported op: %s", op)
	}
}

// 函数可以返回多个值
func div(a, b int) (int, int) {
	return a / b, a % b
}

// 可以给多个返回值起名
func divByName(a, b int) (q, r int) {
	return a / b, a % b
}

// 函数式编程 函数的参数、返回值、体内都可以是函数
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args"+"(%d, %d)\n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// 可变参数
func sum(nums ...int) int {
	res := 0
	for i := range nums {
		res += nums[i]
	}
	return res
}

// 交换指针
func swap(a, b *int) {
	*a, *b = *b, *a
}

// 直接返回交换后的值
func swap2(a, b int) (int, int) {
	return b, a
}
func main() {
	/**
	go 语言没有函数重载，lambda表达式
	go 和 java一样都只有值传递，参数的值都会被拷贝一份，可以通过拷贝指针来改变值
	*/
	if res, err := eval2(3, 4, "("); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	println(eval2(2, 3, "+"))
	fmt.Println(div(13, 3))
	name, r := divByName(13, 5)
	fmt.Println(name, r)
	// 也可以传匿名函数
	res := apply(func(a, b int) int { return a + b }, 3, 4)
	fmt.Println(res)
	// 参数为函数
	fmt.Println(apply(pow, 2, 3))

	fmt.Println(sum(1, 2, 3, 4))

	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)
	fmt.Println(swap2(a, b))
}
