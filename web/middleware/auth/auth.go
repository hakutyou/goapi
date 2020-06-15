package auth

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/hakutyou/goapi/web/utils"
	"net/http"
	"strings"
)

// 用户登录验证
func TokenCheckMiddleware(c *gin.Context) {
	var (
		err   error
		user  User
		token UserToken
	)
	reqToken := c.Request.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")

	if len(splitToken) != 2 {
		utils.Response(c, http.StatusUnauthorized, -1, "登录信息无效")
		c.Abort()
		return
	}
	reqToken = splitToken[1]
	token.Token = reqToken

	xclient := Client.DoConnect("Account")
	defer xclient.Close()
	if err = xclient.Call(context.Background(), "ParseToken", token, &user); err != nil {
		utils.Response(c, http.StatusBadRequest, 1, err.Error())
		return
	}
	c.Set("user", user)
	c.Next()
}
