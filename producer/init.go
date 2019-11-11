package producer

import (
	"github.com/Shopify/sarama"
	"gowatcher/go_spider/exceptions"
	"log"
)

var kafkaSender sarama.SyncProducer

func InitProducer() {
	sender, err := NewKafkaSender()
	if err != nil {
		panic(err)
	}
	kafkaSender = sender
}

func NewKafkaSender() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	// 等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 随机的分区类型：返回一个分区器，该分区器每次选择一个随机分区
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	// 是否等待成功和失败后的响应
	config.Producer.Return.Successes = true

	sender, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("init kafka sender failed,err:%v\n", err)
		return nil, exceptions.ErrKafkaHandle
	}

	return sender, nil
}
