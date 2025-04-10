//go:build js && wasm

package main

import (
	"syscall/js"
)

func generateWorkout(this js.Value, args []js.Value) any {
	workout := []string{
		"Push-ups:10",
		"Squats:15",
		"Jumping Jacks:20",
		"Plank (sec):30",
	}

	// Convert Go slice to JS array
	jsArray := js.Global().Get("Array").New(len(workout))
	for i, item := range workout {
		jsArray.SetIndex(i, item)
	}

	return jsArray
}

func registerCallbacks() {
	js.Global().Set("generateWorkout", js.FuncOf(generateWorkout))
}

func main() {
	c := make(chan struct{}, 0)
	registerCallbacks()
	<-c
}
