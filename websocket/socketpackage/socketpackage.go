package socketpackage

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

// IpMain 验证ip地址是否合法
func IpMain() {

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The addres is ", addr.String())
	}
	os.Exit(0)
}

// SimulationMain  模拟基于http协议的客户端
func SimulationMain() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	_, err = conn.Write([]byte("timestamp"))
	checkError(err)

	result := make([]byte, 256)
	_, err = conn.Read(result)
	checkError(err)

	// 按照格式输出同步的时间
	tmp := string(result)
	fmt.Printf("%T %v\n", tmp, tmp)
	atoi, _ := strconv.ParseInt(tmp, 10, 64)
	//atoi,_ := strconv.Atoi(tmp)
	fmt.Printf("%T %v\n", atoi, atoi)
	fmt.Println(time.Unix(atoi, 0).Format("2006-04-02 15:04:05"))
	//fmt.Println(time.Parse("2006-01-02 15:4:5",string(result)))

	os.Exit(0)
}

// ServerMain 通过 net 包来创建一个服务端的程序，在服务端我们需要绑定服务到指定的非激活端口并监听
func ServerMain() {

	service := "127.0.0.1:9091"

	// ResolveTCPAddr 将service作为tcp地址解析并返回; 返回 网络名和端口名
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	// ListenTCP 类似于监听tcp网络
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		// Accept实现listener接口中的Accept方法;
		// 它等待下一个调用并返回一个泛型conn
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		// 单任务情况
		//daytime := time.Now().String()
		//conn.Write([]byte(daytime))
		//conn.Close()

		// 支持多并发
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	//defer conn.Close()
	//daytime := time.Now().String()
	//conn.Write([]byte(daytime))	// don't care about return value
	//// we're finished with this client

	conn.SetReadDeadline(time.Now().Add(2 * time.Minute)) // set 2 minutes timeout
	request := make([]byte, 128)                          // set maxium request length to 128b to prevent flood attack
	defer conn.Close()                                    // close connection before exit
	for {
		read_len, err := conn.Read(request)
		fmt.Println(string(request))
		if err != nil {
			fmt.Println(err)
			break
		}
		if read_len == 0 {
			break // connection already closed by client
		} else if strings.TrimSpace(string(request[:read_len])) == "timestamp" {
			daytime := strconv.FormatInt(time.Now().Unix(), 10)
			conn.Write([]byte(daytime))
		} else {
			daytime := time.Now().String()
			conn.Write([]byte(daytime))
		}

		request = make([]byte, 128) // clear last read content
	}
}

func checkError(e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", e.Error())
		os.Exit(1)
	}
}
