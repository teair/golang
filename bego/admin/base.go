package admin

import (
	"github.com/beego/beego/v2/core/utils"
	"github.com/beego/beego/v2/server/web"
)

// 基类控制器
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

// 公共分页操作
func (this *baseController) SetPaginator(per int, nums int64) *utils.Paginator {
	p := utils.NewPaginator(this.Ctx.Request, per, nums)
	this.Data["paginator"] = p
	return p
}
