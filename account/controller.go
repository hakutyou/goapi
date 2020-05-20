package account

import (
	"net/http"

	"github.com/hakutyou/goapi/utils"

	"github.com/gin-gonic/gin"
)

func createAccount(c *gin.Context) {
	user := User{
		Status: true,
		Salt:   generateSalt(32),
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		utils.Response(c, http.StatusBadRequest, 1, "参数格式错误")
		return
	}

	// 密码加密
	if err := user.Password.doHash(user.Salt); err != nil {
		utils.Response(c, http.StatusBadRequest, 1, "服务器繁忙")
		return
	}

	if dbc := db.Create(&user); dbc.Error != nil {
		driverErr := dbc.Error.Error()
		utils.Response(c, http.StatusBadRequest, 100, driverErr)
		return
	}
	c.JSON(http.StatusOK, user)
	return
}

func loginAccount(c *gin.Context) {
	var (
		user  User
		err   error
		token string
		ret   bool
	)

	if err := c.ShouldBindJSON(&user); err != nil {
		utils.Response(c, http.StatusBadRequest, 1, "参数格式错误")
		return
	}

	user, ret = user.Login()
	if ret == false {
		utils.Response(c, http.StatusUnauthorized, -1, "用户名或密码错误")
		return
	}

	// 生成 jwt token
	if token, err = user.GenerateToken(); err != nil {
		utils.Response(c, http.StatusBadRequest, 1, "服务器繁忙")
		return
	}
	utils.ResponseWithData(c, http.StatusOK, 0, "操作成功", gin.H{
		"token": token,
	})
	return
}

func getUserinfo(c *gin.Context) {
	var (
		user User
		ret  bool
	)
	userID := c.MustGet("user_id").(uint)
	user, ret = GetUserInfo(userID)
	if ret == false {
		utils.Response(c, http.StatusUnauthorized, -1, "登录信息无效")
		return
	}
	utils.ResponseWithData(c, http.StatusOK, 0, "操作成功", user)
	return
}
