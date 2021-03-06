package main

import (
	"net/http"
	"log"
	"github.com/squiidz/bone"
)

func main() {
	// New mux instance
	mux := bone.New()
	// Custom 404
	mux.NotFound(Handler404)
	// Handle with any http method, Handle takes http.Handler as argument.
	mux.Handle("/index", http.HandlerFunc(homeHandler))
	// Get, Post etc... takes http.HandlerFunc as argument.
	mux.Post("/home", homeHandler)

	// Start Listening
	log.Fatal(http.ListenAndServe(":8080", mux));
}

func homeHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("WELCOME HOME"))
}

func Handler404(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("These are not the droids you're looking for ..."))
}
