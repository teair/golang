package home

import "github.com/beego/beego/v2/server/web"

type TestController struct {
	web.Controller
}

func (t *TestController) Test() {
	t.Ctx.WriteString("home test!")
}
