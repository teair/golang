package must

import (
	"fmt"
	"text/template"
)

func MustMain() {
	tok := template.New("first")
	template.Must(tok.Parse(" some static text /* and a comment */"))
	fmt.Println("The first one parsed ok. ")

	template.Must(template.New("second").Parse("some static text {{ .Name}}"))
	fmt.Println("The second one parsed ok. ")

	fmt.Println("The next one ought to fail.")
	tErr := template.New("check parse error with must. ")
	template.Must(tErr.Parse(" some static text {{ .Name}}"))
	fmt.Println("finished!")
}
