package commands

import (
	"github.com/exybore/ssh-book/book"
	"github.com/exybore/ssh-book/connections"
	"github.com/urfave/cli"
	"log"
)

// ListConnections is a CLI command that lists all the saved connections in user's book.
// It is called when the "ssh-book list" command is triggered by the user.
func ListConnections(c *cli.Context) error {
	connectionsList, err := book.Open()
	if err != nil {
		log.Fatal("error while opening the book: ", err)
		return err
	}

	println("List of all SSH connections :")
	for _, connection := range connections.PrettifyList(connectionsList) {
		println(connection)
	}

	return nil
}
