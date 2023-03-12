package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/theapemachine/am/agent"
	"github.com/theapemachine/am/bloom"
	"github.com/theapemachine/am/prompt"
	"github.com/theapemachine/wrkspc/tweaker"
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
		errnie.Trace()

		var (
			history    string
			scratchpad string
		)

		manager := agent.NewManager(bloom.NewLLM())

		executor := prompt.NewExecutor(strings.Join([]string{
			tweaker.GetString("prompts.bootstrap.prefix"),
			tweaker.GetString("prompts.bootstrap.instructions"),
			tweaker.GetString("prompts.bootstrap.suffix"),
		}, "\n"))

		var input string

		for {
			fmt.Print("> ")
			reader := bufio.NewReader(os.Stdin)
			if input, err = reader.ReadString('\n'); errnie.Handles(err) != nil {
				continue
			}

			// remove the delimeter from the string
			history += strings.TrimSuffix(input, "\n")

			builder := prompt.NewBuilder(
				history,
				input,
				scratchpad,
			)

			output := manager.Predict(executor.Execute(builder))
			history += output + "\n"
			fmt.Println(output)
		}
		// if _, err := tea.NewProgram(
		// 	tui.NewScreen(), tea.WithAltScreen(),
		// ).Run(); err != nil {
		// 	errnie.Handles(err)
		// 	os.Exit(1)
		// }
		return nil
	},
}

/*
runtxt lives here to keep the command definition section cleaner.
*/
var runtxt = `
Run the Agent System.
`
