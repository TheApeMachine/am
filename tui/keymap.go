package tui

import "github.com/charmbracelet/bubbles/key"

type Keymap struct {
	NEXT, QUIT key.Binding
}

func NewKeymap() *Keymap {
	return &Keymap{
		NEXT: key.NewBinding(
			key.WithKeys("tab"),
		),
		QUIT: key.NewBinding(
			key.WithKeys("esc", "ctrl+c"),
		),
	}
}
