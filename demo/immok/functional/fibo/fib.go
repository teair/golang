package fibo

import (
	"bufio"
	"demo/immok/functional/fib"
	"fmt"
	"io"
	"strings"
)

// 计算前两个数之和
func fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// 可一给函数实现接口,也就是说函数是一等公民
type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 1000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

//FibMain 斐波那契数列
func FibMain() {
	var f intGen = fib.Fibonacci()
	printFileContents(f)
}
