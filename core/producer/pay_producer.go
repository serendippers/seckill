package producer

import (
	"errors"
	"fmt"
	"github.com/streadway/amqp"
	"seckill/config"
	"seckill/global"
)

var PAY_PRODUCER *PayProducer

type PayProducer struct {
	channel         *amqp.Channel
	queueName       string
	exchangeName    string
	exchangeType    string
	dlxQueueName    string
	dlxExchangeName string
	dlxExchangeType string
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
	producer.dlxQueueName = config.DlxQueueName
	producer.dlxExchangeName = config.DlxExchangeName
	producer.dlxExchangeType = config.DlxExchangeType

	PAY_PRODUCER = producer
}

func (producer *PayProducer) SendMessage(message []byte) {

	argsQue := amqp.Table{"x-dead-letter-exchange": producer.dlxExchangeName}

	//设置存储订单的队列，并设置队列的死信交换机为 dlx_exchange
	q, err := producer.channel.QueueDeclare(
		producer.queueName, //名称
		true,               //持久性
		false,              //不用时删除
		false,              //排他性
		false,              //不等待
		argsQue,            //参数
	)

	if err != nil {
		global.LOG.Errorf("Failed to declare a queue, queue is %s ,error is %v", producer.queueName, err)
	}

	//TODO 不应该在这里声明交换机，应把声明交换机这部分代码在优化一下
	if err = producer.prepareExchange(producer.dlxExchangeName, producer.dlxExchangeType); err != nil {
		global.LOG.Errorf("Failed to declare a exchange, exchange is %s ,error is %v", producer.dlxExchangeName, err)
	}

	err = producer.channel.QueueBind(
		producer.dlxQueueName,//延时队列
		"", // routing key
		producer.dlxExchangeName,
		false,
		nil,
	)

	err = producer.channel.Publish(
		producer.exchangeName,
		q.Name,
		false, // 如果为true, 会根据exchange类型和routkey规则，如果无法找到符合条件的队列那么会把发送的消息返回给发送者
		false, // 如果为true, 当exchange发送消息到队列后发现队列上没有绑定消费者，则会把消息发还给发送者
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
			//Expiration:  string(30 * time.Minute.Milliseconds()),//设置消息30分钟过期
			Expiration: "6000",
		})
	if err != nil {
		global.LOG.Errorf("Failed to send message, queue is %s ,error is %v", producer.queueName, err)
	}
}

//初始化交换机
//TODO 这个方法还需要整理，不应该放在这里，应该提取出来作为一个公共的方法
func (producer *PayProducer) prepareExchange(exchangeName, exchangeType string) error {
	if exchangeName == "" || exchangeType == "" {
		return errors.New(fmt.Sprintf("invalid param, exchangeName is %s, exchangeType is %s", exchangeName, exchangeType))
	}
	return producer.channel.ExchangeDeclare(
		exchangeName,
		exchangeType,
		true,
		false,
		false,
		false,
		nil,
	)
}
