package routers

import (
	"bego/home"
	"github.com/beego/beego/v2/server/web"
)

// 前台路由配置文件
func init() {
	web.Router("/", &home.TestController{}, "get:Test")
}
