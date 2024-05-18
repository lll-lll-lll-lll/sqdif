package main

import (
	"log"
	"os"

	"github.com/sql-diff-bot/gpt"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "sqdif"
	app.Usage = "Generate SQL diff using GPT-4 API"
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
		&cli.StringFlag{
			Name:  "prompt",
			Usage: "prompt to use for GPT-4 API",
		},
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	generator, err := gpt.NewSQLDiffGenerator(apiKey)
	if err != nil {
		log.Fatal(err)
	}

	app.Action = func(c *cli.Context) error {
		sqlFile := c.String("sql-file")
		outputFile := c.String("output-file")
		override := c.Bool("override")
		prompt := c.String("prompt")

		err := generator.GenerateDiff(sqlFile, outputFile, prompt, override)
		if err != nil {
			log.Fatal(err)
		}

		return nil
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
