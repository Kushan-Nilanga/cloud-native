package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// gorilla mux
	r := mux.NewRouter()

	// routes

	// start server
	http.ListenAndServe(":8080", r)
}
