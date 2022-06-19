package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dukex/ely/server"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "ely",
		Usage: "the project description here",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Value:   "ely.yml",
				Usage:   "Load configuration from `FILE`",
			},
		},
		Action: func(c *cli.Context) error {
			configFile := c.String("config")

			server.Server()

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
