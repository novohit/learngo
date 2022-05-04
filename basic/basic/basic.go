package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var aa = 3
var ss = "kkk"
var bb = true

// var()集中定义变量
var (
	tt = 3
	yy = "kkk"
	kk = true
)

func variableZeroValue() {
	// 局部变量一定要被使用否则会报错
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

// 编译器可推断变量类型
func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

// 原生支持复数
func euler() {
	// 定义复数
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
	// 欧拉公式 e^iπ + 1 = 0
	// complex64的实部和虚部都是float32 complex128的实部和虚部都是float64
	// 浮点数在任何语言都是不精确的
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}
func variableShort() {
	// := 只能在函数内使用，定义局部变量，在函数外不可以:=,只能用var定义包内部变量
	a, b, c, s := 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func triangle() {
	var a, b int = 3, 4
	var c int
	var f float64 = float64(a)
	fmt.Println(f)
	// go 只有强制类型转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

const filename = "abc.txt"

func consts() {
	const filename = "abc.txt"
	// go 常量名字一般不会使用大写
	const (
		name = "zwx"
		age  = 1
	)

	const a, b = 3, 4
	var c int
	// 常量在不规定类型的时候相当于文本 是不确定类型的 所sqrt不用转float
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

// go 没有特定的枚举关键字，用一组const实现枚举
func enums() {
	const (
		cpp    = 0
		java   = 1
		python = 2
		golang = 3
	)
	fmt.Println(cpp, java, python, golang)
	// 自增关键字iota简化
	const (
		a = iota
		r
		c
		d
	)
	fmt.Println(a, r, c, d)
	// iota做自增值种子
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}
func main() {
	fmt.Println("hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShort()
	euler()
	triangle()
	consts()
	enums()
}
