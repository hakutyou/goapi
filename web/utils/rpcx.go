package utils

import (
	"fmt"
	rpcx "github.com/smallnest/rpcx/client"
)

type RpcxClient struct {
	Remote string
	Port   int
}

func (c RpcxClient) DoConnect(servicePath string) rpcx.XClient {
	rpcxService := rpcx.NewPeer2PeerDiscovery(fmt.Sprintf("tcp@%s:%d", c.Remote, c.Port), "")
	return rpcx.NewXClient(servicePath, rpcx.Failtry, rpcx.RandomSelect,
		rpcxService, rpcx.DefaultOption)
}
