package dynamodb

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"identity-hub/packages/formats"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/eventbridge"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var tableName = "PersonsTable"

func SavePersonInfo(person formats.PersonRequest) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-1")},
	)

	if err != nil {
		log.Error().Err(err).Msg("Got error creating session")
	}

	svc := dynamodb.New(sess)

	item, err := dynamodbattribute.MarshalMap(person)

	if err != nil {
		log.Error().Err(err).Msg("Got error marshalling item")
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
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-1")},
	)

	if err != nil {
		log.Error().Err(err).Msg("Got error creating session")
	}

	svc := dynamodb.New(sess)
	params := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}

	result, err := svc.Scan(params)
	if err != nil {
		log.Error().Err(err).Msg("Got error calling Scan")
		return nil, fmt.Errorf("Error scanning Items: %s", err)
	}

	item := []formats.PersonRequest{}
	for _, i := range result.Items {
		var pr formats.PersonRequest

		err = dynamodbattribute.UnmarshalMap(i, &pr)

		if err != nil {
			log.Error().Err(err).Msg("Got error unmarshalling")
		}
		item = append(item, pr)

	}

	return item, nil
}

func PublishToEventBridge(eventDetail string) error {
	sess := session.Must(session.NewSession())
	svc := eventbridge.New(sess)

	input := &eventbridge.PutEventsInput{
		Entries: []*eventbridge.PutEventsRequestEntry{
			{
				EventBusName: aws.String("PersonEventBus"),
				Source:       aws.String("com.example.identity_hub"),
				DetailType:   aws.String("PersonCreated"),
				Detail:       aws.String(eventDetail),
			},
		},
	}

	// Publish the event
	result, err := svc.PutEvents(input)
	if err != nil {
		log.Error().Err(err).Msg("failed to send event")
		return fmt.Errorf("failed to send event: %v", err)
	}

	log.Info().Msgf("Event sent successfully: %v\n", result)
	return nil
}
