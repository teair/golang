package tree

import "fmt"

// Traverse 遍历结构体(中序遍历)
func (node *Node) Traverse() {
	node.TraverseFunc(func(n *Node) {
		n.Print() // 打印这个操作变为动态的【此处是打印每一个树节点】
	})
	fmt.Println()
}

func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}
