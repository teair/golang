package go_2

import (
	"fmt"
	"math"
	"math/cmplx" // math/complx 包提供了复数的常用常数和常用函数
)

// 用 go 验证欧拉公式
func euler() {

	/*// go 表示复数
	c := 3 + 4i
	fmt.Println(c)

	// 取绝对值(模)
	fmt.Println(cmplx.Abs(c))*/

	// cmplx.Pow(0,±0)		return 1 + 0i
	// cmplx.Pow(0,c)		如果 image(c) == 0,则当 real(c) < 0 时返回 Inf+0i;否则返回 Inf+Inf i
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)

	// cmplx.Exp(x)			返回e**x
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

func triangle() {
	var a, b int = 3, 4
	fmt.Println(CalcTriangle(a, b))
}

func CalcTriangle(a, b int) (c int) {
	// 强制类型转换(将返回的float64强制转换a为int类型)
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

func consts() {
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)

	var c int

	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	/*
		// 普通枚举类型
		const (
			cpp		=	iota
			_
			golang
			python
		)
		fmt.Println(cpp,golang,python)*/

	// 自增枚举类型
	// b kb mb gb tb pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}
