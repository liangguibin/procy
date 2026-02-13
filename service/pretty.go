package service

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"os"
)

// table writer
var t = table.NewWriter()

// SetProcessStyle set process style
func SetProcessStyle() {
	t.SetColumnConfigs([]table.ColumnConfig{
		{
			Number:           7,
			WidthMax:         100,
			WidthMaxEnforcer: text.WrapSoft,
		},
	})
}

// Output output
func Output(data [][]string) {
	t.SetOutputMirror(os.Stdout)

	asRow := func(s []string) table.Row {
		row := make(table.Row, len(s))
		for i, v := range s {
			row[i] = v
		}
		return row
	}

	t.AppendHeader(asRow(data[0]))

	if len(data) > 1 {
		for _, r := range data[1:] {
			t.AppendRow(asRow(r))
		}
	}

	t.Render()
}
