package consumer

import (
	"seckill/global"
	"time"
)

type OrderConsumer struct {
}

func (order *OrderConsumer) ConsumerInit(queueName *string, consumer *string) {

	ch, err := global.MQ.Channel()

	if err != nil {
		global.LOG.Errorf("open channel fail, error is %v\n", err)
		return
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		*queueName,
		true,  //是否持久
		false, //不使用时删除
		false, //是否具有排他性
		false, //是否阻塞
		nil,   //额外属性
	)

	if err != nil {
		global.LOG.Errorf("Failed to declare a queue, queue is %s ,error is %v", queueName, err)
	}

	err = ch.Qos(
		1, //预取计数
		0, //预取大小
		false,
	)

	if err != nil {
		global.LOG.Errorf("Failed to set QoS, err is %v", err)
	}

	msgs, err := ch.Consume(
		q.Name,
		*consumer, //消费者名字
		false,     //是否自动应答
		false,     //是否具有排他性
		false,     // 如果设置为true，表示不能将同一个connection中发送的消息传递给这个connection中的消费者
		false,     // 队列消费是否阻塞
		nil,
	)

	if err != nil {
		global.LOG.Errorf("Failed to register a consumer, consumer is %s, err is %v", consumer, err)
	}

	global.LOG.Infof("init %s success", *consumer)
	//forever := make(chan bool)

	for d := range msgs {
		global.LOG.Info("Received a message: %s, consumer is %s", d.Body, *consumer)
		time.Sleep(10 * time.Second)
		_ = d.Ack(false)
	}

	//<-forever
}
