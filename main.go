//go:build js && wasm

package main

import (
	"syscall/js"
)

type Table struct {
	Headers []string
	Rows    [][]string
}

func getTable(this js.Value, args []js.Value) any {
	table := Table{
		Headers: []string{"Column A", "Column B", "Column C"},
		Rows: [][]string{
			{"Row 1A", "Row 1B", "Row 1C"},
			{"Row 2A", "Row 2B", "Row 2C"},
		},
	}

	result := js.Global().Get("Object").New()

	// Set headers
	headers := js.Global().Get("Array").New(len(table.Headers))
	for i, h := range table.Headers {
		headers.SetIndex(i, h)
	}
	result.Set("headers", headers)

	// Set rows
	rows := js.Global().Get("Array").New(len(table.Rows))
	for i, row := range table.Rows {
		jsRow := js.Global().Get("Array").New(len(row))
		for j, col := range row {
			jsRow.SetIndex(j, col)
		}
		rows.SetIndex(i, jsRow)
	}
	result.Set("rows", rows)

	return result
}

func main() {
	js.Global().Set("getTable", js.FuncOf(getTable))
	select {}
}
