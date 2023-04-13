package tool

import (
	"strings"
)

type Agent struct {
	input []string
}

func NewAgent(input string) *Agent {
	return &Agent{strings.Split(input, " ")}
}

func (agent *Agent) Use() string {
	return ""
}
