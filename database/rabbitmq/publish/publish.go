package publish

import (
	Rabbitmq "database/database/rabbitmq/publish/server"
	"strconv"
	"time"
)

// Publish 生产者
func Publish() {

	// 获取生产者对象
	rabbitmq := Rabbitmq.NewRabbitMQPubSub("newProduct")

	for i := 0; i < 10; i++ {

		rabbitmq.PublishPub("订阅模式生产第:" + strconv.Itoa(i) + "条数据!")

		// fmt.Println("订阅模式生产第:" + strconv.Itoa(i) + "条数据！")

		time.Sleep(1 * time.Second)
	}
}

// Customer1 消费者
func Customer1() {

	rabbitmq := Rabbitmq.NewRabbitMQPubSub("newProduct")

	rabbitmq.RecieveSub()
}

// Customer2 消费者
func Customer2() {

	rabbitmq := Rabbitmq.NewRabbitMQPubSub("newProduct")

	rabbitmq.RecieveSub()
}
