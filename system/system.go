package system

import (
	"figin/config"

	"figin/app/providers/cache"
	"figin/app/providers/database"
	"figin/app/providers/logger"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

var (
	// Conf 配置挂载
	Conf  *config.Configuration
	// DB 数据库挂载
	DB  *gorm.DB
	// Logger 日志挂载
	Logger  *logrus.Logger
	// Cache 缓存挂载
	Cache  *redis.Client
)

// Setup 初始化
func Setup() {
	Conf = config.GetConf()
	DB = database.GetDB()
	Logger = logger.GetLogger()
	Cache = cache.GetCache()
}
