package web

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"shentong/demo/beeku"
	"strings"
	"time"
)

// Hello 搭建web服务
func Hello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 解析参数,默认是不会解析的
	// fmt.Println(r.Form) // 这些信息是输出到服务器端的打印信息
	// fmt.Println("path", r.URL.Path)
	// fmt.Println("scheme", r.URL.Scheme)
	// fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") // 这个写入到w的是输出到客户端的
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) // 获取请求的方法
	if r.Method == "GET" {
		cstZone := time.FixedZone("CST", 8*3600) // 东八
		fmt.Println(time.Now().In(cstZone).Format("2006-01-02 15:04:05"))

		// 判断日期和时间是否合法
		// t := time.Date(2021, time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), 0, time.UTC)
		// fmt.Printf("Go launched at %s\n", t.Local())

		// t, _ := template.ParseFiles("login.html")
		// fmt.Println(t.Execute(w, nil))
		// fmt.Println(t)

	} else {

		// 请求的是登陆数据，那么执行登陆的逻辑判断
		r.ParseForm()

		// if len(r.FormValue("username")) == 0 {
		// 	fmt.Fprintf(w, "请输入用户名")
		// 	return
		// }

		// 正则匹配中文
		if chinese, _ := regexp.MatchString("^\\p{Han}{2,5}$", r.FormValue("username")); !chinese {
			fmt.Fprintln(w, "用户名位2-5位汉字!")
			return
		}

		// 验证用户输入内容
		// _, err := strconv.Atoi(r.Form.Get("password"))

		// if err != nil {
		// 	// 字符串转数字出错，说明不是数字
		// 	fmt.Fprint(w, "密码必须为数字")
		// 	return
		// }

		// if len(r.FormValue("password")) < 3 {
		// 	fmt.Fprintf(w, "密码长度不得小于3位")
		// 	return
		// }

		// if len(r.FormValue("password")) > 6 {
		// 	// 密码长度不得大于6位
		// 	fmt.Fprint(w, "密码长度不得大于6位")
		// 	return
		// }

		// 正则匹配数字
		if m, _ := regexp.MatchString("^[0-9]{3,6}$", r.FormValue("password")); !m {
			fmt.Fprintln(w, "密码为3-6位数字!")
			return
		}

		// 判断喜爱的水果是否存在
		exist := Exists(r.FormValue("favorite"))
		if exist == false {
			fmt.Fprintln(w, "干点正事吧铁铁!")
			return
		}

		// 判断用户是否选择喜爱的运动
		fs := Sports(r.Form["interest"])

		if fs == false {
			fmt.Fprintln(w, "未知的爱好？")
			return
		}

		fmt.Fprintln(w, r.FormValue("username"))
		fmt.Fprintln(w, r.FormValue("password"))
		fmt.Fprintln(w, r.FormValue("favorite"))
	}

}

// Sports 用户喜爱的运动
func Sports(favorite []string) bool {
	slice := []string{"football", "basketball", "tennis"}
	a := beeku.SliceDiff(favorite, slice)
	if a == nil {
		return true
	}
	return false
}

// Exists 判断喜爱的水果类型是否存在
func Exists(s string) bool {

	// 判断下拉选项值是否存在
	slice := []string{"apple", "pear", "banana", "pineapple", "strawberry", "peach", "cherry"}

	for _, fruit := range slice {
		if fruit == s {
			return true
		}
	}

	return false
}

// Webmain web服务入口函数
func Webmain() {

	http.HandleFunc("/", Hello) // 设置访问路由

	http.HandleFunc("/login", login) // 设置访问的路由

	err := http.ListenAndServe(":9090", nil) // 设置监听的端口

	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
