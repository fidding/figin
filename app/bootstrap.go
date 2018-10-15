package bootstrap

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"figin/app/routers"
	_ "log"
	"os"
	// 初始化服务提供者
	"figin/app/providers/database"

	"figin/system"
	log "github.com/sirupsen/logrus"
)

func Init() {
	// 加载配置信息
	var configFile string = "config.yaml"
	if err := system.LoadConfiguration(configFile); err != nil {
		fmt.Println("load configuration err:", err)
		return
	}
	config := system.Config()

	// 初始化提供者
	dbConfig := &database.DbConfig{ SqlConnect: config.SqlConnect }
	database.DbInit(dbConfig)
	// 日志
	log.SetFormatter(&log.JSONFormatter{})
	filename := "figin.log"
	f, err := os.OpenFile(filename, os.O_WRONLY | os.O_CREATE | os.O_APPEND, 0755)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(f)
	log.SetLevel(log.WarnLevel)
	log.WithFields(log.Fields{
		"animal": "walrus",
	  }).Warning("A walrus appears")

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