package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type RController struct {
	beego.Controller
}

func (c *RController) Get() {
	id := c.Ctx.Input.Param(":id")
	c.Data["id"] = id
	c.TplName = "admin/index.html"
}

func (c *RController) Post() {
	id := c.Ctx.Input.Param(":id")
	c.Data["id"] = id
	c.TplName = "admin/index.html"
}
