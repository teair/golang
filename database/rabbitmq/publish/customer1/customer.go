package customer1

import (
	RabbitMQ "shentong/database/rabbitmq/publish/server"
)

// Customer1 广播模式消费者1
func Customer1() {
	rabbitmq := RabbitMQ.NewRabbitMQPubSub("newProduct")
	rabbitmq.RecieveSub()
}
