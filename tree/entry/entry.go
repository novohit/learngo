package main

import (
	"fmt"
	"learngo/tree"
)

/**
一个目录下只能放一个包，所以entry的main包要新建一个entry目录
*/

// go语言没有继承， 用组合扩充已有类型
type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	// myNode结构体可以也会包装了一个空的node
	if myNode == nil || myNode.node == nil {
		return
	}
	/*
		这个是编译器的限制。
		我们这个postOrder需要一个指针接收者
		如果是修改后的写成left.postOrder()，left是一个变量，它会自动取left的地址，作为接收者
		但是我们不定义变量，myTreeNode{...}.postOrder()，它就取不了myTreeNode{...}的地址。这只是一个编译器的限制。
		指针接收者和值接收者，区别与指针参数与值参数一样。只是编译器在处理left.postOrder()等的时候，会帮我们自动把left取地址，然后再调用postOrder。
		错误写法：
		myTreeNode{myNode.node.Left}.postOrder()
		myTreeNode{myNode.node.Right}.postOrder()
	*/
	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}
	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}
func main() {
	var root tree.Node // 注意初始结点不是nil 是一个有零值的结构体{0 <nil> <nil>}
	root = tree.Node{Val: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left = tree.CreateNode(4)

	nodes := []tree.Node{
		{Val: 3},
		{Left: nil},
		{Val: 3, Left: nil, Right: &root},
		{6, nil, &root},
	}
	fmt.Println(nodes)
	fmt.Println(root)
	root.Print()
	tree.Print(root)
	// 会自动传指针
	root.SetValue(100)
	root.Print()

	pRoot := &root
	pRoot.Print()
	pRoot.SetValue(99)
	pRoot.Print()

	// 传nil
	var ptr *tree.Node
	fmt.Println(ptr)
	ptr.SetValue(10)
	// 中序遍历
	root.InOrder()
	fmt.Println()

	//
	myNode := myTreeNode{&root}
	myNode.postOrder()
}
