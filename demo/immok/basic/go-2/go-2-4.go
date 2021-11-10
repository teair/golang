package go_2

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"
)

func read() {
	/*const filename = "immok/abc.txt"
	if contents,err := ioutil.ReadFile(filename);err != nil {
		fmt.Println(err)
	} else {
		//fmt.Printf("%s\n",contents)
		fmt.Printf(string(contents))
	}*/

	score := 102

	switch {
	case score <= 100:
		fmt.Println("B")
	case score < 90:
		fmt.Println("C")
	case score < 80:
		fmt.Println("D")
	case score < 0 || score > 100:
		panic(fmt.Sprintf("wrong score: %d", score))
	default:
		panic(fmt.Sprintf("wrong score: %d", score))
	}
}

func convertToBin(n int) string {
	res := ""
	if n == 0 {
		return strconv.Itoa(n)
	}
	for ; n > 0; n /= 2 {
		tmp := n % 2
		res = strconv.Itoa(tmp) + res
	}
	return res
}

func PrintFile() {
	filename := "immok/basic/go-2/abc.txt"

	// 打开并获取到关联文件
	file, err := os.Open(filename)
	if err != nil {
		panic("open error")
	}

	printFileContents(file)

	s := `aaa
		bbbb
		ccc

		ddd
		我是自定义字符串
		`
	// 字符串转Reader类型
	printFileContents(strings.NewReader(s))
}

// printFileContents 读取文件内容
func printFileContents(reader io.Reader) {
	// 创建并返回一个对操作文件的扫描器,默认的分割函数是 ScanLines
	scanner := bufio.NewScanner(reader)

	// 逐行扫描文件
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func tconvertToBin() {
	fmt.Println(
		convertToBin(5),
		convertToBin(13),
		convertToBin(15),
		convertToBin(0),
		convertToBin(1),
		convertToBin(2),
	)
}

func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("unsupported operation:" + op)
	}
}

// 返回 a 的 b 次方(相当于幂函数)
func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

//func apply(op func(int,int,string)  int, a,b int, c string) int {
func apply(op func(int, int) int, a, b int) int {
	// 通过 reflect 包获取到该函数的指针
	p := reflect.ValueOf(op).Pointer()
	// 通过 runtime 包与指针获取到该函数的名称
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args "+
		"(%d,%d)\n", opName, a, b)
	return op(a, b)
	//return op(a,b,c)
}

func sum(numbers ...int) (n int) {
	for _, v := range numbers {
		n += v
	}
	return
}
