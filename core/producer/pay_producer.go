package producer

import (
	"github.com/streadway/amqp"
	"seckill/config"
	"seckill/global"
	"time"
)

var PAY_PRODUCER *PayProducer

type PayProducer struct {
	channel      *amqp.Channel
	queueName    string
	exchangeName string
	exchangeType string
	dlxQueueName string
}

//初始化生产者（处理支付的延迟队列）
func (producer *PayProducer) ProducerInit(config *config.ConsumerConfig) {
	ch, err := global.MQ.Channel()

	if err != nil {
		global.LOG.Errorf("OrderProducer open channel fail, error is %v\n", err)
		return
	}
	producer.channel = ch
	producer.queueName = config.OrderQueueName
	producer.dlxQueueName = config.PayDlxQueueName

	PAY_PRODUCER = producer
}

func (producer *PayProducer) SendMessage(message []byte) {

	q, err := producer.channel.QueueDeclare(
		producer.queueName, //名称
		true,               //持久性
		false,              //不用时删除
		false,              //排他性
		false,              //不等待
		nil,                //参数
	)

	if err != nil {
		global.LOG.Errorf("Failed to declare a queue, queue is %s ,error is %v", producer.queueName, err)
	}

	//if err = producer.prepareExchange(); err != nil {
	//	global.LOG.Errorf("Failed to declare a exchange, exchange is %s ,error is %v", producer.exchangeName, err)
	//}

	err = producer.channel.Publish(
		producer.exchangeName,
		q.Name,
		false, // 如果为true, 会根据exchange类型和routkey规则，如果无法找到符合条件的队列那么会把发送的消息返回给发送者
		false, // 如果为true, 当exchange发送消息到队列后发现队列上没有绑定消费者，则会把消息发还给发送者
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
			Expiration:  string(30 * time.Minute.Milliseconds()),//设置消息30分钟过期
		})
	if err != nil {
		global.LOG.Errorf("Failed to send message, queue is %s ,error is %v", producer.queueName, err)
	}

}

func (producer *PayProducer) prepareExchange() error {
	return producer.channel.ExchangeDeclare(
		producer.exchangeName,
		producer.exchangeType,
		true,
		false,
		false,
		false,
		nil,
	)
}
