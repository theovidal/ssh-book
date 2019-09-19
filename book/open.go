package book

import (
	"github.com/exybore/ssh-book/connections"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

// Open reads the book file (or creates it) and parses connections into a single array
func Open() (connectionsList []*connections.Connection, err error) {
	filename, err := Path()
	if err != nil {
		return
	}

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			return connectionsList, err
		}
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(data, &connectionsList)
	if err != nil {
		return
	}

	return
}
