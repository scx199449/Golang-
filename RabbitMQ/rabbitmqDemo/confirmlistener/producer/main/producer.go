package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

const (
	EXCHANGENAME = "test_confirm_exchange"
	ROUTINGKEY   = "confirmlistener"
)

type Producer struct {
	name          string           //routing key
	conn          *amqp.Connection //连接
	ch            *amqp.Channel    //通道
	notifyconfirm chan amqp.Confirmation
}

//生产者
func main() {

	//1、创建connection
	conn, err := amqp.Dial("amqp://scx199449:123456@192.168.0.102:5672/")
	FailOnError(err, "failed to connect to RabbitMQ")
	defer conn.Close()

	//2、通过connection创建一个新的channel
	ch, err := conn.Channel()
	FailOnError(err, "failed to open a channel")
	defer ch.Close()

	p := &Producer{
		name:          ROUTINGKEY,
		conn:          conn,
		ch:            ch,
		notifyconfirm: make(chan amqp.Confirmation),
	}

	//3、声明使用confirm消息确认机制
	p.ch.Confirm(false)
	//4、注册监听
	p.ch.NotifyPublish(p.notifyconfirm)

	//5、发送一条消息
	body := "hello, rabbitmq confirmlistener to golang  "
	err = p.ch.Publish(
		EXCHANGENAME,
		ROUTINGKEY,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	FailOnError(err, "Failed to publish a message")

	forever := make(chan bool)
	select {
	case confirm := <-p.notifyconfirm:
		fmt.Println("ack is :", confirm.Ack)
		if confirm.Ack {
			fmt.Println("confirmlistener is true")
		}
	}
	<-forever
}

//创建一个返回错误打印日志的函数
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
