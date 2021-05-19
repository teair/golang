package routers

import (
	"bego/admin"
	"github.com/beego/beego/v2/server/web"
)

// 后台路由配置文件
func init() {
	// 自动匹配路由
	web.AutoRouter(&admin.LoginController{})
	// 登录页路由
	//web.Router("/login",&admin.LoginController{})
}
