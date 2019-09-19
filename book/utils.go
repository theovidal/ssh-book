package book

import (
	"fmt"
	"os"
)

// Path returns the absolute path of the book
func Path() (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/.ssh-book.yml", homedir), nil
}
