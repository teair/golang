package go_3

import "fmt"

func mapMain() {

	m := map[string]int{
		"PHP":    1,
		"Golang": 2,
		"Java":   3,
		"React":  4,
		"NodeJs": 5,
	}

	// map 的声明
	m1 := make(map[string]int) // empty map
	m2 := map[string]int{}     // nil
	fmt.Println(m1, m2)

	fmt.Println("Traversing map")
	for k, v := range m {
		fmt.Println(k, v)
	}

	java, ok := m["Java"]
	fmt.Println(ok)
	if ok {
		fmt.Println(java)
	}
	fmt.Println("Deleting values")
	delete(m, "Java")
	fmt.Println(m["Java"])
}

var lastOccured = make([]int, 0xffff) // 六进制

// 寻找最长不重复子串
func lengthOfNonRepeatingSubStr(s string) int {

	start := 0 // 最长不重复子串的起始位置

	maxLength := 0 // 最长不重复子串的长度

	//var lastOccured = make(map[rune]int) // 最后一次遇到起始字符的位置
	//lastOccured := make([]int,0xffff)	// 优化

	for j := range lastOccured {
		lastOccured[j] = -1
	}

	for i, chr := range []rune(s) { // rune => int32

		// 如果最后遇到的字符索引大于最长不重复子串的起始位置
		//if lastI, ok := lastOccured[chr]; ok && lastI >= start {
		if lastI := lastOccured[chr]; lastI != -1 && lastI >= start { // 优化
			//start = lastOccured[chr] + 1
			start = lastI + 1 // 优化
		}

		// 如果该不重复子串的长度大于最大不重复子串则更新
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}

		lastOccured[chr] = i
	}

	return maxLength
}

func testLength() {
	fmt.Println(
		lengthOfNonRepeatingSubStr("abcabcdabc"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("bbbb"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("pwwkew"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("c"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("abcdefg"))
}
