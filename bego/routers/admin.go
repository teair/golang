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
		web.NSRouter("/test", &admin.ApiController{}, "get:TestApi"),
		// 登录页
		web.NSRouter("/", &admin.LoginController{}, "get:Index"),
		web.NSRouter("/login", &admin.LoginController{}, "post:AdminLogin"),
		// 注销登陆
		web.NSRouter("/logout", &admin.LoginController{}, "get:AdminLogout"),
		// 首页
		web.NSRouter("/index/?:mark_online/?:p/?:app_name", &admin.IndexController{}, "get:Index"),
		web.NSRouter("/listen", &admin.IndexController{}, "post:Listen"),
		web.NSRouter("/gamesel", &admin.IndexController{}, "post:Gamesel"),
		web.NSRouter("/gamedel", &admin.IndexController{}, "post:Gamedel"),
		web.NSRouter("/markonline", &admin.IndexController{}, "post:MarkOnline"),
		web.NSRouter("/gameadd", &admin.GameController{}, "post:GameAdd"),
		// 上传页
		web.NSRouter("/upindex/:gid", &admin.UploadController{}, "get:UpIndex"),
		web.NSRouter("/upfile", &admin.UploadController{}, "post:UpFile"),
		web.NSRouter("/upsubmit", &admin.UploadController{}, "post:UpSubmit"),
		// 区服页
		web.NSRouter("/serverlist", &admin.ServerController{}, "get:ServerIndex"),
		web.NSRouter("/editserver", &admin.ServerController{}, "get,post:EditServer"),
		web.NSRouter("/addserver", &admin.ServerController{}, "get,post:AddServer"),
	)
	web.AddNamespace(ns)
}
