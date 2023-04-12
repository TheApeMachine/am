package capital

import (
	"github.com/theapemachine/wrkspc/tweaker"
	"github.com/wrk-grp/please"
)

type Client struct {
	conn *please.Request
}

func NewClient() *Client {
	return &Client{
		please.NewRequest(
			tweaker.GetString("capital.demo"),
		).AddHeaders(map[string]string{
			"Content-Type":  "application/json",
			"X-CAP-API-KEY": tweaker.GetString("capital.token"),
		}),
	}
}
