package interfaces

import (
	"demo/immok/interfaces/testing"
	"fmt"
)

// retriever->Something that can get
type retriever interface {
	Get(string) string
}

// getRetriever 接口实现
func getRetriever() retriever {
	return testing.Retriever{}
	//return infra.Retriever{}
}

// InterfaceMain 接口入口
func InterfaceMain() {
	var r retriever = getRetriever()
	fmt.Println(r.Get("https://www.imooc.com"))
}
