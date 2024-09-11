package dynamodb

import (
	"fmt"
	"identity-hub/packages/formats"

	// "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	// "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	// "github.com/google/uuid"
)


var sess = session.Must(session.NewSessionWithOptions(session.Options{
	SharedConfigState: session.SharedConfigEnable,
}))

// Create DynamoDB client
var svc = dynamodb.New(sess)

var tableName = "PersonsTable"

func SavePersonInfo(person formats.PersonRequest) error {
	item, err := dynamodbattribute.MarshalMap(person)
	// item["id"] = &types.AttributeValueMemberS{Value: uuid.New().String()} 
	// item["id"] = &dynamodb.AttributeValue{Value: "1"} 

	if err != nil {
		return fmt.Errorf("error marshalling map: %s", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(tableName),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		return fmt.Errorf("error inserting Item: %s", err)
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
