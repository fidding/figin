package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Index 首页
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Hello Figin",
	})
}
