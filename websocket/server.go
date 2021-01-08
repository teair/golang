
package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter()
    go h.run()
    router.HandleFunc("/stt", myws)
    if err := http.ListenAndServe("0.0.0.0:7978", router); err != nil {
        fmt.Println("err:", err)
    }
}
