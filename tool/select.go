package tool

import "github.com/theapemachine/am/agent"

type Element struct {
	Type string
	Text string
}

type Select struct{}

func NewSelect() *Select {
	return &Select{}
}

func (sel *Select) Pick(elements []Element) string {
	var t agent.Tool

	for _, element := range elements {
		switch element.Type {
		case "[SHELL]":
			t = NewShell(element.Text)
		case "[SEARCH]":
			t = NewSearch(element.Text)
		default:
			t = NewInquire(element.Text)
		}
	}

	return t.Use()
}
