package demo

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup) {
	r.GET("/cache", getCache)
	r.POST("/cache", setCache)
	r.POST("/delay", runAsynq)

	r.POST("/id_card", idCardRecognize)
}
