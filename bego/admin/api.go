package admin

import (
	"bego/models/admin"
	"github.com/beego/beego/v2/server/web"
)

type ApiController struct {
	web.Controller
}

func (this *ApiController) TestApi() {
	lists := admin.GameInfoSelect()
	this.Data["json"] = lists
	this.ServeJSON()
}
