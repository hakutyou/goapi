package account

import (
	"github.com/hakutyou/goapi/web/services"
	"go.uber.org/zap"
)

var (
	sugar      *zap.SugaredLogger
	tencentSms *services.TencentSms
)

func SetLogger(sugarLogger *zap.SugaredLogger) {
	sugar = sugarLogger
}

func SetTencentSms(t *services.TencentSms) {
	tencentSms = t
}
