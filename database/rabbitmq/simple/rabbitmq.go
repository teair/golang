package simple

import (
	"database/database/rabbitmq/simple/receive"
	"database/database/rabbitmq/simple/send"
	"fmt"
	"strconv"
	"time"
)

// Rmqsend 测试rabbitmq发送
func Rmqsend() {

	rabbitmq := send.NewRabbitMQSimple("test")

	for i := 0; i <= 10; i++ {
		rabbitmq.PublishSimple("测试 test 队列能否收到消息!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}

// Rmqreceive 测试rabbitmq接收
func Rmqreceive() {
	receive.Receive()
}

// Rmqreceive2 第二个消费者
func Rmqreceive2() {
	receive.Receive()
}
