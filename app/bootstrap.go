package bootstrap

import (
	router "figin/app/routers"
	"fmt"

	"github.com/gin-gonic/gin"

	"figin/app/providers/logger"
	"figin/config"

	"figin/app/providers/cache"
	"figin/app/providers/database"
	migration "figin/database/migrations"
	"figin/system"
)

// Init 初始化
func Init() {
	// 配置初始化
	config.Setup()
	// 数据库初始化
	database.Setup()
	// 数据迁移
	migration.Setup()
	// 日志
	logger.Setup()
	// 缓存
	cache.Setup()
	// 将组件挂载到sys
	system.Setup()

	conf := system.Conf
	// 线上环境
	if conf.Server.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.Default()
	// 设置静态资源访问目录
	r.Static("/assets", "./assets")
	// 设置WEB模板目录
	r.LoadHTMLGlob("app/views/*")
	// 初始化路由
	r = router.Init(r)

	// 启动
	r.Run(fmt.Sprintf(":%d", conf.Server.Port))
}
