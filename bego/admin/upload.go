package admin

import (
	"bego/common"
	"bego/models"
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
	GameType, GameTicai []models.TagInfo
	Curtype, Curtheme   []models.GameLabel
	Slogan, Gid         string
	FileInfo            []models.FileInfo
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
		var gameinfo []models.GameInfo
		// 查看该游戏来源
		gameinfo, _ = models.FindGame("Gid", gid, 0, 0)
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
		data.GameType = models.GetGameTag(87)
		data.GameTicai = models.GetGameTag(88)

		// 查看当前游戏是否存在标签
		data.Curtype = models.CurLabel(gid, common.GAME_TYPE)
		data.Curtheme = models.CurLabel(gid, common.GAME_THEME)

		data.Gid = gid

		// 获取该游戏的图片信息
		_, data.FileInfo = models.FindFile(gid)
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
		num, filelist := models.FindFile(gid)

		switch operate {
		case common.GAME_LOGO:

			// 删除旧文件
			attachment.DelOldFile(num, filelist, common.GAME_LOGO)

			// 上传文件信息
			var fileinfo models.FileInfo

			// 接收文件
			f, h, err := this.Ctx.Request.FormFile(common.GAME_LOGO)
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
			data["type"] = common.GAME_LOGO

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
		case common.GAME_THUMB:
			// 删除该游戏对应的图片信息
			attachment.DelOldFile(num, filelist, common.GAME_THUMB)

			// 上传文件信息
			var fileinfo models.FileInfo

			// 接收文件
			f, h, err := this.Ctx.Request.FormFile(common.GAME_THUMB)
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
			data["type"] = common.GAME_THUMB

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
		case common.GMAE_LUNBO:
			// 删除该游戏对应的图片信息
			attachment.DelOldFile(num, filelist, common.GMAE_LUNBO)

			// 上传文件信息
			var fileinfo models.FileInfo

			// 接收文件
			f, h, err := this.Ctx.Request.FormFile(common.GMAE_LUNBO)
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
			data["type"] = common.GMAE_LUNBO

			if returnFileinfo, err = attachment.SaveImage(&fileinfo, f, mime, h.Filename, t, data, h.Size); err != nil {
				this.Data["json"] = ReturnData{
					Code: 413,
					Data: nil,
					Info: fmt.Sprintf("%v", err),
				}
				this.ServeJSON()
			}
			// 保存文件到指定位置
			this.SaveToFile(common.GMAE_LUNBO, strings.Trim(fileinfo.Path, "/"))
			break
		case common.GAME_HOTREC:
			// 删除该游戏对应的图片信息
			attachment.DelOldFile(num, filelist, common.GAME_HOTREC)

			// 上传文件信息
			var fileinfo models.FileInfo

			// 接收文件
			f, h, err := this.Ctx.Request.FormFile(common.GAME_HOTREC)
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
			data["type"] = common.GAME_HOTREC

			if returnFileinfo, err = attachment.SaveImage(&fileinfo, f, mime, h.Filename, t, data, h.Size); err != nil {
				this.Data["json"] = ReturnData{
					Code: 413,
					Data: nil,
					Info: fmt.Sprintf("%v", err),
				}
				this.ServeJSON()
			}
			// 保存文件到指定位置
			this.SaveToFile(common.GAME_HOTREC, strings.Trim(fileinfo.Path, "/"))
			break
		case common.GAME_NORMREC:
			// 删除该游戏对应的图片信息
			attachment.DelOldFile(num, filelist, common.GAME_NORMREC)

			// 上传文件信息
			var fileinfo models.FileInfo

			// 接收文件
			f, h, err := this.Ctx.Request.FormFile(common.GAME_NORMREC)
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
			data["type"] = common.GAME_NORMREC

			if returnFileinfo, err = attachment.SaveImage(&fileinfo, f, mime, h.Filename, t, data, h.Size); err != nil {
				this.Data["json"] = ReturnData{
					Code: 413,
					Data: nil,
					Info: fmt.Sprintf("%v", err),
				}
				this.ServeJSON()
			}
			// 保存文件到指定位置
			this.SaveToFile(common.GAME_NORMREC, strings.Trim(fileinfo.Path, "/"))
			break
		case common.GAME_OFFICALRBG:
			// 删除该游戏对应的图片信息
			attachment.DelOldFile(num, filelist, common.GAME_OFFICALRBG)

			// 上传文件信息
			var fileinfo models.FileInfo

			// 接收文件
			f, h, err := this.Ctx.Request.FormFile(common.GAME_OFFICALRBG)
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
			data["type"] = common.GAME_OFFICALRBG

			if returnFileinfo, err = attachment.SaveImage(&fileinfo, f, mime, h.Filename, t, data, h.Size); err != nil {
				this.Data["json"] = ReturnData{
					Code: 413,
					Data: nil,
					Info: fmt.Sprintf("%v", err),
				}
				this.ServeJSON()
			}
			// 保存文件到指定位置
			this.SaveToFile(common.GAME_OFFICALRBG, strings.Trim(fileinfo.Path, "/"))
			break
		case common.GAME_CUT1:
			// 删除该游戏对应的图片信息
			attachment.DelOldFile(num, filelist, common.GAME_CUT1)

			// 上传文件信息
			var fileinfo models.FileInfo

			// 接收文件
			f, h, err := this.Ctx.Request.FormFile(common.GAME_CUT1)
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
			data["type"] = common.GAME_CUT1

			if returnFileinfo, err = attachment.SaveImage(&fileinfo, f, mime, h.Filename, t, data, h.Size); err != nil {
				this.Data["json"] = ReturnData{
					Code: 413,
					Data: nil,
					Info: fmt.Sprintf("%v", err),
				}
				this.ServeJSON()
			}
			// 保存文件到指定位置
			this.SaveToFile(common.GAME_CUT1, strings.Trim(fileinfo.Path, "/"))
			break
		case common.GAME_CUT2:
			// 删除该游戏对应的图片信息
			attachment.DelOldFile(num, filelist, common.GAME_CUT2)

			// 上传文件信息
			var fileinfo models.FileInfo

			// 接收文件
			f, h, err := this.Ctx.Request.FormFile(common.GAME_CUT2)
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
			this.SaveToFile(common.GAME_CUT2, strings.Trim(fileinfo.Path, "/"))
			break
		case common.GAME_CUT3:
			// 删除该游戏对应的图片信息
			attachment.DelOldFile(num, filelist, common.GAME_CUT3)

			// 上传文件信息
			var fileinfo models.FileInfo

			// 接收文件
			f, h, err := this.Ctx.Request.FormFile(common.GAME_CUT3)
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
			data["type"] = common.GAME_CUT3

			if returnFileinfo, err = attachment.SaveImage(&fileinfo, f, mime, h.Filename, t, data, h.Size); err != nil {
				this.Data["json"] = ReturnData{
					Code: 413,
					Data: nil,
					Info: fmt.Sprintf("%v", err),
				}
				this.ServeJSON()
			}
			// 保存文件到指定位置
			this.SaveToFile(common.GAME_CUT3, strings.Trim(fileinfo.Path, "/"))
			break
		case common.GAME_CUT4:
			// 删除该游戏对应的图片信息
			attachment.DelOldFile(num, filelist, common.GAME_CUT4)

			// 上传文件信息
			var fileinfo models.FileInfo

			// 接收文件
			f, h, err := this.Ctx.Request.FormFile(common.GAME_CUT4)
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
			data["type"] = common.GAME_CUT4

			if returnFileinfo, err = attachment.SaveImage(&fileinfo, f, mime, h.Filename, t, data, h.Size); err != nil {
				this.Data["json"] = ReturnData{
					Code: 413,
					Data: nil,
					Info: fmt.Sprintf("%v", err),
				}
				this.ServeJSON()
			}
			// 保存文件到指定位置
			this.SaveToFile(common.GAME_CUT4, strings.Trim(fileinfo.Path, "/"))
			break
		case common.GAME_CUT5:
			// 删除该游戏对应的图片信息
			attachment.DelOldFile(num, filelist, common.GAME_CUT5)

			// 上传文件信息
			var fileinfo models.FileInfo

			// 接收文件
			f, h, err := this.Ctx.Request.FormFile(common.GAME_CUT5)
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
			data["type"] = common.GAME_CUT5

			if returnFileinfo, err = attachment.SaveImage(&fileinfo, f, mime, h.Filename, t, data, h.Size); err != nil {
				this.Data["json"] = ReturnData{
					Code: 413,
					Data: nil,
					Info: fmt.Sprintf("%v", err),
				}
				this.ServeJSON()
			}
			// 保存文件到指定位置
			this.SaveToFile(common.GAME_CUT5, strings.Trim(fileinfo.Path, "/"))
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
	gameinfo := models.GameInfo{Id: id, Publicity: publicity}
	(*models.GameInfo).Update(&gameinfo, "Publicity")

	// 删除当前游戏类型与主题
	curtype := models.CurLabel(gid, common.GAME_TYPE)
	curtheme := models.CurLabel(gid, common.GAME_THEME)
	for _, v := range curtype {
		gamelabel := models.GameLabel{Id: v.Id}
		(*models.GameLabel).Delete(&gamelabel)
	}
	for _, v := range curtheme {
		gamelabel := models.GameLabel{Id: v.Id}
		(*models.GameLabel).Delete(&gamelabel)
	}

	for _, v := range appType {
		labelSlice := strings.Split(v, ",")
		labelId, _ := strconv.Atoi(labelSlice[0])

		// 添加当前游戏类型
		labelinfo := models.GameLabel{
			Gid:        gid,
			Type:       common.GAME_TYPE,
			LabelId:    strconv.Itoa(labelId),
			LabelName:  labelSlice[1],
			CreateTime: time.Time{},
		}
		(*models.GameLabel).Insert(&labelinfo)
	}

	for _, v := range appTheme {
		labelSlice := strings.Split(v, ",")
		labelId, _ := strconv.Atoi(labelSlice[0])

		// 添加当前游戏主题
		labelinfo := models.GameLabel{
			Gid:        gid,
			Type:       common.GAME_THEME,
			LabelId:    strconv.Itoa(labelId),
			LabelName:  labelSlice[1],
			CreateTime: time.Time{},
		}
		(*models.GameLabel).Insert(&labelinfo)
	}

	this.Redirect("/admin/index", 302)
}
