package test

import (
	"fmt"
	"github.com/streadway/amqp"
	"seckill/global"
	"testing"
)

func TestRabbitMQInit(t *testing.T) {
	fmt.Println("test")
	global.LOG.Info("global.MQ is %T\n", global.MQ)
	//创建通道
	ch, err := global.MQ.Channel()
	if err != nil {
		global.LOG.Errorf("无法打开通道，error is %s\n", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil)

	if err != nil {
		global.LOG.Errorf("无法声明队列，err is %s\n", err)
	}

	body := "hello world!"

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	global.LOG.Errorf("Failed to publish a message, error is %s\n", err)
}
