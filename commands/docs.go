// Package commands contains all the functions for CLI commands, such as listing, adding or deleting SSH connections.
// See an example of an integration with the CLI app :
//
//   app.Commands = []cli.Command{
//		{
//			Name:    "list",
//			Aliases: []string{"l"},
//			Usage:   "list all SSH connections",
//			Action: func(c *cli.Context) error {
//				return commands.ListConnections(c)
//			},
//		},
//	}
package commands
