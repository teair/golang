package regstring

import (
	"fmt"
	"regexp"
)

func RegstringMain() {

	str := "I am learning go language!"

	reg, _ := regexp.Compile("[a-z]{2,4}")
	// reg, _ := regexp.CompilePOSIX("[a-z]{2,4}")

	// 查找符合正则的第一个
	one := reg.Find([]byte(str))
	fmt.Println("Find:", string(one))

	// 查找符合正则的所有slice，n小于0表示返回全部符合的字符串，不然就是返回指定的长度
	all := reg.FindAll([]byte(str), -1)
	fmt.Println("FindAll", fmt.Sprintf("%c", all))

	// 查找符合条件的index位置，开始位置和结束位置
	index := reg.FindIndex([]byte(str))
	fmt.Println("FindIndex", index)

	// 查找符合条件的所有index的位置，n同上
	allindex := reg.FindAllIndex([]byte(str), -1)
	fmt.Println("FindAllIndex", allindex)

	re2, _ := regexp.Compile("am(.*)lang(.*)")

	// 查找Submatch，返回数组，第一个元素是匹配的全部元素，第二个元素是第一个()里面的，第三个是第二个()里面的
	// 下面的输出第一个元素是 "am learning Go language"
	// 第二个元素是 " learning Go "，注意包含空格的输出
	// 第三个元素是 "uage"
	submatch := re2.FindSubmatch([]byte(str))
	fmt.Println("FindSubmatch", fmt.Sprintf("%c", submatch))
	for _, v := range submatch {
		fmt.Println(string(v))
	}

	// 定义和上面的FindIndex一样
	submatchindex := re2.FindSubmatchIndex([]byte(str))
	fmt.Println(submatchindex)

	// FindAllSubmatch,查找所有符合条件的子匹配
	submatchall := re2.FindAllSubmatch([]byte(str), -1)
	fmt.Println(fmt.Sprintf("%c", submatchall))

	// FindAllSubmatchIndex,查找所有符合条件子匹配的index
	submatchallindex := re2.FindAllSubmatchIndex([]byte(str), -1)
	fmt.Println(submatchallindex)

}
