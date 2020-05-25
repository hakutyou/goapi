package services

import (
	"errors"
	"fmt"
	"github.com/asmcos/requests"

	"github.com/hakutyou/goapi/utils"
)

var baiduApiBaseUrl = "https://aip.baidubce.com"

type BaiduApi struct {
	ApiKey    string
	SecretKey string
}

func (api BaiduApi) getAccessToken(requestId string) (accessToken string, err error) {
	var (
		retJson map[string]interface{}
	)
	// TODO: 使用 Redis 缓存

	retJson, err = utils.ServiceRequest(requestId,
		fmt.Sprintf(
			"%s/oauth/2.0/token?grant_type=%s&client_id=%s&client_secret=%s",
			baiduApiBaseUrl, "client_credentials", api.ApiKey, api.SecretKey), nil)

	if err != nil {
		accessToken = ""
		return
	}

	accessToken = retJson["access_token"].(string)
	return
}

func (api BaiduApi) IdCardRecognition(requestId string, image string, idCardSide string) (retJson map[string]interface{}, err error) {
	var accessToken string

	accessToken, err = api.getAccessToken(requestId)
	if err != nil {
		return
	}

	retJson, err = utils.ServiceRequest(requestId,
		fmt.Sprintf(
			"%s/rest/2.0/ocr/v1/idcard?access_token=%s",
			baiduApiBaseUrl, accessToken),
		requests.Datas{
			"image":        image,
			"id_card_side": idCardSide,
		})
	if err != nil {
		return
	}

	errorMsg := retJson["error_msg"]
	if errorMsg != nil {
		err = errors.New(errorMsg.(string))
		return
	}

	err = nil
	return
}
