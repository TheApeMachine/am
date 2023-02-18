package openai

import (
	"context"

	gogpt "github.com/sashabaranov/go-gpt3"
	"github.com/theapemachine/wrkspc/tweaker"
	"github.com/wrk-grp/errnie"
)

type Client struct {
	conn *gogpt.Client
}

func NewClient() *Client {
	return &Client{
		conn: gogpt.NewClient(tweaker.GetString("models.openai.key")),
	}
}

func (client *Client) Do(prompt string) string {
	res, err := client.conn.CreateCompletion(
		context.Background(),
		gogpt.CompletionRequest{
			Model:     gogpt.GPT3Ada,
			MaxTokens: 5,
			Prompt:    prompt,
		},
	)

	errnie.Handles(err)
	return res.Choices[0].Text
}
