package main

import (
	"github.com/streadway/amqp"
	"log"
	"strconv"
	"time"
)

const (
	URL        = "amqp://scx199449:123456@192.168.0.102:5672/"
	EXCHANGE   = "test_ack_exchange"
	ROUTINGKEY = "ack.save"
)

//生产者
func main() {
	//创建一个连接
	conn, err := amqp.Dial(URL)
	FailOnError(err, "failed to connect to RabbitMQ")
	defer conn.Close()

	//开启通道
	ch, err := conn.Channel()
	FailOnError(err, "failed to open a channel")
	defer ch.Close()

	//发布消息
	forever := make(chan bool)
	for i := 0; i < 5; i++ {
		body := "hello, RabbitMQ ack Message :  " + strconv.Itoa(i)
		headers := make(map[string]interface{})
		headers["num"] = int64(i)
		err = ch.Publish(
			EXCHANGE,
			ROUTINGKEY,
			false,
			false,
			amqp.Publishing{
				ContentType:     "text/plain",
				Body:            []byte(body),
				DeliveryMode:    2,
				ContentEncoding: "utf-8",
				Headers:         headers,
			})
		FailOnError(err, "Failed to publish a message")
	}
	time.Sleep(1 * time.Second)
	<-forever
}

//创建一个返回错误打印日志的函数
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
