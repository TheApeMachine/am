package prompt

import (
	"strings"
)

type Builder struct {
	buf strings.Builder
}

func NewBuilder() *Builder {
	return &Builder{}
}
