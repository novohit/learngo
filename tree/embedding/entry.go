package main

import (
	"fmt"
	"learngo/tree"
)

//  用内嵌扩充已有类型 内嵌其实就是一个go语言特殊的语法糖，它与继承很相似，但是本质上是不一样的
// 可以通过myTreeNode. 出来*tree.Node中的成员变量和成员方法
type myTreeNode struct {
	*tree.Node // Embedding
}

// InOrder 与继承很相似，可以实现类似重载的功能 Shadowed
func (myNode *myTreeNode) InOrder() {
	fmt.Println("Is Shadowed Methods")
}

func (myNode *myTreeNode) postOrder() {
	// myNode结构体可以也会包装了一个空的node
	if myNode == nil || myNode.Node == nil {
		return
	}

	left := myTreeNode{myNode.Left}
	right := myTreeNode{myNode.Right}
	left.postOrder()
	right.postOrder()
	myNode.Print()
}

func main() {
	var root myTreeNode // 通过内嵌可以直接使用myTreeNode
	root = myTreeNode{&tree.Node{Val: 3}}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left = tree.CreateNode(4)

	root.Print()

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
	root.InOrder()      // 此时调用的是myTreeNode的Shadowed方法
	root.Node.InOrder() // 这样就可以拿到node的InOrder
	fmt.Println()

	// 与继承的本质区别就是，继承可以使用多态，父类引用子类对象，而内嵌是不可以的 go语言是通过接口来实现这种能力的
	/* 这样是不行
	var son *tree.Node
	root := &son
	*/

	root.postOrder()
}
