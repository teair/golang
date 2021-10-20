package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

// Print 为结构体定义方法,值传递
func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

// SetValue 为结构体定义方法,带*号代表指针传递
func (node *Node) SetValue(value int) {
	node.Value = value
}

// CreateNode 工厂函数
func CreateNode(value int) *Node {
	return &Node{Value: value}
}
