// Package connections contains all the models and operations related to SSH connections stored in the app
package connections

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