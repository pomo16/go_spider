package main

import "lhx-github/go_spider/consts"

func main() {
	S := NewAppleCommentSpider()
	G := NewCommentGraph()
	K := consts.GetAppMap()["抖音"]
	StartCrawl(S, G, K)
}
