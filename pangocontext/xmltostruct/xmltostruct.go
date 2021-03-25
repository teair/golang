package xmltostruct

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Recurlyservers struct {
	XMLName     xml.Name `xml:"servers"`
	Version     string   `xml:"version,attr"`
	Svs         []server `xml:"server"`
	Description string   `xml:",innerxml"`
}

type server struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"ServerName"`
	ServerIP   string   `xml:"serverIP"`
}

func XMLMain() {

	// 获取xml对象
	file, err := os.Open("xmltostruct/static/servers.xml")
	defer file.Close()
	if err != nil {
		fmt.Printf("get xml object failed! %v", err)
		return
	}

	// 读取xml数据
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("Read xml failed! %v", err)
		return
	}

	// xml 转 struct
	v := Recurlyservers{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("failed to struct %v!", err)
		return
	}

	fmt.Printf("The struct is %v", v)

}
