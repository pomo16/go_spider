package test

import (
	"gowatcher/go_spider/model"
	"gowatcher/go_spider/producer"
	"testing"
)

func TestKafka(t *testing.T) {
	producer.InitProducer()
	comment := model.Comment{
		CommentId:   "1",
		MainId:      "123",
		Content:     "xxx",
		Rating:      "3",
		Version:     "3.1.2",
		PublishTime: "",
		CrawlTime:   "",
	}
	producer.SendToKafka(comment)
}
