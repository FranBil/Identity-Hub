package dynamodb

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type PersonInfo struct {
	FirstName string `json:"firstName`
	LastName string `json:"firstName`
	PhoneNumber string `json:"phoneNumber`
	Address string `json:"address`
}

var sess = session.Must(session.NewSessionWithOptions(session.Options{
    SharedConfigState: session.SharedConfigEnable,
}))

// Create DynamoDB client
var svc = dynamodb.New(sess)

 var tableName = "PersonsTable"


func SavePersonInfo(person PersonInfo) error {
	item, err := dynamodbattribute.MarshalMap(person)
	if err != nil {
		return fmt.Errorf("Error marshalling map: %s", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(tableName),
	}
	
	_, err = svc.PutItem(input)
	if err != nil {
		return fmt.Errorf("Error inserting Item: %s", err)
	}
	return nil
}

// func GetAllPersonsInfo() ([]PersonInfo, error) {
// 	result, err := svc.GetItem(&dynamodb.GetItemInput{
// 		TableName: aws.String(tableName),
// 		Key: map[string]*dynamodb.AttributeValue{
// 			"Year": {
// 				N: aws.String(movieYear),
// 			},
// 			"Title": {
// 				S: aws.String(movieName),
// 			},
// 		},
// 	})
// 	if err != nil {
// 		log.Fatalf("Got error calling GetItem: %s", err)
// 	}
// }