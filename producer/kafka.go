package producer

import (
	"fmt"
	"github.com/Shopify/sarama"
	"gowatcher/go_spider/consts"
	"time"
)

func SendToKafka() {
	//构建发送的消息，
	msg := &sarama.ProducerMessage{
		Topic: consts.Topic, //包含了消息的主题
	}

	var value string
	value = time.Now().Format(consts.TimeStr)
	msg.Value = sarama.ByteEncoder(value)
	partition, offset, err := kafkaSender.SendMessage(msg)

	if err != nil {
		fmt.Println("Send message Fail")
	}
	fmt.Printf("Partition = %d, offset=%d\n", partition, offset)
	time.Sleep(10 * time.Second)
}
