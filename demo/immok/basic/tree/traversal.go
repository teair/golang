package tree

// Traverse 遍历结构体(中序遍历)
func (node *Node) Traverse() {
	if node == nil {
		return
	}

	// 中序遍历=>先左、再中、再右
	// 遍历左子树
	node.Left.Traverse()

	node.Print()

	// 遍历右子树
	node.Right.Traverse()
}
