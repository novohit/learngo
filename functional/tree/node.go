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
func CreateNode(val int) *Node {
	return &Node{Val: val}
}
