package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

/* Example Usage:
 * $ ./binary -s
 * $ ./binary --start
 */

var (
	rootCmd = &cobra.Command{
		Use:   "cobra",
		Short: "example app to parse flags",
		Long:  "example app to parse flags",
	}
	start bool
	kill  bool
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if start || kill {
		fmt.Println("Successfully parsed flag(s)")
	}
}

func init() {
	// cobra
	rootCmd.PersistentFlags().BoolVarP(&start, "start", "s", false, "start the service")
	rootCmd.PersistentFlags().BoolVarP(&kill, "kill", "k", false, "kill the service")
}
