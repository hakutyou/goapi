package utils

import "go.uber.org/zap"

var (
	jwtSecret []byte
	sugar     *zap.SugaredLogger
)

func SetLogger(sugarLogger *zap.SugaredLogger) {
	sugar = sugarLogger
}
