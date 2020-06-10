package external

import (
	"github.com/hakutyou/goapi/web/services"
	"go.uber.org/zap"
)

var (
	sugar    *zap.SugaredLogger
	baiduOcr services.BaiduApi
)

func SetLogger(sugarLogger *zap.SugaredLogger) {
	sugar = sugarLogger
}

func SetBaiduOcr(b services.BaiduApi) {
	baiduOcr = b
}
