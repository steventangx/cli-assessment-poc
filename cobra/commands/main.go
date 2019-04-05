package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

/* Example Usage:
 * $ ./binary oc config
 */

var (
	rootCmd = &cobra.Command{
		Use:   "cobra",
		Short: "example app to parse flags",
		Long:  "example app to parse flags",
	}
	ocCmd = &cobra.Command{
		Use:   "oc",
		Short: "oc app",
		Long:  "oc app",
		Aliases: []
	}
	configCmd = &cobra.Command{
		Use:   "config",
		Short: "description for config",
		Long:  "description for config",
	}
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// cobra
	ocCmd.AddCommand(configCmd)
	rootCmd.AddCommand(ocCmd)
}
