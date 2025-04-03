//go:build js && wasm

package main

import (
	"syscall/js"
)

func generateTable(this js.Value, p []js.Value) interface{} {
	document := js.Global().Get("document")
	table := document.Call("createElement", "table")
	table.Set("border", "1")

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	hours := []string{"9:00 AM", "12:00 PM", "3:00 PM", "6:00 PM", "9:00 PM"}

	for _, day := range days {
		row := document.Call("createElement", "tr")
		headCell := document.Call("createElement", "th")
		headCell.Set("innerText", day)
		row.Call("appendChild", headCell)
		for _, hour := range hours {
			cell := document.Call("createElement", "td")
			cell.Set("innerText", hour)
			row.Call("appendChild", cell)
		}
		table.Call("appendChild", row)
	}

	container := document.Call("getElementById", "tableContainer")
	container.Set("innerHTML", "") // Clear previous table
	container.Call("appendChild", table)

	return nil
}

func main() {
	cb := js.FuncOf(generateTable)
	defer cb.Release()

	js.Global().Get("document").Call("getElementById", "incrementButton").Call("addEventListener", "click", cb)
	select {}
}
