package mistake

import (
	"errors"
	"fmt"
)

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	return f, nil
}

func MistakeMain() {

	f, err := Sqrt(-1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f)

}

func SliceAppend() {
	/*slice1 := []int{1,2,3,4,5}
	slice2 := []int{6,7,8}
	slice1 =  append(slice1,slice2...)	// 往切片中追加切片
	fmt.Println(slice1)
	slice1 = append(slice1,9,10)	// 往切片中追加元素
	fmt.Println(slice1)*/

	// 切片容量的增长方式
	/*s1 := []int{1}
	fmt.Println("s1=>",s1,cap(s1))
	s2 := append(s1,2)
	fmt.Println("s2=>",s2,cap(s2))
	s6 := append(s2,7,8,9)
	fmt.Println("s6=>",s6,cap(s6))
	s3 := append(s2,3)
	fmt.Println("s3=>",s3,cap(s3))
	s4 := append(s3,4)
	fmt.Println("s4=>",s4,cap(s4))
	s5 := append(s3,6)
	fmt.Println("s5=>",s5,cap(s5))*/

	s5 := []int{3, 4, 5, 6}

	// 切片作为参数传递(Go语言的函数参数传递只有值传递没有引用传递)
	f(s5)
	fmt.Println("f(s5)=>", s5, cap(s5))

	// 这里s虽然改变了,但并不会影响外层函数的s(并不会修改底层的数组)
	fs5 := f1(s5)
	fmt.Println(fs5)
	fmt.Println(s5)

	// 会改变底层数组
	f2(&s5)
	fmt.Println(s5)
}

func f1(s []int) []int {
	// 这里s虽然改变了,但并不会影响外层函数的s(并不会修改底层的数组)
	s = append(s, 100)
	return s
}

func f2(s *[]int) {
	*s = append(*s, 200)
	return
}

func f(s []int) {
	for i := range s {
		s[i] += 1
	}
}
