package middlewar

import (
	"log"
	"time"
	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc{
	return func(c *gin.Context) {
		println("开始记录")
		t := time.Now()

		if 1 == 1 {
			c.Next()
		} else {
			c.Abort()
		}
		
		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
		println("结束记录")
	}
}
