package main

import (
	"gowatcher/go_spider/crons"
	"gowatcher/go_spider/service"
	"gowatcher/go_spider/service/database"
	"gowatcher/go_spider/service/redis"
)

func Init() {
	database.InitDB()
	redis.InitRedis()
	service.InitTaskService()
	crons.InitCrons()
}

func main() {
	Init()
}
