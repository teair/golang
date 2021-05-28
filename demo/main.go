package main

import (
	"fmt"
	"sort"
)

func main() {

	// slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// typefunc.TypeFunc(slice)

	// fmt.Println("=============")

	// 结构体初始化
	// gostruct.TestStruct()

	// 结构体对象
	// gostruct.CalcArea()

	// 盒子对象
	// gostruct.ExtendMethod()

	a := []int{1,2,3,6}
	b := []int{1,3,4}

	dfa,dfb := test(a,b)
	fmt.Println(dfa,dfb)
	//web.Webmain()
}

func test(a,b []int) (diffA, diffB []int) {
	sort.Ints(a)
	sort.Ints(b)
	i,n := 0, len(a)
	j,m := 0, len(b)
	for {
		if i == n {
			diffB = append(diffB,b[j:]...)
			return
		}
		if j == m {
			diffA = append(diffA,a[i:]...)
			return
		}
		x , y := a[i], b[j]
		if x < y {
			diffA = append(diffA,x)
			i++
		} else if x > y {
			diffB = append(diffB,y)
			j++
		} else {
			i++
			j++
		}
	}
}
