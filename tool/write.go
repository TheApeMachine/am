package tool

import (
	"strings"
)

type Write struct {
	input []string
}

func NewWrite(input string) *Write {
	return &Write{strings.Split(input, " ")}
}

func (write *Write) Use() string {
	return ""
}
