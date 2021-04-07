package nested

import (
	"fmt"
	"net/http"
	"text/template"
)

func NestedMain(w http.ResponseWriter, r *http.Request) {
	s1, _ := template.ParseFiles("template/html/header.html", "template/html/content.html", "template/html/footer.html")
	s1.ExecuteTemplate(w, "header", nil)
	fmt.Fprintln(w, nil)
	s1.ExecuteTemplate(w, "content", nil)
	fmt.Fprintln(w, nil)
	s1.ExecuteTemplate(w, "footer", nil)
	fmt.Fprintln(w, nil)
	s1.Execute(w, nil)
}
