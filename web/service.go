package main

import (
	"github.com/hakutyou/goapi/web/external"
	"github.com/hakutyou/goapi/web/services"
)

var cfgBaiduOcr services.BaiduApi

func openBaiduOcrService() (err error) {
	// API 服务配置
	if err = v.UnmarshalKey("BAIDU_OCR", &cfgBaiduOcr); err != nil {
		return
	}

	external.SetBaiduOcr(cfgBaiduOcr)
	return
}
