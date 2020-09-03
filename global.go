/*
Package wasm provides some helpers to facilitate communication between
JavaScript DOM/globals and Go, in the context of WebAssembly.
For now it is an incomplete collection of experiments and it is not meant
to be used in production.
*/
package wasm

import (
	"reflect"
	"syscall/js"
)

// JSFunc type
type JSFunc func(this js.Value, args []js.Value) interface{}

// JSFuncMap type
type JSFuncMap map[string]JSFunc

// JSPropMap type
type JSPropMap map[string]interface{}

// Define sets a JavaScript property with given name and value.
// The value cannot be a function and panics in this case,
// use DefineFunc instead.
func Define(name string, value interface{}) {
	if reflect.TypeOf(value).Kind() == reflect.Func {
		panic("cannot use Define to set a function, use DefineFunc instead")
	}
	js.Global().Set(name, value)
}

// DefineMap sets a JavaScript properties according to the given map.
// Map values cannot be a function and panics in this case,
// use DefineFuncMap instead.
func DefineMap(propMap map[string]interface{}) {
	for k, v := range propMap {
		Define(k, v)
	}
}

// DefineFunc sets a JavaScript function with given name in the global scope.
func DefineFunc(name string, function JSFunc) {
	js.Global().Set(name, js.FuncOf(function))
}

// DefineFuncMap sets JavaScript functions in the global scope
// according to the input JSFuncMap
func DefineFuncMap(funcMap JSFuncMap) {
	for k, v := range funcMap {
		DefineFunc(k, v)
	}
}

// DefineAll sets JavaScript functions in the global scope
// according to the current *JSFuncMap
func (fm *JSFuncMap) DefineAll() *JSFuncMap {
	DefineFuncMap(*fm)
	return fm
}

// DefineAll sets JavaScript properties in the global scope
// according to the current *JSPropMap. It panics when trying
// to define a function, use DefineFuncMap() or *JSFuncMap.DefineAll()
// instead.
func (pm *JSPropMap) DefineAll() *JSPropMap {
	DefineMap(*pm)
	return pm
}
