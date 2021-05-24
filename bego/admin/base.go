package admin

import (
	"github.com/beego/beego/v2/server/web"
)

// 基类控制器
type baseController struct {
	web.Controller
}

// 定义返回数据格式
type ReturnData struct {
	Code int
	Data []interface{}
	Info string
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
