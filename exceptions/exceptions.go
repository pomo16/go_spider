package exceptions

import "errors"

var (
	//ErrDBHandle DB构建失败
	ErrDBHandle = errors.New("handle db error")
	//ErrKafkaHandle Kafka构建失败
	ErrKafkaHandle = errors.New("handle kafka error")
	//ErrFileRead 文件读取失败
	ErrFileRead = errors.New("read file error")
)
