package receive

import (
	"database/database/rabbitmq/send"
)

// Receive 接收消息
func Receive() {
	rabbitmq := send.NewRabbitMQSimple("test")
	rabbitmq.ConsumeSimple()
}
