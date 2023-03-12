package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/theapemachine/am/agent"
	"github.com/theapemachine/am/bloom"
	"github.com/theapemachine/am/prompt"
	"github.com/theapemachine/wrkspc/tweaker"
	"github.com/wrk-grp/errnie"
)

type Screen struct {
	width      int
	height     int
	keymap     *Keymap
	inputs     []textarea.Model
	focus      int
	manager    *agent.Manager
	executor   *prompt.Executor
	working    bool
	history    string
	scratchpad string
}

func NewScreen() *Screen {
	errnie.Trace()
	screen := &Screen{
		inputs:  make([]textarea.Model, 3),
		keymap:  NewKeymap(),
		manager: agent.NewManager(bloom.NewLLM()),
		executor: prompt.NewExecutor(strings.Join([]string{
			tweaker.GetString("prompts.bootstrap.prefix"),
			tweaker.GetString("prompts.bootstrap.instructions"),
			tweaker.GetString("prompts.bootstrap.suffix"),
		}, "\n")),
	}

	for i := 0; i < 3; i++ {
		screen.inputs[i] = NewTextArea()
	}

	screen.inputs[screen.focus].Focus()

	return screen
}

func (screen *Screen) Init() tea.Cmd {
	errnie.Trace()
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
			for i := range screen.inputs {
				screen.inputs[i].Blur()
			}

			return screen, tea.Quit
		case key.Matches(msg, screen.keymap.NEXT):
			screen.inputs[screen.focus].Blur()
			screen.focus++

			if screen.focus > len(screen.inputs)-1 {
				screen.focus = 0
			}

			cmd = screen.inputs[screen.focus].Focus()
			cmds = append(cmds, cmd)

			builder := prompt.NewBuilder(
				screen.history,
				screen.inputs[screen.focus].Value(),
				screen.scratchpad,
			)

			screen.history += fmt.Sprintf(
				"\n\nHUMAN: %s",
				screen.inputs[screen.focus].Value(),
			)

			screen.inputs[screen.focus].SetValue(
				screen.executor.Execute(builder),
			)
		}
	case tea.WindowSizeMsg:
		screen.height = msg.Height
		screen.width = msg.Width
	}

	screen.sizeInputs()

	for i := range screen.inputs {
		screen.inputs[i], cmd = screen.inputs[i].Update(msg)
		cmds = append(cmds, cmd)
	}

	return screen, tea.Batch(cmds...)
}

func (screen *Screen) View() string {
	var views []string

	for i := range screen.inputs {
		views = append(views, screen.inputs[i].View())
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, views...)
}

func (screen *Screen) sizeInputs() {
	for i := range screen.inputs {
		screen.inputs[i].SetWidth(screen.width / len(screen.inputs))
		screen.inputs[i].SetHeight(screen.height - 2)
	}
}
