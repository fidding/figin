package router

import (
	"github.com/gin-gonic/gin"

	middlewar "figin/app/middlewares"
	"figin/app/routers/api"
	"figin/app/routers/web"
)

// Init 路由初始化
func Init(router *gin.Engine) *gin.Engine {
	// 接口
	router.GET("/", api.User)
	router.Use(middlewar.Auth())
	// HTML
	router.GET("/index", web.Index)

	return router
}
