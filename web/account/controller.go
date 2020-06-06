package account

import (
	"net/http"

	"github.com/hakutyou/goapi/web/utils"
	"github.com/hakutyou/goapi/web/utils/ucaptcha"

	"github.com/dchest/captcha"
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

// @Summary	获取验证码
// @Tags 用户
// @Produce	json
// @Router	/go/captcha	[get]
func captchaGet(c *gin.Context) {
	captchaId := captcha.New()
	if captchaId == "" {
		utils.Response(c, http.StatusBadRequest, 1, "服务器繁忙")
		return
	}
	utils.ResponseWithData(c, http.StatusOK, 0, "操作成功", gin.H{
		"path": captchaId + ".png",
	})
	return
}

// @Summary	检验验证码
// @Tags 用户
// @Accept	mpfd
// @Produce	json
// @Router	/go/captcha	[post]
func captchaVerify(c *gin.Context) {
	var captchaVerifyRequest = struct {
		CaptchaId string `binding:"required" form:"captcha_id" json:"captcha_id"`
		Value     string `binding:"required" form:"value" json:"value"`
	}{}
	// 获取参数
	if err := c.ShouldBind(&captchaVerifyRequest); err != nil {
		utils.Response(c, http.StatusBadRequest, 1, "参数格式错误")
		return
	}

	if captcha.VerifyString(captchaVerifyRequest.CaptchaId, captchaVerifyRequest.Value) {
		utils.Response(c, http.StatusOK, 0, "验证成功")
	} else {
		utils.Response(c, http.StatusBadRequest, 1, "验证失败")
	}
}

// @Summary	查看验证码图片
// @Tags 用户
// @Router	/go/captcha/:source	[get]
func captchaGetPng(c *gin.Context) {
	ucaptcha.ServeHTTP(c.Writer, c.Request)
}

// @Summary	获取手机短信验证码
// @Tags 用户
// @Produce	json
// @Router	/go/sms_captcha	[get]
func smsCaptchaGet(c *gin.Context) {
	var smsCaptchaGetRequest = struct {
		Phone string `binding:"required" form:"phone" json:"phone"`
	}{}
	// 获取参数
	if err := c.ShouldBind(&smsCaptchaGetRequest); err != nil {
		utils.Response(c, http.StatusBadRequest, 1, "参数格式错误")
		return
	}

	// 发送短信
	if err := tencentSms.SendSms(smsCaptchaGetRequest.Phone); err != nil {
		utils.Response(c, http.StatusBadRequest, 1, "发送失败")
		return
	}
	utils.Response(c, http.StatusOK, 0, "发送成功")
	return
}

// @Summary	检验手机短信验证码
// @Tags 用户
// @Accept	mpfd
// @Produce	json
// @Router	/go/sms_captcha	[post]
func smsCaptchaVerify(c *gin.Context) {
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
