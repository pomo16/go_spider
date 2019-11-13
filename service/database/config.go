package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/yaml.v2"
	"gowatcher/go_spider/exceptions"
	"gowatcher/go_spider/model"
	"gowatcher/go_spider/utils"
	"log"
	"os"
	"path/filepath"
)

var (
	dbReader *gorm.DB
)

func InitDB() {
	dbLink, err := ReadYamlConfig()
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open("mysql", dbLink)
	if err == nil {
		db.SingularTable(true)
		dbReader = db
	} else {
		panic(err)
	}
}

//ReadYamlConfig 读取yaml配置文件返回数据库链接
func ReadYamlConfig() (string, error) {
	path, _ := filepath.Abs("../config/config.yaml")
	conf := &model.Config{}
	if f, err := os.Open(path); err != nil {
		return "", exceptions.ErrFileRead
	} else {
		yaml.NewDecoder(f).Decode(conf)
	}

	dbConfig := conf.Mysql
	link := dbConfig.UserName + ":" + dbConfig.Password +
		"@tcp(" + dbConfig.Host + ":" + dbConfig.Port +
		")/gowatcher?charset=utf8&parseTime=True&loc=Local"
	return link, nil
}

//QueryTasks 获取爬虫任务列表
func QueryTasks(lastTime string) (*model.TaskTable, error) {
	rows, err := dbReader.Table("gowatcher.crawl_task_table").Debug().
		Select("id, app_name, app_id, status, create_time, modify_time").
		Where("modify_time > ?", lastTime).Order("modify_time").Rows()

	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("Recovered in QueryTasks: %v\n", r)
		}

		//官方文档示例未考虑指针为空调用Close会panic的情形
		if rows != nil {
			rows.Close()
		}
	}()

	if err != nil {
		log.Printf("QueryTasks error, err: %v\n", err.Error())
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
