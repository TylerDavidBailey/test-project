//go:build js && wasm

package main

import (
	"log/slog"
	"syscall/js"
)

type Args struct {
	Input string
}

func getArgsFromDOM() Args {
	document := js.Global().Get("document")
	selected := document.Call("getElementById", "inputSelect").Get("value").String()
	return Args{Input: selected}
}

func handleClick(this js.Value, args []js.Value) any {
	argData := getArgsFromDOM()

	slog.Info("selected value", slog.Any("input", argData.Input))

	document := js.Global().Get("document")
	outputElem := document.Call("getElementById", "output")
	outputElem.Set("innerHTML", argData.Input)

	return nil
}

func main() {
	slog.Info("WASM app starting")

	c := make(chan struct{}, 0)

	document := js.Global().Get("document")
	enterBtn := document.Call("getElementById", "enter")
	enterBtn.Call("addEventListener", "click", js.FuncOf(handleClick))

	<-c
}
