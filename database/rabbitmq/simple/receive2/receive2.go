package receive2

import (
	"shentong/rabbitmq/simple/send"
)

// Receive 接收消息
func Receive() {
	rabbitmq := send.NewRabbitMQSimple("test")
	rabbitmq.ConsumeSimple()
}
