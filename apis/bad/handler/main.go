package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (response events.APIGatewayProxyResponse, err error) {
	return response, nil
}

func main() {
	lambda.Start(Handler)
}
