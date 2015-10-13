package main

import (
	"net/http"

	"github.com/codegangsta/negroni"
	_ "github.com/lib/pq"
)

func main() {
	mux := http.NewServeMux()
	neg := negroni.Classic()

	mux.HandleFunc("/api", RootHandler)
	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	neg.UseHandler(mux)
	neg.Run(":9494")
}
