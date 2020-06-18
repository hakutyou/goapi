package internal

import (
	"github.com/hakutyou/goapi/web/utils"
	"github.com/hibiken/asynq"
	"go.uber.org/zap"
)

var (
	Client          utils.RpcxClient
	MoonlightClient utils.RpcxClient
	aclient         *asynq.Client
	sugar           *zap.SugaredLogger
)

func SetClient(remote string, port int) {
	Client = utils.RpcxClient{
		Remote: remote,
		Port:   port,
	}
	return
}

func SetMoonlightClient(remote string, port int) {
	MoonlightClient = utils.RpcxClient{
		Remote: remote,
		Port:   port,
	}
	return
}

func SetAsynq(c *asynq.Client) {
	aclient = c
}

func SetLogger(sugarLogger *zap.SugaredLogger) {
	sugar = sugarLogger
}
