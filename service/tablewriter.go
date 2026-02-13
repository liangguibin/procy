package service

import (
	"github.com/olekukonko/tablewriter"
	"os"
)

// Output output
func Output(data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.Header(data[0])
	_ = table.Bulk(data[1:])
	_ = table.Render()
}
