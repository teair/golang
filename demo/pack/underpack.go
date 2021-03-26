package pack

import (
	"fmt"
	_ "os"
)

var name = "申通通"

func UnderPack() {
	fmt.Println(name)
	/*buf := make([]byte,1024)
	f,_ := os.Open("/test.txt")
	fmt.Println(buf)
	fmt.Println(f)
	os.Exit(0)
	defer f.close()
	for {
		n,_ := f.Read(buf)
		if n == 0 {
			break
		}
		os.Stdout.Write(buf[:n])
	}*/
}
