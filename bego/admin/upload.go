package admin

import (
	"bego/common"
	"bego/models/adminmodel"
	"bego/service/attachment"
	"fmt"
	"html/template"
	"strconv"
	"strings"
	"time"
)

type UploadController struct {
	baseController
}

type returnUpload struct {
	Code int
	Data interface{}
	Info string
}

type UploadInfo struct {
	Id                  int
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
			data.Id = v.Id
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
			tmp[v.LabelId] = strconv.Itoa(v.Id)
		}
		mp["curtype"] = tmp
	}
	if len(data.Curtheme) != 0 {
		tmp := make(map[string]string)
		for _, v := range data.Curtheme {
			tmp[v.LabelId] = strconv.Itoa(v.Id)
		}
		mp["curtheme"] = tmp
	}
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
	this.Data["isActive"] = mp
	this.Data["Data"] = data
	this.TplName = "admin/upload/index.html"
}

// UpFile 异步上传图片
func (this *UploadController) UpFile() {

	if this.Ctx.Request.Method == "POST" {

		// 返回数据
		var returnFileinfo interface{}

		// 接收参数
		gid := this.Ctx.Input.Query("gid")
		operate := this.Ctx.Input.Query("operate")

		// 判断当前游戏是否存在上传文件
		num, filelist := adminmodel.FindFile(gid)

		switch operate {
		case "gamelogo":

			// 删除旧文件
			attachment.DelOldFile(num, filelist, "gamelogo")

			// 上传文件信息
			var fileinfo adminmodel.FileInfo

			// 接收文件
			f, h, err := this.Ctx.Request.FormFile("gamelogo")
			if err != nil {
				return
			}
			defer f.Close()

			// get mime type
			mime := h.Header.Get("Content-Type")
			t := time.Now()

			// FileInfo 数据
			data := make(map[string]string)
			data["gid"] = gid
			data["type"] = "gamelogo"

			if returnFileinfo, err = attachment.SaveImage(&fileinfo, f, mime, h.Filename, t, data, h.Size); err != nil {
				this.Data["json"] = ReturnData{
					Code: 413,
					Data: nil,
					Info: fmt.Sprintf("%v", err),
				}
				this.ServeJSON()
			}

			// 保存文件到指定位置
			this.SaveToFile("gamelogo", strings.Trim(fileinfo.Path, "/"))
			break
		case "gamethumb":
			// 删除该游戏对应的图片信息
			attachment.DelOldFile(num, filelist, "gamethumb")

			// 上传文件信息
			var fileinfo adminmodel.FileInfo

			// 接收文件
			f, h, err := this.Ctx.Request.FormFile("gamethumb")
			if err != nil {
				return
			}
			defer f.Close()

			// get mime type
			mime := h.Header.Get("Content-Type")

			t := time.Now()

			// FileInfo 数据
			data := make(map[string]string)
			data["gid"] = gid
			data["type"] = "gamethumb"

			if returnFileinfo, err = attachment.SaveImage(&fileinfo, f, mime, h.Filename, t, data, h.Size); err != nil {
				this.Data["json"] = ReturnData{
					Code: 413,
					Data: nil,
					Info: fmt.Sprintf("%v", err),
				}
				this.ServeJSON()
			}
			// 保存文件到指定位置
			this.SaveToFile("gamethumb", strings.Trim(fileinfo.Path, "/"))
			break
		case "gamelunbo":
			// 删除该游戏对应的图片信息
			attachment.DelOldFile(num, filelist, "gamelunbo")

			// 上传文件信息
			var fileinfo adminmodel.FileInfo

			// 接收文件
			f, h, err := this.Ctx.Request.FormFile("gamelunbo")
			if err != nil {
				return
			}
			defer f.Close()

			// get mime type
			mime := h.Header.Get("Content-Type")

			t := time.Now()

			// FileInfo 数据
			data := make(map[string]string)
			data["gid"] = gid
			data["type"] = "gamelunbo"

			if returnFileinfo, err = attachment.SaveImage(&fileinfo, f, mime, h.Filename, t, data, h.Size); err != nil {
				this.Data["json"] = ReturnData{
					Code: 413,
					Data: nil,
					Info: fmt.Sprintf("%v", err),
				}
				this.ServeJSON()
			}
			// 保存文件到指定位置
			this.SaveToFile("gamelunbo", strings.Trim(fileinfo.Path, "/"))
			break
		case "gamehotrec":
			// 删除该游戏对应的图片信息
			attachment.DelOldFile(num, filelist, "gamehotrec")

			// 上传文件信息
			var fileinfo adminmodel.FileInfo

			// 接收文件
			f, h, err := this.Ctx.Request.FormFile("gamehotrec")
			if err != nil {
				return
			}
			defer f.Close()

			// get mime type
			mime := h.Header.Get("Content-Type")

			t := time.Now()

			// FileInfo 数据
			data := make(map[string]string)
			data["gid"] = gid
			data["type"] = "gamehotrec"

			if returnFileinfo, err = attachment.SaveImage(&fileinfo, f, mime, h.Filename, t, data, h.Size); err != nil {
				this.Data["json"] = ReturnData{
					Code: 413,
					Data: nil,
					Info: fmt.Sprintf("%v", err),
				}
				this.ServeJSON()
			}
			// 保存文件到指定位置
			this.SaveToFile("gamehotrec", strings.Trim(fileinfo.Path, "/"))
			break
		case "gamenormrec":
			// 删除该游戏对应的图片信息
			attachment.DelOldFile(num, filelist, "gamenormrec")

			// 上传文件信息
			var fileinfo adminmodel.FileInfo

			// 接收文件
			f, h, err := this.Ctx.Request.FormFile("gamenormrec")
			if err != nil {
				return
			}
			defer f.Close()

			// get mime type
			mime := h.Header.Get("Content-Type")

			t := time.Now()

			// FileInfo 数据
			data := make(map[string]string)
			data["gid"] = gid
			data["type"] = "gamenormrec"

			if returnFileinfo, err = attachment.SaveImage(&fileinfo, f, mime, h.Filename, t, data, h.Size); err != nil {
				this.Data["json"] = ReturnData{
					Code: 413,
					Data: nil,
					Info: fmt.Sprintf("%v", err),
				}
				this.ServeJSON()
			}
			// 保存文件到指定位置
			this.SaveToFile("gamenormrec", strings.Trim(fileinfo.Path, "/"))
			break
		case "gameofficalbg":
			// 删除该游戏对应的图片信息
			attachment.DelOldFile(num, filelist, "gameofficalbg")

			// 上传文件信息
			var fileinfo adminmodel.FileInfo

			// 接收文件
			f, h, err := this.Ctx.Request.FormFile("gameofficalbg")
			if err != nil {
				return
			}
			defer f.Close()

			// get mime type
			mime := h.Header.Get("Content-Type")

			t := time.Now()

			// FileInfo 数据
			data := make(map[string]string)
			data["gid"] = gid
			data["type"] = "gameofficalbg"

			if returnFileinfo, err = attachment.SaveImage(&fileinfo, f, mime, h.Filename, t, data, h.Size); err != nil {
				this.Data["json"] = ReturnData{
					Code: 413,
					Data: nil,
					Info: fmt.Sprintf("%v", err),
				}
				this.ServeJSON()
			}
			// 保存文件到指定位置
			this.SaveToFile("gameofficalbg", strings.Trim(fileinfo.Path, "/"))
			break
		case "gamecut1":
			// 删除该游戏对应的图片信息
			attachment.DelOldFile(num, filelist, "gamecut1")

			// 上传文件信息
			var fileinfo adminmodel.FileInfo

			// 接收文件
			f, h, err := this.Ctx.Request.FormFile("gamecut1")
			if err != nil {
				return
			}
			defer f.Close()

			// get mime type
			mime := h.Header.Get("Content-Type")

			t := time.Now()

			// FileInfo 数据
			data := make(map[string]string)
			data["gid"] = gid
			data["type"] = "gamecut1"

			if returnFileinfo, err = attachment.SaveImage(&fileinfo, f, mime, h.Filename, t, data, h.Size); err != nil {
				this.Data["json"] = ReturnData{
					Code: 413,
					Data: nil,
					Info: fmt.Sprintf("%v", err),
				}
				this.ServeJSON()
			}
			// 保存文件到指定位置
			this.SaveToFile("gamecut1", strings.Trim(fileinfo.Path, "/"))
			break
		case "gamecut2":
			// 删除该游戏对应的图片信息
			attachment.DelOldFile(num, filelist, "gamecut2")

			// 上传文件信息
			var fileinfo adminmodel.FileInfo

			// 接收文件
			f, h, err := this.Ctx.Request.FormFile("gamecut2")
			if err != nil {
				return
			}
			defer f.Close()

			// get mime type
			mime := h.Header.Get("Content-Type")

			t := time.Now()

			// FileInfo 数据
			data := make(map[string]string)
			data["gid"] = gid
			data["type"] = "gamecut2"

			if returnFileinfo, err = attachment.SaveImage(&fileinfo, f, mime, h.Filename, t, data, h.Size); err != nil {
				this.Data["json"] = ReturnData{
					Code: 413,
					Data: nil,
					Info: fmt.Sprintf("%v", err),
				}
				this.ServeJSON()
			}
			// 保存文件到指定位置
			this.SaveToFile("gamecut2", strings.Trim(fileinfo.Path, "/"))
			break
		case "gamecut3":
			// 删除该游戏对应的图片信息
			attachment.DelOldFile(num, filelist, "gamecut3")

			// 上传文件信息
			var fileinfo adminmodel.FileInfo

			// 接收文件
			f, h, err := this.Ctx.Request.FormFile("gamecut3")
			if err != nil {
				return
			}
			defer f.Close()

			// get mime type
			mime := h.Header.Get("Content-Type")

			t := time.Now()

			// FileInfo 数据
			data := make(map[string]string)
			data["gid"] = gid
			data["type"] = "gamecut3"

			if returnFileinfo, err = attachment.SaveImage(&fileinfo, f, mime, h.Filename, t, data, h.Size); err != nil {
				this.Data["json"] = ReturnData{
					Code: 413,
					Data: nil,
					Info: fmt.Sprintf("%v", err),
				}
				this.ServeJSON()
			}
			// 保存文件到指定位置
			this.SaveToFile("gamecut3", strings.Trim(fileinfo.Path, "/"))
			break
		case "gamecut4":
			// 删除该游戏对应的图片信息
			attachment.DelOldFile(num, filelist, "gamecut4")

			// 上传文件信息
			var fileinfo adminmodel.FileInfo

			// 接收文件
			f, h, err := this.Ctx.Request.FormFile("gamecut4")
			if err != nil {
				return
			}
			defer f.Close()

			// get mime type
			mime := h.Header.Get("Content-Type")

			t := time.Now()

			// FileInfo 数据
			data := make(map[string]string)
			data["gid"] = gid
			data["type"] = "gamecut4"

			if returnFileinfo, err = attachment.SaveImage(&fileinfo, f, mime, h.Filename, t, data, h.Size); err != nil {
				this.Data["json"] = ReturnData{
					Code: 413,
					Data: nil,
					Info: fmt.Sprintf("%v", err),
				}
				this.ServeJSON()
			}
			// 保存文件到指定位置
			this.SaveToFile("gamecut4", strings.Trim(fileinfo.Path, "/"))
			break
		case "gamecut5":
			// 删除该游戏对应的图片信息
			attachment.DelOldFile(num, filelist, "gamecut5")

			// 上传文件信息
			var fileinfo adminmodel.FileInfo

			// 接收文件
			f, h, err := this.Ctx.Request.FormFile("gamecut5")
			if err != nil {
				return
			}
			defer f.Close()

			// get mime type
			mime := h.Header.Get("Content-Type")

			t := time.Now()

			// FileInfo 数据
			data := make(map[string]string)
			data["gid"] = gid
			data["type"] = "gamecut5"

			if returnFileinfo, err = attachment.SaveImage(&fileinfo, f, mime, h.Filename, t, data, h.Size); err != nil {
				this.Data["json"] = ReturnData{
					Code: 413,
					Data: nil,
					Info: fmt.Sprintf("%v", err),
				}
				this.ServeJSON()
			}
			// 保存文件到指定位置
			this.SaveToFile("gamecut5", strings.Trim(fileinfo.Path, "/"))
			break
		default:
			this.Data["json"] = returnUpload{
				Code: 0,
				Data: nil,
				Info: "系统繁忙...",
			}
			this.ServeJSON()
			return
		}

		this.Data["json"] = returnUpload{
			Code: 1,
			Data: returnFileinfo,
			Info: "上传成功!",
		}
		this.ServeJSON()
	}
}

