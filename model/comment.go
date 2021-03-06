package model

type CommentSpider struct {
	CommentId   string `json:"userReviewId"`
	Title       string `json:"title"`
	Content     string `json:"body"`
	Rating      string `json:"rating"`
	PublishTime string `json:"date"`
}

type Comment struct {
	CommentId        string `json:"comment_id"`
	MainId           string `json:"main_id"`
	AppID            string `json:"app_id"`
	AppName          string `json:"app_name"`
	Title            string `json:"title"`
	Content          string `json:"content"`
	Rating           string `json:"rating"`
	Version          string `json:"version"`
	PublishTime      string `json:"publish_time"`
	PublishTimeStamp int64  `json:"publish_timestamp"`
	CrawlTime        string `json:"crawl_time"`
	CrawlTimeStamp   int64  `json:"crawl_timestamp"`
}

type WholeComment struct {
	CommentId        string  `json:"comment_id"`
	MainId           string  `json:"main_id"`
	AppID            string  `json:"app_id"`
	AppName          string  `json:"app_name"`
	Title            string  `json:"title"`
	Content          string  `json:"content"`
	Rating           string  `json:"rating"`
	Version          string  `json:"version"`
	Polarity         string  `json:"polarity"`
	Score            float64 `json:"score"`
	PublishTime      string  `json:"publish_time"`
	PublishTimeStamp int64   `json:"publish_timestamp"`
	CrawlTime        string  `json:"crawl_time"`
	CrawlTimeStamp   int64   `json:"crawl_timestamp"`
}
