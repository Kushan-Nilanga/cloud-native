package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gorilla/mux"
)

var (
	svc *dynamodb.DynamoDB
)

// notes struct
type Note struct {
	Title   string `json:"title"`
	UserID  string `json:"userid,omitempty"`
	Content string `json:"content,omitempty"`
	Created string `json:"created"`
	Updated string `json:"updated"`
}

// this function will send dynamodb query to get all notes for a user
func ListNotes(w http.ResponseWriter, r *http.Request) {
	// get user id from url
	userID := mux.Vars(r)["userid"]

	// query dynamodb to get all notes for user
	// query result should only contain userid, created, updated and title
	result, err := svc.Query(&dynamodb.QueryInput{
		TableName: aws.String("notes"),
		KeyConditions: map[string]*dynamodb.Condition{
			"userid": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(userID),
					},
				},
			},
		},
		ProjectionExpression: aws.String("userid, created, updated, title"),
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	// create a slice of notes
	notes := []Note{}

	// loop through dynamodb result and append to notes slice
	for _, i := range result.Items {
		notes = append(notes, Note{
			Title:   *i["title"].S,
			Created: *i["created"].S,
			Updated: *i["updated"].S,
		})
	}

	// send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}

// this function will send dynamodb query to create a new note
type CreateNoteRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func CreateNote(w http.ResponseWriter, r *http.Request) {
	// get user id from url
	userID := mux.Vars(r)["userid"]

	// get request body
	var req CreateNoteRequest
	json.NewDecoder(r.Body).Decode(&req)

	// create a new note
	// send dynamodb query to create a new note
	result, err := svc.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String("notes"),
		Item: map[string]*dynamodb.AttributeValue{
			"userid": {
				S: aws.String(userID),
			},
			"title": {
				S: aws.String(req.Title),
			},
			"content": {
				S: aws.String(req.Content),
			},
			"updated": {
				S: aws.String(time.Now().Format("2006-01-02-15:04:05")),
			},
			"created": {
				S: aws.String(time.Now().Format("2006-01-02-15:04:05")),
			},
		},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	// send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// this function will send dynamodb query to get a note
func GetNote(w http.ResponseWriter, r *http.Request) {
	// get user id and created from url
	userID := mux.Vars(r)["userid"]
	created := mux.Vars(r)["created"]

	// query dynamodb to get a note
	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("notes"),
		Key: map[string]*dynamodb.AttributeValue{
			"userid": {
				S: aws.String(userID),
			},
			"created": {
				S: aws.String(created),
			},
		},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	// create a note
	note := Note{
		Title:   *result.Item["title"].S,
		Content: *result.Item["content"].S,
		Created: *result.Item["created"].S,
		Updated: *result.Item["updated"].S,
	}

	// send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

// this function will send dynamodb query to update a note
type UpdateNoteRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func UpdateNote(w http.ResponseWriter, r *http.Request) {
	// get user id and created from url
	userID := mux.Vars(r)["userid"]
	created := mux.Vars(r)["created"]

	// get request body
	var req UpdateNoteRequest
	json.NewDecoder(r.Body).Decode(&req)

	// update a note
	// send dynamodb query to update a note
	result, err := svc.UpdateItem(&dynamodb.UpdateItemInput{
		TableName: aws.String("notes"),
		Key: map[string]*dynamodb.AttributeValue{
			"userid": {
				S: aws.String(userID),
			},
			"created": {
				S: aws.String(created),
			},
		},
		UpdateExpression: aws.String("set title = :t, content = :c, updated = :u"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":t": {
				S: aws.String(req.Title),
			},
			":c": {
				S: aws.String(req.Content),
			},
			":u": {
				S: aws.String(time.Now().Format("2006-01-02-15:04:05")),
			},
		},
		ReturnValues: aws.String("UPDATED_NEW"),
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	// send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// this function will send dynamodb query to delete a note
func DeleteNote(w http.ResponseWriter, r *http.Request) {
	// get user id and created from url
	userID := mux.Vars(r)["userid"]
	created := mux.Vars(r)["created"]

	// delete a note
	// send dynamodb query to delete a note
	result, err := svc.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String("notes"),
		Key: map[string]*dynamodb.AttributeValue{
			"userid": {
				S: aws.String(userID),
			},
			"created": {
				S: aws.String(created),
			},
		},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	// send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func main() {
	// create dynamodb connection
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-2")},
	)

	if err != nil {
		log.Fatal(err)
		return
	}

	// create a dynamodb client
	svc = dynamodb.New(sess)

	// new gorilla mux router
	r := mux.NewRouter()

	// routes
	r.HandleFunc("/api/notes/{userid}", ListNotes).Methods("GET")
	r.HandleFunc("/api/notes/{userid}", CreateNote).Methods("POST")
	r.HandleFunc("/api/notes/{userid}/{created}", GetNote).Methods("GET")
	r.HandleFunc("/api/notes/{userid}/{created}", UpdateNote).Methods("PUT")
	r.HandleFunc("/api/notes/{userid}/{created}", DeleteNote).Methods("DELETE")

	// start server
	log.Fatal(http.ListenAndServe(":80", r))
}
