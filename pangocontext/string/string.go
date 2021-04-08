package string

import (
	"fmt"
	"strconv"
	"strings"
)

func StringMain() {

	// 判断字符串 s 中是否包含字符 a(包括汉字)
	s := "abcde我fghi"
	bol1 := strings.Contains(s, "是")
	fmt.Println("有没有这个东西:", bol1) // true

	// 查找字符串 s 中 我 出现位置的索引值并返回
	index := strings.Index(s, "我")
	fmt.Println("终于找到你:", index)

	// 重复字符串count次，并返回重复的字符串
	repeat := strings.Repeat(s[index:(index+3)], 3)
	fmt.Println("重复三次找到的字符串:", repeat)

	// 在字符 s 中将 old 字符串替换为 new 字符串，n 表示替换的次数，小于 0 表示全部替换
	fmt.Println("将你替换为我:", strings.Replace("你是一个大大的帅哥,你吃了吗,你是猪吗？", "你", "我", -1))

	// 将字符串按照指定字符分割为slice
	fmt.Println("把这个字符串给我炸开:", strings.Split("a.b.c.d.e", "."))

	// 在字符串的头部和尾部切除指定的字符串
	fmt.Println("在字符串的头部和尾部切除指定的字符串:", strings.Trim(" 我前后都有个空格 ", " "))

	// 字符串去除空格并按照空格分割为slice
	fmt.Println("字符串去除空格并按照空格分割为slice:", strings.Fields("我是 根 据 空 格 分 割 的"))

	// 把切片slice通过-连接成字符串
	slice := []string{"I", "Love", "You"}
	fmt.Println("给我连起来:", strings.Join(slice, "-"))
}

// 字符串转换
func StrconvMain() {

	// Append 系列函数：把整数等其他类型转换为字符串追加到现有的字节数组中
	// str := make([]byte, 0, 100)
	// str = strconv.AppendInt(str, 456789, 10)  // 追加整数456789的字符串形式到str中，因为str为byte类型
	// str = strconv.AppendBool(str, false)      // 追加布尔的字符串形式到str中
	// str = strconv.AppendQuote(str, "abcdefg") // 追加字符类型数据到str中
	// str = strconv.AppendQuoteRune(str, 'z')   // AppendQuoteRune将由QuoteRune生成的表示符文的单引号Go字符文字附加到dst，并返回扩展缓冲区。
	// fmt.Printf("%s", string(str))

	// Format 系列函数：把其他数据类型转换为字符串
	// a := strconv.FormatBool(false)
	// b := strconv.FormatFloat(123.23, 'g', 12, 64)
	// c := strconv.FormatInt(1234, 10)
	// d := strconv.FormatUint(12345, 10)
	// e := strconv.Itoa(1023)
	// fmt.Printf("%T\n %T\n %T\n %T\n %T\n", a, b, c, d, e)

	// Parse 系列函数：把字符串转换为其他类型
	a, err := strconv.ParseBool("false")
	checkError(err)
	b, err := strconv.ParseFloat("123.64", 64)
	checkError(err)
	c, err := strconv.ParseInt("12345", 10, 64)
	checkError(err)
	d, err := strconv.ParseUint("12345", 10, 64)
	checkError(err)
	e, err := strconv.Atoi("1023")
	checkError(err)
	fmt.Printf("%T %T %T %T %T", a, b, c, d, e)
}

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