// UpSubmit 提交修改标语与游戏类型、游戏主题
func (this *UploadController) UpSubmit() {

	gid := this.GetString("gid")
	id, _ := this.GetInt("id")
	appType := this.GetStrings("appType")
	appTheme := this.GetStrings("appTheme")
	publicity := this.GetString("publicity")

	// 修改游戏标语
	gameinfo := adminmodel.GameInfo{Id: id, Publicity: publicity}
	(*adminmodel.GameInfo).Update(&gameinfo, "Publicity")

	// 删除当前游戏类型与主题
	curtype := adminmodel.CurLabel(gid, common.GAME_TYPE)
	curtheme := adminmodel.CurLabel(gid, common.GAME_THEME)
	for _, v := range curtype {
		gamelabel := adminmodel.GameLabel{Id: v.Id}
		(*adminmodel.GameLabel).Delete(&gamelabel)
	}
	for _, v := range curtheme {
		gamelabel := adminmodel.GameLabel{Id: v.Id}
		(*adminmodel.GameLabel).Delete(&gamelabel)
	}

	for _, v := range appType {
		labelSlice := strings.Split(v, ",")
		labelId, _ := strconv.Atoi(labelSlice[0])

		// 添加当前游戏类型
		labelinfo := adminmodel.GameLabel{
			Gid:        gid,
			Type:       common.GAME_TYPE,
			LabelId:    strconv.Itoa(labelId),
			LabelName:  labelSlice[1],
			CreateTime: time.Time{},
		}
		(*adminmodel.GameLabel).Insert(&labelinfo)
	}

	for _, v := range appTheme {
		labelSlice := strings.Split(v, ",")
		labelId, _ := strconv.Atoi(labelSlice[0])

		// 添加当前游戏主题
		labelinfo := adminmodel.GameLabel{
			Gid:        gid,
			Type:       common.GAME_THEME,
			LabelId:    strconv.Itoa(labelId),
			LabelName:  labelSlice[1],
			CreateTime: time.Time{},
		}
		(*adminmodel.GameLabel).Insert(&labelinfo)
	}

	this.Ctx.WriteString("submit!")
}
