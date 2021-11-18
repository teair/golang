package filelistingserver

import (
	"demo/immok/basic/errhandling/filelistingserver/filelisting"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

type userError interface {
	error
	Message() string
}

func errorWrapper(handle appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {

		// 【Panic】
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic:%v", r)
				http.Error(
					writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError,
				)
			}
		}()

		err := handle(writer, request)

		if err != nil {
			// 控制台打印错误
			log.Warn("Error handling request:", err.Error())

			// 对外可访问错误【userError】
			if userErr, ok := err.(userError); ok {
				http.Error(
					writer,
					userErr.Message(),
					http.StatusBadRequest,
				)
				return
			}

			code := http.StatusOK

			// 【SystemError】
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
	http.HandleFunc("/", errorWrapper(filelisting.HandleFileList))
	err := http.ListenAndServe(
		":8888",
		nil,
	)
	if err != nil {
		panic(err)
	}
}
