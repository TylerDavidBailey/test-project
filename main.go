//go:build js && wasm

package main

import (
	"encoding/base64"
	"syscall/js"
)

func encodeBase64(this js.Value, args []js.Value) any {
	document := js.Global().Get("document")

	input := document.Call("getElementById", "inputText").Get("value").String()
	encoded := base64.StdEncoding.EncodeToString([]byte(input))

	outputElem := document.Call("getElementById", "output")
	outputElem.Set("innerHTML", encoded)

	return nil
}

func main() {
	c := make(chan struct{}, 0)

	document := js.Global().Get("document")

	encodeBtn := document.Call("getElementById", "encode")
	encodeBtn.Call("addEventListener", "click", js.FuncOf(encodeBase64))

	<-c
}
