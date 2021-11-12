package filelist

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/myfile/"

type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

func MyFileHandle(writer http.ResponseWriter, request *http.Request) error {

	url := request.URL.Path

	// 如果不存在 prefix 前缀
	if strings.Index(url, prefix) != 0 {
		return userError("path must start with " + prefix)
	}

	path := request.URL.Path[len(prefix):] // 获取 /myfile/ 后的路径也就是指定的文件名

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file) // 读取文件内容
	if err != nil {
		return err
	}

	writer.Write(all) // 输出文件内容
	return nil
}
