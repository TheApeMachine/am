package bloom

import (
	"encoding/json"

	"github.com/wrk-grp/errnie"
)

type Parameters struct {
	DoSample       bool `json:"do_sample"`
	EarlyStopping  bool `json:"early_stopping"`
	LengthPenalty  int  `json:"length_penalty"`
	MaxNewTokens   int  `json:"max_new_tokens"`
	MaxTime        int  `json:"max_time"`
	Seed           int  `json:"seed"`
	ReturnFullText bool `json:"return_full_text"`
}

type Msg struct {
	Inputs     string     `json:"inputs"`
	Parameters Parameters `json:"parameters"`
}

func NewMsg(input string) *Msg {
	errnie.Trace()

	return &Msg{
		Inputs: input,
		Parameters: Parameters{
			DoSample:       false,
			EarlyStopping:  true,
			LengthPenalty:  0,
			MaxNewTokens:   25,
			MaxTime:        120,
			ReturnFullText: false,
		},
	}
}

func (msg *Msg) Marshal() []byte {
	errnie.Trace()

	buf, err := json.Marshal(msg)
	errnie.Handles(err)
	return buf
}

type Result struct {
	GeneratedText string `json:"generated_text"`
}
