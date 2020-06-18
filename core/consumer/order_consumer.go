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
		true,  //持久
		false, //不使用时删除
		false,
		false,
		nil,
	)

	if err != nil {
		global.LOG.Errorf("Failed to declare a queue, queue is %s ,error is %v", queueName, err)
	}

	err = ch.Qos(
		1,
		0,
		false,
	)

	if err != nil {
		global.LOG.Errorf("Failed to set QoS, err is %v", err)
	}

	msgs, err := ch.Consume(
		q.Name,
		*consumer,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		global.LOG.Errorf("Failed to register a consumer, consumer is %s, err is %v", consumer, err)
	}
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			global.LOG.Info("Received a message: %s", d.Body)
			time.Sleep(10 * time.Second)
			_ = d.Ack(false)
		}
	}()

	<-forever
}
