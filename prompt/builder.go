package prompt

import "github.com/wrk-grp/errnie"

type Builder struct {
	History    string
	Input      string
	Scratchpad string
}

func NewBuilder(history, input, scratchpad string) *Builder {
	errnie.Trace()
	return &Builder{history, input, scratchpad}
}
