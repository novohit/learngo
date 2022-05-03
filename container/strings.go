package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "abc你好世界!!" // UTF-8 三字节一个中文
	fmt.Println(len(str))
	for _, b := range []byte(str) { // []byte()获取字节数组
		fmt.Printf("%X ", b) // %X 16进制输出
	}
	fmt.Println()
	// 将str进行了utf-8的解码，解出来每一个字符转成Unicode存放到4个字节的rune类型里
	for i, ch := range str { // ch is a rune
		fmt.Printf("(%d, %X)", i, ch)
	}
	fmt.Println()
	// a  b  c     你       好        世       界     !  !
	// 61 62 63 E4 BD A0 E5 A5 BD E4 B8 96 E7 95 8C 21 21
	// (0, 61)(1, 62)(2, 63)(3, 4F60)(6, 597D)(9, 4E16)(12, 754C)(15, 21)(16, 21)
	// 你：E4 BD A0 UTF-8转成Unicode就是4F60两字节

	fmt.Println(utf8.RuneCountInString(str)) // RuneCountInString是通过解码去数的过程获取字符数，len()只是获取的字节数

	bytes := []byte(str)
	// while
	for len(bytes) > 0 {
		decodeRune, size := utf8.DecodeRune(bytes) // 将一个一个字符转换为rune
		bytes = bytes[size:]
		fmt.Printf("%d %c ", size, decodeRune)
	}
	fmt.Println()

	// 这里的i拿到的是字符的起始索引
	for i, ch := range str { // ch is a rune
		fmt.Printf("(%d, %c)", i, ch)
	}
	fmt.Println()

	// 把str转成rune再遍历拿到的i才是每一个字符的位置
	// 转成rune实际上是decode出来的一个结果，存放在新的内存，每个rune为4字节
	for i, ch := range []rune(str) {
		fmt.Printf("(%d, %c)", i, ch)
	}
}
