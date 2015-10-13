package utils

import "github.com/gopherjs/gopherjs/js"

type window struct {
	JS      *js.Object
	Console *js.Object
}

func (w window) Alert(msg string) {
	w.JS.Call("alert", msg)
}

func (w window) Log(msg string) {
	w.Console.Call("log", msg)
}

func (w window) Call(obj, function, args ...string) {
	object := w.JS.Get(obj)
	object.Call(function, args)
}

var global = js.Global.Get("window")

// Window represents the JS global window.
var Window = window{
	JS:      global,
	Console: global.Get("console"),
}
