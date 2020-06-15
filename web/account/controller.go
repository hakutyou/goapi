package account

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/hakutyou/goapi/web/middleware/auth"
	"github.com/hakutyou/goapi/web/utils"
	"net/http"
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
	var (
		user auth.User
		err  error
	)
	if err = c.ShouldBind(&user); err != nil {
		utils.Response(c, http.StatusBadRequest, 1, "参数格式错误")
		return
	}
	xclient := auth.Client.DoConnect("Account")
	defer xclient.Close()
	err = xclient.Call(context.Background(), "CreateAccount", user, &user)
	if err != nil {
		sugar.Errorw("RPCx服务调用错误",
			"error", err.Error())
		utils.Response(c, http.StatusBadRequest, 1, err.Error())
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
		user  auth.User
		err   error
		reply auth.UserToken
	)

	if err = c.ShouldBind(&user); err != nil {
		utils.Response(c, http.StatusBadRequest, 1, "参数格式错误")
		return
	}
	xclient := auth.Client.DoConnect("Account")
	defer xclient.Close()
	if err = xclient.Call(context.Background(), "LoginAccount", user, &reply); err != nil {
		sugar.Errorw("RPCx服务调用错误",
			"error", err.Error())
		utils.Response(c, http.StatusBadRequest, 1, err.Error())
		return
	}
	utils.ResponseWithData(c, http.StatusOK, 0, "操作成功", reply)
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
func userinfoAccount(c *gin.Context) {
	user := c.MustGet("user").(auth.User)
	utils.ResponseWithData(c, http.StatusOK, 0, "操作成功", user)
	return
}
