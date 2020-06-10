package DFA

import "context"

type DFA int

type Args struct {
	Sensitive string
}

type Reply struct {
	Filter string
}

func (t *DFA) DFAFilter(ctx context.Context, args *Args, reply *Reply) error {
	reply.Filter = ChangeSensitiveWords(args.Sensitive)
	return nil
}
