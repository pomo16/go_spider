package model

type Config struct {
	Mysql         DBConfig `yaml:"mysql"`
	Kafka         KFConfig `yaml:"kafka"`
	Redis         RDConfig `yaml:"redis"`
	ElasticSearch ESConfig `yaml:"elasticsearch"`
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

type RDConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
}

type ESConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}
