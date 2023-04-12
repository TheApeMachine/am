package tool

import (
	"os/exec"
	"strings"
)

type Shell struct {
	input []string
}

func NewShell(input string) *Shell {
	return &Shell{strings.Split(input, " ")}
}

func (shell *Shell) Use() string {
	cmd := exec.Command(shell.input[0], shell.input[1:]...)

	output, err := cmd.Output()
	if err != nil {
		return err.Error()
	}

	return string(output)
}
