package admin

import (
	"bego/common"
	"bego/models/adminmodel"
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"strings"
)

type UploadController struct {
	baseController
}

type UploadInfo struct {
	AppName, AppRename  string
	Source              string
	GameType, GameTicai []adminmodel.TagInfo
	Curtype, Curtheme   []adminmodel.GameLabel
	Slogan, Gid         string
	FileInfo            []adminmodel.FileInfo
}

func (this *UploadController) UpIndex() {
	// 公共控制
	var index IndexData
	index.Active = 0
	this.Data["Index"] = index

	// 上传页信息
	var data UploadInfo

	gid := this.Ctx.Input.Param(":gid")
	if gid != "" {
		var slogan string
		var gameinfo []adminmodel.GameInfo
		// 查看该游戏来源
		gameinfo, _ = adminmodel.FindGame("Gid", gid, 0, 0)
		for _, v := range gameinfo {
			data.Source = v.Source
			data.AppName = v.AppName
			data.AppRename = v.AppRename
			slogan = v.Publicity
			data.Slogan = slogan
		}
		// 存储标语
		this.Ctx.SetCookie("slogan", slogan, 0)
		// 获取所有游戏题材和游戏类型标签
		data.GameType = adminmodel.GetGameTag(87)
		data.GameTicai = adminmodel.GetGameTag(88)

		// 查看当前游戏是否存在标签
		data.Curtype = adminmodel.CurLabel(gid, common.GAME_TYPE)
		data.Curtheme = adminmodel.CurLabel(gid, common.GAME_THEME)

		data.Gid = gid

		// 获取该游戏的图片信息
		_, data.FileInfo = adminmodel.FindFile(gid)
	}

	// 游戏题材与类型多选是否选中
	mp := make(map[string]map[string]string)
	if len(data.Curtype) != 0 {
		tmp := make(map[string]string)
		for _, v := range data.Curtype {
			tmp[v.LabelId] = v.LabelId
		}
		mp["curtype"] = tmp
	}
	if len(data.Curtheme) != 0 {
		tmp := make(map[string]string)
		for _, v := range data.Curtheme {
			tmp[v.LabelId] = v.LabelId
		}
		mp["curtheme"] = tmp
	}
	this.Data["isactive"] = mp
	this.Data["Data"] = data
	this.TplName = "admin/upload/index.html"
}

func (this *UploadController) UpFile() {
	if this.Ctx.Request.Method == "POST" {
		// 接收参数
		gid := this.Ctx.Input.Query("gid")
		operate := this.Ctx.Input.Query("operate")

		// 判断当前游戏是否存在上传文件
		num, filelist := adminmodel.FindFile(gid)

		switch operate {
		case "gamelogo":
			if num != 0 {
				for _, v := range filelist {
					adminmodel.DelFile(v.Id)
				}
			}

			// 上传文件信息
			var fileinfo adminmodel.FileInfo

			// 接收文件
			f, h, _ := this.GetFile("gamelogo")
			fmt.Println("the filename is:", h.Filename)
			fileinfo.Filename = h.Filename
			arr := strings.Split(h.Filename, ":")
			if len(arr) > 1 {
				index := len(arr) - 1
				fileinfo.Filename = arr[index]
			}

			// 关闭上传的文件,不然会出现临时文件不能删除的情况
			f.Close()

			// 保存文件到指定位置
			upath, _ := web.AppConfig.String("upload::upDir")
			fmt.Println("上传文件的目录为:", upath)
			break
		}
		fmt.Println("the gid is:", gid)
		fmt.Println("the operate is:", operate)
		this.Ctx.WriteString("?????")
	}
}
