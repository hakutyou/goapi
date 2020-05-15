package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/hakutyou/goapi/utils"
	"net/http"
	"strings"
	"time"
)

func TokenCheckMiddleware(c *gin.Context) {
	reqToken := c.Request.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	claims, err := utils.ParseToken(reqToken)
	if err != nil {
		utils.Response(c, http.StatusUnauthorized, -1, "登录信息无效")
		return
	}
	if time.Now().Unix() > claims.ExpiresAt {
		utils.Response(c, http.StatusUnauthorized, -1, "登录信息过期")
	}

	c.Set("user_id", claims.UserID)
	c.Next()
}
