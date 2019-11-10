package main

import (
	"gowatcher/go_spider/consts"
	"gowatcher/go_spider/model"
	"gowatcher/go_spider/service"
)

func Init() {
	//database.InitDB()
	//service.InitTaskService()
	//crons.InitCrons()
}

func main() {
	S := service.NewAppleSpiders()
	G := service.NewGraph()

	//任务列表
	//K := service.GlobalTaskLoader.GetTaskMap()
	//service.StartCrawl(S, G, K)
	tasks := make(service.TaskDict)
	t1 := &model.Task{
		AppID:         "1142110895",
		LastCrawlTime: "2019-11-10 00:55:07",
		Status:        consts.Normal,
	}
	t2 := &model.Task{
		AppID:         "414478124",
		LastCrawlTime: "2019-11-10 00:55:07",
		Status:        consts.Normal,
	}
	tasks[t1.AppID] = t1
	tasks[t2.AppID] = t2
	service.StartCrawl(S, G, tasks)

	//crons.CronJobs()
}
