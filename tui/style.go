package tui

import "github.com/charmbracelet/lipgloss"

var (
	cursorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))

	cursorlineStyle = lipgloss.NewStyle().Background(
		lipgloss.Color("57"),
	).Foreground(
		lipgloss.Color("230"),
	)

	placeholderstyle = lipgloss.NewStyle().Foreground(
		lipgloss.Color("238"),
	)

	endOfBufferStyle = lipgloss.NewStyle().Foreground(
		lipgloss.Color("235"),
	)

	focusedPlaceholderStyle = lipgloss.NewStyle().Foreground(
		lipgloss.Color("99"),
	)

	focusedBorderStyle = lipgloss.NewStyle().Border(
		lipgloss.RoundedBorder(),
	).BorderForeground(
		lipgloss.Color("238"),
	)

	blurredBorderStyle = lipgloss.NewStyle().Border(
		lipgloss.HiddenBorder(),
	)
)
