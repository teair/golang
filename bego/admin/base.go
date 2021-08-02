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
	nc := []string{"IndexController", "GameController", "UploadController"}
	na := []string{"Gamesel", "Listen", "Gamedel", "GameAdd", "UpFile", "UpSubmit", "MarkOnline"}
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

// SetPaginator 公共分页操作
func (this *baseController) SetPaginator(per int, nums int64) *pagination.Paginator {
	p := pagination.NewPaginator(this.Ctx.Request, per, nums)
	this.Data["paginator"] = p
	return p
}

// Paginator 分页
func (this *baseController) Paginator() (p, per int) {
	per, _ = web.AppConfig.Int("page::perPage")
	p, _ = utils.StrTo(this.Ctx.Input.Query("p")).Int()

	switch {
	case per > 100:
		per = 100
	case per <= 0:
		per = 10
	}

	if p == 0 || p < 0 {
		p = 1
	}

	if p > 0 {
		p = (p - 1) * per
	}
	return p, per
}
