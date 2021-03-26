package jsontostruct

import (
	"encoding/json"
	"fmt"

	"github.com/bitly/go-simplejson"
)

type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
	Servers []Server
}

func ToStructMain() {
	var s Serverslice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)
}

func SimpleJson() {

	json := `{"test": {"array": [1, "2", 3],"int": 10,"float": 5.150,"bignum": 9223372036854775807,"string": "simplejson","bool": true}}`

	js, _ := simplejson.NewJson([]byte(json))

	arr, _ := js.Get("test").Get("array").Array()
	i, _ := js.Get("test").Get("int").Int()
	ms := js.Get("test").Get("string").MustString()

	fmt.Printf("The data's array is: %s\n", arr)
	fmt.Printf("The data's int is: %v\n", i)
	fmt.Printf("The data's string is: %s\n", ms)
}
