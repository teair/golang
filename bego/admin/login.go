package admin

import (
	"bego/models/admin"
	"encoding/base64"
	"fmt"
	"github.com/beego/beego/v2/core/validation"
	"github.com/beego/beego/v2/server/web"
	"html/template"
)

type LoginController struct {
	web.Controller
}

// 定义返回数据格式
type ReturnData struct {
	Code int
	Data []interface{}
	Info string
}

func (this *LoginController) LoginPrepare() {
	this.Data["login"] = "login"
}

func (this *LoginController) Index() {
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
	this.TplName = "admin/login/index.html"
}

func (this *LoginController) AdminLogin() {
	// 接收参数
	username := this.GetString("username")
	password := this.GetString("password")

	// 参数验证
	valid := validation.Validation{}

	// 验证用户名
	valid.Required(username, "username").Message("请输入用户名!")
	valid.Alpha(username, "username").Message("用户名含有非法字符!")
	valid.MinSize(username, 3, "username").Message("用户名在3-6字符之间!")
	valid.MaxSize(username, 6, "username").Message("用户名在3-6字符之间!")

	// 验证密码
	valid.Required(password, "password").Message("请输入密码!")
	valid.AlphaNumeric(password, "password").Message("密码中含有非法字符!")
	valid.MinSize(password, 6, "password").Message("密码在6-12字符数字之间!")
	valid.MaxSize(password, 12, "password").Message("密码在6-12字符数字之间!")

	// 判断验证结果
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

	flg, err := admin.CheckAdmin(username, password)

	if err != nil && flg {
		this.Data["json"] = ReturnData{Code: 400, Data: nil, Info: fmt.Sprintf("%x", err)}
		this.ServeJSON()
	}

	// 储存session
	this.SetSession("username", username)
	b64 := base64.StdEncoding.EncodeToString([]byte(password))
	this.SetSession("password", b64)

	this.Data["json"] = ReturnData{Code: 200, Data: nil, Info: fmt.Sprintf("%x", "登陆成功!")}
	this.ServeJSON()
}
