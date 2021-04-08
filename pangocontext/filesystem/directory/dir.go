package directory

import (
	"fmt"
	"os"
)

func DirMain() {
	os.Mkdir("ShenTong", 0777)
	os.MkdirAll("shen/tong/tong", 0777)
	err := os.Remove("ShenTong")
	if err != nil {
		fmt.Println(err)
	}
	os.RemoveAll("shen")
}
