package main

import (
	"gowatcher/go_spider/crons"
	"gowatcher/go_spider/service"
	"gowatcher/go_spider/service/database"
)

func Init() {
	database.InitDB()
	service.InitTaskService()
	crons.InitCrons()
}

func main() {
	Init()
}
