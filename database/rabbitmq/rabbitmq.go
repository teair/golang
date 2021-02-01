package rabbitmq

import (
	"database/database/rabbitmq/receive"
	"database/database/rabbitmq/send"
)

// Rmqsend 测试rabbitmq发送
func Rmqsend() {
	rabbitmq := send.NewRabbitMQSimple("test")
	rabbitmq.PublishSimple("go test queue啊啊!")
}

// Rmqreceive 测试rabbitmq接收
func Rmqreceive() {
	receive.Receive()
}
