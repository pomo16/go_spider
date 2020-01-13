package consts

import "time"

//Url 爬虫URL
const (
	CommentUrlPrefix = "https://itunes.apple.com/WebObjects/MZStore.woa/wa/userReviewsRow?cc=cn&displayable-kind=11&sort=0&appVersion=all"
	VersionUrlPrefix = "https://itunes.apple.com/rss/customerreviews/"
	VersionUrlSuffix = "/sortby=mostrecent/json?l=en&&cc=cn"
)

//TimeStr 时间格式化模板
const TimeStr = "2006-01-02 15:04:05"

//Status 数据状态
const (
	Normal = 1 //启用中
	Unused = 2 //未启用
)

//PageSize 爬虫偏移量
const PageSize = 100

//Timing 定时时机
const Timing = "0 0 */6 * * *"
const TestTiming = "*/1 * * * * *"

//Topic Kafka topic
const Topic = "test"

//ConfFile 配置文件路径
const ConfFile = "config/config.yaml"

//log文件相关
const (
	LogFilePath    = "go_spider_log"
	LogFilePrefix  = "spider"
	LogFileTimeStr = "20060102150405"
	LogFileSuffix  = ".log"
)

//Redis
const (
	RedisCTPrefix = "spider:ct:"
	CrawlTimeExpired = 48 * time.Hour
)