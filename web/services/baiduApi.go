package services

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/hakutyou/goapi/web/utils"

	"github.com/garyburd/redigo/redis"
)

const baiduApiBaseUrl string = "https://aip.baidubce.com"

type BaiduApi struct {
	ApiKey    string `yaml:"ApiKey"`
	SecretKey string `yaml:"SecretKey"`
}

// 获取 accessToken
func (api BaiduApi) getAccessToken(requestId string) (accessToken string, err error) {
	var (
		retJson map[string]interface{}
	)
	// 查询 Redis 缓存
	accessToken, err = redis.String(conn.Do("GET", "BAIDU_OCR_ACCESS_TOKEN"))
	if err == nil {
		return
	}

	_, retJson, err = utils.ServiceRequestJson(requestId, "post",
		fmt.Sprintf(
			"%s/oauth/2.0/token?grant_type=%s&client_id=%s&client_secret=%s",
			baiduApiBaseUrl, "client_credentials", api.ApiKey, api.SecretKey), nil)
	if err != nil {
		accessToken = ""
		sugar.Error(err.Error())
		return
	}

	accessToken = retJson["access_token"].(string)
	expiresIn := retJson["expires_in"].(float64)
	// 存 Redis
	if _, err := conn.Do("SET", "BAIDU_OCR_ACCESS_TOKEN", accessToken, "EX", int(expiresIn)); err != nil {
		sugar.Warnw("Redis 连接错误",
			"message", err.Error())
	}
	return
}

// 清除缓存的 accessToken
func (api BaiduApi) clearAccessToken(err error) {
	_, err = conn.Do("DEL", "BAIDU_OCR_ACCESS_TOKEN")
	return
}

// 身份证识别
func (api BaiduApi) IdCardRecognition(requestId string, image string, idCardSide string) (retJson map[string]interface{}, err error) {
	var (
		accessToken string
		formByte    []byte
	)
	accessToken, err = api.getAccessToken(requestId)
	if err != nil {
		return
	}

	formByte = []byte(url.Values{
		"image":        []string{image},
		"id_card_side": []string{idCardSide},
	}.Encode())
	_, retJson, err = utils.ServiceRequestJson(
		requestId,
		"post-form",
		fmt.Sprintf(
			"%s/rest/2.0/ocr/v1/idcard?access_token=%s",
			baiduApiBaseUrl, accessToken),
		formByte)
	if err != nil {
		return
	}

	errorMsg := retJson["error_msg"]
	if errorMsg != nil {
		err = errors.New(errorMsg.(string))
		return
	}
	return
}
