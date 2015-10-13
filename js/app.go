package main

import "github.com/buddhamagnet/gopherj/js/utils"

// Window represents the JS global window.
var Window = utils.Window

func main() {
	Window.Alert("Hello Dave!")
	Window.Log("And I'm here too!")
	Window.Call("console", "log", "And I am clever!")
}
