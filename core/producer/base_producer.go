package producer

//生产者接口
type IMessageProducer interface {

	//初始化生产者
	ProducerInit(queueName string)

	SendMessage(message []byte)
}
