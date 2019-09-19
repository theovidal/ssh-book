// Package connections contains all the models and operations related to SSH connections stored in the app
package connections

import "fmt"

// Connection represents SSH credentials for a connexion
type Connection struct {
	// Name given fby the user to identify the connexion
	Name string
	// Hostname to connect to, e.g "google.com" or "192.168.1.27"
	Hostname string
	// Port to use for the connexion (default is 22)
	Port int
	// Username to connect with, e.g "larrypage"
	Username string
	// Options to parse when executing the command
	Options string
}

// String returns the connection formatted as a string
func (c Connection) String() string {
	return fmt.Sprintf("%s : %s@%s (with port %d and options \"%s\")",
		c.Name,
		c.Username,
		c.Hostname,
		c.Port,
		c.Options,
	)
}
