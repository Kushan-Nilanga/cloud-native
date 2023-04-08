package main

import (
	"io/ioutil"
	"net/http"
	"os"
)

func signupHandler(w http.ResponseWriter, r *http.Request) {
	// do a post request to /api/signup with the request body
	// if successful, return the response
	// if not, return the error
	resp, err := http.Post(os.Getenv("auth-service")+"/api/signup", "application/json", r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// write response to response writer
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(respBody)
}

func signinHandler(w http.ResponseWriter, r *http.Request) {
	// do a post request to /api/signin with the request body
	// if successful, return the response
	// if not, return the error
	resp, err := http.Post(os.Getenv("auth-service")+"/api/signin", "application/json", r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// write response to response writer
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(respBody)
}

func verifyEmailHandler(w http.ResponseWriter, r *http.Request) {
	// do a post request to /api/verify-email with the request body
	// if successful, return the response
	// if not, return the error
	resp, err := http.Post(os.Getenv("auth-service")+"/api/verify-email", "application/json", r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// write response to response writer
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(respBody)
}
