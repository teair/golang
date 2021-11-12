package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {

	url := request.URL.Path
	// 非前缀请求
	if strings.Index(url, prefix) != 0 {
		return userError("request must be prefix:" + prefix)
	}

	// 获取文件路径
	path := request.URL.Path[len(prefix):] // 匹配/list/后边的参数,也就是指定文件的路径

	// 根据文件路径打开文件
	file, err := os.Open(path) // 打开存在的文件
	if err != nil {
		return err
	}
	defer file.Close()

	// 读取文件列表
	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	writer.Write(all) // 浏览器打印文件列表
	return nil
}
