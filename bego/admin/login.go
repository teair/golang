package admin

type LoginController struct {
	baseController
}

func (l *LoginController) LoginPrepare() {
	l.Data["login"] = "login"
}

func (l *LoginController) Get() {
	l.TplName = "admin/login/index.html"
}

func (l *LoginController) Post() {
	l.Ctx.WriteString("登陆操作")
}
