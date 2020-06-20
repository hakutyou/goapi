package utils

import "go.uber.org/zap"

var (
	sugar *zap.SugaredLogger
)

func SetLogger(sugarLogger *zap.SugaredLogger) {
	sugar = sugarLogger
}
