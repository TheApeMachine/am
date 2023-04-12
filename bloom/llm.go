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
	errnie.Trace()

	endpoint := tweaker.GetString("models.bloom.endpoint")
	key := tweaker.GetString("models.bloom.key")
	req := network.NewRequest(network.POST, endpoint)
	req.AddHeader("Authorization", "Bearer "+key)

	return &LLM{req}
}

func (llm *LLM) Predict(input string) string {
	errnie.Trace()

	res := []Result{}
	msg := llm.req.Do(NewMsg(input).Marshal())

	errnie.Handles(json.Unmarshal(
		msg,
		&res,
	))

	return res[0].GeneratedText
}
