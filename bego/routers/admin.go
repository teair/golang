package routers

import (
	"bego/admin"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

// 后台路由配置文件
func init() {

	/*web.Get("/", func(ctx *context.Context) {
		if ctx.Input.Domain() != "admin.wsst.vip" {
			ctx.Redirect(401,"/")
		}
	})*/

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
			ctx.CheckXSRFCookie()
		}),
		// 测试控制器
		web.NSRouter("/test", &admin.TestController{}, "get:Test"),
		// 登录页
		web.NSRouter("/", &admin.LoginController{}, "get:Index"),
		web.NSRouter("/login", &admin.LoginController{}, "post:Post"),
	)
	web.AddNamespace(ns)
}
