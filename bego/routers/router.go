package routers

import (
	"bego/admin"
	"bego/controllers"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func init() {

	// beego 基本Get路由
	beego.Get("/", func(ctx *context.Context) {
		ctx.Output.Body([]byte("Hello,Get"))
	})

	// beego 基本Post路由
	beego.Post("/alice", func(ctx *context.Context) {
		ctx.Output.Body([]byte("Hello,Post"))
	})

	// beego 注册一个可以相应任何HTTP的路由
	beego.Any("/foo", func(ctx *context.Context) {
		ctx.Output.Body([]byte("来者不拒!"))
	})

	// beego 固定路由
	beego.Router("/home", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/admin", &admin.MainController{})
	beego.Router("/admin/article", &admin.ArticleController{})

	// 正则路由
	beego.Router("api/?:id", &controllers.RController{}) // api/12与api/ 都可匹配
	beego.Router("api/:id", &controllers.RController{})  // api/12可匹配,api/无法匹配
}
