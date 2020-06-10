package external

import "github.com/gin-gonic/gin"

func Routes(r *gin.RouterGroup) {
	// controller.go
	r.POST("/proxy", doProxy)

	// baiduApi.go
	r.POST("/id_card", idCardRecognize)
}
