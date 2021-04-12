package replace

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"
)

type Friend struct {
	Fname string
}

type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

func EmailDealWith(args ...interface{}) string {

	ok := false

	var s string
	if len(args) == 1 {
		s, ok = args[0].(string)
	}

	if !ok {
		s = fmt.Sprint(args...)
	}

	// find the @ symbol
	substrs := strings.Split(s, "@")
	if len(substrs) != 2 {
		return s
	}

	// replace the @ by " at "
	return substrs[0] + "at" + substrs[1]
}

func ReplacesMain(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		f1 := Friend{Fname: "刘备"}

		f2 := Friend{Fname: "关羽"}

		f3 := Friend{Fname: "张飞"}

		// 控制台输出模板内容
		// t := template.New("fieldname example")
		// t = t.Funcs(template.FuncMap{"emailDeal": EmailDealWith})
		// t.Parse(`
		// 	hello {{.UserName}}!
		// 	{{range .Emails}}
		// 		an emails {{.|emailDeal}}
		// 	{{end}}
		// 	{{with .Friends}}
		// 		{{range .}}my friend name is {{.Fname}}
		// 		{{end}}
		// 	{{end}}
		// `)

		// p := Person{
		// 	UserName: "ShenTongtong",
		// 	Emails:   []string{"1111@qq.com", "2222@qq.com", "3333@qq.com"},
		// 	Friends:  []*Friend{&f1, &f2, &f3},
		// }

		// t.Execute(os.Stdout, &p)

		// 浏览器访问模板页
		t := template.New(`template/html/replace.html`)
		t = t.Funcs(template.FuncMap{"emailDeal": EmailDealWith})
		t, _ = t.ParseFiles(`template/html/replace.html`)
		p := Person{
			UserName: "ShenTongtong",
			Emails:   []string{"1111@qq.com", "2222@qq.com", "3333@qq.com"},
			Friends:  []*Friend{&f1, &f2, &f3},
		}
		if err := t.ExecuteTemplate(w, "replace.html", p); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	} else {
		fmt.Println("Sorry....")
	}
}
