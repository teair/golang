package routers

import (
	"bego/home"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

// 前台路由配置文件
func init() {
	ns := web.NewNamespace("/home",
		web.NSCond(func(ctx *context.Context) bool {
			if ctx.Input.Domain() == "www.wsst.vip" {
				return true
			}
			return false
		}),
		web.NSBefore(func(ctx *context.Context) {

		}),
		web.NSRouter("/test/:string", &home.TestController{}, "get:Test"),
	)
	web.AddNamespace(ns)
}
