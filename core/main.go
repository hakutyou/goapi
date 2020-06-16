package main

import (
	"fmt"
	"github.com/hakutyou/goapi/core/DFA"
	"github.com/hakutyou/goapi/core/Excel"
	"github.com/hakutyou/goapi/core/Mail"
	"github.com/smallnest/rpcx/server"
)

func init() {
	var err error

	if err = LoadConfigure(); err != nil {
		panic(err)
	}
}

func main() {
	var (
		err error
		s   *server.Server
	)

	s = server.NewServer()
	if err = s.RegisterName("Mail", new(Mail.Mail), ""); err != nil {
		panic(err)
	}
	if err = s.RegisterName("DFA", new(DFA.DFA), ""); err != nil {
		panic(err)
	}
	if err = s.RegisterName("Excel", new(Excel.Excel), ""); err != nil {
		panic(err)
	}
	if err = s.Serve("tcp",
		fmt.Sprintf("localhost:%d", cfg.Port)); err != nil {
		panic(err)
	}
}
