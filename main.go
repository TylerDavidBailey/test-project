//go:build js && wasm

package main

import (
	"strings"
	"syscall/js"
)

func generateWorkout(this js.Value, args []js.Value) any {
	plan := []string{
		"10 push-ups",
		"15 squats",
		"20 jumping jacks",
		"30-second plank",
	}

	// Join as a string
	workout := strings.Join(plan, "\\n")

	// Set the text to the DOM
	js.Global().Get("document").
		Call("getElementById", "workout").
		Set("innerText", workout)

	return nil
}

func registerCallbacks() {
	js.Global().Set("generateWorkout", js.FuncOf(generateWorkout))
}

func main() {
	c := make(chan struct{}, 0)
	registerCallbacks()
	<-c
}
