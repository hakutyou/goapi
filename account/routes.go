package account

import (
	"github.com/hakutyou/goapi/middleware/auth"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup) {
	r.POST("/", createAccount)
	r.POST("/login", loginAccount)

	// 需要登录的接口
	r_userinfo := r.Group("/userinfo")
	r_userinfo.Use(auth.TokenCheckMiddleware)
	r_userinfo.GET("/", getUserinfo)
}
