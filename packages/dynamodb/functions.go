package dynamodb

import (
	"fmt"
	"identity-hub/packages/formats"
	"github.com/rs/zerolog/log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)


var sess = session.Must(session.NewSessionWithOptions(session.Options{
	SharedConfigState: session.SharedConfigEnable,
}))

var svc = dynamodb.New(sess)

var tableName = "PersonsTable"

func SavePersonInfo(person formats.PersonRequest) error {
	item, err := dynamodbattribute.MarshalMap(person)

	if err != nil {
		return fmt.Errorf("error marshalling map: %s", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(tableName),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		log.Error().Err(err).Msg("Error saving person info")
		return fmt.Errorf("error inserting Item: %s", err)
	}
	return nil
}

func GetAllPersonsInfo() ([]formats.PersonRequest, error) {
	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
	})
	if err != nil {
		log.Error().Msg("Got error calling GetItem: %s", err)
		return nil, fmt.Errorf("Error getting Items: %s", err)
	}

	if result.Item == nil {
		return nil, nil
	}
	var persons []formats.PersonRequest
	err = dynamodbattribute.UnmarshalMap(result.Item, &persons)
	if err != nil {
		log.Error().Msg("Error unmarshalling items: %s", err)
		return nil, fmt.Errorf("error unmarshalling map: %s", err)
	}
	return persons, nil
}
