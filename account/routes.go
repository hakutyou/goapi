package account

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Routes(r *gin.RouterGroup) {
	r.POST("/", createPeople)
}
