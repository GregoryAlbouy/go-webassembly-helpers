package wasm

import (
	"syscall/js"
)

// Object wraps a javascript object js.Value, providing it some
// extended methods
type Object struct {
	Value js.Value
}

// PropMap is a map of JavaScript properties
type PropMap map[string]interface{}

// AttributeMap is a map of HTML attributes
type AttributeMap map[string]string

// StyleMap is a map of style values
type StyleMap map[string]string

// QuerySelector func
func (o *Object) QuerySelector(sel string) js.Value {
	return o.GetCall("querySelector", sel)
}

// QuerySelectorAll func
func (o *Object) QuerySelectorAll(sel string) js.Value {
	return o.GetCall("querySelectorAll", sel)
}

// ObjectFrom returns a wrapper of the input object js.Value
// to access some extended methods
func ObjectFrom(v js.Value) *Object {
	return &Object{Value: v}
}

// NewElement creates a new wrapped object from input tagname
func NewElement(tagname string) *Object {
	return &Object{Value: CreateElement(tagname)}
}

// Get returns the queried JavaScript property p of value v.
func (o *Object) Get(name string) js.Value {
	return o.Value.Get(name)
}

// SetProp method
func (o *Object) SetProp(name string, value interface{}) *Object {
	o.Value.Set(name, value)
	return o
}

// Call calls an object's method and returns the current object. The result
// of the called method is is ignored. Example:
// o.Call("setAttribute", "role", "search").SetProp(...)
func (o *Object) Call(name string, args ...interface{}) *Object {
	o.Value.Call(name, args...)
	return o
}

// GetCall calls an object's method and return its return value.
func (o *Object) GetCall(name string, args ...interface{}) js.Value {
	return o.Value.Call(name, args...)
}

// Invoke calls the current object if it is a function and panics otherwise.
// It ignores its return value and return the current *Object instead.
// (Use GetInvoke() if that return value is needed)
func (o *Object) Invoke(args ...interface{}) *Object {
	o.Value.Invoke(args...)
	return o
}

// GetInvoke calls the current object if it is a function and panics otherwise.
// It returns its return value.
func (o *Object) GetInvoke(args ...interface{}) js.Value {
	return o.Value.Invoke(args...)
}

// SetProps allows to set several properties in a single declaration
// using a map. Example: o.SetProps(wasm.PropMap{"textContent": "click me",
// "href": "./home.html"}
func (o *Object) SetProps(pm PropMap) *Object {
	for k, v := range pm {
		o.Value.Set(k, v)
	}
	return o
}

// GetAttribute returns the attribute of current *Object's value.
func (o *Object) GetAttribute(name string) string {
	return o.Value.Call("getAttribute", name).String()
}

// SetAttribute sets HTML attribute name to value and returns
// the current *Object.
func (o *Object) SetAttribute(name string, value string) *Object {
	o.Value.Call("setAttribute", name, value)
	return o
}

// SetAttributes sets HTML attribute from given AttributeMap and returns
// the current *Object.
func (o *Object) SetAttributes(am AttributeMap) *Object {
	for k, v := range am {
		o.SetAttribute(k, v)
	}
	return o
}

// AddClass adds classes to the *Object value and returns the *Object.
func (o *Object) AddClass(names ...string) *Object {
	return o.manageClass("add", names...)
}

// RemoveClass removes classes to the *Object value and returns the *Object.
func (o *Object) RemoveClass(names ...string) *Object {
	return o.manageClass("remove", names...)
}

func (o *Object) manageClass(method string, names ...string) *Object {
	// "Convert" []string to []interface{}
	values := make([]interface{}, len(names))
	for i, v := range names {
		values[i] = v
	}

	o.Value.Get("classList").Call(method, values...)
	return o
}

// SetStyle sets current *Object value style name to value.
func (o *Object) SetStyle(name, value string) *Object {
	o.Value.Get("style").Set(name, value)
	return o
}

// SetStyles sets the *Object current value styles from a wadom.StyleMap.
func (o *Object) SetStyles(styleMap StyleMap) *Object {
	for k, v := range styleMap {
		o.SetStyle(k, v)
	}
	return o
}
