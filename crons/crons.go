package crons

import (
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
	"gowatcher/go_spider/consts"
	"gowatcher/go_spider/utils"
)

//InitCrons 初始化定时更新任务
func InitCrons() {
	CronJobs()
}

//CronJobs 定时任务
func CronJobs() {
	//直接先运行一次任务
	StartSpiders()

	//定时循环
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
	utils.InitLogger()
	defer utils.CloseLogFile()
	//service.GlobalTaskLoader.Load()
	//S := service.NewAppleSpiders()
	//G := service.NewGraph()
	//T := service.GlobalTaskLoader.GetTaskMap()
	//service.StartCrawl(S, G, T)
	logrus.Error("test")

	//Banned.This method will kill the program.
	//logrus.Exit(0)
}
