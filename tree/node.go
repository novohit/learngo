package tree

import (
	"fmt"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// Print (node TreeNode) 是一个值接收者相当于this 值接收者会拷贝一份
func (node Node) Print() {
	fmt.Print(node.Val, " ")
}

// SetValue 指针接收者才能改变值 go语言无论是结构体还是指针都是用.调用，结构过大也考虑用指针接收者
func (node *Node) SetValue(val int) {
	if node == nil {
		fmt.Println("arg is nil")
		return
	}
	node.Val = val
}

// Print 这两种写法都是一样的，只不过是调用的方式不同
func Print(node Node) {
	fmt.Println(node.Val)
}

// CreateNode go语言没有构造函数，自己定义工厂函数 go语言可以返回局部变量地址 C/C++是不行的
// go语言局部变量在哪分配是不确定的，如果局部变量不会被外部使用就在栈分配，如果需要被外部使用则用堆分配
func CreateNode(val int) *Node {
	return &Node{Val: val}
}
