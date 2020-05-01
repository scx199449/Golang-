package main

import (
	"github.com/streadway/amqp"
	"log"
	"time"
)
//消费者
func main() {
	//创建连接
	conn,err := amqp.Dial("amqp://scx199449:123456@192.168.0.102:5672")
    FailOnError(err,"Failed to connect to RabbitMQ")
	defer conn.Close()

	//创建通道
	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	//创建队列
	q, err := ch.QueueDeclare(
		"test number", // name
		false,         // durable
		false,         // delete when usused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	FailOnError(err, "Failed to declare a queue")

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

	/*wg := sync.WaitGroup{}
	wg.Add(60)*/
	forever  := make(chan bool)
	for i := 0; i < 60; i++ {
		go func() {
			for d := range msgs {
				log.Printf("Received a message: %s", d.Body)
			}
			//wg.Done()
		}()
		time.Sleep(5 * time.Second)
	}
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<- forever //这个地方在channel中取出值之前一直阻塞，从而起到等待多个goroutine执行完的效果
	//wg.Wait()


}

//创建一个返回错误打印日志的函数
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}