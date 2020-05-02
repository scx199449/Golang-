package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

const (
	URL             = "amqp://scx199449:123456@192.168.0.102:5672/"
	EXCHANGE        = "test_return_exchange"
	ROUTINGKEY      = "return.save"
	ROUTINGKEYERROR = "abc.save"
)

type Produce struct {
	name         string           //routing key
	conn         *amqp.Connection //连接
	ch           *amqp.Channel    //通道
	notifyReturn chan amqp.Return
}

func main() {

	//1、创建连接
	conn, err := amqp.Dial(URL)
	FailOnError(err, "failed to connect to RabbitMQ")
	defer conn.Close()

	//2、创建通道
	ch, err := conn.Channel()
	FailOnError(err, "failed to open a channel")
	defer ch.Close()

	p := &Produce{
		name:         ROUTINGKEYERROR,
		conn:         conn,
		ch:           ch,
		notifyReturn: make(chan amqp.Return),
	}

	p.ch.NotifyReturn(p.notifyReturn)

	//3、发送消息
	msg := "Hello RabbitMQ Return Message"
	ch.Publish(
		EXCHANGE,
		ROUTINGKEYERROR,
		true,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})

	forever := make(chan bool)
	select {
	case r := <-p.notifyReturn:
		fmt.Printf("replycode is :%d; replytext is :%s; replyto is :%s\n", r.ReplyCode, r.ReplyText, r.ReplyTo)
		fmt.Printf("exchange is :%s; routingkey is :%s\n", r.Exchange, r.RoutingKey)
		fmt.Printf("body is :%s", string(r.Body))
	}
	<-forever
}

//创建一个返回错误打印日志的函数
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
