//go:build js && wasm

package main

import (
	"syscall/js"
)

var count int

func updateCounterDisplay() {
	document := js.Global().Get("document")
	counterText := document.Call("getElementById", "counter")
	counterText.Set("innerHTML", count)
}

func increment(this js.Value, args []js.Value) any {
	count++
	updateCounterDisplay()
	return nil
}

func clear(this js.Value, args []js.Value) any {
	count = 0
	updateCounterDisplay()
	return nil
}

func main() {
	c := make(chan struct{}, 0)

	document := js.Global().Get("document")

	// Bind increment button
	incrBtn := document.Call("getElementById", "increment")
	incrBtn.Call("addEventListener", "click", js.FuncOf(increment))

	// Bind clear button
	clearBtn := document.Call("getElementById", "clear")
	clearBtn.Call("addEventListener", "click", js.FuncOf(clear))

	// Initialize display
	updateCounterDisplay()

	<-c // Prevent main from exiting
}
