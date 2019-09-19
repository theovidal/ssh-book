// Package connections provides all the models and tools to interact with SSH connections.
//
// Prettify SSH connections into a list of strings :
//   rawList := []*Connection{...}
//   for _, connection := range connections.PrettifyList(rawList) {
//	   println(connection)
//	 }
//   // Output :
//   // (0) MyComputer : me@localhost (with port 22 and options "")
//   // (1) MyCloud : me@cloud.me.com (with port 256 and options "-g")
package connections
