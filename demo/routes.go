package demo

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup) {
	r.GET("/cache", GetCache)
	r.POST("/cache", SetCache)
}
