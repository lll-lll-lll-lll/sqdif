package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sql-diff-bot/gpt"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "hello"
	app.Usage = "Generate SQL diff"
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "sql-file",
			Usage: "Path to the SQL file",
		},
		&cli.StringFlag{
			Name:  "output-file",
			Usage: "Path to the output file",
		},
		&cli.BoolFlag{
			Name:  "override",
			Usage: "Override the output file if it already exists",
		},
	}
	app.Action = func(c *cli.Context) error {
		gptClient, err := gpt.NewClient(os.Getenv("OPENAI_API_KEY"))
		if err != nil {
			log.Fatal(err)
		}
		sqlFile := c.String("sql-file")
		outputFile := c.String("output-file")
		override := c.Bool("override")

		if !strings.HasSuffix(sqlFile, ".sql") {
			log.Fatal("Invalid SQL file extension")
		}
		if _, err := os.Stat(sqlFile); os.IsNotExist(err) {
			log.Fatal("SQL file does not exist: ", sqlFile)
		}

		promptFile, err := os.ReadFile("prompt.txt")
		if err != nil {
			log.Fatal(err)
		}
		sqlContent, err := os.ReadFile(sqlFile)
		if err != nil {
			log.Fatal(err)
		}

		res, err := gptClient.Do(context.Background(), string(promptFile)+"\n"+string(sqlContent))
		if err != nil {
			log.Fatal(err)
		}

		if override {
			if _, err := os.Stat(outputFile); os.IsNotExist(err) {
				log.Fatal("outputFile does not exist: ", sqlFile)
			}
			outputFile, err := os.Open(outputFile)
			if err != nil {
				log.Fatal(err)
			}
			defer outputFile.Close()
			if _, err := outputFile.WriteString(res.Choices[0].Text); err != nil {
				log.Fatal(err)
			}
			fmt.Println("Output file updated")
		} else {
			f, err := os.Create(outputFile)
			if err != nil {
				log.Fatal(err)
			}
			f.WriteString(res.Choices[0].Text)
			defer f.Close()
			fmt.Println("Output file created")
		}

		fmt.Printf("Generated SQL: %s\n", res.Choices[0].Text)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}