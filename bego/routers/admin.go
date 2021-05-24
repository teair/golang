package routers

import (
	"bego/admin"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

// 后台路由配置文件
func init() {
	// 注册后台命名空间
	ns := web.NewNamespace("/admin",
		web.NSCond(func(ctx *context.Context) bool {
			if ctx.Input.Domain() == "www.wsst.vip" {
				return true
			}
			return false
		}),
		web.NSBefore(func(ctx *context.Context) {
			web.SetStaticPath("/static", "static")
		}),
		// 重置已经注册的模型struct
		web.NSGet("/reset", func(ctx *context.Context) {
			orm.ResetModelCache()
			ctx.WriteString("注册模型已重置!")
		}),
		// 测试控制器
		web.NSRouter("/test", &admin.TestController{}, "get:Test"),
		// 登录页
		web.NSRouter("/", &admin.LoginController{}, "get:Index"),
		web.NSRouter("/login", &admin.LoginController{}, "post:AdminLogin"),
		// 首页
		web.NSRouter("/Index", &admin.IndexController{}, "get:Index"),
	)
	web.AddNamespace(ns)
}
