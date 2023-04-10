package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
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

		// resp has following format in json:
		// {"MFAOptions":null,"PreferredMfaSetting":null,"UserAttributes":[{"Name":"sub","Value":"17f4366f-9fa9-422b-8ee3-2611446020b3"},{"Name":"email_verified","Value":"true"},{"Name":"email","Value":"dathalage99@gmail.com"}],"UserMFASettingList":null,"Username":"17f4366f-9fa9-422b-8ee3-2611446020b3"}

		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var data map[string]interface{}
		err = json.Unmarshal(respBody, &data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// get Username from data
		Username, ok := data["Username"].(string)
		if !ok {
			http.Error(w, "Username not found", http.StatusUnauthorized)
			return
		}

		// set username in context
		ctx := context.WithValue(r.Context(), "username", Username)
		r = r.WithContext(ctx)

		// if response is 200, call next handler
		next.ServeHTTP(w, r)
	})
}

func SearchNotes(w http.ResponseWriter, r *http.Request) {
	// get username from context
	username := r.Context().Value("username").(string)

	// get query from url
	query := r.URL.Query().Get("query")

	// do a get request to /api/search
	// send username and query in url
	resp, err := http.Get(os.Getenv("search-service") + "/api/search?username=" + username + "&query=" + query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// if response is not 200, return error
	if resp.StatusCode != http.StatusOK {
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}

	// resp has following format in json:
	// [{"created":"2021-04-20T11:24:25.000Z","title":"hello","content":"world"}]

	// read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// write response body
	w.Write(respBody)
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
	notesSubRouter.HandleFunc("/{created}", GetNote).Methods("GET")
	notesSubRouter.HandleFunc("/{created}", UpdateNote).Methods("PUT")
	notesSubRouter.HandleFunc("/{created}", DeleteNote).Methods("DELETE")

	// search service
	searchSubRouter := r.PathPrefix("/api/search").Subrouter()
	searchSubRouter.Use(authenticateHandler)
	searchSubRouter.HandleFunc("/", SearchNotes).Methods("POST")

	// start server
	http.ListenAndServe(":80", r)
}
