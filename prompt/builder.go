package prompt

type Builder struct {
	History    string
	Input      string
	Scratchpad string
}

func NewBuilder(input string) *Builder {
	return &Builder{Input: input}
}
