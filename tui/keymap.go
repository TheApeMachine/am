package tui

import "github.com/charmbracelet/bubbles/key"

type Keymap struct {
	NEXT, PREV, ADD, REMOVE, PROMPT, QUIT key.Binding
}

func NewKeymap() *Keymap {
	return &Keymap{
		NEXT: key.NewBinding(
			key.WithKeys("ctrl+l"),
		),
		PREV: key.NewBinding(
			key.WithKeys("ctrl+h"),
		),
		ADD: key.NewBinding(
			key.WithKeys("ctrl+n"),
		),
		REMOVE: key.NewBinding(
			key.WithKeys("ctrl+w"),
		),
		PROMPT: key.NewBinding(
			key.WithKeys("tab"),
		),
		QUIT: key.NewBinding(
			key.WithKeys("esc", "ctrl+c"),
		),
	}
}
