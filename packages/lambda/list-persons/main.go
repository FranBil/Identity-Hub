package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"identity-hub/packages/dynamodb"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type response events.APIGatewayProxyResponse

func handler(request events.APIGatewayV2HTTPRequest) (response, error) {
	items, err := dynamodb.GetAllPersonsInfo()
	if err != nil {
		log.Error().Err(err).Msg("Error Getting Persons")
		return response{StatusCode: 500}, err
	}

	log.Info().Msg("Successfully fetched Persons: " + fmt.Sprint(items))
	return response{
		Body:       fmt.Sprint(items),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
