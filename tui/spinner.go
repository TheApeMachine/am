package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Loader struct {
	spinner spinner.Model
}

func NewLoader() *Loader {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return &Loader{s}
}

func (loader *Loader) Init() tea.Cmd {
	return loader.spinner.Tick
}

func (loader *Loader) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	loader.spinner, cmd = loader.spinner.Update(msg)
	return loader, cmd
}

func (loader *Loader) View() string {
	return fmt.Sprintf(
		"\n\n   %s Working...\n\n",
		loader.spinner.View(),
	)
}
