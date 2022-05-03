package main

import "fmt"

// 不加长度就是Slice
func updateSlice(s []int) {
	s[0] = 100
}
func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	// 切片是一个视图 左闭右开索引[2, 6)
	// Slice本身没有数据，是对底层array的一个view
	s := arr[2:6]
	fmt.Println(s)
	fmt.Println(arr[:6])
	fmt.Println(arr[2:])
	fmt.Println(arr[:])
	s1 := arr[2:]
	s2 := arr[:]
	fmt.Println("After update s1")
	updateSlice(s1)
	fmt.Println("s1:", s1)   // [100 3 4 5 6 7 8]
	fmt.Println("arr:", arr) // arr也会改变 [0 1 100 3 4 5 6 7 8]
	fmt.Println("After update s2")
	updateSlice(s2)
	fmt.Println("s2:", s2)

	fmt.Println("多次切片")
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	arr2 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s3 := arr2[2:6] // [2 3 4 5]
	s4 := s3[3:5]   // Slice向后拓展 len越界的话是取不到值的，但是只要容量没超cap，可以继续进行切片
	fmt.Printf("s3 = %v, len(s3) = %d, 容量cap(s3) = %d \n", s3, len(s3), cap(s3))
	fmt.Printf("s4 = %v, len(s4) = %d, 容量cap(s4) = %d \n", s4, len(s4), cap(s4))

	s5 := append(s4, 10)
	s6 := append(s5, 11)
	s7 := append(s6, 12)
	// s6 and s7 不再是对 arr的一个view 添加元素时如果超越cap，系统会进行动态扩容，重新分配更大的底层数组
	// append操作len会改变，如果动态扩容可能会造成cap值和ptr的改变，所以必须要接收append的返回值
	fmt.Println("s5, s6, s7 =", s5, s6, s7)
	fmt.Println(arr2) // [0 1 2 3 4 5 6 10]
}
