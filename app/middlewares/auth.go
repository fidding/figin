package middlewar

import (
	"github.com/gin-gonic/gin"
)

// Auth 认证
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		println("用户认证开始")
		c.Next()
		println("用户认证结束")
	}
}
