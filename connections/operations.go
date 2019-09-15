package connections

import "fmt"

func ListConnections(rawList []*Connection) (prettifiedList []string) {
	for index, con := range rawList {
		prettifiedList = append(prettifiedList,
			fmt.Sprintf("(%d) %s : %s@%s (with port %d and options %s)",
				index,
				con.Name,
				con.Username,
				con.Hostname,
				con.Port,
				con.Options,
			),
		)
	}
	return
}

func AddConnection(list *[]*Connection, element *Connection) {
	*list = append(*list, element)
}

func DeleteConnection(list *[]*Connection, index int) {

}

func Connect(list *[]*Connection, index int) {

}
