package main

import (
	"github.com/exybore/ssh-book/commands"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "ssh-book"
	app.Usage = "stores all your SSH connections"
	app.Version = "0.1.0"
	app.Authors = []cli.Author{
		{
			Name:  "Exybore",
			Email: "exybore@becauseofprog.fr",
		},
	}
	app.UsageText = "ssh-book [flags] command [command options] [arguments]"

	app.Commands = []cli.Command{
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "List all SSH connections",
			Action: func(c *cli.Context) error {
				return commands.ListConnections(c)
			},
		},
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "Add a SSH connection",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:     "name, n",
					Value:    "connection",
					Usage:    "name for the connection",
					Required: true,
				},
			},
			Action: func(c *cli.Context) error {
				return commands.AddConnection(c)
			},
		},
		{
			Name:    "remove",
			Aliases: []string{"r"},
			Usage:   "Remove a SSH connection",
			Action: func(c *cli.Context) error {
				return commands.RemoveConnection(c)
			},
		},
		{
			Name:    "connect",
			Aliases: []string{"c"},
			Usage:   "Connect to a SSH connection",
			Action: func(c *cli.Context) error {
				return commands.Connect(c)
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
