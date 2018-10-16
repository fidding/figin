package system

import (
	"figin/app/providers/database"
	"figin/app/providers/log"

	"figin/config"
	"database/sql"
)

// 集成配置
func Config() *config.Configuration {
	return config.Config()
}

// 集成数据库服务
func Db(dbType string) *sql.DB {
	return database.Db(dbType)
}

// 集成日志服务
func Log(msg string) {
	log.Logger(msg)
}