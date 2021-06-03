package home

import (
	"fmt"
	"github.com/beego/beego/v2/server/web"
)

type TestController struct {
	web.Controller
}

func (t *TestController) Test() {
	//mark_online := t.Ctx.Input.Param(":mark_online")
	//app_name := t.Ctx.Input.Param(":app_name")
	//fmt.Println("mark_online:",mark_online)
	//fmt.Println("app_name:",app_name)
	data := t.Ctx.Input.Params()
	fmt.Println("data=>", data)
	t.TplName = "home/test/test.html"
}
