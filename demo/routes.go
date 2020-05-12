package demo

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup) {
	r.GET("/ping/", doRequest)
}

func doRequest(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
