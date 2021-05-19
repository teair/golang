package admin

import (
	"github.com/beego/beego/v2/server/web"
)

type LoginPrepare interface {
	LoginPrepare()
}

type baseController struct {
	web.Controller
}

func (b *baseController) Prepare() {
	// 后台验证登陆
	if b.GetSession("username") == nil || b.GetSession("password") == nil {
		// 没有登陆跳转至登录页
		b.Redirect("/login/index", 401)
		return
	}

	if app, ok := b.AppController.(LoginPrepare); ok {
		app.LoginPrepare()
	}
}
