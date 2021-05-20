package admin

import (
	"github.com/beego/beego/v2/server/web"
)

type baseController struct {
	web.Controller
}

func (b *baseController) Prepare() {
	// 后台验证登陆
	if b.GetSession("username") == nil || b.GetSession("password") == nil {
		// 没有登陆跳转至登录页
		//b.Ctx.WriteString("please login!")
		b.Ctx.Redirect(302, "/admin")
		return
	}
}
