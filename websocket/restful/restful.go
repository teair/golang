package restful

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Welcome REST!\n")
}
func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName("name")
	fmt.Fprintf(w, "Hello,%s", name)
}
func adduser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	fmt.Fprintf(w, "You are add user %s", uid)
}
func deleteuser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	fmt.Fprintf(w, "You are delete user %s", uid)
}
func modifyuser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	fmt.Fprintf(w, "You are modify user %s", uid)
}
func user(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uname := ps.ByName("uname")
	fmt.Fprintf(w, "You are find user %s", uname)
}

// RestMain REST架构路由接口设计
func RestMain() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello", Hello)

	router.GET("/user/:uname", user)
	router.GET("/deleteuser/:uid", deleteuser)
	router.GET("/modifyuser/:uid", modifyuser)
	router.GET("/adduser/:uid", adduser)

	log.Fatal(http.ListenAndServe(":8080", router))
}
