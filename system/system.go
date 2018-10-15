package system

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"figin/app/providers/database"
	"figin/app/providers/log"

	"database/sql"
)

// 配置结构
type Configuration struct {
	AppName string `yaml:"AppName"`
	Env string `yaml:"Env"`
	Port string `yaml:"Port"`
	SqlConnect string `yaml:"SqlConnect"`
}

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

// 集成数据库服务
func Db(dbType string) *sql.DB {
	return database.Db(dbType)
}

// 集成日志服务
func Log(msg string) {
	log.Logger(msg)
}

func init() {

}