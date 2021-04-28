package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {

	// 获取应用全局配置文件
	//appname,_ := beego.AppConfig.String("runmode")
	//c.Ctx.WriteString(appname)

	// 动态设置运行模式
	/*err := beego.AppConfig.Set("runmode","prod")
	if err == nil {
		runmod, _ := beego.AppConfig.String("runmode")
		c.Ctx.WriteString(runmod)
	}*/

	c.Data["Pinduoduo"] = "拼多多真坑!"
	c.TplName = "home/user.html"
}
