package sessioncookie

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/astaxie/beego/session"
	_ "github.com/astaxie/beego/session/redis"
)

var globalSessions *session.Manager

func init() {
	sessionConfig := &session.ManagerConfig{
		CookieName:      "gosessionid",
		EnableSetCookie: true,
		Gclifetime:      60,
		Maxlifetime:     60,
		Secure:          false,
		CookieLifeTime:  60,
		// ProviderConfig:  "./tmp",
		ProviderConfig: "RedisIP:端口,连接池(最大连接数),密码,数据库",
	}
	globalSessions, _ = session.NewManager("redis", sessionConfig)
	go globalSessions.GC()
}

// SessionMain go session 入口
func SessionMain(w http.ResponseWriter, r *http.Request) {
	provider := globalSessions.GetProvider()
	sess, _ := globalSessions.SessionStart(w, r)
	defer sess.SessionRelease(w)
	if r.Method == "GET" {
		fmt.Println(provider)
		t, _ := template.ParseFiles("sessioncookie/view/login.html")
		w.Header().Set("Content-Type", "text/html")
		t.Execute(w, nil)
	} else {
		val := sess.Get("username")
		if val == nil {
			sess.Set("username", r.FormValue("username"))
		} else {
			fmt.Println(val)
		}
		// http.Redirect(w, r, "http://www.baidu.com", 302)
	}
}

// CookieMain go cookie 入口
func CookieMain() {

	http.HandleFunc("/setcookie", cookieparts) // 设置路由

	http.HandleFunc("/getcookie", getcookie)

	http.HandleFunc("/delcookie", DelCookie)

	http.HandleFunc("/setsession", SessionMain)

	err := http.ListenAndServe(":9090", nil) // 设置监听的端口

	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

// cookieparts 设置cookie
func cookieparts(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		expiration := time.Now()
		expiration = expiration.AddDate(1, 0, 0)
		cookie := http.Cookie{Name: "username", Value: "ShenTongtong", Expires: expiration}
		http.SetCookie(w, &cookie)
	}
}

// getcookie 获取cookie
func getcookie(w http.ResponseWriter, r *http.Request) {

	// cookie, _ := r.Cookie("username")
	// fmt.Fprint(w, cookie)

	for _, cookie := range r.Cookies() {
		fmt.Fprintln(w, cookie)
		fmt.Fprintln(w, cookie.Name)
		fmt.Fprintln(w, cookie.Value)
	}
}

// DelCookie 删除Cookie
func DelCookie(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		now := time.Now()
		now = now.AddDate(0, 0, -1)
		cookie := http.Cookie{Name: "username", Expires: now}
		http.SetCookie(w, &cookie) // 删除cookie
	}
}

// Getime
func Getime() {
	now := time.Now()
	fmt.Println(now)
}

// Test 测试  os/signal 包中signal方法(Signal方法向进程发送一个信号,目前在windows中
// 向进程发送Interrupt信号尚未实现)
func Test() {

	m := make(chan int)
	defer close(m)

	sigABRT := make(chan os.Signal, 1)
	defer close(sigABRT)

	// Notify函数让signal包将输入信号转发到sigABRT
	// 如果没有列出要传递的信号，会将所有信号传递到sigABRT，否则只传递列出的信号
	// signal包不会为了向sigABRT发送信息而阻塞(就是说如果发送时sigABRT阻塞了，signal包会直接放弃)
	// 调用者应该保证sigABRT有足够的缓存空间可以跟上期望的信号频率
	// 对于使用单一信号用于通知的通道，缓存为1就够了
	// 可以使用同一通道多次调用Notify：每一次都会扩展该通道接收的信号集。唯一从信号集去除信号的方法是调用Stop。
	// 可以使用同一信号不同通道多次调用Notify：每个通道都会独立接收到该信号的一个拷贝
	signal.Notify(sigABRT, syscall.SIGABRT)

	sigINT := make(chan os.Signal, 1)
	defer close(sigINT)
	signal.Notify(sigINT, syscall.SIGINT)

	sigTERM := make(chan os.Signal, 1)
	defer close(sigTERM)
	signal.Notify(sigTERM, syscall.SIGTERM)

	go func() {
		for {
			select {
			case sign := <-sigABRT:
				fmt.Println("sigABRT:", sign)
				m <- 1
			case sign := <-sigINT:
				fmt.Println("sigINT:", sign)
				m <- 1
			case sign := <-sigTERM:
				fmt.Println("sigTERM:", sign)
				m <- 1
			default:
				fmt.Println("sleep....")
				time.Sleep(time.Second)
			}
		}
	}()
	fmt.Println("waiting...")
	<-m
	fmt.Println("end")
}
