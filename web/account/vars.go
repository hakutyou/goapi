package account

import (
	"github.com/hakutyou/goapi/web/services"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

var (
	db         *gorm.DB
	sugar      *zap.SugaredLogger
	tencentSms *services.TencentSms
)

func SetDatabase(gormdb *gorm.DB) {
	db = gormdb
}

func SetLogger(sugarLogger *zap.SugaredLogger) {
	sugar = sugarLogger
}

func SetTencentSms(t *services.TencentSms) {
	tencentSms = t
}
