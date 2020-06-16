package internal

import "github.com/gin-gonic/gin"

func Routes(r *gin.RouterGroup) {
	r.POST("/sendmail", sendMail)
	r.POST("/sensitive", sensitiveFilter)
	r.GET("/excel", excelDemo)

	//	r.POST("/delay", runAsynq)
}
