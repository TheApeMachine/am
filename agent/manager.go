package agent

import (
	gogpt "github.com/sashabaranov/go-openai"
	"github.com/theapemachine/am/prompt"
	"github.com/theapemachine/wrkspc/tweaker"
	"github.com/wrk-grp/errnie"
)

type Manager struct {
	LLMs       []LLM
	builder    *prompt.Builder
	scratchpad []map[string]string
	history    []map[string]string
}

func NewManager(llms ...LLM) *Manager {
	errnie.Trace()

	return &Manager{
		LLMs: llms,
		builder: prompt.NewBuilder(
			tweaker.GetString("ai.prompt.preset.system"),
			tweaker.GetString("ai.prompt.preset.current"),
		),
		scratchpad: make([]map[string]string, 0),
		history:    make([]map[string]string, 0),
	}
}

func (manager *Manager) System(system string) {
	sys := make(map[string]string)
	sys[gogpt.ChatMessageRoleSystem] = system

	manager.scratchpad = make([]map[string]string, 0)
	manager.scratchpad = append(manager.scratchpad, sys)
}

/*
Predict the response of all the Large Language Models that have been
loaded into the Manager.

This allows for multiple AI "personas" to be joined into the session.
Given some form of shared memory object, the response of one can be
used as (part of) the prompt for the next one(s).
*/
func (manager *Manager) Predict(input map[string]string) chan string {
	errnie.Trace()
	out := make(chan string)

	manager.scratchpad = append(manager.scratchpad, input)

	go func() {
		defer close(out)

		for _, llm := range manager.LLMs {
			for chunk := range llm.Predict(manager.scratchpad) {
				out <- chunk
			}
		}
	}()

	return out
}
