package main

import (
	"fmt"
	"identity-hub/packages/dynamodb"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type response events.APIGatewayProxyResponse

func handler(request events.APIGatewayV2HTTPRequest) (response, error) {
	items, err := dynamodb.GetAllPersonsInfo()
	if err != nil {
		return response{StatusCode: 500}, err
	}

	return response{
		Body: fmt.Sprint(items),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
