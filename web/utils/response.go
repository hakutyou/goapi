package utils

import "github.com/gin-gonic/gin"

type ResponseResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ResponseDataResult struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Response(c *gin.Context, code int, errcode int, message string) {
	c.JSON(code, ResponseResult{errcode, message})
}

func ResponseWithData(c *gin.Context, code int, errcode int, message string, data interface{}) {
	c.JSON(code, ResponseDataResult{errcode, message, data})
}
