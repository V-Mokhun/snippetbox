package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	// new way to handle different methods, introduced in Go 1.22:
	// mux.HandleFunc("POST /snippet/create", snippetCreate)

	log.Print("Starting server on :4000")
	err := http.ListenAndServe("localhost:4000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
