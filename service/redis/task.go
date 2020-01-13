package redis

import (
	"gowatcher/go_spider/consts"
	"gowatcher/go_spider/exceptions"
)

//SetCrawlTime 设置爬取时间
func SetCrawlTime(appID string, crawlTime string) error {
	key := consts.RedisCTPrefix + appID
	err := redisClient.Set(key, crawlTime, consts.CrawlTimeExpired).Err()
	if err != nil {
		return exceptions.ErrRedisHandle
	}
	return nil
}

//GetCrawlTime 获取爬取时间
func GetCrawlTime(appID string) (string, error) {
	key := consts.RedisCTPrefix + appID
	check, err := redisClient.Get(key).Result()
	if err != nil {
		return "", exceptions.ErrRedisHandle
	}

	return check, nil
}