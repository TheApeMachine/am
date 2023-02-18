package agent

type Manager struct {
	LLMs []LLM
}

func NewManager(llms ...LLM) *Manager {
	return &Manager{
		LLMs: llms,
	}
}

func (manager *Manager) Predict(input string) string {
	for _, llm := range manager.LLMs {
		input = llm.Predict(input)
	}

	return input
}
