package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	gogpt "github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
	"github.com/theapemachine/am/agent"
	"github.com/theapemachine/am/openai"
	"github.com/theapemachine/am/tool"
	"github.com/theapemachine/wrkspc/tweaker"
	"github.com/wrk-grp/errnie"
)

type Element struct {
	Type string
	Text string
}

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
		client := openai.NewClient()
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

			fmt.Println()
			chunks := " "

			in := make(map[string]string)
			in[gogpt.ChatMessageRoleUser] = input

			for chunk := range manager.Predict(in) {
				chunks += chunk
				fmt.Print(chunk)
			}

			fmt.Println()
			fmt.Println("ENTITIES:")
			elements := parseElements(chunks)

			var t agent.Tool

			for _, element := range elements {
				switch element.Type {
				case "[SHELL]":
					t = tool.NewShell(element.Text)
				case "[SEARCH]":
					t = tool.NewSearch(element.Text)
				}

				input = t.Use()
			}

			fmt.Println(input)

			// cmd := exec.Command("say", chunks)
			// errnie.Handles(cmd.Run())

			fmt.Println()
		}

		return nil
	},
}

func parseElements(input string) []Element {
	lines := strings.Split(input, "\n")
	var elements []Element

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if isAction(line) {
			chunks := strings.Split(line, " ")

			elements = append(elements, Element{
				Type: chunks[0],
				Text: strings.TrimSpace(strings.Join(chunks[1:], " ")),
			})
		}
	}

	return elements
}

func isAction(line string) bool {
	errnie.Debugs("isAction <-", line)

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
