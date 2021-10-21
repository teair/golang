package infra

import (
	"io/ioutil"
	"net/http"
)

type Retriever struct{}

// Get 通过接口实现下载文件
func (Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
