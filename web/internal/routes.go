package internal

import "github.com/gin-gonic/gin"

func Routes(r *gin.RouterGroup) {
	r.POST("/sendmail", sendMail)
	r.POST("/sensitive", sensitiveFilter)

	r.GET("/moonlight/bang", moonlightBang)
	//	r.POST("/delay", runAsynq)
}
