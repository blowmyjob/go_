package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/zngw/log"
	"strings"
)

var producer sarama.AsyncProducer

// 初始化生产者
func InitProducer(hosts string) {
	config := sarama.NewConfig()
	client, err := sarama.NewClient(strings.Split(hosts, ","), config)
	if err != nil {
		log.Error("unable to create kafka client: ", err)
	}
	producer, err = sarama.NewAsyncProducerFromClient(client)
	if err != nil {
		log.Error(err)
	}
}

// 发送消息
func Send(topic, data string) {
	producer.Input() <- &sarama.ProducerMessage{Topic: topic, Key: nil, Value: sarama.StringEncoder(data)}
	log.Trace("kafka", "Produced message: ["+data+"]")
}

func Close() {
	if producer != nil {
		producer.Close()
	}
}
