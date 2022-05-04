package tree

//为一个结构定义的方法必须放在同一个包内，但是可以是不同文件

// InOrder 值接收值是不能接收nil的，所以要用指针接收者
func (node *Node) InOrder() {
	if node == nil {
		return
	}
	node.Left.InOrder()
	node.Print()
	node.Right.InOrder()
}
