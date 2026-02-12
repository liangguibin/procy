package version

import (
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
)

// RootCmd version
var RootCmd = &cobra.Command{
	Use:   "version",
	Short: "version number",
	Long:  `the version number of procy`,
	Run: func(cmd *cobra.Command, args []string) {
		data := [][]string{
			{"version"},
			{"v0.0.1"},
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.Header(data[0])
		_ = table.Bulk(data[1:])
		_ = table.Render()
	},
}
