package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type NoteController struct {
	web.Controller
}

func (this *NoteController) Prepare() {
	this.EnableXSRF = false
}

// 注解路由【修改注解路由后需要重启方可生效,当beego框架版本为2.X的时候xsrf只支持HTTPS协议】
func (this *NoteController) URLMapping() {
	this.Mapping("StaticBlock", this.StaticBlock)
	this.Mapping("AllBlock", this.AllBlock)
}

// @router	/staticblock/:key [get]
func (this *NoteController) StaticBlock() {
	this.TplName = "common/xsrf.html"
	//this.Ctx.WriteString("???????")
}

// @router /all/:key [post]
func (this *NoteController) AllBlock() {
	this.Ctx.WriteString("???")
}
