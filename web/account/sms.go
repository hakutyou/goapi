package account

import (
	"github.com/gin-gonic/gin"
	"github.com/hakutyou/goapi/web/utils"
	"net/http"
)

// @Summary	获取手机短信验证码
// @Description	获取手机短信验证码
// @Tags 短信
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
// @Description	检验手机短信验证码
// @Tags 短信
// @Accept	mpfd
// @Produce	json
// @Router	/go/sms_captcha	[post]
func smsCaptchaVerify(c *gin.Context) {
	return
}
