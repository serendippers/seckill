package consumer

//消费者接口
type IMessageConsumer interface {

	//初始化消费者
	ConsumerInit(queueName *string, consumer *string)

}

