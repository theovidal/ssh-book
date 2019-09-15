package main

import (
	"github.com/exybore/ssh-book/commands"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "list all SSH connections",
			Action: func(c *cli.Context) error {
				return commands.ListConnections(c)
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
