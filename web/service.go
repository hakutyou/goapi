package main

import (
	"github.com/hakutyou/goapi/web/external"
	"github.com/hakutyou/goapi/web/services"
)

var cfgBaiduOcr services.BaiduApi

func openBaiduOcrService() {
	// API 服务配置
	if err := v.UnmarshalKey("BAIDU_OCR", &cfgBaiduOcr); err != nil {
		panic(err)
	}

	external.SetBaiduOcr(cfgBaiduOcr)
}
