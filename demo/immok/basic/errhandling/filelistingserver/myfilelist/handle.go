package myfilelist

import (
	"demo/immok/basic/errhandling/filelistingserver/myfilelist/filelist"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

type appHandle func(writer http.ResponseWriter, request *http.Request) error

func wrapperHandle(handle appHandle) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := handle(writer, request)
		if err != nil {
			//log.Warn("Error handle request:", err.Error())
			// 默认的log
			log.Printf("Error occured "+"handle request:%s", err.Error())
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

	http.HandleFunc("/myfile/", wrapperHandle(filelist.MyFileHandle))

	http.ListenAndServe(":9999", nil)
}
