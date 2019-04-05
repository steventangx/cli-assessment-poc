package main

import (
	"fmt"

	flags "github.com/jessevdk/go-flags"
)

/* Example Usage:
 * $ ./binary oc config
 */
const description = "start the app"

type ocCmd struct {
	config bool `long:"config"`
}

func main() {
	ocCmd := &ocCmd{}
	parser := flags.NewParser(nil, flags.Default)
	parser.AddCommand("oc", description, description, ocCmd)
	parser.AddCommand("opencollector", description, description, ocCmd)
	val, _ := parser.Parse()

	if val[0] == "config" {
		fmt.Println("successfully parsed command")
	}
}
