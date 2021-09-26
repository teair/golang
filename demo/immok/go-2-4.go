package immok

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func printFile() {
	filename := "immok/abc.txt"

	// 打开并获取到关联文件
	file, err := os.Open(filename)
	if err != nil {
		panic("open error")
	}

	// 创建并返回一个对操作文件的扫描器,默认的分割函数是 ScanLines
	scanner := bufio.NewScanner(file)

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
