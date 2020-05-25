package utils

import "github.com/asmcos/requests"

func ServiceRequest(requestId string, url string, data requests.Datas) (retJson map[string]interface{}, err error) {
	var resp *requests.Response

	sugar.Infow("发送请求",
		"request_id", requestId,
		"path", url,
		"method", "POST",
		"body", data)

	if data == nil {
		resp, err = requests.Post(url)
	} else {
		resp, err = requests.Post(url, data)
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
