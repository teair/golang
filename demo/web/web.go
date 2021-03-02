package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strings"
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
		t, _ := template.ParseFiles("login.html")
		fmt.Println(t.Execute(w, nil))
		// fmt.Println(t)

	} else {
		v := url.Values{}
		v.Set("aa", "11")
		v.Add("bb", "22")

		// 请求的是登陆数据，那么执行登陆的逻辑判断
		r.ParseForm()
		fmt.Println(r.Form)
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])

		// fmt.Println(v)
		// fmt.Println(v.Get("aa"))
		// fmt.Println(v.Get("bb"))
	}

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
