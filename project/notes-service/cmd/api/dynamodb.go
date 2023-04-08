package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// parameters for dynamodb table
const (
	TableName = "notes"
	Region    = "ap-southeast-2"
)

func ListNotesFromDB(userID string) ([]Note, error) {
	session := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(Region),
	}))

	// get notes from dynamodb
	connection := dynamodb.New(session)

	// get all notes with userID
	params := &dynamodb.ScanInput{
		TableName:        aws.String(TableName),
		FilterExpression: aws.String("user_id = :user_id"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":user_id": {
				S: aws.String(userID),
			},
		},
	}

	result, err := connection.Scan(params)
	if err != nil {
		fmt.Println(err.Error())
		return []Note{}, err
	}

	print(result)

	fmt.Println("ListNotesFromDB: userID:", userID)

	return []Note{}, nil
}

func CreateNoteInDB(note Note) (Note, error) {
	// create note in dynamodb
	// ...

	fmt.Println("CreateNoteInDB: title:", note.Title, "userID:", note.UserID, "content:", note.Content)

	return Note{}, nil
}

func GetNoteFromDB(noteID string) (Note, error) {
	// get note from dynamodb
	// ...

	fmt.Println("GetNoteFromDB: noteID:", noteID)

	return Note{}, nil
}

func UpdateNoteInDB(note Note) (Note, error) {
	// update note in dynamodb
	// ...

	fmt.Println("UpdateNoteInDB: noteID:", note.ID, "title:", note.Title, "content:", note.Content)

	return Note{}, nil
}
