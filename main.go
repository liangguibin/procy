package main

import (
	"fmt"
	"github.com/liangguibin/procy/cmd/version"
	"github.com/spf13/cobra"
	"os"
)

// main
func main() {
	// root command
	var rootCmd = &cobra.Command{
		Use:   "procy",
		Short: "a process management tool",
		Long:  "a process management tool for linux",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Thank you for using procy and enjoying it.")
		},
	}
	// sub command
	rootCmd.AddCommand(version.RootCmd)
	// execute
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
