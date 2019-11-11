package model

type Config struct {
	Mysql DBConfig `yaml:"mysql"`
	Kafka KFConfig `yaml:"kafka"`
}

type DBConfig struct {
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
}

type KFConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}
