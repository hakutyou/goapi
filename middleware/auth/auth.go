package auth

import (
	"net/http"
	"strings"
	"time"

	"github.com/hakutyou/goapi/utils"

	"github.com/gin-gonic/gin"
)

func TokenCheckMiddleware(c *gin.Context) {
	reqToken := c.Request.Header.Get("Authorization")

	splitToken := strings.Split(reqToken, "Bearer ")

	if len(splitToken) != 2 {
		utils.Response(c, http.StatusUnauthorized, -1, "登录信息无效")
		c.Abort()
		return
	}
	reqToken = splitToken[1]

	claims, err := utils.ParseToken(reqToken)
	if err != nil {
		utils.Response(c, http.StatusUnauthorized, -1, "登录信息无效")
		c.Abort()
		return
	}
	if time.Now().Unix() > claims.ExpiresAt {
		utils.Response(c, http.StatusUnauthorized, -1, "登录信息过期")
		c.Abort()
		return
	}

	c.Set("user_id", claims.UserID)
	c.Next()
}
