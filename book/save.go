package book

import (
	"github.com/exybore/ssh-book/connections"
	"gopkg.in/yaml.v2"
	"os"
)

// Save writes into the book modifications made by the user : adding, updating or deleting connections.
func Save(list *[]*connections.Connection) (err error) {
	filename, err := Path()
	if err != nil {
		return
	}

	err = os.Truncate(filename, 0)
	if err != nil {
		return
	}

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return
	}

	data, err := yaml.Marshal(&list)
	if err != nil {
		return
	}

	_, err = file.Write(data)
	if err != nil {
		return
	}

	err = file.Close()
	return
}
