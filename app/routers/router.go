package router

import (
	"github.com/gin-gonic/gin"

	"figin/app/routers/api"
	"figin/app/routers/web"
	"figin/app/middlewares"
)

/**
 * 路由初始化
 */
func Init(router *gin.Engine) *gin.Engine {

	// 中间件
	router.Use(middlewar.Logger())
	// 接口
	router.GET("/bus_pv", api.BusPv)
	router.Use(middlewar.Auth())
	// HTML
	router.GET("/index", web.Index)
	
	return router
}