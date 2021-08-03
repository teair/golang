package main

import (
	test2 "demo/test"
	"fmt"
	"reflect"
	"sort"
)

func main() {
	test2.TestMain()
}

func app() func(string) string {
	t := "Hi"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	fmt.Println(reflect.TypeOf(t))
	fmt.Println(reflect.TypeOf(c))
	return c
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

