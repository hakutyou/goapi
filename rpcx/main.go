package main

import (
	"flag"

	"github.com/hakutyou/goapi/rpcx/DFA"
	"github.com/hakutyou/goapi/rpcx/Excel"

	"github.com/smallnest/rpcx/server"
)

var addr = flag.String("addr", "localhost:8972", "server address")

func main() {
	flag.Parse()

	s := server.NewServer()
	_ = s.RegisterName("DFA", new(DFA.DFA), "")
	_ = s.RegisterName("Excel", new(Excel.Excel), "")
	// _ = s.Register(new(Arith), "")
	_ = s.Serve("tcp", *addr)
}
