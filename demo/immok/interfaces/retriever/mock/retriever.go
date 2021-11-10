package mock

import "fmt"

type Retriever struct {
	Contents string
}

// String 常用系统接口
func (r *Retriever) String() string {
	return fmt.Sprintf("Retriever: {Content=%s}",r.Contents)
}

func (r *Retriever) Post(url string, form map[string]string) string {
	r.Contents = form["contents"]
	return "ok"
}

func (r *Retriever) Get(url string) string {
	return r.Contents
}
