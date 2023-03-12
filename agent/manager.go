package agent

import "github.com/wrk-grp/errnie"

type Manager struct {
	LLMs []LLM
}

func NewManager(llms ...LLM) *Manager {
	errnie.Trace()

	return &Manager{
		LLMs: llms,
	}
}

func (manager *Manager) Predict(input string) string {
	errnie.Trace()

	for _, llm := range manager.LLMs {
		input = llm.Predict(input)
	}

	return input
}
