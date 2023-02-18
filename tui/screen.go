package tui

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/theapemachine/am/agent"
	"github.com/theapemachine/am/bloom"
	"github.com/theapemachine/am/prompt"
	"github.com/theapemachine/wrkspc/tweaker"
)

type Screen struct {
	width    int
	height   int
	keymap   *Keymap
	outputs  []textarea.Model
	input    textarea.Model
	focus    int
	manager  *agent.Manager
	executor *prompt.Executor
}

func NewScreen() *Screen {
	screen := &Screen{
		outputs: make([]textarea.Model, 1),
		input:   textarea.New(),
		keymap:  NewKeymap(),
		manager: agent.NewManager(bloom.NewLLM()),
		executor: prompt.NewExecutor(strings.Join([]string{
			tweaker.GetString("prompts.bootstrap.prefix"),
			tweaker.GetString("prompts.bootstrap.instructions"),
			tweaker.GetString("prompts.bootstrap.suffix"),
		}, "\n")),
	}

	for i := 0; i < 1; i++ {
		screen.outputs[i] = NewTextArea()
	}

	screen.outputs[screen.focus].Focus()
	screen.updateKeyBindings()

	return screen
}

func (screen *Screen) Init() tea.Cmd {
	return textarea.Blink
}

func (screen *Screen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmds []tea.Cmd
		cmd  tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, screen.keymap.QUIT):
			for i := range screen.outputs {
				screen.outputs[i].Blur()
			}

			return screen, tea.Quit
		case key.Matches(msg, screen.keymap.NEXT):
			screen.outputs[screen.focus].Blur()
			screen.focus++

			if screen.focus > len(screen.outputs)-1 {
				screen.focus = 0
			}

			cmd = screen.outputs[screen.focus].Focus()
			cmds = append(cmds, cmd)

			builder := prompt.NewBuilder(
				screen.manager.Predict(screen.input.Value()),
			)

			screen.outputs[screen.focus].SetValue(
				screen.executor.Execute(builder),
			)
		case key.Matches(msg, screen.keymap.PREV):
			screen.outputs[screen.focus].Blur()
			screen.focus--

			if screen.focus < 0 {
				screen.focus = len(screen.outputs) - 1
			}

			cmd = screen.outputs[screen.focus].Focus()
			cmds = append(cmds, cmd)
		case key.Matches(msg, screen.keymap.ADD):
			screen.outputs = append(screen.outputs, NewTextArea())
		case key.Matches(msg, screen.keymap.REMOVE):
			screen.outputs = screen.outputs[:len(screen.outputs)-1]

			if screen.focus > len(screen.outputs)-1 {
				screen.focus = len(screen.outputs) - 1
			}
		case key.Matches(msg, screen.keymap.PROMPT):
			screen.outputs[screen.focus].Blur()
			cmd = screen.input.Focus()
			cmds = append(cmds, cmd)
		}
	case tea.WindowSizeMsg:
		screen.height = msg.Height
		screen.width = msg.Width
	}

	screen.updateKeyBindings()
	screen.sizeOutputs()

	for i := range screen.outputs {
		screen.outputs[i], cmd = screen.outputs[i].Update(msg)
		cmds = append(cmds, cmd)
		screen.input, cmd = screen.input.Update(msg)
		cmds = append(cmds, cmd)
	}

	return screen, tea.Batch(cmds...)
}

func (screen *Screen) View() string {
	var views []string

	for i := range screen.outputs {
		views = append(views, screen.outputs[i].View())
	}

	return lipgloss.JoinVertical(lipgloss.Top,
		lipgloss.JoinHorizontal(lipgloss.Top, views...),
		screen.input.View(),
	)
}

func (screen *Screen) sizeOutputs() {
	for i := range screen.outputs {
		screen.outputs[i].SetWidth(screen.width / len(screen.outputs))
		screen.outputs[i].SetHeight(screen.height - 8)
	}
}

func (screen *Screen) updateKeyBindings() {
	screen.keymap.ADD.SetEnabled(len(screen.outputs) < 2)
	screen.keymap.REMOVE.SetEnabled(len(screen.outputs) > 1)
}
