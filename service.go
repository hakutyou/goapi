package main

import (
	"github.com/hakutyou/goapi/demo"
	"github.com/hakutyou/goapi/services"
)

var cfgBaiduOcr services.BaiduApi

func openBaiduOcrService() {
	// API 服务配置
	if err := v.UnmarshalKey("BAIDU_OCR", &cfgBaiduOcr); err != nil {
		panic(err)
	}

	demo.SetBaiduOcr(cfgBaiduOcr)
}
