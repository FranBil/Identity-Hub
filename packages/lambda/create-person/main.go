package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"identity-hub/packages/dynamodb"
)

type response events.APIGatewayProxyResponse

func handler(request events.APIGatewayV2HTTPRequest) (response, error) {
	var person dynamodb.PersonInfo

	err := json.Unmarshal([]byte(request.Body), &person)

	if err != nil {
		return response{
			StatusCode: 400,
			Body:  fmt.Sprintf("Invalid request body: %s", err),
		}, nil
	}

	err = dynamodb.SavePersonInfo(person)
	if err != nil {
		return response{
			StatusCode: 500,
			Body:  fmt.Sprintf("Error saving person info: %s", err),
			}, nil
			}
			return response{
				StatusCode: 200,
				Body:  "Person info saved successfully",
				}, nil
}

func main() {
	lambda.Start(handler)
}