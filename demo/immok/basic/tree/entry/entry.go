package entry

import (
	"demo/immok/basic/tree"
	"fmt"
)

// myTreeNode 组合扩展业务
type myTreeNode struct {
	//node *tree.Node
	*tree.Node // 通过内嵌扩展已有类型
}

// postOrder 后序遍历
func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.Node == nil {
		return
	}

	leftNode := myTreeNode{myNode.Left}
	leftNode.postOrder()
	rightNode := myTreeNode{myNode.Right}
	rightNode.postOrder()
	myNode.Print()
}

func (node *myTreeNode) Traverse() {
	fmt.Printf("this is entry Traverse!")
}

func TreeMain() {

	// 创建结构
	//var root tree.Node
	//root = tree.Node{Value: 3}
	root := myTreeNode{&tree.Node{Value: 3}}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{Value: 5} // 创建并初始化
	root.Right.Left = new(tree.Node)  // new 返回 Node 的地址
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)

	// 包中实现中序遍历
	//root.Traverse() // root 本身的 Traverse
	//fmt.Println()

	root.Node.Traverse() // root.Node的 Traverse,他们虽然看起来一样,其实是个语法糖

	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++ // 遍历树节点的同时计算节点数量
	})
	fmt.Println(nodeCount)

	//fmt.Println()

	// 自定义实现后序遍历
	//myRoot := myTreeNode{&root}
	//myRoot.postOrder()
	//root.postOrder()
}
