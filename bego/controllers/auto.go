package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type AutoController struct {
	beego.Controller
}

func (c *AutoController) Login() {
	c.Ctx.WriteString("this is login")
}

func (c *AutoController) Logout() {
	c.Ctx.WriteString("this is logout")
}
