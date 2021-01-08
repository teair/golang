// 导入GO程序的入门包
package main
// net/http包-->HTPP的基础封装和访问
import(
	"net/http"
)
// 程序执行的入口函数
func main(){
	// 使用http.FileServer文件服务器将当前目录作为根目录(/目录)的处理器，访问根目录就会进入当前目录
	http.Handle("/",http.FileServer(http.Dir("../demo")))
	// 默认的HTTP服务侦听在本机80端口
	http.ListenAndServe(":9797",nil)
}
