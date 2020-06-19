package initialize

import (
	"fmt"
	"github.com/streadway/amqp"
	"seckill/core/consumer"
	"seckill/core/producer"
	"seckill/global"
)

func InitRabbitMQ() {

	rabbitConfig := global.CONFIG.RabbitMQ

	conn, err := amqp.Dial(rabbitConfig.Path)

	if err != nil {
		global.LOG.Errorf("init rabbitMQ fail, error is %s\n", err)
	} else {
		global.MQ = conn
		global.LOG.Info("init RabbitMQ success")
	}

	orderConsumer := new(consumer.OrderConsumer)
	startTask(orderConsumer, global.CONFIG.ConsumerConfig.OrderQueueName, global.CONFIG.ConsumerConfig.OrderPoolSize)
	orderProducer := new(producer.OrderProducer)
	orderProducer.ProducerInit(global.CONFIG.ConsumerConfig.OrderQueueName)

}

func startTask(consumer consumer.IMessageConsumer, queueNqme string, poolSize int) {

	for i := 0; i < poolSize; i++ {
		consumerName := fmt.Sprintf("%s:%d", queueNqme, i)
		go consumer.ConsumerInit(&queueNqme, &consumerName)

	}
}

//func TestRabbitMQInit() {
//	//创建通道
//	ch, err := global.MQ.Channel()
//	if err != nil {
//		global.LOG.Errorf("无法打开通道，error is %s\n", err)
//	}
//	defer ch.Close()
//
//	q, err := ch.QueueDeclare(
//		"hello",
//		false,
//		false,
//		false,
//		false,
//		nil)
//
//	if err != nil {
//		global.LOG.Errorf("无法声明队列，err is %s\n", err)
//	}
//
//	body := "hello world!"
//
//	err = ch.Publish(
//		"",
//		q.Name,
//		false,
//		false,
//		amqp.Publishing{
//			ContentType: "text/plain",
//			Body:        []byte(body),
//		})
//	if err != nil {
//		global.LOG.Errorf("Failed to publish a message, error is %s\n", err)
//	}
//}
//
//func seckillConsumer() {
//
//}
