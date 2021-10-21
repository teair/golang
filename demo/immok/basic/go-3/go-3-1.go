package go_3

import "fmt"

// 数组
func arrays() {

	var ar1 [5]int
	ar1 = [5]int{1, 2, 3, 4, 5}
	ar3 := [...]int{2, 4, 5, 6, 7, 8, 9}

	fmt.Println(ar1)
	fmt.Println(ar3)
	ar4 := [3][2]int{{1, 10}, {2}, {3}} // 3 个长度为 2 的 int 数组
	fmt.Println(ar4)
}

func updateSlice(s []int) {
	s[0] = 88
}

// 切片 slice
func Slices() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("origin arr:%v, len(arr):%d, cap(arr):%d\n", arr, len(arr), cap(arr))
	//fmt.Println("arr[2:6] = ", arr[2:7])
	//fmt.Println("arr[2:] = ", arr[2:])
	//fmt.Println("arr[:6] = ", arr[:7])
	//fmt.Println("arr[:] = ", arr[:])

	s1 := arr[:]
	fmt.Println("originSlice:", s1)
	updateSlice(s1)
	fmt.Println("updateSlice:", s1)
	s2 := s1[:5]
	fmt.Println("reSlice s2:", s2)
	s2 = s2[2:]
	fmt.Printf("reSlice s2 s2[2:] %v, len(s1)=%d, cap(s1)=%d\n",
		s2, len(s2), cap(s2))
	s3 := s2[2:5]
	fmt.Printf("extending slice: %v, len(s3)=%d cap(s3)=%d\n",
		s3, len(s3), cap(s3)) // 它指向的始终是原始数组,切片可以向后扩展不可以向前

	s4 := s3[1:3]
	fmt.Printf("the s4: %v, len(s4):%d, cap(s4):%d\n",
		s4, len(s4), cap(s4))
	// 向切片追加元素时,如果该切片元素数量小于切片cap容量,则追加的元素会覆盖掉原来该位置的元素,此操作
	// 会对底层数组造成影响;
	// 如果切片元素数量等于切片容量cap,追加元素操作会导致该切片扩容,也就是分配一块新的内存,不会对原切片的
	// 底层数组造成影响;
	s5 := append(s4, 10)
	s6 := append(s5, 11)
	s7 := append(s6, 12)
	fmt.Println(s5, s6, s7)
	fmt.Printf("appended s4:%v, len(s4):%d, cap(s4):%d\n", s4, len(s4), cap(s4))
	fmt.Printf("appended arr:%v, len(arr):%d, cap(arr):%d\n", arr, len(arr), cap(arr))
}

func printSlice(slice []int) {
	fmt.Printf("%v, len:%d,cap:%d\n", slice, len(slice), cap(slice))
}

func definedSlice() {

	// 定义切片
	var slice []int // Zero value for slice is nil

	for i := 0; i < 100; i++ {
		//printSlice(slice)
		slice = append(slice, 2*i+1)
	}

	slice1 := []int{1, 2, 3, 4} // 创建一个数组并通过指针指向这个数组
	printSlice(slice1)

	slice2 := make([]int, 2) // 创建一个长度为2的切片
	printSlice(slice2)

	slice3 := make([]int, 5, 10) // 创建一个长度为5容量为10的切片
	printSlice(slice3)

	fmt.Println("Copy Slice-->被拷贝的切片在后")
	copy(slice3, slice1)
	printSlice(slice3)

	fmt.Println("Deleting elements from slice")
	slice3 = append(slice3[:2], slice3[3:]...) // 删除下标为2的切片元素
	printSlice(slice3)

	fmt.Println("Poping from front-->从头部删除元素")
	front := slice3[0]  // 头部元素
	slice3 = slice3[1:] // 修改切片
	fmt.Println(front)
	printSlice(slice3)

	fmt.Println("Poping from back-->从尾部删除元素")
	tail := slice3[len(slice3)-1]   // 尾部元素
	slice3 = slice3[:len(slice3)-1] // 修改切片
	fmt.Println(tail)
	printSlice(slice3)

	//fmt.Println(slice)
}

// 数组、切片、容器入口
func Amain() {

	// 数组定义
	//arrays()

	// 切片操作
	//Slices()

	definedSlice()
}
