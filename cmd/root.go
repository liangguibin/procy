package cmd

import "github.com/spf13/cobra"

// root command
var rootCmd = &cobra.Command{
	Use:   "procy",
	Short: "a process management tool",
	Long:  "a process management tool for linux",
}

// Execute execute
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
