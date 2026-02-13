package cmd

import (
	"github.com/liangguibin/procy/service"
	"github.com/spf13/cobra"
)

// version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version number",
	Long:  `the version number of procy`,
	Run: func(cmd *cobra.Command, args []string) {
		data := [][]string{
			{"version"},
			{"v0.0.1"},
		}

		service.Output(data)
	},
}

// init
func init() {
	rootCmd.AddCommand(versionCmd)
}
