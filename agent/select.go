package agent

import "github.com/wrk-grp/errnie"

/*
Select implements the Tool interface to allow one large language model
to determine which large larguage model should perform the current
task.

It also manages any dynamic promting that may be useful to the selected
model.
*/
type Select struct {
	manager *LLM
}

func NewSelect(manager *LLM) *Select {
	errnie.Trace()

	return &Select{
		manager: manager,
	}
}
