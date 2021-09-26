package immok

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
