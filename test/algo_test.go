package test

import (
	"fmt"
	"gowatcher/go_spider/algoml"
	"gowatcher/go_spider/service/text_analyze"
	"testing"
)

func TestAlgo(t *testing.T) {
	algoml.InitAlgoModel()
	res, _ := text_analyze.GetSentimentPolarity("很不错")
	fmt.Println(res.Polarity)
	fmt.Println(res.Score)
}
