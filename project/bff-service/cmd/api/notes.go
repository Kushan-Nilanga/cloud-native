package main

import (
	"io"
	"net/http"
	"os"
)

func ListNotes(w http.ResponseWriter, r *http.Request) {
	// do a get request to /api/notes
	// if successful, return the response
	// if not, return the error
	res, err := http.Get(os.Getenv("notes-service") + "/api/notes")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(resBody)
}

func CreateNote(w http.ResponseWriter, r *http.Request) {
	// do a post request to /api/notes
	// if successful, return the response
	// if not, return the error
	res, err := http.Post(os.Getenv("notes-service")+"/api/notes", "application/json", r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(resBody)
}

func UpdateNote(w http.ResponseWriter, r *http.Request) {
	// do a put request to /api/notes
	// if successful, return the response
	// if not, return the error
	res, err := http.Post(os.Getenv("notes-service")+"/api/notes", "application/json", r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(resBody)
}

func GetNote(w http.ResponseWriter, r *http.Request) {
	// do a get request to /api/notes
	// if successful, return the response
	// if not, return the error
	res, err := http.Get(os.Getenv("notes-service") + "/api/notes")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(resBody)
}
