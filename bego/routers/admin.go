package routers

import (
	"bego/admin"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

// 后台路由配置文件
func init() {
	ns := web.NewNamespace("/admin",
		web.NSCond(func(ctx *context.Context) bool {
			if ctx.Input.Domain() == "www.wsst.vip" {
				return true
			}
			return false
		}),
		web.NSRouter("/test/:string", &admin.TestController{}, "get:Test"),
	)
	web.AddNamespace(ns)
}
