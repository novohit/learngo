package main

import (
	"fmt"
)

type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
}

// (node treeNode) 是一个值接收者相当于this 值接收者会拷贝一份
func (node treeNode) print() {
	fmt.Println(node.val)
}

// 指针接收者才能改变值 go语言无论是结构体还是指针都是用.调用，结构过大也考虑用指针接收者
func (node *treeNode) setValue(val int) {
	if node == nil {
		fmt.Println("arg is nil")
		return
	}
	node.val = val
}

// 值接收值是不能接收nil的，所以要用指针接收者
func (node *treeNode) inOrder() {
	if node == nil {
		return
	}
	node.left.inOrder()
	node.print()
	node.right.inOrder()
}

// 这两种写法都是一样的，只不过是调用的方式不同
func print(node treeNode) {
	fmt.Println(node.val)
}

// go语言没有构造函数，自己定义工厂函数 go语言可以返回局部变量地址 C/C++是不行的
// go语言局部变量在哪分配是不确定的，如果局部变量不会被外部使用就在栈分配，如果需要被外部使用则用堆分配
func createNode(val int) *treeNode {
	return &treeNode{val: val}
}
func main() {
	var root treeNode // 注意初始结点不是nil 是一个有零值的结构体{0 <nil> <nil>}
	root = treeNode{val: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.left = createNode(2)

	nodes := []treeNode{
		{val: 3},
		{left: nil},
		{val: 3, left: nil, right: &root},
		{6, nil, &root},
	}
	fmt.Println(nodes)
	fmt.Println(root)
	root.print()
	print(root)
	// 会自动传指针
	root.setValue(100)
	root.print()

	pRoot := &root
	pRoot.print()
	pRoot.setValue(99)
	pRoot.print()

	// 传nil
	var ptr *treeNode
	fmt.Println(ptr)
	ptr.setValue(10)
	root.inOrder()
}
