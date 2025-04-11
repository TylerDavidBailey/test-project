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

	jsArray := js.Global().Get("Array").New(len(workout))
	for i, item := range workout {
		jsArray.SetIndex(i, item)
	}
	return jsArray
}

func main() {
	js.Global().Set("generateWorkout", js.FuncOf(generateWorkout))
	select {}
}
