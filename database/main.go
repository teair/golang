package main

import (
	"database/database/rabbitmq"
	"fmt"
)

func main() {

	fmt.Println("11")

	// Rmqsend 测试rabbitmq发送
	// rabbitmq.Rmqsend()

	//Rmqsend 测试rabbitmq发送
	rabbitmq.Rmqsend()

	//Rmqreceive 测试rabbitmq接收
	rabbitmq.Rmqreceive()

	//mysql.Mysql()
}
