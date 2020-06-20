package account

import (
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/hakutyou/goapi/web/utils"
	"github.com/hakutyou/goapi/web/utils/ucaptcha"
	"net/http"
)

// @Summary	获取验证码
// @Description	获取验证码
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
// @Description	检验验证码
// @Tags 验证码
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
// @Description	查看验证码图片
// @Tags 验证码
// @Router	/go/captcha/:source	[get]
func captchaGetPng(c *gin.Context) {
	ucaptcha.ServeHTTP(c.Writer, c.Request)
}
