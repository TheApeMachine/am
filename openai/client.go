package openai

import (
	"context"
	"io"

	gogpt "github.com/sashabaranov/go-openai"
	"github.com/theapemachine/am/audio"
	"github.com/theapemachine/wrkspc/tweaker"
	"github.com/wrk-grp/errnie"
)

type Client struct {
	conn       *gogpt.Client
	microphone *audio.Microphone
}

func NewClient() *Client {
	return &Client{
		conn:       gogpt.NewClient(tweaker.GetString("models.openai.key")),
		microphone: audio.NewMicrophone(),
	}
}

func (client *Client) SpeechInput() string {
	client.microphone.Record()

	req := gogpt.AudioRequest{
		Model:    gogpt.Whisper1,
		FilePath: "out.wav",
	}

	resp, err := client.conn.CreateTranscription(context.Background(), req)
	errnie.Handles(err)

	return resp.Text
}

func (client *Client) Predict(input []map[string]string) chan string {
	out := make(chan string)

	go func() {
		defer close(out)

		msgs := make([]gogpt.ChatCompletionMessage, 0)

		for _, in := range input {
			for key, val := range in {
				msgs = append(msgs, gogpt.ChatCompletionMessage{
					Role: key, Content: val,
				})
			}
		}

		stream, err := client.conn.CreateChatCompletionStream(
			context.Background(),
			gogpt.ChatCompletionRequest{
				Model:       gogpt.GPT3Dot5Turbo,
				Temperature: 0,
				Messages:    msgs,
				Stream:      true,
			},
		)

		errnie.Handles(err)

		defer stream.Close()

		for {
			resp, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if errnie.Handles(err) != nil {
				break
			}

			out <- resp.Choices[0].Delta.Content
		}
	}()

	return out
}
