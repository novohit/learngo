package tree

//为一个结构定义的方法必须放在同一个包内，但是可以是不同文件

// InOrder 值接收值是不能接收nil的，所以要用指针接收者
func (node *Node) InOrder() {
	node.InOrderFunc(func(n *Node) {
		n.Print()
	})
}

// 参数传一个函数
func (node *Node) InOrderFunc(f func(node *Node)) {
	if node == nil {
		return
	}
	node.Left.InOrderFunc(f)
	// 执行该函数
	f(node)
	node.Right.InOrderFunc(f)
}
