package routers

func init() {
	//// beego 基本Get路由
	///*beego.Get("/", func(ctx *context.Context) {
	//	ctx.Output.Body([]byte("Hello,Get"))
	//})*/
	//
	//beego.Get("/", func(ctx *context.Context) {
	//	ctx.WriteString("?????????")
	//})
	//
	//// beego 基本Post路由
	///*beego.Post("/alice", func(ctx *context.Context) {
	//	ctx.Output.Body([]byte("Hello,Post"))
	//})
	//
	//// beego 注册一个可以相应任何HTTP的路由
	//beego.Any("/foo", func(ctx *context.Context) {
	//	ctx.Output.Body([]byte("来者不拒!"))
	//})
	//
	//// beego 固定路由
	//beego.Router("/controllers", &controllers.MainController{})
	//beego.Router("/user", &controllers.UserController{})
	//beego.Router("/admin", &admin.MainController{})
	//beego.Router("/admin/article", &admin.ArticleController{})*/
	//
	//// 正则路由
	////beego.Router("api/?:id", &controllers.RController{}) // api/12与api/ 都可匹配
	////beego.Router("api/:id", &controllers.RController{})  // api/12可匹配,api/无法匹配
	////beego.Router("/api/:id([0-9]+)",&controllers.RController{})	// api/:id([0-9]+) 限制id为数字
	////beego.Router("/api/:id([0-9]+)/:username([\\w]+)",&controllers.RController{})	// api/:username([\\w]+) 限制username为字母下划线组合
	////beego.Router("/static/*.*",&controllers.RController{})	// *.* 匹配方式
	////beego.Router("/static/*",&controllers.RController{})	// * 全匹配方式
	////beego.Router("/test/:id:int",&controllers.RController{})	// int 类型设置方式,匹配:id为int类型,框架帮你实现了正则([0-9]+)
	////beego.Router("/test/:username:string",&controllers.RController{})	// 匹配:username 为 string 类型,框架帮你实现了正则([\w]+)
	////beego.Router("shen_:id([0-9]+).html",&controllers.RController{})	// 带有前缀的自定义正则,匹配:id为正则类型,类似:shen_123.html
	//
	//beego.Router("/test/:username:string", &controllers.RController{}, "*:RcGet")
	//beego.Router("/cmn", &controllers.CmnController{}, "*:CmnTest;get,post:CmnGet")
	//
	//// 自动匹配路由
	//beego.AutoRouter(&controllers.AutoController{})
	//
	//// 注解路由
	//beego.Include(&controllers.NoteController{})
	//
	//// namespace
	//ns :=
	//	beego.NewNamespace("v1",
	//		// 支持满足条件的就执行该namespace
	//		beego.NSCond(func(b *context.Context) bool  {
	//			if b.Input.Domain() == "www.wsst.vip" {
	//				b.WriteString("1")
	//				return true
	//			}
	//			b.WriteString("0")
	//			return false
	//		}),
	//		beego.NSBefore(func(this *context.Context) {
	//			this.WriteString("this is before")
	//		}),
	//		beego.NSGet("/notallowed", func(ctx *context.Context) {
	//			//ctx.Output.Body([]byte("notAllowed"))
	//			ctx.WriteString("notAllowed")
	//		}),
	//		beego.NSRouter("/version", &controllers.WatermelonController{}, "get:ShowAPIVersion"),
	//		beego.NSRouter("/changepassword", &controllers.WatermelonController{}, "get:ChangePassword"),
	//		beego.NSRouter("/xsrfhtml", &controllers.WatermelonController{}, "get:XsrfHtml"),
	//		beego.NSRouter("/xsrfpost", &controllers.WatermelonController{}, "post:XsrfPost"),
	//		beego.NSNamespace("/shop",
	//			beego.NSGet("/:id", func(ctx *context.Context) {
	//				ctx.Output.Body([]byte("shop/int"))
	//			}),
	//		),
	//		beego.NSNamespace("/cms",
	//			beego.NSInclude(
	//				&controllers.WatermelonController{},
	//				&controllers.NoteController{},
	//				&controllers.UserController{},
	//			),
	//		),
	//		beego.NSRouter("/watermelon", &controllers.WatermelonController{}, "get:WaterFruit"),
	//	)
	//
	//// 注册 namespace
	//beego.AddNamespace(ns)
}
