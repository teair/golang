package admin

import (
	"bego/models/admin"
	"github.com/beego/beego/v2/server/web"
)

type ApiController struct {
	web.Controller
}

func (this *ApiController) TestApi() {
	var num int64
	num, lists := admin.GameInfoSelect()
	this.Data["num"] = num
	this.Data["json"] = lists
	this.ServeJSON()
}
