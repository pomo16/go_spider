package elasticsearch

import (
	"context"
	"github.com/sirupsen/logrus"
	"gowatcher/go_spider/consts"
	"gowatcher/go_spider/exceptions"
	"gowatcher/go_spider/model"
	"gowatcher/go_spider/service"
)

//InsertNewComment 插入新的反馈数据
func InsertNewComment(ctx context.Context, comment *model.Comment) error {
	if comment == nil {
		return exceptions.ErrValueEmpty
	}

	_, err := elasticClient.Index().
		Index(consts.ESTempIndex).
		Id(comment.MainId).
		BodyJson(comment).
		Do(ctx)

	if err != nil {
		logrus.Error("insert comment error:", err)
		return err
	}

	return nil
}

//SinkGraph 评论集合落地
func SinkGraph(graph service.Graph) {
	for _, comment := range graph {
		err := InsertNewComment(context.Background(), comment)
		if err == nil {
			logrus.Infof("%+v", comment)
		} else {
			logrus.Errorf("%+v", comment)
		}
	}
}
