package admin

import (
	"bego/models/adminmodel"
	"strings"
	"time"
)

type GameController struct {
	baseController
}

func (this *baseController) GameAdd() {
	data := this.Ctx.Input.Query("msg")
	exp := strings.Split(data, "^")
	if data == "" || len(exp) != 8 {
		this.Data["json"] = ReturnData{
			Code: 402,
			Data: nil,
			Info: "参数异常...",
		}
		this.ServeJSON()
	}

	gid, app_name, app_rename, game_size, publicity, fuli, source, game_introduce := exp[0], exp[1], exp[2], exp[3], exp[4], exp[5], exp[6], exp[7]

	gamedata := adminmodel.GameInfo{
		Gid:           gid,
		AppName:       app_name,
		AppRename:     app_rename,
		GameSize:      game_size,
		Source:        source,
		Publicity:     publicity,
		Fuli:          fuli,
		GameIntroduce: game_introduce,
		CreateTime:    time.Now(),
	}

	// 入库

	id := adminmodel.GameAdd(gamedata)

	if id == 0 {
		this.Data["json"] = ReturnData{
			Code: 500,
			Data: nil,
			Info: "创建失败...",
		}
		this.ServeJSON()
	}
	this.Redirect("/admin/index", 302)
}
