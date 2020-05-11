package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("./go/ping/", doRequest)
	r.Run(":8080")
}

func doRequest(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
