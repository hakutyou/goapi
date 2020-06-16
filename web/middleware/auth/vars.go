package auth

import (
	"github.com/hakutyou/goapi/web/utils"
)

var (
	Client utils.RpcxClient
)

func SetClient(remote string, port int) {
	Client = utils.RpcxClient{
		Remote: remote, // "localhost",
		Port:   port,   // 8971,
	}
	return
}
