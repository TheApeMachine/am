package tool

import (
	"strings"
)

type Respond struct {
	input []string
}

func NewRespond(input string) *Respond {
	return &Respond{strings.Split(input, " ")}
}

func (respond *Respond) Use() string {
	return ""
}
