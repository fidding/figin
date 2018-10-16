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
	db := database.DataObject{}
	return db.Db(dbType)
}

// 集成数据库服务
var DB *database.DataObject
func DbInit() {
	db := database.DataObject{}
	DB = db.Load(config.Config())
}

// 集成日志服务
func Log(msg string) {
	log.Logger(msg)
}