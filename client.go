package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/sashabaranov/go-openai"
)

const defaultPrompt = `
Please create SQL that generates 100 rows of dummy data based solely on table information. 
Do not generate anything other than SQL.`

type SQLDiffGenerator struct {
	gptClient *openai.Client
}

func newGPTGenerator(apiKey string) (*SQLDiffGenerator, error) {
	if apiKey == "" {
		return nil, errors.New("apiKey is required")
	}
	c := openai.NewClient(apiKey)
	return &SQLDiffGenerator{
		gptClient: c,
	}, nil
}

func (g *SQLDiffGenerator) Do(sqlFile, outputFile, prompt string, override bool) error {
	if prompt == "" {
		prompt = defaultPrompt
	}

	if !strings.HasSuffix(sqlFile, ".sql") {
		return fmt.Errorf("invalid SQL file extension")
	}

	if _, err := os.Stat(sqlFile); os.IsNotExist(err) {
		return fmt.Errorf("SQL file does not exist: %s", sqlFile)
	}

	if _, err := os.Stat(outputFile); os.IsNotExist(err) {
		return fmt.Errorf("output file does not exist: %s", outputFile)
	}

	sqlContent, err := os.ReadFile(sqlFile)
	if err != nil {
		return err
	}

	resp, err := g.doToGPT4(context.Background(), prompt+"\n"+string(sqlContent))
	if err != nil {
		return err
	}

	if override {
		outputFile, err := os.Create(outputFile)
		if err != nil {
			return err
		}
		defer outputFile.Close()

		if _, err := outputFile.WriteString(resp.Choices[0].Message.Content); err != nil {
			return err
		}

		fmt.Println("Output file updated")
	} else {
		fmt.Println(resp.Choices[0].Message.Content)
	}

	return nil
}

func (c *SQLDiffGenerator) doToGPT4(ctx context.Context, prompt string) (openai.ChatCompletionResponse, error) {
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
