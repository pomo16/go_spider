package test

import (
	"gowatcher/go_spider/producer"
	"testing"
)

func TestKafka(t *testing.T) {
	producer.InitProducer()
	producer.SendToKafka()
}
