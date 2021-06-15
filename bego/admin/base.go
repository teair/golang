package admin

import (
	"github.com/beego/beego/v2/core/utils"
	"github.com/beego/beego/v2/core/utils/pagination"
	"github.com/beego/beego/v2/server/web"
)

// 基类控制器
type baseController struct {
	web.Controller
}

func (this *baseController) Prepare() {
	// 取消xsrf验证的控制器和方法
	nc := []string{"IndexController", "GameController"}
	na := []string{"Gamesel", "Listen", "Gamedel", "GameAdd"}
	c, a := this.GetControllerAndAction()
	if utils.InSlice(c, nc) && utils.InSlice(a, na) {
		this.EnableXSRF = false
	}

	// 后台验证登陆
	if this.GetSession("username") == nil || this.GetSession("password") == nil {
		// 没有登陆跳转至登录页
		//b.Ctx.WriteString("please login!")
		this.Ctx.Redirect(302, "/admin")
		return
	}
}

// 公共分页操作
func (this *baseController) SetPaginator(per int, nums int64) *pagination.Paginator {
	p := pagination.NewPaginator(this.Ctx.Request, per, nums)
	this.Data["paginator"] = p
	return p
}
