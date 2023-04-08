package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type AuthRequest struct {
	Token string `json:"token"`
}

func authenticateHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// get bearer token from header
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		// access token string after "Bearer "
		token = token[7:]

		// do a post request to /api/authenticate
		// send token in body as 'token'
		req := AuthRequest{Token: token}
		reqBody, err := json.Marshal(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp, err := http.Post(os.Getenv("auth-service")+"/api/authenticate", "application/json", bytes.NewBuffer(reqBody))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// if response is not 200, return error
		if resp.StatusCode != http.StatusOK {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		// attach resp to request context
		r = r.WithContext(context.WithValue(r.Context(), "user", resp))

		// if response is 200, call next handler
		next.ServeHTTP(w, r)
	})
}

func main() {
	// gorilla mux
	r := mux.NewRouter()

	// / get route
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("bff service running"))
	})

	// routes
	// auth service
	authSubRouter := r.PathPrefix("/api/auth").Subrouter()
	authSubRouter.HandleFunc("/signup", signupHandler).Methods("POST")
	authSubRouter.HandleFunc("/signin", signinHandler).Methods("POST")
	authSubRouter.HandleFunc("/verify-email", verifyEmailHandler).Methods("POST")

	// notes service
	notesSubRouter := r.PathPrefix("/api/notes").Subrouter()
	notesSubRouter.Use(authenticateHandler)
	notesSubRouter.HandleFunc("/", ListNotes).Methods("GET")
	notesSubRouter.HandleFunc("/", CreateNote).Methods("POST")
	notesSubRouter.HandleFunc("/{id}", GetNote).Methods("GET")
	notesSubRouter.HandleFunc("/{id}", UpdateNote).Methods("PUT")

	// search service
	searchSubRouter := r.PathPrefix("/api/search").Subrouter()
	searchSubRouter.Use(authenticateHandler)

	// start server
	http.ListenAndServe(":80", r)
}
