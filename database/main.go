package main

import (
	my "shentong/mysql"
)

func main() {

	// fmt.Println("11")

	// Rmqsend 测试rabbitmq发送
	// rabbitmq.Rmqsend()

	// simple(一对一)模式推送消息

	//Rmqsend 测试rabbitmq发送
	// rabbitmq.Rmqsend()

	// //Rmqreceive 测试rabbitmq接收
	// rabbitmq.Rmqreceive()

	// //Rmqreceive2 测试多个消费者
	// rabbitmq.Rmqreceive2()

	/**
	 * fanout 类型
	 */
	// rabbitmq.Publish()

	// rabbitmq.Customer1()

	// rabbitmq.Customer2()

	// mysql.Mysql()

	my.OrmMain()

	// sessioncookie.CookieMain()

	//beego.OrmMain()

	//fmt.Println("111111111")
}
