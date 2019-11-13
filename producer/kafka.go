package producer

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"gowatcher/go_spider/consts"
	"gowatcher/go_spider/model"
)

func SendToKafka(comment model.Comment) {
	msg := &sarama.ProducerMessage{
		Topic: consts.Topic,
	}

	xx, _ := json.Marshal(comment)
	msg.Value = sarama.ByteEncoder(xx)
	_, _, err := kafkaSender.SendMessage(msg)
	if err != nil {
		fmt.Println("Send message Fail")
	}
}
