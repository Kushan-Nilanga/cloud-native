package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/tsawler/toolbox"
)

// App Config struct
type Config struct{}

// LogPayload struct
type LogPayload struct {
	Message   string `json:"message"`
	TimeStamp string `json:"timestamp,omitempty"`
}

// routes function
func (app *Config) routes() http.Handler {
	// instantiating router
	mux := chi.NewRouter()

	// who is allowed to connect
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// health check
	mux.Use(middleware.Heartbeat("/ping"))

	// routes
	mux.Post("/api/logs", app.Logger)

	// returns
	return mux
}

func (app *Config) Logger(w http.ResponseWriter, r *http.Request) {
	var tools toolbox.Tools

	// get the payload
	var logPayload LogPayload
	err := tools.ReadJSON(w, r, &logPayload)
	if err != nil {
		payload := toolbox.JSONResponse{
			Error:   true,
			Message: "Invalid request formatting",
		}

		tools.WriteJSON(w, http.StatusBadRequest, payload)
		return
	}

	// TODO functionality log data to cloudwatch logs

	// sending response to client service
	payload := toolbox.JSONResponse{
		Error:   false,
		Message: logPayload.Message + " has been logged",
	}

	_ = tools.WriteJSON(w, http.StatusAccepted, payload)
}

// main function
func main() {
	app := Config{}

	fmt.Printf("Starting Service")

	srv := &http.Server{
		Addr:    ":80",
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
