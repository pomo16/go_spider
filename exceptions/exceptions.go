package exceptions

import "errors"

var (
	//ErrDBHandle DB构建失败
	ErrDBHandle = errors.New("handle db error")

	//ErrKafkaHandle Kafka构建失败
	ErrKafkaHandle = errors.New("handle kafka error")
	//ErrKafkaSead Kafka发送失败
	ErrKafkaSend = errors.New("send kafka error")

	//ErrConfigRead 配置读取失败
	ErrConfigRead = errors.New("read config error")

	//ErrRedisHandle Redis执行失败
	ErrRedisHandle = errors.New("handle redis error")

	ErrValueEmpty = errors.New("value empty")
)
