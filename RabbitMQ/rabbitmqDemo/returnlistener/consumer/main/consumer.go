package main

import (
	"github.com/streadway/amqp"
	"log"
)

const (
	URL             = "amqp://scx199449:123456@192.168.0.102:5672/"
	EXCHANGE        = "test_return_exchange"
	QUEUENAME       = "test_return_queue"
	ROUTINGKEY      = "return.#"
	ROUTINGKEYERROR = "abc.save"
)

func main() {
	//创建连接
	conn, err := amqp.Dial(URL)
	FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	//创建通道
	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	defer ch.Close()

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

	ch.QueueBind(q.Name, ROUTINGKEY, EXCHANGE, false, nil)

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()
	<-forever //这个地方在channel中取出值之前一直阻塞，从而起到等待多个goroutine执行完的效果
}

//创建一个返回错误打印日志的函数
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
