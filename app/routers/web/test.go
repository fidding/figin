package web

import (
	_ "fmt"
	_ "figin/app/providers/database"

	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Main website",
	})
}