package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"figin/system"
)

func BusPv(c *gin.Context) {

	fmt.Println(system.Db("sql"))

	c.JSON(200, gin.H{
		"message": "BusPv:123",
	})
}