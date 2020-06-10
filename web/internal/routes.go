package internal

import "github.com/gin-gonic/gin"

func Routes(r *gin.RouterGroup) {
	r.POST("/sensitive", sensitiveFilter)

	r.POST("/rpcx", rpcxDemo)
	r.POST("/delay", runAsynq)
}
