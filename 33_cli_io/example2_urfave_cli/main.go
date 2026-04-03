package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "mycli",
		Usage: "A simple CLI with urfave/cli",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "name",
				Aliases: []string{"n"},
				Usage:   "Name to greet",
				Value:   "World",
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "greet",
				Usage: "Greet someone",
				Action: func(c *cli.Context) error {
					name := c.String("name")
					fmt.Printf("Hello, %s!\n", name)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
