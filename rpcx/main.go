package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/smallnest/rpcx/server"
)

var addr = flag.String("addr", "localhost:8972", "server address")

type Arith int

type Args struct {
	A int
	B int
}

type Reply struct {
	C int
}

func (t *Arith) Mul(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A * args.B
	fmt.Printf("call: %d * %d = %d\n", args.A, args.B, reply.C)
	return nil
}

func (t *Arith) Add(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A + args.B
	fmt.Printf("call: %d + %d = %d\n", args.A, args.B, reply.C)
	return nil
}

func (t *Arith) Say(ctx context.Context, args *string, reply *string) error {
	*reply = "hello " + *args
	return nil
}

func main() {
	flag.Parse()

	s := server.NewServer()
	// s.RegisterName("Arith", new(example.Arith), "")
	_ = s.Register(new(Arith), "")
	_ = s.Serve("tcp", *addr)
}
