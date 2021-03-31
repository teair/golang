package structojson

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	ServerName string `json:"servername"`
	ServerIP   string `json:"serverip"`
	ServerTest int    // 结构体数据类型如果为int则可以string<=>int
}

type ServerSlice struct {
	Servers []Server `json:"servers"` // 可通过tag标签操作json的字段名
}

func ToJson(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var s ServerSlice
		s.Servers = append(s.Servers, Server{ServerName: "ShangHai_VPN", ServerIP: "127.0.0.1", ServerTest: 0})
		s.Servers = append(s.Servers, Server{ServerName: "BeiJing_VPN", ServerIP: "127.0.0.1", ServerTest: 0})
		b, err := json.Marshal(s)
		if err != nil {
			fmt.Printf("failed to json,err:%s", err)
		}

		fmt.Fprintln(w, string(b))
	} else {
		fmt.Fprintln(w, "????????")
	}
}

func ToJsonMain() {

	http.HandleFunc("/tojson", ToJson)

	err := http.ListenAndServe("127.0.0.1:9091", nil) // 设置监听的端口

	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
