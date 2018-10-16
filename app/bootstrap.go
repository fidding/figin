package bootstrap

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"figin/app/routers"
	_ "log"
	// 初始化服务提供者

	"figin/system"
	"figin/config"
	"figin/app/providers/log"

	"figin/database/migrations"
)

func Init() {
	// 加载配置信息
	var configFile string = "config.yaml"
	if err := config.LoadConfiguration(configFile); err != nil {
		fmt.Println("load configuration err:", err)
		return
	}
	config := system.Config()

	// 初始化提供者
	system.DbInit()
	
	// 数据迁移
	migration.Migrate()

	// 日志
	log.LogInit()

	// 线上环境
	if config.Env == "prod" {
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
	r.Run(":" + config.Port)
}

func main() {
	
}