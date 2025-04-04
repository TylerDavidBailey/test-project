//go:build js && wasm

package main

import (
	"encoding/base64"
	"syscall/js"
)

// Go function to encode text to base64
func encodeBase64(this js.Value, p []js.Value) interface{} {
	text := p[0].String() // Get the text from JavaScript
	encoded := base64.StdEncoding.EncodeToString([]byte(text))
	return js.ValueOf(encoded) // Return the base64 encoded string to JavaScript
}

func main() {
	// Register Go function so it can be called from JavaScript
	js.Global().Set("wasmEncodeBase64", js.FuncOf(encodeBase64))

	// Block the Go program from exiting
	select {}
}
