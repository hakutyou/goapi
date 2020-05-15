package utils

import "github.com/gin-gonic/gin"

func Response(c *gin.Context, code int, errcode int, message string) {
	c.JSON(code, gin.H{
		"code":    errcode,
		"message": message,
	})
}

func ResponseWithData(c *gin.Context, code int, errcode int, message string, data interface{}) {
	c.JSON(code, gin.H{
		"code":    errcode,
		"message": message,
		"data":    data,
	})
}
