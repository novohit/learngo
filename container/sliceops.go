package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("s = %v, len(s) = %d, 容量cap(s) = %d \n", s, len(s), cap(s))
}
func main() {
	// 创建slice方式一
	var s []int // Zero value for slice is nil
	fmt.Println(s)
	for i := 0; i < 100; i++ {
		fmt.Printf("len(s) = %d, cap = %d \n", len(s), cap(s)) // 可以看出每次扩容为原来的两倍
		s = append(s, i)
	}
	fmt.Println(s)
	// make函数创建 方式二
	s2 := make([]int, 10)      // 10 为len
	s3 := make([]int, 10, 128) // 128 为cap
	printSlice(s2)
	printSlice(s3)

	// Slice的copy
	fmt.Println("Copying slice")
	copy(s2, s)    // 拷贝值 dst <-- src
	printSlice(s2) // [0 1 2 3 4 5 6 7 8 9]

	fmt.Println("Deleting elements from slice")
	// 将s2的4删除掉
	s2 = append(s2[:4], s2[5:]...) // arg2为可变参数，将s2[5:]的值一个一个的append到s2[:4]
	printSlice(s2)                 // [0 1 2 3 5 6 7 8 9]

	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println("front =", front)
	printSlice(s2)

	fmt.Println("Popping from tail")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println("tail =", tail)
	printSlice(s2)
}
