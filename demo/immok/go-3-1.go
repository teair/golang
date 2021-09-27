package immok

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
	fmt.Println("arr[2:6] = ", arr[2:7])
	fmt.Println("arr[2:] = ", arr[2:])
	fmt.Println("arr[:6] = ", arr[:7])
	fmt.Println("arr[:] = ", arr[:])

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

	s4 := s3[1:4]
	fmt.Printf("the s4: %v, len(s4):%d, cap(s4):%d",
		s4, len(s4), cap(s4))
}

// 数组、切片、容器入口
func Amain() {

	// 数组定义
	//arrays()

	// 切片操作
	Slices()
}
