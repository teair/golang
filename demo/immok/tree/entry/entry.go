package entry

import (
	"demo/immok/tree"
	"fmt"
)

// myTreeNode 组合扩展业务
type myTreeNode struct {
	node *tree.Node
}

// postOrder 后序遍历
func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	leftNode := myTreeNode{myNode.node.Left}
	leftNode.postOrder()
	rightNode := myTreeNode{myNode.node.Right}
	rightNode.postOrder()
	myNode.node.Print()
}

func TreeMain() {

	// 创建结构
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{Value: 5} // 创建并初始化
	root.Right.Left = new(tree.Node)  // new 返回 Node 的地址
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)

	// 包中实现中序遍历
	root.Traverse()
	fmt.Println()

	// 自定义实现后序遍历
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
}
