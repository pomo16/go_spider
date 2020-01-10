package redis

import (
	"gowatcher/go_spider/consts"
	"gowatcher/go_spider/exceptions"
)

//SetCrawlTime 设置爬取时间
func SetCrawlTime(appName string, crawlTime string) error {
	key := consts.RedisCTPrefix + appName
	err := redisClient.Set(key, crawlTime, consts.CrawlTimeExpired).Err()
	if err != nil {
		return exceptions.ErrRedisHandle
	}
	return nil
}

//GetCrawlTime 获取爬取时间
func GetCrawlTime(appName string) (string, error) {
	key := consts.RedisCTPrefix + appName
	check, err := redisClient.Get(key).Result()
	if err != nil {
		return "", exceptions.ErrRedisHandle
	}

	return check, nil
}