package go_2

import "fmt"

// 指针类型典例:交换数据
//func swap(a,b *int) {
func swap(a, b int) (int, int) {
	//*a,*b = *b,*a		// 将a所指向的类型改为b所指向的类型,b所指向的类型指向a所指向的类型
	return b, a
}

// Pmain 测试指针入口
func Pmain() {
	a, b := 3, 4
	//swap(&a,&b)	// 通过指针的方式交换数据
	a, b = swap(a, b) // 这种方式相比于指针更好
	fmt.Println(a, b)
}
