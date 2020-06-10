package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func ServiceRequestJson(requestId string, method string, url string, jsonByte []byte) (statusCode int, retJson map[string]interface{}, err error) {
	var retPage []byte
	statusCode, retPage, err = ServiceRequest(requestId, method, url, jsonByte)
	if err != nil {
		return
	}
	err = json.Unmarshal(retPage, &retJson)
	if err != nil {
		return
	}
	return
}

func ServiceRequest(requestId string, method string, urls string, jsonByte []byte) (statusCode int, retPage []byte, err error) {
	var (
		jsonStr string
		resp    *http.Response
	)

	if len(jsonByte) > 100 {
		jsonStr = "<data>"
	} else {
		jsonStr = string(jsonByte)
	}
	sugar.Infow("发送请求",
		"request_id", requestId,
		"path", urls,
		"method", method,
		"body", jsonStr)

	if method == "post" {
		resp, err = http.Post(urls, "application/json", bytes.NewBuffer(jsonByte))
	} else if method == "post-form" {
		resp, err = http.Post(urls, "application/x-www-form-urlencoded", bytes.NewBuffer(jsonByte))
	} else {
		resp, err = http.Get(urls)
	}
	if err != nil {
		sugar.Errorw("发送请求错误",
			"request_id", requestId,
			"error", err.Error())
		return
	}
	defer resp.Body.Close()
	// 获取返回
	statusCode = resp.StatusCode
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		sugar.Errorw("发送请求-返回错误",
			"request_id", requestId,
			"error", err.Error())
		return
	}
	retPage = buf.Bytes()
	sugar.Infow("发送请求-返回",
		"request_id", requestId,
		"body", string(retPage))
	return
}
