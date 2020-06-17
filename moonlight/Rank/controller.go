package Rank

import "context"

type Args struct{}
type Reply struct{}

func (Rank) GetAccountRank(_ context.Context,
	args Args, reply *Reply) (err error) {
	return
}
