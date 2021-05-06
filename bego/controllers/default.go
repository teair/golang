package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	dir := "/static"
	path := "/img/a.jpg"
	imgpath := dir + path
	c.Ctx.WriteString(imgpath)
}
