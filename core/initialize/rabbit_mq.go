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

	//TODO 定义消费者的函数可以独立出来，将要执行的方法当成参数传入
	//订单生产者，请求通过redis扣减库存，通过orderProducer生产数据
	orderProducer := new(producer.OrderProducer)
	orderProducer.ProducerInit(&global.CONFIG.ConsumerConfig)
	//消费orderProducer产生的消息，检测库存，生成订单，购买记录
	orderConsumer := new(consumer.OrderConsumer)
	startTask(orderConsumer, global.CONFIG.ConsumerConfig.OrderQueueName, global.CONFIG.ConsumerConfig.OrderPoolSize)


	dlxConsumer := new(consumer.DlxConsumer)
	startTask(dlxConsumer, global.CONFIG.ConsumerConfig.DlxQueueName, global.CONFIG.ConsumerConfig.DlxPoolSize)

}

func startTask(consumer consumer.IMessageConsumer, queueNqme string, poolSize int) {

	for i := 0; i < poolSize; i++ {
		consumerName := fmt.Sprintf("%s:%d", queueNqme, i)
		go consumer.ConsumerInit(&queueNqme, &consumerName)

	}
}

