package main

import (
	"fmt"
)

func printArray(arr [5]int) {
	fmt.Println()
	for i, v := range arr {
		fmt.Printf("arrs[%d] = %d ", i, v)
	}
	fmt.Println()
}

// go 数组的参数传递是拷贝的是整个数组的值而不是地址值，java的数组参数传递则是拷贝了引用值
func swap3(arr [5]int) {
	// 只会改变副本的值
	arr[0] = 100
	arr[1] = 100
}

// 通过数组的指针来改变值
func swap3ByPoint(arr *[5]int) {
	// 数组可以直接arr[i]使用指针改值
	arr[0] = 100
	arr[1] = 100
}

// 传切片
func swap3BySlice(arr []int) {
	arr[0] = 900
	arr[1] = 800
}
func main() {
	/**
	go 语言一般不直接使用数组，因为数组不是很方便
	*/
	var arr1 [5]int
	fmt.Println(arr1)

	// 使用:= 定义数组一定要花括号赋初值
	arr2 := [3]int{}
	fmt.Println(arr2)
	// 切片 让比编译器识别数组长度
	arr3 := [...]int{2, 4, 6, 8, 10}
	fmt.Println(arr3)

	grid := [4][5]bool{}
	fmt.Println(grid)

	// 常规遍历数组
	for i := 0; i < len(arr3); i++ {
		fmt.Printf("%d ", arr3[i])
	}
	// range 获得下标 遍历数组
	for i := range arr3 {
		fmt.Printf("%d ", arr3[i])
	}
	fmt.Println()
	// range 获取下标和值
	for i, v := range arr3 {
		fmt.Printf("arrs[%d] = %d ", i, v)
	}
	// range 只获取值
	fmt.Println()
	for _, v := range arr3 {
		fmt.Printf("%d ", v)
	}
	printArray(arr3)
	swap3(arr3)
	printArray(arr3)
	swap3ByPoint(&arr3)
	printArray(arr3)
	swap3BySlice(arr3[:]) // 传切片
	printArray(arr3)
}
