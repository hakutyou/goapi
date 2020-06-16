package DFA

import "context"

type Args struct {
	Sensitive string
}

type Reply struct {
	Filter string
}

func (DFA) DFAFilter(ctx context.Context, args *Args, reply *Reply) error {
	reply.Filter = changeSensitiveWords(args.Sensitive)
	return nil
}
