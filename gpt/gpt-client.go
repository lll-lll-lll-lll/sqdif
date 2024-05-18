package gpt

import (
	"context"
	"errors"

	"github.com/sashabaranov/go-openai"
)

type Client struct {
	gptClient *openai.Client
}

func NewClient(apiKey string) (*Client, error) {
	if apiKey == "" {
		return nil, errors.New("apiKey is required")
	}
	client := openai.NewClient(apiKey)
	return &Client{gptClient: client}, nil
}

func (c *Client) Do(ctx context.Context, prompt string) (openai.ChatCompletionResponse, error) {
	return c.gptClient.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT4,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)
}
