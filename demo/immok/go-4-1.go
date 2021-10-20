package immok

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

// 为结构体定义方法,值传递
func (node treeNode) print() {
	fmt.Println(node.value)
}

// 为结构体定义方法,带*号代表指针传递
func (node *treeNode) setValue(value int) {
	node.value = value
}

// 遍历结构体
func (node *treeNode) traverse() {
	if node == nil {
		return
	}

	// 中序遍历=>先左、再中、再右
	// 遍历左子树
	node.left.traverse()

	node.print()

	// 遍历右子树
	node.right.traverse()
}

// 工厂函数
func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

func structMain() {

	// 创建结构
	var root treeNode
	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil} // 创建并初始化
	root.right.left = new(treeNode)     // new 返回 treeNode 的地址
	root.left.right = createNode(2)
	root.right.left.setValue(4)

	root.traverse()

	//fmt.Println(root)

	//root.print()

	// 由此可见为值传递
	//root.setValue(4)
	//root.print()

	// 那么如何实现引用传递呢(对于传递的结构体加个 * 号即可)
	//root.setValue(4)	// 此时传递的是 root 的地址
	//root.print()			// 此时传递的是值也就是拷贝
	//
	//fmt.Println()

}
