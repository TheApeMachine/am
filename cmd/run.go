package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	gogpt "github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
	"github.com/theapemachine/am/agent"
	"github.com/theapemachine/am/bloom"
	"github.com/theapemachine/am/tool"
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

		// Connect to LLM interfaces so we can send and receive messages.
		// client := openai.NewClient()
		client := bloom.NewLLM()

		// Pass the clients to the agent manager so they can be sequenced.
		manager := agent.NewManager(client)
		manager.System(tweaker.GetString("prompt.preset.system"))

		var input string

		fmt.Print("> ")
		reader := bufio.NewReader(os.Stdin)

		if input, err = reader.ReadString('\n'); errnie.Handles(err) != nil {
			return err
		}

		for {
			if input == "quit" {
				break
			}

			chunks := " "

			in := make(map[string]string)
			in[gogpt.ChatMessageRoleUser] = input

			for chunk := range manager.Predict(in) {
				chunks += chunk
				fmt.Print(chunk)
			}

			fmt.Println()
			input = tool.NewSelect().Pick(parseElements(chunks))
			fmt.Println(input)

			// cmd := exec.Command("say", chunks)
			// errnie.Handles(cmd.Run())
		}

		return nil
	},
}

func parseElements(input string) []tool.Element {
	lines := strings.Split(input, "\n")
	var elements []tool.Element

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if isAction(line) {
			chunks := strings.Split(line, " ")

			elements = append(elements, tool.Element{
				Type: chunks[0],
				Text: strings.TrimSpace(strings.Join(chunks[1:], " ")),
			})
		}
	}

	return elements
}

func isAction(line string) bool {
	if strings.HasPrefix(line, "[THINK]") ||
		strings.HasPrefix(line, "[REASON]") ||
		strings.HasPrefix(line, "[INQUIRE]") ||
		strings.HasPrefix(line, "[RESPOND]") ||
		strings.HasPrefix(line, "[SHELL]") ||
		strings.HasPrefix(line, "[AGENT]") ||
		strings.HasPrefix(line, "[WRITE]") ||
		strings.HasPrefix(line, "[SEARCH]") {
		return true
	}

	return false
}

/*
runtxt lives here to keep the command definition section cleaner.
*/
var runtxt = `
Run the Agent System.
`
