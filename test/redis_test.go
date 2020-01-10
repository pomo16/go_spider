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
	appName := "TestApp"
	err := redis.SetCrawlTime(appName, utils.FillLastCrawlTime())
	if err != nil {
		fmt.Println("set_err:", err)
		return
	}

	res, err := redis.GetCrawlTime(appName + "1")
	if err != nil {
		fmt.Println("get_err:", err)
		return
	}
	fmt.Println(res)
}