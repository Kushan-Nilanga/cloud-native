package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func ListNotes(w http.ResponseWriter, r *http.Request) {
	// do a get request to /api/notes/{username}
	// if successful, return the response
	// if not, return the error
	// write user context to request header
	// get user from context
	username := r.Context().Value("username")
	if username == nil {
		http.Error(w, "username not found", http.StatusUnauthorized)
		return
	}

	fmt.Println("username: ", username.(string))

	// do a get request to /api/notes
	// attach username from context to url
	res, err := http.Get(os.Getenv("notes-service") + "/api/notes/" + username.(string))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("response body: ", string(resBody))

	w.Write(resBody)
}

func CreateNote(w http.ResponseWriter, r *http.Request) {
	// do a post request to /api/notes
	// if successful, return the response
	// if not, return the error
	// write user context to request header
	// get user from context
	username := r.Context().Value("username")
	if username == nil {
		http.Error(w, "username not found", http.StatusUnauthorized)
		return
	}

	fmt.Println("username: ", username.(string))

	// do a post request to /api/notes
	// attach username from context to url
	res, err := http.Post(os.Getenv("notes-service")+"/api/notes/"+username.(string), "application/json", r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("response body: ", string(resBody))

	w.Write(resBody)
}

func GetNote(w http.ResponseWriter, r *http.Request) {
	// do a get request to /api/notes/{username}/{created}
	// if successful, return the response
	// if not, return the error
	// write user context to request header
	// get user from context
	// get created from request params
	username := r.Context().Value("username")
	if username == nil {
		http.Error(w, "username not found", http.StatusUnauthorized)
		return
	}

	created := mux.Vars(r)["created"]
	fmt.Println("username: ", username.(string), "created: ", created)

	// do a get request to /api/notes
	// attach username from context to url
	res, err := http.Get(os.Getenv("notes-service") + "/api/notes/" + username.(string) + "/" + created)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("response body: ", string(resBody))

	w.Write(resBody)
}

func UpdateNote(w http.ResponseWriter, r *http.Request) {
	// do a put request to /api/notes/{username}/{created}
	// if successful, return the response
	// if not, return the error
	// write user context to request header
	// get user from context
	// get created from request params
	username := r.Context().Value("username")
	if username == nil {
		http.Error(w, "username not found", http.StatusUnauthorized)
		return
	}

	created := mux.Vars(r)["created"]
	fmt.Println("username: ", username.(string), "created: ", created)

	// do a put request to /api/notes
	// attach username from context to url
	req, err := http.NewRequest("PUT", os.Getenv("notes-service")+"/api/notes/"+username.(string)+"/"+created, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("response body: ", string(resBody))

	w.Write(resBody)
}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	// do a delete request to /api/notes/{username}/{created}
	// if successful, return the response
	// if not, return the error
	// write user context to request header
	// get user from context
	// get created from request params
	username := r.Context().Value("username")
	if username == nil {
		http.Error(w, "username not found", http.StatusUnauthorized)
		return
	}

	created := mux.Vars(r)["created"]
	fmt.Println("username: ", username.(string), "created: ", created)

	// do a delete request to /api/notes
	// attach username from context to url
	req, err := http.NewRequest("DELETE", os.Getenv("notes-service")+"/api/notes/"+username.(string)+"/"+created, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("response body: ", string(resBody))

	w.Write(resBody)
}
