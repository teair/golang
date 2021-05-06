package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type RController struct {
	beego.Controller
}

func (c *RController) RcGet() {
	//c.Data["id"] = c.Ctx.Input.Param(":id")
	c.Data["username"] = c.Ctx.Input.Param(":username")

	/*c.Data["ext"] = c.Ctx.Input.Param(":ext")
	c.Data["path"] = c.Ctx.Input.Param(":path")
	staDir,_ := beego.AppConfig.String("WebConfig.StaticDir")
	c.Data["imgpath"] = staDir*/

	/*username := c.Ctx.Input.Param(":username")
	c.Data["username"] = username*/

	c.TplName = "admin/index.html"
}

func (c *RController) Post() {
	data := c.Ctx.Input.Params()
	c.Data["id"] = data["id"]
	c.Data["username"] = data["username"]
	c.TplName = "admin/index.html"
}
