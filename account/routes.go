package account

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup) {
	r.POST("/", createAccount)
	r.POST("/login", loginAccount)
}
