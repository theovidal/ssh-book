package book

import (
	"fmt"
	"github.com/exybore/ssh-book/connections"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func Open() (connectionsList []*connections.Connection, err error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return
	}

	filename := fmt.Sprintf("%s/.ssh-book.yml", homedir)
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
