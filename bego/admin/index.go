package admin

import (
	"bego/models/admin"
	"github.com/beego/beego/v2/core/validation"
	"strings"
)

type IndexController struct {
	baseController
}

// 后台首页
func (this *IndexController) Index() {
	// 接收参数
	data := this.Ctx.Input.Params()

	// 获取全部分发游戏
	gameList := admin.GameInfoSelect()
	alreadyGid := []string{}
	for i := range gameList {
		alreadyGid = append(alreadyGid, gameList[i].Gid)
	}

	// 可创建分发游戏列表
	var listenFirst bool
	if this.GetSession("listenfirst") == nil {
		listenFirst = false
	} else {
		listenFirst = true
	}

	gamemsg := admin.GetGid(listenFirst, strings.Join(alreadyGid, ","))

	this.Data["gamelist"] = gamemsg

	// 获取全部游戏数据
	if data["mark_online"] == "" || data["app_name"] == "" {
		// 获取全部创建分发的游戏与分页
		fenfa := admin.FindAll()
		this.Data["data"] = fenfa
	} else {
		if data["mark_online"] != "" {
			// 验证 mark_online 参数是否合法
			valid := validation.Validation{}
			valid.Required(data["mark_online"], "mark_online").Message("非法操作!")
			valid.Range(data["mark_online"], 0, 2, "mark_online").Message("参数异常!")

			if valid.HasErrors() {
				for _, err := range valid.Errors {
					this.Data["json"] = ReturnData{
						Code: 400,
						Data: nil,
						Info: err.Message,
					}
					this.ServeJSON()
				}
			}

			if data["mark_online"] == "2" {
				// 获取全部游戏
				datainfo := admin.FindAll()
				this.Data["data"] = datainfo
			} else {
				// 获取上线/下线游戏
				datainfo := admin.FindGame("mark_online", data["mark_online"])
				this.Data["data"] = datainfo
			}
			this.Data["mark_onlilne"] = data["mark_online"]
		}

		if data["app_name"] != "" {
			// 验证游戏名称
			valid := validation.Validation{}
			valid.Required(data["app_name"], "app_name").Message("请求参数异常!")
			valid.Alpha(data["app_name"], "app_name").Message("请游戏拼音!")
			if valid.HasErrors() {
				for _, err := range valid.Errors {
					this.Data["json"] = ReturnData{
						Code: 400,
						Data: nil,
						Info: err.Message,
					}
					this.ServeJSON()
				}
			}
			gamedata := admin.FindGame("app_name", data["app_name"])
			this.Data["data"] = gamedata
			this.Data["selname"] = data["app_name"]
		}
	}

	this.TplName = "admin/index/index.html"
}
