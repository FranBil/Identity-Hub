package dynamodb

import (
	"errors"
	"fmt"
	"identity-hub/packages/formats"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func SavePersonInfoWithMock(person formats.PersonRequest, svc dynamodbiface.DynamoDBAPI) error {
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
		return fmt.Errorf("error inserting Item: %s", err)
	}

	return nil
}

func GetAllPersonsInfoWithMock(svc dynamodbiface.DynamoDBAPI) ([]formats.PersonRequest, error) {
	params := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}

	result, err := svc.Scan(params)
	if err != nil {
		return nil, fmt.Errorf("error scanning items: %s", err)
	}

	persons := []formats.PersonRequest{}
	for _, i := range result.Items {
		var pr formats.PersonRequest
		err = dynamodbattribute.UnmarshalMap(i, &pr)
		if err != nil {
			continue
		}
		persons = append(persons, pr)
	}

	return persons, nil
}

func TestSavePersonInfoSuccess(t *testing.T) {
	mockDynamoDB := new(MockDynamoDB)

	person := formats.PersonRequest{
		FirstName:   "John",
		LastName:    "Doe",
		PhoneNumber: "123456789",
		Address:     "123 Main St",
	}

	mockDynamoDB.On("PutItem", mock.AnythingOfType("*dynamodb.PutItemInput")).
		Return(&dynamodb.PutItemOutput{}, nil)

	err := SavePersonInfoWithMock(person, mockDynamoDB)

	assert.NoError(t, err)

	mockDynamoDB.AssertExpectations(t)
}

func TestSavePersonInfoErrorInPutItem(t *testing.T) {
	mockDynamoDB := new(MockDynamoDB)

	person := formats.PersonRequest{
		FirstName:   "Jane",
		LastName:    "Doe",
		PhoneNumber: "987654321",
		Address:     "456 Main St",
	}

	mockDynamoDB.On("PutItem", mock.AnythingOfType("*dynamodb.PutItemInput")).
		Return(nil, errors.New("DynamoDB PutItem error"))

	err := SavePersonInfoWithMock(person, mockDynamoDB)

	assert.Error(t, err)
	assert.Equal(t, fmt.Errorf("error inserting Item: DynamoDB PutItem error"), err)

	mockDynamoDB.AssertExpectations(t)
}

func TestGetAllPersonsInfoSuccess(t *testing.T) {
	mockDynamoDB := new(MockDynamoDB)

	person := formats.PersonRequest{
		FirstName:   "John",
		LastName:    "Doe",
		PhoneNumber: "123456789",
		Address:     "123 Main St",
	}

	item, _ := dynamodbattribute.MarshalMap(person)
	mockDynamoDB.On("Scan", mock.AnythingOfType("*dynamodb.ScanInput")).
		Return(&dynamodb.ScanOutput{
			Items: []map[string]*dynamodb.AttributeValue{item},
		}, nil)

	persons, err := GetAllPersonsInfoWithMock(mockDynamoDB)

	assert.NoError(t, err)

	assert.Equal(t, 1, len(persons))
	assert.Equal(t, "John", persons[0].FirstName)

	mockDynamoDB.AssertExpectations(t)
}

func TestGetAllPersonsInfoScanError(t *testing.T) {
	mockDynamoDB := new(MockDynamoDB)

	mockDynamoDB.On("Scan", mock.AnythingOfType("*dynamodb.ScanInput")).
		Return(nil, errors.New("DynamoDB Scan error"))

	persons, err := GetAllPersonsInfoWithMock(mockDynamoDB)

	assert.Error(t, err)
	assert.Equal(t, "error scanning items: DynamoDB Scan error", err.Error())

	assert.Nil(t, persons)

	mockDynamoDB.AssertExpectations(t)
}
