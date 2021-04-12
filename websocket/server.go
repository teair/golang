package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// SocketMain 测试
func SocketMain() {
	router := mux.NewRouter()
	go h.run()
	router.HandleFunc("/stt", myws)
	if err := http.ListenAndServe("127.0.0.1:7978", router); err != nil {
		fmt.Println("err:", err)
	}
}
