package filelistingserver

import (
	"demo/immok/basic/errhandling/filelistingserver/filelisting"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errorWrapper(handle appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := handle(writer, request)
		if err != nil {

			// 控制台打印错误
			log.Warn("Error handling request:", err.Error())

			code := http.StatusOK

			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}

			http.Error(
				writer,
				http.StatusText(code),
				code,
			)
		}
	}
}

func Main() {
	http.HandleFunc("/list/", errorWrapper(filelisting.HandleFileList))
	err := http.ListenAndServe(
		":8888",
		nil,
	)
	if err != nil {
		panic(err)
	}
}
