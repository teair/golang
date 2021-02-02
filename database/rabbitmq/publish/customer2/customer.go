package customer2

import (
	RabbitMQ "database/database/rabbitmq/publish/server"
)

// Customer2 广播模式消费者1
func Customer2() {
	rabbitmq := RabbitMQ.NewRabbitMQPubSub("newProduct")
	rabbitmq.RecieveSub()
}
