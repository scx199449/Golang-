package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

const (
	URL        = "amqp://scx199449:123456@192.168.0.102:5672/"
	EXCHANGE   = "test_qos_exchange"
	ROUTINGKEY = "qos.#"
	QUEUENAME  = "test_qos_queue"
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
	q, err := ch.QueueDeclare(
		QUEUENAME, // name
		true,      // durable
		false,     // delete when usused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	FailOnError(err, "Failed to declare a queue")

	//交换机与队列绑定
	ch.QueueBind(q.Name, ROUTINGKEY, EXCHANGE, false, nil)

	//1、限流第一步设置一次性消费的消息数量
	ch.Qos(1, 0, false)

	//2、限流方式 第二件事就是autoAck设置为false
	msgs, err := ch.Consume(
		QUEUENAME, // queue
		"",        // consumer
		false,     // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)
	for d := range msgs {
		fmt.Println("consumertag is :", d.ConsumerTag)
		log.Printf("Received a message: %s", d.Body)
		d.Ack(false)
	}
	<-forever

}

//创建一个返回错误打印日志的函数
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
