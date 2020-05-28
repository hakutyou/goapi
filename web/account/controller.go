package account

import (
	"net/http"

	"github.com/hakutyou/goapi/web/utils"

	"github.com/gin-gonic/gin"
)

// @Summary	创建用户
// @Description	创建用户
// @Tags 用户
// @Accept	mpfd
// @Produce	json
// @Param	name		formData	string	true	"用户名"
// @Param	password	formData	string	true	"密码"
// @success	200	{object}	utils.ResponseDataResult	"code 为 0 表示成功"
// @success	400	{object}	utils.ResponseResult		"message 返回错误信息"
// @Router	/go/account/	[post]
func createAccount(c *gin.Context) {
	user := User{
		Status: true,
		Salt:   generateSalt(32),
	}

	if err := c.ShouldBind(&user); err != nil {
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
	utils.ResponseWithData(c, http.StatusOK, 0, "操作成功", user)
	return
}

// @Summary	用户登录
// @Description	用户登录
// @Tags 用户
// @Accept	mpfd
// @Produce	json
// @Param	name		formData	string	true	"用户名"
// @Param	password	formData	string	true	"密码"
// @success	200	{object}	utils.ResponseDataResult	"code 为 0 表示成功"
// @success	400	{object}	utils.ResponseResult		"message 返回错误信息"
// @Router	/go/account/login	[post]
func loginAccount(c *gin.Context) {
	var (
		user  User
		err   error
		token string
		ret   bool
	)

	if err := c.ShouldBind(&user); err != nil {
		utils.Response(c, http.StatusBadRequest, 1, "参数格式错误")
		return
	}

	user, ret = user.login()
	if ret == false {
		utils.Response(c, http.StatusUnauthorized, -1, "用户名或密码错误")
		return
	}

	// 生成 jwt token
	if token, err = user.generateToken(); err != nil {
		utils.Response(c, http.StatusBadRequest, 1, "服务器繁忙")
		return
	}
	utils.ResponseWithData(c, http.StatusOK, 0, "操作成功", gin.H{
		"token": token,
	})
	return
}

// @Summary	查看用户信息
// @Description	查看用户信息
// @Tags 用户
// @Security ApiKeyAuth
// @Accept	mpfd
// @Produce	json
// @success	200	{object}	utils.ResponseDataResult	"code 为 0 表示成功"
// @success	400	{object}	utils.ResponseResult		"message 返回错误信息"
// @Router	/go/account/userinfo/	[get]
func getUserinfo(c *gin.Context) {
	var (
		user User
		ret  bool
	)
	userID := c.MustGet("user_id").(uint)
	user, ret = getUserInfo(userID)
	if ret == false {
		utils.Response(c, http.StatusUnauthorized, -1, "登录信息无效")
		return
	}
	utils.ResponseWithData(c, http.StatusOK, 0, "操作成功", user)
	return
}
