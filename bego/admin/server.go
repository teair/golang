package admin

import (
	"bego/models"
	"fmt"
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

	p, per := this.Paginator()

	serverData.ServerList, total = models.Select(p, per)

	// 前台分页
	this.Data["paginator"] = this.SetPaginator(per, total)

	this.Data["Index"] = serverData

	this.TplName = "admin/server/serverlist.html"
}

func (this *ServerController) EditServer() {

	// 模板数据
	serverData := ServerData{Active: 1, Open: 0}
	this.Data["Index"] = serverData

	if this.Ctx.Request.Method == "POST" {

	} else {
		id, _ := this.GetInt("id")
		gid := this.GetString("gid")
		var m = models.GameServer{Id: id}
		(*models.GameServer).Find(&m)

		this.Data["ServerData"] = m

		fmt.Println("the gameserver is:", m)
		fmt.Println("the id is:", id)
		fmt.Println("the gid is:", gid)

		this.TplName = "admin/server/editserver.html"
	}

}
