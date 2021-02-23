package spider

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

var (
	// w代表大小写字母+数字+下划线
	reEmail = `\w+@\w+\.\w+`
	// s?有或者没有s
	// +代表出1次或多次
	//\s\S各种字符
	// +?代表贪婪模式
	reLinke  = `href="(https?://[\s\S]+?)"`
	rePhone  = `1[3456789]\d\s?\d{4}\s?\d{4}`
	reIdcard = `[123456789]\d{5}((19\d{2})|(20[01]\d))((0[1-9])|(1[012]))((0[1-9])|([12]\d)|(3[01]))\d{3}[\dXx]`
	reImg    = `https?://[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`
)

// HandleError 错误异常处理
func HandleError(err error, s string) {
	if err != nil {
		fmt.Println(err, s)
	}
}

// GetEmail 爬取邮箱目标网站内容
func GetEmail(url string) {

	// 获取目标网站内容
	pageStr := GetPageStr(url)

	// 过滤邮箱
	re := regexp.MustCompile(reEmail)

	result := re.FindAllStringSubmatch(pageStr, -1)

	for _, result := range result {
		fmt.Println(result)
	}
}

// GetMobile 爬取手机数据列表
func GetMobile(url string) {

	// 获取目标网站body中的所有内容
	pageStr := GetPageStr(url)

	// 过滤手机号
	re := regexp.MustCompile(rePhone)

	// 获取过滤后的手机列表
	result := re.FindAllStringSubmatch(pageStr, -1)

	for _, res := range result {
		fmt.Println(res[0])
	}
}

// GetIDCard 爬取身份证
func GetIDCard(url string) {

	// 获取目标网站body中的所有内容
	pageStr := GetPageStr(url)

	re := regexp.MustCompile(reIdcard)

	result := re.FindAllStringSubmatch(pageStr, -1)

	for _, re := range result {
		fmt.Println(re[0])
	}
}

// Spider Go爬取页面数据
func Spider() {

	// 爬取邮箱列表数据
	// GetEmail("https://tieba.baidu.com/p/6051076813?red_tag=1573533731")

	// 爬取手机号
	// GetMobile("https://www.zhaohaowang.com/")

	// 爬取身份证号
	GetIDCard("https://henan.qq.com/a/20171107/069413.htm")
}

// GetPageStr 获取指定页面内容
func GetPageStr(url string) (pageStr string) {

	// 1. 请求指定url连接
	resp, err := http.Get(url)

	HandleError(err, "http error!")

	defer resp.Body.Close()

	// -1 全部读取
	bytes, err := ioutil.ReadAll(resp.Body)

	HandleError(err, "ReadAll error!")

	// 字节转字符串
	pageStr = string(bytes)

	return pageStr
}
