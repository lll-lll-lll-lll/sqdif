package main

import (
	"log"
	"os"

	"github.com/sqdif/gpt"
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
		&cli.StringFlag{
			Name:  "api-key",
			Usage: "API key for GPT-4 AP",
		},
	}

	app.Action = func(c *cli.Context) error {

		apiKey := c.String("api-key")
		generator, err := gpt.NewSQLDiffGenerator(apiKey)
		if err != nil {
			log.Fatal(err)
		}
		sqlFile := c.String("sql-file")
		outputFile := c.String("output-file")
		override := c.Bool("override")
		prompt := c.String("prompt")

		if err := generator.GenerateDiff(sqlFile, outputFile, prompt, override); err != nil {
			log.Fatal(err)
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
