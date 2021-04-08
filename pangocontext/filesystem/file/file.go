package file

import (
	"fmt"
	"os"
)

func WriteFile() {
	userfile := "test.txt"
	fout, err := os.Create(userfile)
	if err != nil {
		fmt.Println(userfile, err)
		return
	}
	defer fout.Close()
	fout.WriteString("鸥.外。金额）佛士大夫立刻四大发明里卡多什么!")
}

func ReadFile() {
	userFile := "test.txt"
	fl, err := os.Open(userFile)
	if err != nil {
		fmt.Println(userFile, err)
	}
	defer fl.Close()

	buf := make([]byte, 1024)

	// 读取全部文本内容
	// for {
	// 	n, _ := fl.Read(buf)
	// 	if n == 0 {
	// 		break
	// 	}
	// 	os.Stdout.Write(buf[:n])
	// }s

	// 读取部分文本内容(一个汉字占用3个字节)
	_, err = fl.ReadAt(buf, 25)

	fmt.Printf("%s", buf)

	// fmt.Printf("%s", buf[25:])

	// os.Stdout.Write(buf[n:])
}

func RemoveFile() {
	userfile := "test.txt"
	os.Remove(userfile)
}
