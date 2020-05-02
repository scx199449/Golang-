package main

import (
	"github.com/streadway/amqp"
	"log"
)

const (
	URL        = "amqp://scx199449:123456@192.168.0.102:5672/"
	EXCHANGE   = "test_dlx_exchange"
	ROUTINGKEY = "dlx.#"
	QUEUENAME  = "test_dlx_queue"
)

//消费者
func main() {
	//创建连接
	conn, err := amqp.Dial(URL)
	FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	//创建通道
	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	//创建交换机
	ch.ExchangeDeclare(EXCHANGE, "topic", true, false, false, false, nil)

	//创建队列
	arguments := make(map[string]interface{})
	arguments["x-dead-letter-exchange"] = "dlx.exchange"

	q, err := ch.QueueDeclare(
		QUEUENAME, // name
		true,      // durable
		false,     // delete when usused
		false,     // exclusive
		false,     // no-wait
		arguments, // arguments
	)
	FailOnError(err, "Failed to declare a queue")
	//交换机与队列绑定
	ch.QueueBind(q.Name, ROUTINGKEY, EXCHANGE, false, nil)

	//进行死信队列声明
	ch.ExchangeDeclare("dlx.exchange", "topic", true, false, false, false, nil)
	ch.QueueDeclare("dlx.queue", true, false, false, false, nil)
	ch.QueueBind("dlx.queue", "#", "dlx.exchange", false, nil)

	msgs, err := ch.Consume(
		QUEUENAME, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	FailOnError(err, "Failed to register a consumer")

	for d := range msgs {
		log.Printf("Received a message: %s", d.Body)
	}

}

//创建一个返回错误打印日志的函数
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
