package main

import (
	"github.com/streadway/amqp"
	"log"
	"strconv"
	"time"
)

//生产者
func main() {
	//创建一个连接
	conn,err := amqp.Dial("amqp://scx199449:123456@192.168.0.102:5672/")
	FailOnError(err,"failed to connect to RabbitMQ")
	defer conn.Close()

	//开启通道
	ch,err := conn.Channel()
	FailOnError(err,"failed to open a channel")
	defer ch.Close()

	//声明队列
	q,err := ch.QueueDeclare(
		"test number",
		false,
		false,
		false,
		false,
		nil,
		)
	FailOnError(err,"Failed to declare a queue")

	//发布消息
	for i := 0; i < 60; i++ {
		body := "hello, ZhangShouFu  " + strconv.Itoa(i)
		err = ch.Publish(
			"", //使用默认的exchange
			q.Name,
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		FailOnError(err, "Failed to publish a message")
		time.Sleep(1 * time.Second)
	}

}

//创建一个返回错误打印日志的函数
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}