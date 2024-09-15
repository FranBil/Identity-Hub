package dynamodb

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/stretchr/testify/mock"
)

type MockDynamoDB struct {
	dynamodbiface.DynamoDBAPI
	mock.Mock
}

// Mock PutItem function
func (m *MockDynamoDB) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	args := m.Called(input)
	if args.Get(0) != nil {
		return args.Get(0).(*dynamodb.PutItemOutput), args.Error(1)
	}
	return nil, args.Error(1)
}

// Mock Scan function
func (m *MockDynamoDB) Scan(input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	args := m.Called(input)
	if args.Get(0) != nil {
		return args.Get(0).(*dynamodb.ScanOutput), args.Error(1)
	}
	return nil, args.Error(1)
}
