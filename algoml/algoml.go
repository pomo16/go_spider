package algoml

import (
	"gowatcher/go_spider/algoml/sentiment"
	"gowatcher/go_spider/algoml/tokenizer"
)

//模型句柄
var (
	SentiML   *sentiment.Sentiment
	Tokenizer tokenizer.Tokenizer
)

//InitAlgoModel 初始化算法模型
func InitAlgoModel() {
	Tokenizer = tokenizer.NewJiebaTokenizer(true)
	SentiML = sentiment.NewSentiment(Tokenizer)
}
