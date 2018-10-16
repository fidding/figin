package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"figin/system"
)

func BusPv(c *gin.Context) {

	fmt.Println(system.DB)

	c.JSON(200, gin.H{
		"message": "BusPv:123",
	})
}