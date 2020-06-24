package producer

import "seckill/config"

//生产者接口
type IMessageProducer interface {

	//初始化生产者
	ProducerInit(config *config.ConsumerConfig)

	SendMessage(message []byte)
}
