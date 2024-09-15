Identity-Hub
============

[![Build Status](https://github.com/FranBil/Identity-Hub/actions/workflows/deploy.yml/badge.svg?branch=main)]

Identity-Hub is a serverless application built with GoLang and TypeScript using [Serverless Stack](https://sst.dev) (SST). The service exposes endpoints to manage Person Information.

## Project structure

![Architecture diagram](./IHUB-1.jpg)

## Folder Structure

The following directories are the most important ones in the project:

- `packages/lambda`: Contains the Lambda function that is triggered by API Gateway.
- `packages/dynamodb`: Contains DB functions.
- `packages/formats`: Contains the data types and functions for validation.
- `stacks`: Contains the CDK stacks that define the infrastructure of the application.

## Setup

In order to run the project, you can run:

- Install the dependencies with: `npm install`.
- Setup AWS credentials with: `aws sso login`.

## Testing 
To run the tests, you can use the following command:

- ```curl -X POST https://hygomisjoi.execute-api.eu-west-1.amazonaws.com/v1/persons \
  -H "Content-Type: application/json" \
  -d '{
    "firstName": "John",
    "lastName": "Doe",
    "phoneNumber": "1234567890",
    "address": "123 Main St"
  }'```

- `curl -X GET https://hygomisjoi.execute-api.eu-west-1.amazonaws.com/v1/persons`