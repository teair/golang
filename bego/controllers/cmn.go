package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type CmnController struct {
	beego.Controller
}

func (cmn *CmnController) CmnGet() {
	//c.Data["username"] = "shentongtong"
	//c.Ctx.WriteString("什么东西？")
	cmn.TplName = "admin/cmn.html"
}

func (cmn *CmnController) CmnTest() {
	cmn.Ctx.WriteString("快到碗里来!")
}
