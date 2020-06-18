package producer

import (
	"github.com/streadway/amqp"
	"seckill/global"
)

var ORDER_PRODUCER *OrderProducer

type OrderProducer struct {
	channel   *amqp.Channel
	queueName string
	exchange  string
}

func (producer *OrderProducer) ProducerInit(queueName string) {
	ch, err := global.MQ.Channel()

	if err != nil {
		global.LOG.Errorf("OrderProducer open channel fail, error is %v\n", err)
		return
	}
	producer.channel = ch
	producer.queueName = queueName

	ORDER_PRODUCER = producer
}

func (producer *OrderProducer) SendMessage(message []byte) {

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

	err = producer.channel.Publish(
		producer.exchange,
		q.Name,
		false, // 如果为true, 会根据exchange类型和routkey规则，如果无法找到符合条件的队列那么会把发送的消息返回给发送者
		false, // 如果为true, 当exchange发送消息到队列后发现队列上没有绑定消费者，则会把消息发还给发送者
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		})
	if err != nil {
		global.LOG.Errorf("Failed to send message, queue is %s ,error is %v", producer.queueName, err)
	}
}
