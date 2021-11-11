package testdefer

import (
	"bufio"
	"demo/immok/functional/fib"
	"fmt"
	"os"
)

func tryDefer() {
	defer fmt.Println(1) // defer 类似栈,先进后出(FILO)
	defer fmt.Println(2)
	fmt.Println(3)
	return
	//panic("error occured")	// defer 会在panic之前执行
	fmt.Println(4)
}

func calcDefer() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("printed too many")
		}
	}
}

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file) // 如果直接写入文本会比较慢,所以可以先写入buffer缓冲取(内存),满了之后在写入文本
	defer writer.Flush()            // 写入内存

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f()) // 写入文本
	}
}

// Main 测试defer
func Main() {
	tryDefer()
	//writeFile("fib.txt")
	//calcDefer()
}
