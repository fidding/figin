package config

// 配置结构
type Configuration struct {
	AppName string `yaml:"AppName"`
	Env string `yaml:"Env"`
	Port string `yaml:"Port"`
	SqlConnect string `yaml:"SqlConnect"`
}