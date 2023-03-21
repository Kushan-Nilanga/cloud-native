package main

import (
	"fmt"
	"log"
	"net/http"
)

type Config struct{}

func main() {
	app := Config{}

	fmt.Println("Starting server")

	// define http server and server handler
	srv := &http.Server{
		Addr:    ":80",
		Handler: app.routes(),
	}

	// start server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
