package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

const version = "0.2.6"

func main() {
	app := cli.NewApp()
	app.Version = version
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
			Name:  "prompt-path",
			Usage: "prompt file path to use for GPT-4 API",
		},
		&cli.StringFlag{
			Name:  "api-key",
			Usage: "API key for GPT-4 AP",
		},
	}

	app.Action = func(c *cli.Context) error {

		apiKey := c.String("api-key")
		generator, err := newGPTGenerator(apiKey)
		if err != nil {
			log.Fatal(err)
		}
		sqlFile := c.String("sql-file")
		outputFile := c.String("output-file")
		override := c.Bool("override")
		promptPath := c.String("prompt-path")

		if err := generator.Do(sqlFile, outputFile, promptPath, override); err != nil {
			log.Fatal(err)
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
