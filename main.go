//go:build js && wasm

package main

import (
	"encoding/base64"
	"syscall/js"
)

func encodeBase64(this js.Value, p []js.Value) interface{} {
	text := p[0].String() // Get the text from JavaScript
	encoded := base64.StdEncoding.EncodeToString([]byte(text))
	return js.ValueOf(encoded) // Return the encoded base64 string back to JavaScript
}

func main() {
	c := make(chan struct{}, 0)

	// Register the Go function to be callable from JavaScript
	js.Global().Set("runWasm", js.FuncOf(encodeBase64))

	<-c // Keep the Go program running
}
