package bloom

import (
	"encoding/json"

	"github.com/theapemachine/am/network"
	"github.com/theapemachine/wrkspc/tweaker"
	"github.com/wrk-grp/errnie"
)

type LLM struct {
	req *network.Request
}

func NewLLM() *LLM {
	endpoint := tweaker.GetString("models.bloom.endpoint")
	key := tweaker.GetString("models.bloom.key")
	req := network.NewRequest(network.POST, endpoint)
	req.AddHeader("Authorization", "Bearer "+key)

	return &LLM{req}
}

func (llm *LLM) Predict(input string) string {
	res := Result{}

	errnie.Handles(json.Unmarshal(
		llm.req.Do(NewMsg(input).Marshal()),
		&res,
	))

	return res.GeneratedText
}
