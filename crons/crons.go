package crons

import (
	"github.com/robfig/cron"
	"gowatcher/go_spider/consts"
	"gowatcher/go_spider/service"
)

//InitCrons 初始化定时更新任务
func InitCrons() {
	CronJobs()
}

//CronJobs 定时任务
func CronJobs() {
	StartSpiders()
	c := cron.New()
	spec := consts.Timing
	err := c.AddFunc(spec, StartSpiders)
	if err != nil {
		panic(err)
	}
	c.Start()

	select {}
}

//StartSpiders 定时爬虫任务
func StartSpiders() {
	service.GlobalTaskLoader.Load()
	S := service.NewAppleSpiders()
	G := service.NewGraph()
	//任务列表
	K := service.GlobalTaskLoader.GetTaskMap()
	service.StartCrawl(S, G, K)
}
