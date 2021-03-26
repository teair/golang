package web

import (
	"bytes"
	"crypto/md5"
	"demo/beeku"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"text/template"
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

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// 生成token值
		curtime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(curtime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("web/view/upload.html")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666) // 此处假设当前目录下已存在test目录
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

// UploadSpeessd 模拟客户端上传文件
func UploadSpeessd(filename, targeturl string) error {

	// bytes包实现了操作[]byte的常用函数(类似string包)
	// Buffer是一个实现了读写方法的可变大小的字节缓冲。本类型的零值是一个空的可用于读写的缓冲
	bodyBuf := &bytes.Buffer{}

	// multipart 实现了MIME的multipart解析，适用于http和常见浏览器生成的multipart主题
	// NewWrite 函数返回了一个设定了一个随机边界的Writer，数据写入bodyBuf
	bodyWriter := multipart.NewWriter(bodyBuf)

	// CreatePart 方法使用提供的header创建创建一个新的multipart记录，该记录的主体应该写入返回的Writer接口。调用本方法之后，
	// 任何之前的记录都不能再写入
	// CreateFormFile 是CreatePart方法的包装，使用给出的属性名和文件名创建一个新的form-data头
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)

	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}

	// os 包提供了操作系统函数的不依赖平台的接口。设计为Unix风格的，虽然错误处理是Go风格的;失败的调用会返回错误值而非错误码。
	// 通常错误值里包含更多信息。例如：如果某个使用一个文件名的调用(如Open、Stat)失败了，打印错误时会包含该文件名，错误类型将
	// 为 *PathError ，其内部可以解包获得更多信息

	// os 包的接口规定为再所有操作系统中都是一致的。非公用的属性可以从操作系统特定的syscall包获取
	fh, err := os.Open(filename)

	if err != nil {
		fmt.Println("error opening file")
		return err
	}

	defer fh.Close()

	// io.Copy 将 fh 的数据拷贝到fielWriter，直到在fh上到达EOF或发生错误，返回拷贝的字节数和第一个错误
	// 对成功的调用，返回值err为nil而非EOF,因为Copy定义为从fh读取直到EOF,它不会将读取到EOF视为应报告的错误。
	// 如果fh实现了WriterTo接口，本函数会调用fh.WriterTo(fileWriter)进行拷贝;否则如果fileWriter实现了ReaderFrom
	// 接口，本函数会调用fileWriter.ReaderFrom(fh)进行拷贝
	_, err = io.Copy(fileWriter, fh)

	if err != nil {
		return err
	}

	// 方法返回bodyWriter对应的HTTP multipart请求的Content-type的值，多以multipart/form-data起始
	contentType := bodyWriter.FormDataContentType()

	// Close 方法结束multipart信息，并将结尾的边界写入底层io.Writer接口
	bodyWriter.Close()

	// Post 向指定的 targeturl 发出一个POST请求。contentType 为POST数据的类型，bodyBuf 为POST数据作为请求的主体。
	// 如果参数bodyBuf实现了io.Closer接口，它会在发送请求后关闭。调用者有责任在读取完返回值resp的主体后关闭它。
	resp, err := http.Post(targeturl, contentType, bodyBuf)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	// ReadAll 从resp.Body读取数据直到EOF或者遇到error，返回读取的数据和遇到的错误。成功的调用返回的是nil而非EOF。
	// 因为本函数定义为读取resp.Body直到EOF，它不会将读取返回的EOF视为应报告的错误
	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	fmt.Println(resp.Status)

	fmt.Println(string(respBody))

	return nil
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) // 获取请求的方法
	if r.Method == "GET" {
		cstZone := time.FixedZone("CST", 8*3600) // 东八
		fmt.Println(time.Now().In(cstZone).Format("2006-01-02 15:04:05"))

		// 判断日期和时间是否合法
		// t := time.Date(2021, time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), 0, time.UTC)
		// fmt.Printf("Go launched at %s\n", t.Local())

		// 防止重复提交的唯一token
		curtime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(curtime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil)) // 得到token并存储到session中

		t, _ := template.ParseFiles("web/view/login.html")

		//给静态页赋值
		t.Execute(w, token)

	} else {

		// 请求的是登陆数据，那么执行登陆的逻辑判断
		r.ParseForm()

		// if len(r.FormValue("username")) == 0 {
		// 	fmt.Fprintf(w, "请输入用户名")
		// 	return
		// }

		// func HTMLEscape(w io.Writer, b []byte) 把b进行转义后写入w
		template.HTMLEscape(w, []byte(r.Form.Get("xss")))

		// func HTMLEscapeString(s string) string 转义s之后返回字符串
		fmt.Fprintln(w, template.HTMLEscapeString(r.FormValue("xss")))

		// 输出看起来正常的js信息
		fmt.Fprintln(w, r.FormValue("xss"))

		token := r.Form.Get("token")

		if token != "" {
			// 验证token的合法性(与session中的token进行比对)

		} else {
			fmt.Fprintln(w, "token验证失败!")
			return
		}

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

	http.HandleFunc("/upload", upload) // 设置上传页路由

	err := http.ListenAndServe(":9090", nil) // 设置监听的端口

	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

	// 模拟客户端表单功能支持文件上传
	// filename := "web/static/2.jpg"
	// UploadSpeessd(filename, "http://127.0.0.1:9090")

}
