package external

import "github.com/gin-gonic/gin"

func Routes(r *gin.RouterGroup) {
	r.POST("/proxy", doProxy)
}
