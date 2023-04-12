package capital

import "github.com/wrk-grp/please"

type Ping struct {
	client  *please.Request
	session *SessionPresenter
}

func NewPing(session *SessionPresenter) *Ping {
	return &Ping{NewClient().conn, session}
}

func (ping *Ping) Do() {
	ping.client.AddHeaders(map[string]string{
		"CST":              ping.session.CST,
		"X-SECURITY-TOKEN": ping.session.Token,
	})

	ping.client.Get("/api/v1/ping", map[string]interface{}{})
}
