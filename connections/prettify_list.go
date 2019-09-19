package connections

import "fmt"

// PrettifyList formats a list of connections into a nice-looking list
func PrettifyList(rawList []*Connection) (prettifiedList []string) {
	for index, con := range rawList {
		prettifiedList = append(prettifiedList, fmt.Sprintf("(%d) %s", index, con))
	}
	return
}
