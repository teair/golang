package home

import (
	"github.com/beego/beego/v2/server/web"
)

type IndexController struct {
	web.Controller
}

func (this *IndexController) Get() {
	this.Ctx.WriteString("hello world!")
}
