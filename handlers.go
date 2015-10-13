package main

import (
	"fmt"
	"net/http"
)

// RootHandler serves the main page.
func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, alert)
}
