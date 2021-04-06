package template

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Friend struct {
	Fname string
}

type Person struct {
	UserName string
	Emails   []string // 未导出的字段首字母是小写的
	Friends  []*Friend
}

type Cemat struct {
	Id                           int
	Uname, Company, Duty, Mobile string
}

type Data struct {
	Cemats []*Cemat
	Count  int
}

func init() {

	// 注册mysql驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// 注册默认数据库
	orm.RegisterDataBase("default", "mysql", "连接参数", 30)
	// 注册表对应的模型
	orm.RegisterModel(new(Cemat))
}

// 取出数据库数据并渲染至页面
func DbTemplate(w http.ResponseWriter, r *http.Request) {

	// 开启sql调试模式
	orm.Debug = true

	// 获取orm操作对象
	o := orm.NewOrm()

	// 储存cemat表数据
	var c []*Cemat
	num, err := o.QueryTable("Cemat").OrderBy("Id").All(&c)

	if err != nil {
		fmt.Println("Query cemat failed! ")
	}

	// 模板数据
	var d Data

	d = Data{
		Cemats: c,
		Count:  int(num),
	}

	t, err := template.ParseFiles("template/html/cemat.html")

	if err != nil {
		fmt.Println("ParseFiles failed!")
	}

	t.Execute(w, &d)

}

func TemplateTest(w http.ResponseWriter, r *http.Request) {

	// t := template.New("some template")                // 创建一个模板
	// t, _ = t.ParseFiles("template/html/welcome.html") // 解析模板文件
	// user := "11"                                      // 获取当前用户信息
	// t.Execute(w, user)

	// t := template.New("field name example! ")
	// t.Parse("hello, this is {{.UserName}}, the password is: {{.Password}}!")
	// p := Person{UserName: "ShenTongtong", Password: "12345"}
	// t.Execute(os.Stdout, p)

	// 返回一个模板对象解析后在控制台输出
	// t := template.New("fieldname example!")

	// t, _ = t.Parse(`hello, {{.UserName}}!
	// 	{{range .Emails}}
	// 		an email {{.}}
	// 	{{end}}
	// 	{{with .Friends}}
	// 	{{range .}}
	// 		my friend name is {{.Fname}}
	// 	{{end}}
	// 	{{end}}
	// `)

	f1 := Friend{Fname: "minux.ma"}
	f2 := Friend{Fname: "xushiyou"}

	t, _ := template.ParseFiles("template/html/welcome.html")

	p := Person{
		UserName: "<script>alert('1111');</script>",
		Emails:   []string{"111@qq.com", "222@qq.com"},
		Friends:  []*Friend{&f1, &f2},
	}

	t.Execute(w, p)
}

func TemplateMain() {

	http.HandleFunc("/template", TemplateTest)

	http.HandleFunc("/dbtemplate", DbTemplate)

	err := http.ListenAndServe("127.0.0.1:9091", nil)

	if err != nil {
		fmt.Println("failed to listen! ")
	}
}
