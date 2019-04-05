package main

import (
	"fmt"
	"os"

	flags "github.com/jessevdk/go-flags"
)

/* Example Usage:
 * $ ./binary -s
 * $ ./binary --start
 */

type goFlag struct {
	Start bool `short:"s" long:"start" description:"start the service"`
	Kill  bool `short:"k" long:"kill" description:"kill the service"`
}

// Flag parsing
func main() {
	cmd := &goFlag{}
	_, err := flags.Parse(cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if cmd.Start || cmd.Kill {
		fmt.Println("Successfully parsed flag(s)")
	}
}
