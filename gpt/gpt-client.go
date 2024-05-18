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

func (c *Client) Do(ctx context.Context, prompt string) (openai.CompletionResponse, error) {
	return c.gptClient.CreateCompletion(
		ctx,
		openai.CompletionRequest{
			Model: openai.GPT4o, Prompt: prompt,
		},
	)
}
