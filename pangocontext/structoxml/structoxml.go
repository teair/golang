package structoxml

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Servers struct {
	XMLName xml.Name `xml:"servers"`
	Version string   `xml:"version,attr"`
	Svs     []server `xml:"server"`
}

type server struct {
	ServerName string `xml:"ServerName"`
	ServerIP   string `xml:"ServerIP"`
	// ServerTest string `xml:",chardata"`
	// ServerTest string `xml:",innerxml"`
	ServerTest string `xml:",comment"` // 注释中不能含有 --
	// ServerTest string `xml:",omitempty"`
	// Aa string `xml:"aa>bb>cc"`
}

func StructMain() {

	v := &Servers{Version: "1"}

	v.Svs = append(v.Svs, server{"ShangHai_VPN", "127.0.0.1", "aa"})

	v.Svs = append(v.Svs, server{"BeiJing_VPN", "127.0.0.1", "aa"})

	output, err := xml.MarshalIndent(v, " ", "  ")

	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write([]byte(xml.Header))

	os.Stdout.Write(output)
}
