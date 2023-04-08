package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// notes struct
type Note struct {
	ID      string `json:"id,omitempty"`
	Title   string `json:"title"`
	UserID  string `json:"user_id"`
	Content string `json:"content"`
	Created string `json:"created,omitempty"`
	Updated string `json:"updated,omitempty"`
}

// ListNotes request struct
type ListNotesRequest struct {
	UserID string `json:"user_id"`
}

// ListNotes response struct
type ListNotesResponse struct {
	Notes []Note `json:"notes"`
}

// ListNotes handler
func ListNotes(w http.ResponseWriter, r *http.Request) {
	// get request body
	var req ListNotesRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// get notes
	notes, err := ListNotesFromDB(req.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return response
	res := ListNotesResponse{Notes: notes}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// CreateNoteRequest request struct
type CreateNoteRequest struct {
	Note Note `json:"note"`
}

// CreateNoteResponse response struct
type CreateNoteResponse struct {
	Note Note `json:"note"`
}

// CreateNote handler
func CreateNote(w http.ResponseWriter, r *http.Request) {
	// get request body
	var req CreateNoteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// create note
	note, err := CreateNoteInDB(req.Note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return response
	res := CreateNoteResponse{Note: note}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GetNoteRequest request struct
type GetNoteRequest struct {
	ID string `json:"id"`
}

// GetNoteResponse response struct
type GetNoteResponse struct {
	Note Note `json:"note"`
}

// GetNote handler
func GetNote(w http.ResponseWriter, r *http.Request) {
	// get request body
	var req GetNoteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// get note
	note, err := GetNoteFromDB(req.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return response
	res := GetNoteResponse{Note: note}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// UpdateNoteRequest request struct
type UpdateNoteRequest struct {
	Note Note `json:"note"`
}

// UpdateNoteResponse response struct
type UpdateNoteResponse struct {
	Note Note `json:"note"`
}

// UpdateNote handler
func UpdateNote(w http.ResponseWriter, r *http.Request) {
	// get request body
	var req UpdateNoteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// update note
	note, err := UpdateNoteInDB(req.Note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return response
	res := UpdateNoteResponse{Note: note}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	// new gorilla mux router
	r := mux.NewRouter()

	// routes
	r.HandleFunc("/notes", ListNotes).Methods("GET")
	r.HandleFunc("/notes", CreateNote).Methods("POST")
	r.HandleFunc("/notes/{id}", GetNote).Methods("GET")
	r.HandleFunc("/notes/{id}", UpdateNote).Methods("POST")

	// start server
	log.Fatal(http.ListenAndServe(":80", r))
}
