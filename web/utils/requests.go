package utils

import (
	"bytes"
	"github.com/asmcos/requests"
	"net/http"
)

func ServiceRequest(requestId string, method string, url string, data requests.Datas) (retJson map[string]interface{}, err error) {
	var resp *requests.Response

	sugar.Infow("发送请求",
		"request_id", requestId,
		"path", url,
		"method", "POST",
		"body", data)

	if method == "post" {
		resp, err = requests.Post(url, data)
	} else {
		resp, err = requests.Get(url, data)
	}

	if err != nil {
		return
	}

	err = resp.Json(&retJson)

	sugar.Infow("返回发送请求",
		"request_id", requestId,
		"body", resp.Text())
	return
}

func ServiceProxy(_ string, method string, url string, jsonStr string) (statusCode int, retPage string, err error) {
	var (
		resp *http.Response
		req  *http.Request
	)

	if method == "post" {
		req, err = http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonStr)))
		if err != nil {
			return
		}
		req.Header.Set("Content-Type", "application/json")
	} else {
		req, err = http.NewRequest("GET", url, bytes.NewBuffer([]byte(jsonStr)))
		if err != nil {
			return
		}
	}

	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	// 获取返回
	statusCode = resp.StatusCode
	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(resp.Body)
	retPage = buf.String()
	return
}
