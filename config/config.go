package config

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/spf13/viper"
)

var conf = &Configuration{
	Server: server{},
	DB:     database{},
	Cache:  cache{},
	Logger: logger{},
}

// GetConf 获取配置信息
func GetConf() *Configuration {
	return conf
}

// Setup 配置初始化
func Setup() {
	viper.SetConfigType("YAML")
	// 读取配置文件
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("Read 'config.yaml' fail: %v\n", err)
	}
	// 解析配置文件
	viper.ReadConfig(bytes.NewBuffer(data))
	viper.UnmarshalKey("server", &conf.Server)
	viper.UnmarshalKey("database", &conf.DB)
	viper.UnmarshalKey("cache", &conf.Cache)
	viper.UnmarshalKey("logger", &conf.Logger)
	fmt.Println(conf)
}
