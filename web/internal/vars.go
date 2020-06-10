package internal

import (
	"github.com/hibiken/asynq"
	rpcx "github.com/smallnest/rpcx/client"
	"go.uber.org/zap"
)

var (
	rpcxService rpcx.ServiceDiscovery
	client      *asynq.Client
	sugar       *zap.SugaredLogger
)

func SetAsynq(c *asynq.Client) {
	client = c
}

func SetLogger(sugarLogger *zap.SugaredLogger) {
	sugar = sugarLogger
}
