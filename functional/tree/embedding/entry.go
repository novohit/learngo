package main

import (
	"fmt"
	"learngo/functional/tree"
)

type myTreeNode struct {
	*tree.Node // Embedding
}

func main() {
	var root myTreeNode // 通过内嵌可以直接使用myTreeNode
	root = myTreeNode{&tree.Node{Val: 3}}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left = tree.CreateNode(4)

	// 中序遍历
	root.InOrder()
	fmt.Println()

	// 利用函数式编程，在遍历时计算结点个数
	nodeCount := 0
	root.InOrderFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println("Node Count :", nodeCount)
}
