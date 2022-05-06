package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// 1, 1, 2, 3, 5, 8, 13
//    a1 a2
//       a1 a2
// 斐波那契生成器
func fibonacci() intGen {
	a1, a2 := 0, 1
	return func() int {
		a1, a2 = a2, a1+a2
		return a1
	}
}

// 函数也能实现接口
type intGen func() int

// 函数也能作为接收者，接收者只是一种特殊的参数而已
func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	// TODO: incorrect if p is too small
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}
func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
func main() {
	f := fibonacci()

	/*fmt.Println(f()) // 1
	fmt.Println(f()) // 1
	fmt.Println(f()) // 2
	fmt.Println(f()) // 3
	fmt.Println(f()) // 5
	fmt.Println(f()) // 8
	fmt.Println(f()) // 13
	fmt.Println(f()) // 21*/

	printFileContents(f)
}
