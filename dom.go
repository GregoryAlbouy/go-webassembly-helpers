package wasm

import (
	"syscall/js"
)

// Document returns the JS global document
func Document() js.Value {
	return js.Global().Get("document")
}

// CreateElement return a new js.Value element with the input tagname
// (shortcut to Document().Call("createElement", tagname).
func CreateElement(tagname string) js.Value {
	return Document().Call("createElement", tagname)
}

// GetElementByID retrieves an element in the DOM by its id
// (shortcut to Document().Call("getElementById", id)).
func GetElementByID(id string) js.Value {
	return Document().Call("getElementById", id)
}
