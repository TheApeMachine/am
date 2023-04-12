package capital

import (
	"encoding/json"

	"github.com/theapemachine/wrkspc/tweaker"
	"github.com/wrk-grp/errnie"
	"github.com/wrk-grp/please"
)

type AccountInfo struct {
	Balance    float64 `json:"balance"`
	Deposit    float64 `json:"deposit"`
	ProfitLoss float64 `json:"profitLoss"`
	Available  float64 `json:"available"`
}

type SessionPresenter struct {
	Account AccountInfo `json:"accountInfo"`
	Stream  string      `json:"streamingHost"`
	CST     string
	Token   string
}

type Session struct {
	client *please.Request
	token  string
	cst    string
}

func NewSession() *Session {
	return &Session{client: NewClient().conn}
}

func (session *Session) Open() SessionPresenter {
	res := session.client.Post("/api/v1/session", map[string]interface{}{
		"identifier":        tweaker.GetString("capital.login"),
		"password":          tweaker.GetString("capital.password"),
		"encryptedPassword": false,
	})

	data := SessionPresenter{}
	errnie.Handles(json.Unmarshal(res, &data))
	data.CST = session.client.GetHeader("CST")
	data.Token = session.client.GetHeader("X-SECURITY-TOKEN")

	return data
}
