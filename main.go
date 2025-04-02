package main

import (
	"syscall/js"
)

var counter int

func increment(this js.Value, p []js.Value) interface{} {
	counter++
	js.Global().Get("document").Call("getElementById", "counter").Set("innerText", counter)
	return nil
}

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("increment", js.FuncOf(increment))
	<-c
}
