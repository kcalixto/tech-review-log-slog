package main

import (
	"log/slog"
	"net/http"
	"os"
	"runtime"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") != "" {
		lambda.Start(LogSlogExample)
	} else {
		LogSlogExample()
	}
}

func LogSlogExample() (events.APIGatewayProxyResponse, error) {
	logger := NewTextHandler()
	// logger := NewJSONHandler()
	// logger := NewJSONHandlerWithOptions()

	// Levels
	PrintLevels()

	// Set a logger variable to be used by the slog package
	slog.SetDefault(logger)

	// Grouping
	PrintWithInlineFields()
	PrintWithSlogGroup()
	PrintWithSlogWith()
	AddToDefaultLogStructure("GOOS", runtime.GOOS)

	// Redact
	PrintWithRedactedFields()

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
	}, nil
}
