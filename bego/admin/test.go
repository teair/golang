package admin

type TestController struct {
	baseController
}

func (t *TestController) Test() {
	t.Ctx.WriteString("admin test!")
}
