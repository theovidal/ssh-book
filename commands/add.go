package commands

import (
	"github.com/exybore/ssh-book/book"
	"github.com/exybore/ssh-book/connections"
	"github.com/urfave/cli"
	"log"
	"strconv"
)

// AddConnection is a CLI command that adds a connection into user's book.
// It is called when the "ssh-book add" command is triggered by the user.
func AddConnection(c *cli.Context) (err error) {
	connectionsList, err := book.Open()
	if err != nil {
		log.Fatal("error while opening the book: ", err)
		return
	}

	name := c.String("name")
	hostname := c.String("hostname")
	port, err := strconv.Atoi(c.String("port"))
	if err != nil {
		log.Fatal("error: given port isn't an integer")
		return
	}
	username := c.String("username")
	options := c.String("options")

	connectionsList = append(connectionsList, &connections.Connection{
		Name:     name,
		Hostname: hostname,
		Port:     port,
		Username: username,
		Options:  options,
	})

	err = book.Save(&connectionsList)
	if err == nil {
		println("Connection saved successfully.")
	}

	return
}
