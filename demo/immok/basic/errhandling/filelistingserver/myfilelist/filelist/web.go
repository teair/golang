package filelist

import (
	"io/ioutil"
	"net/http"
	"os"
)

func MyFileHandle(writer http.ResponseWriter, request *http.Request) error {

	path := request.URL.Path[len("/myfile/"):] // 获取 /myfile/ 后的路径也就是指定的文件名

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
