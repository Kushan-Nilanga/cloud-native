package main

import (
	"encoding/json"
	"log"
	"net/http"

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

func SearchNotes(w http.ResponseWriter, r *http.Request) {
	// get user id from url
	userID := mux.Vars(r)["userid"]
	query := mux.Vars(r)["query"]

	// query dynamodb to get all notes for user where content contains query or title contains query
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
		FilterExpression: aws.String("contains(content, :query) OR contains(title, :query)"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":query": {
				S: aws.String(query),
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
			UserID:  *i["userid"].S,
			Created: *i["created"].S,
			Updated: *i["updated"].S,
			Title:   *i["title"].S,
		})
	}

	// send json response
	json.NewEncoder(w).Encode(notes)
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

	// gorilla mux
	r := mux.NewRouter()

	// routes
	r.HandleFunc("/api/search/{userid}/{query}", SearchNotes).Methods("GET")

	// start server
	http.ListenAndServe(":8080", r)
}
