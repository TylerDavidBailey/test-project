//go:build js && wasm

package main

import (
	"encoding/base64"
	"log/slog"
	"syscall/js"
)

func encodeBase64(this js.Value, args []js.Value) any {
	document := js.Global().Get("document")

	input := document.Call("getElementById", "inputText").Get("value").String()
	slog.Info("encoding this value", slog.Any("input", input))
	encoded := base64.StdEncoding.EncodeToString([]byte(input))
	slog.Info("encoded value", slog.Any("encoded", encoded))

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
