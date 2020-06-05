package account

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

var (
	db    *gorm.DB
	sugar *zap.SugaredLogger
)

func SetDatabase(gormdb *gorm.DB) {
	db = gormdb
}

func SetLogger(sugarLogger *zap.SugaredLogger) {
	sugar = sugarLogger
}
