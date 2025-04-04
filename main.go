//go:build js && wasm

package main

import (
	"encoding/base64"
	"log/slog"
	"syscall/js"
)

type Args struct {
	Input string
}

// Grabs the selected value from the dropdown menu
func getArgsFromDOM() Args {
	document := js.Global().Get("document")
	selected := document.Call("getElementById", "inputSelect").Get("value").String()
	return Args{Input: selected}
}

func encodeBase64(this js.Value, args []js.Value) any {
	argData := getArgsFromDOM()

	slog.Info("encoding selected value", slog.Any("input", argData.Input))
	encoded := base64.StdEncoding.EncodeToString([]byte(argData.Input))
	slog.Info("encoded result", slog.Any("encoded", encoded))

	document := js.Global().Get("document")
	outputElem := document.Call("getElementById", "output")
	outputElem.Set("innerHTML", encoded)

	return nil
}

func main() {
	slog.Info("WASM app starting")

	c := make(chan struct{}, 0)

	document := js.Global().Get("document")
	encodeBtn := document.Call("getElementById", "encode")
	encodeBtn.Call("addEventListener", "click", js.FuncOf(encodeBase64))

	<-c
}
