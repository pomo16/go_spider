package database

import (
	"github.com/sirupsen/logrus"
	"gowatcher/go_spider/exceptions"
	"gowatcher/go_spider/model"
	"gowatcher/go_spider/utils"
)

//QueryTasks 获取爬虫任务列表
func QueryTasks(lastTime string) (*model.TaskTable, error) {
	rows, err := dbReader.Table("gowatcher.crawl_task_table").Debug().
		Select("id, app_name, app_id, status, create_time, modify_time").
		Where("modify_time > ?", lastTime).Order("modify_time").Rows()

	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("Recovered in QueryTasks: %v", r)
		}

		//官方文档示例未考虑指针为空调用Close会panic的情形
		if rows != nil {
			rows.Close()
		}
	}()

	if err != nil {
		logrus.Errorf("QueryTasks error, err: %v", err.Error())
		return nil, exceptions.ErrDBHandle
	}

	tTable := &model.TaskTable{}
	for rows.Next() {
		var tmpFT model.TaskRow
		dbReader.ScanRows(rows, &tmpFT)
		tTable.Rows = append(tTable.Rows, &tmpFT)
		tTable.LastTime = utils.ConvertGoTimeToStd(tmpFT.ModifyTime)
	}
	return tTable, nil
}
