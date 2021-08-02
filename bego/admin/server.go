package admin

import (
	"bego/models"
	"fmt"
	"github.com/beego/beego/v2/core/utils"
	"github.com/beego/beego/v2/server/web"
)

type ServerController struct {
	baseController
}

type ServerData struct {
	Active, Open int
	ServerList   []models.GameServer
}

func (this *ServerController) ServerIndex() {

	// 区服总数
	var total int64

	// 模板数据
	serverData := ServerData{Active: 1, Open: 0}

	per, _ := web.AppConfig.Int("page::perPage")
	p, _ := utils.StrTo(this.Ctx.Input.Query("p")).Int()
	if p > 0 {
		p = (p - 1) * per
	}

	serverData.ServerList, total = models.Select(p, per)

	// 前台分页
	this.Data["paginator"] = this.SetPaginator(per, total)
	this.Data["Index"] = serverData

	this.TplName = "admin/server/serverlist.html"
}

func (this *ServerController) EditServer() {
	id, _ := this.GetInt("id")
	gid := this.GetString("gid")
	fmt.Println("the id is:", id)
	fmt.Println("the gid is:", gid)
	this.Ctx.WriteString("??????????")
	//this.TplName = "admin/server/editserver.html"
}
