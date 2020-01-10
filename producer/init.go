package producer

import (
	"github.com/Shopify/sarama"
	"gopkg.in/yaml.v2"
	"gowatcher/go_spider/consts"
	"gowatcher/go_spider/exceptions"
	"gowatcher/go_spider/model"
	"log"
	"os"
	"path/filepath"
)

var kafkaSender sarama.SyncProducer

func InitProducer() {
	sender, err := newKafkaSender()
	if err != nil {
		panic(err)
	}
	kafkaSender = sender
}

func newKafkaSender() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	// 等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 随机的分区类型：返回一个分区器，该分区器每次选择一个随机分区
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	// 是否等待成功和失败后的响应
	config.Producer.Return.Successes = true

	kfLink, _ := ReadYamlConfig()
	sender, err := sarama.NewSyncProducer([]string{kfLink}, config)
	if err != nil {
		log.Fatalf("init kafka sender failed,err:%v\n", err)
		return nil, exceptions.ErrKafkaHandle
	}

	return sender, nil
}

//ReadYamlConfig 读取yaml配置文件返回kafka链接
func ReadYamlConfig() (string, error) {
	path, _ := filepath.Abs(consts.ConfFile)
	conf := &model.Config{}
	if f, err := os.Open(path); err != nil {
		return "", exceptions.ErrConfigRead
	} else {
		yaml.NewDecoder(f).Decode(conf)
	}

	kfConfig := conf.Kafka
	link := kfConfig.Host + ":" + kfConfig.Port
	return link, nil
}
