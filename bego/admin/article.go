package admin

import beego "github.com/beego/beego/v2/server/web"

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) Get() {
	c.TplName = "admin/article.html"
}
