package text_analyze

import (
	"gowatcher/go_spider/algoml"
	"gowatcher/go_spider/exceptions"
	"gowatcher/go_spider/model"
	"gowatcher/go_spider/service"
)

//GetSentimentPolarity 调用情感分析模型分析文本极性
func GetSentimentPolarity(doc string) (*model.SentimentInfo, error) {
	if doc == "" {
		return nil, exceptions.ErrValueEmpty
	}

	score := algoml.SentiML.Classify(doc)
	polarity := EchoPolarity(score)
	result := &model.SentimentInfo{
		Score:    score,
		Polarity: polarity,
	}

	return result, nil
}

//EchoPolarity 将情感分析分数转换为极性标签
func EchoPolarity(score float64) string {
	polarity := ""
	if score < 0.2 {
		polarity = "neg"
	} else if score > 0.8 {
		polarity = "pos"
	} else {
		polarity = "net"
	}
	return polarity
}

//BatchPolarityAnalyze 批量情感分析
func BatchPolarityAnalyze(graph service.Graph) service.SinkerGraph {
	SG := service.NewSinkerGraph()
	for k, v := range graph {
		whole := &model.WholeComment{
			CommentId:        v.CommentId,
			MainId:           v.MainId,
			AppID:            v.AppID,
			AppName:          v.AppName,
			Title:            v.Title,
			Content:          v.Content,
			Rating:           v.Rating,
			Version:          v.Version,
			Polarity:         "",
			Score:            0,
			PublishTime:      v.PublishTime,
			PublishTimeStamp: v.PublishTimeStamp,
			CrawlTime:        v.CrawlTime,
			CrawlTimeStamp:   v.CrawlTimeStamp,
		}

		senti, _ := GetSentimentPolarity(v.Content)
		whole.Polarity = senti.Polarity
		whole.Score = senti.Score

		SG[k] = whole
	}

	return SG
}
