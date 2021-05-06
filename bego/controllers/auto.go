package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type AutoController struct {
	beego.Controller
}

func (c *AutoController) Login() {
	c.Ctx.WriteString("this is login\n")

	// 接收所有携带的参数
	//mps := c.Ctx.Input.Params()

	// 获取第一个参数
	//c.Ctx.WriteString(mps["0"]+"\n")
	// 获取第二个参数
	//c.Ctx.WriteString(mps["1"]+"\n")
	// 获取第三个参数
	//c.Ctx.WriteString(mps["2"]+"\n")

	// 获取url后缀名
	c.Ctx.WriteString(c.Ctx.Input.Param(":ext"))

	/*for _,v := range mps {
		c.Ctx.WriteString(v+"\n")
	}*/
}

func (c *AutoController) Logout() {
	c.Ctx.WriteString("this is logout")
}
