package config

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

var configuration *Configuration

// 加载配置文件
func LoadConfiguration(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	var config Configuration
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return err
	}
	configuration = &config
	return err
}

// 读取配置信息
func Config() *Configuration {
	return configuration
}
