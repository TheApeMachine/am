package cmd

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/theapemachine/am/tui"
	"github.com/wrk-grp/errnie"
)

/*
runCmd is a proxy for running any terminal command using a container
which is dynamically built from an image in a configured registry.
*/
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the system with the ~/.am.yml config values.",
	Long:  runtxt,
	RunE: func(_ *cobra.Command, _ []string) (err error) {
		if _, err := tea.NewProgram(
			tui.NewScreen(), tea.WithAltScreen(),
		).Run(); err != nil {
			errnie.Handles(err)
			os.Exit(1)
		}
		return nil
	},
}

/*
runtxt lives here to keep the command definition section cleaner.
*/
var runtxt = `
Run the Agent System.
`
