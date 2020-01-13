package test

import (
	"fmt"
	"gowatcher/go_spider/service/redis"
	"gowatcher/go_spider/utils"
	"testing"
)

func TestRedisPing(t *testing.T) {
	redis.InitRedis()
	redis.PingRedis()
}

func TestRedis(t *testing.T) {
	redis.InitRedis()
	appID := "TestApp"
	err := redis.SetCrawlTime(appID, utils.FillLastCrawlTime())
	if err != nil {
		fmt.Println("set_err:", err)
		return
	}

	res, err := redis.GetCrawlTime(appID)
	if err != nil {
		fmt.Println("get_err:", err)
		return
	}
	fmt.Println(res)
}