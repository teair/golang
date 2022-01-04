package main

import (
	"fmt"
	"math"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"
	"strconv"
)

func main() {

	f, _ := os.OpenFile("cpu.profile", os.O_CREATE|os.O_RDWR, 0644)
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	//if huiWen(strconv.Itoa(1437342)) {
	//	fmt.Println("huiwen")
	//} else {
	//	fmt.Println("bushi")
	//}

	fmt.Println(primePalindrome(9989900))
}

func reverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func huiWen(n string) bool {

	originStr := []byte(n)

	reverseStr := []byte(reverseString(n))

	len := len(originStr)

	var oStr, rStr string

	for i := 0; i < len; i++ {
		oStr = fmt.Sprintf("%v", originStr[i])
		rStr = fmt.Sprintf("%v", reverseStr[i])
		if oStr != rStr {
			return false
		}
	}

	return true
}

func suShu(n int) bool {

	if n <= 3 {
		return n > 1
	}

	// 不在6的倍数两侧的一定不是质数
	if n%6 != 1 && n%6 != 5 {
		return false
	}

	sorts := int(math.Sqrt(float64(n)))

	for i := 5; i <= sorts; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}

	return true
}

func primePalindrome(n int) int {
	for !suShu(n) || !huiWen(strconv.Itoa(n)) {
		n++
	}
	return n
}
