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

func getArgsFromDOM() Args {
	document := js.Global().Get("document")
	input := document.Call("getElementById", "inputText").Get("value").String()
	return Args{Input: input}
}

func encodeBase64(this js.Value, args []js.Value) any {
	argData := getArgsFromDOM()

	slog.Info("encodeBase64", slog.Any("argData", argData))

	encoded := base64.StdEncoding.EncodeToString([]byte(argData.Input))

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
