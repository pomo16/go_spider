package test

import (
	"context"
	"fmt"
	"gowatcher/go_spider/consts"
	"gowatcher/go_spider/model"
	"gowatcher/go_spider/service"
	"gowatcher/go_spider/service/elasticsearch"
	"testing"
	"time"
)

func TestConfig(t *testing.T) {
	addr, err := elasticsearch.ReadYamlConfig()
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	fmt.Println(addr)
}

func TestES(t *testing.T) {
	elasticsearch.InitElasticSearch()

	comment := &model.Comment{
		CommentId:   "10000",
		MainId:      "test10000",
		Content:     "test content",
		Rating:      "5",
		Version:     "0.0.0",
		PublishTime: time.Now().Format(consts.TimeStr),
		CrawlTime:   time.Now().Format(consts.TimeStr),
	}

	err := elasticsearch.InsertNewComment(context.Background(), comment)
	if err != nil {
		fmt.Println(err)
	}
}

func TestSinkGraph(t *testing.T) {
	comment1 := &model.Comment{
		CommentId:   "1",
		MainId:      "1",
		Content:     "1",
		Rating:      "1",
		Version:     "1",
		PublishTime: "1",
		CrawlTime:   "1",
	}
	comment2 := &model.Comment{
		CommentId:   "2",
		MainId:      "2",
		Content:     "2",
		Rating:      "2",
		Version:     "2",
		PublishTime: "2",
		CrawlTime:   "2",
	}

	graph := service.Graph{"1": comment1, "2": comment2}
	elasticsearch.SinkGraph(graph)
}
