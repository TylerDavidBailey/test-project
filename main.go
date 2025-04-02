package main

import (
	"strconv"
	"syscall/js"
)

var counter int

func increment(this js.Value, p []js.Value) interface{} {
	counter++
	js.Global().Get("document").Call("getElementById", "counter").Set("innerText", strconv.Itoa(counter))
	return nil
}

func main() {
	// Set initial counter value in the HTML
	js.Global().Get("document").Call("getElementById", "counter").Set("innerText", strconv.Itoa(counter))

	// Bind the increment function to the button's click event
	cb := js.FuncOf(increment)
	defer cb.Release()
	js.Global().Get("document").Call("getElementById", "incrementButton").Call("addEventListener", "click", cb)

	// Prevent the Go program from exiting
	select {}
}
