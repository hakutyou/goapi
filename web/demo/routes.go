package demo

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup) {
	r.GET("/cache", getCache)
	r.POST("/cache", setCache)
	r.POST("/rpcx", rpcxDemo)
	r.POST("/delay", runAsynq)

	// 内部服务
	r.POST("/sensitive", sensitiveFilter)
	// 外部服务
	r.POST("/id_card", idCardRecognize)
}
