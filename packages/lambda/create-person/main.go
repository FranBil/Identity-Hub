package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"identity-hub/packages/dynamodb"
	"identity-hub/packages/formats"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type response events.APIGatewayProxyResponse

func badGateway(body []byte, err error) (response, error) {
	return response{
		StatusCode: http.StatusBadGateway,
		Body:       fmt.Sprintf("Data is not serializable: %s", body),
	}, err
}

func handler(request events.APIGatewayV2HTTPRequest) (response, error) {
	var personRequest formats.PersonRequest
	// var person dynamodb.PersonInfo

	err := json.Unmarshal([]byte(request.Body), &personRequest)

	if err != nil {
		return response{
			StatusCode: 400,
			Body:       fmt.Sprintf("Invalid request body: %s", err),
		}, nil
	}

	item := formats.PersonRequest{
		FirstName:   personRequest.FirstName,
		LastName:    personRequest.LastName,
		PhoneNumber: personRequest.PhoneNumber,
		Address:     personRequest.Address,
	}

	isValid, errors := item.IsValid()
	if !isValid {
		data := []string{}
		for _, err := range errors {
			data = append(data, err.Error())
		}
		errBody, err := json.Marshal(data)
		if err != nil {
			return badGateway(errBody, err)
		}
	} else {
		err = dynamodb.SavePersonInfo(item)
		if err != nil {
			return response{
				StatusCode: 500,
				Body:       fmt.Sprintf("Error saving person info: %s", err),
			}, nil
		}
	}
	return response{
		StatusCode: 200,
		Body:       "Person info saved successfully",
	}, nil
}

func main() {
	lambda.Start(handler)
}
