package myfilelist

import (
	"demo/immok/basic/errhandling/filelistingserver/myfilelist/filelist"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

type appHandle func(writer http.ResponseWriter, request *http.Request) error

type userError interface {
	error
	Message() string
}

func wrapperHandle(handle appHandle) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		// 对于访问未知目录接管异常处理
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
			//log.Warn("Error handle request:", err.Error())
			// 默认的log
			log.Printf("Error occured "+"handle request:%s", err.Error())

			// 判断当前错误是否为对用户可见的错误【用户可见错误】
			if userErr, ok := err.(userError); ok {
				http.Error(
					writer,
					userErr.Message(), // 给用户看
					http.StatusBadRequest,
				)
				return
			}

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

	http.HandleFunc("/", wrapperHandle(filelist.MyFileHandle))

	http.ListenAndServe(":9999", nil)
}
