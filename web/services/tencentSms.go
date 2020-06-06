package services

import (
	"errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20190711"
)

type TencentSms struct {
	SecretID   string `yaml:"SecretID"`
	SecretKey  string `yaml:"SecretKey"`
	Region     string `yaml:"Region"`
	AppId      string `yaml:"AppID"`
	Sign       string `yaml:"Sign"`
	TemplateId string `yaml:"TemplateId"`
}

func (tencentSms TencentSms) initClient() (err error) {
	credential := common.NewCredential(
		tencentSms.SecretID,
		tencentSms.SecretKey,
	)
	cpf := profile.NewClientProfile()
	tencentSmsClient, err = sms.NewClient(credential, tencentSms.Region, cpf)
	return
}

func (tencentSms TencentSms) SendSms(phone string) (err error) {
	var (
		response *sms.SendSmsResponse
	)

	if tencentSmsClient == nil {
		err = tencentSms.initClient()
		if err != nil {
			return
		}
	}

	request := sms.NewSendSmsRequest()
	request.SmsSdkAppid = common.StringPtr(tencentSms.AppId)
	request.Sign = common.StringPtr(tencentSms.Sign)
	// 国际/港澳台短信 senderid, 国内短信空
	// request.SenderId = common.StringPtr("")
	// 用户的 session 内容
	// request.SessionContext = common.StringPtr("xxx")
	// 短信码号扩展号, 默认未开通
	// request.ExtendCode = common.StringPtr("0")
	request.TemplateID = common.StringPtr(tencentSms.TemplateId)
	// TODO: 模板参数
	code := "123456"
	timeout := "3"
	request.TemplateParamSet = common.StringPtrs([]string{code, timeout})
	// phone 格式形如 "+8613711112222"
	request.PhoneNumberSet = common.StringPtrs([]string{phone})

	response, err = tencentSmsClient.SendSms(request)
	if err != nil {
		return
	}
	// TODO: 返回信息并显示发送失败
	// print(*response.Response.SendStatusSet[0].Message)
	return errors.New(*response.Response.SendStatusSet[0].Message)
}
