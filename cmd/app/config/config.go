package config

type Config struct {
	Default DefaultConfig `yaml:"default"`
	Mongo   MongoConfig   `yaml:"mongo"`
}

type DefaultConfig struct {
	// debug mode
	Mode   string `yaml:"mode"`
	Listen int    `yaml:"listen"`

	// 自动创建表结构
	AutoMigrate bool `yaml:"autoMigrate"`
}

type MongoConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}
