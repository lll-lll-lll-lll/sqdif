package gpt

import (
	"context"
	"fmt"
	"os"
	"strings"
)

const defaultPrompt = `
Please generate dummy data for up to 100 rows based on SQL table definitions.
`

type SQLDiffGenerator struct {
	gptClient *Client
}

func NewSQLDiffGenerator(apiKey string) (*SQLDiffGenerator, error) {
	gptClient, err := NewClient(apiKey)
	if err != nil {
		return nil, err
	}
	return &SQLDiffGenerator{
		gptClient: gptClient,
	}, nil
}

func (g *SQLDiffGenerator) GenerateDiff(sqlFile, outputFile, prompt string, override bool) error {
	if prompt == "" {
		prompt = defaultPrompt
	}

	if !strings.HasSuffix(sqlFile, ".sql") {
		return fmt.Errorf("invalid SQL file extension")
	}

	if _, err := os.Stat(sqlFile); os.IsNotExist(err) {
		return fmt.Errorf("SQL file does not exist: %s", sqlFile)
	}

	sqlContent, err := os.ReadFile(sqlFile)
	if err != nil {
		return err
	}

	resp, err := g.gptClient.Do(context.Background(), prompt+"\n"+string(sqlContent))
	if err != nil {
		return err
	}

	if override {
		if _, err := os.Stat(outputFile); os.IsNotExist(err) {
			return fmt.Errorf("output file does not exist: %s", outputFile)
		}

		outputFile, err := os.Open(outputFile)
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
