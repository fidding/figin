package api

import (
	"github.com/gin-gonic/gin"
)

// User 路由
func User(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, I am fidding",
	})
}
