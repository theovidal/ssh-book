package main

import (
	"fmt"
	"github.com/exybore/ssh-book/connections"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("error while getting home directory: %v", err)
	}

	filename := fmt.Sprintf("%s/.ssh-book.yml", homedir)
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			log.Fatalf("error while creating book file: %v", err)
		}
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("error while reading book file: %v", err)
	}

	var connexionsList []*connections.Connection
	err = yaml.Unmarshal(data, &connexionsList)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	for _, con := range connexionsList {
		fmt.Print(con)
	}
}
