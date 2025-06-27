package main

import (
	"fmt"
	"log/slog"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") != "" {
		lambda.Start(LogSlogInMemoryTest)
	} else {
		LogSlogExample()
	}
}

type Process struct {
	Name string
	log  *slog.Logger
}

func NewProcess() *Process {
	name := RandomStr()
	slog.SetDefault(slog.Default().With(
		slog.String("name", name),
		slog.String("timestamp", time.Now().Format(time.RFC3339)),
	))
	return &Process{
		Name: name,
		log:  slog.Default(),
	}
}

func RandomStr() string {
	slog.Info("generating random string...")
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length := 10
	bytes := make([]byte, length)
	for i := range bytes {
		bytes[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(bytes)
}

func (p *Process) ReceiveRequest() {
	p.log.Info("request received")
	previousLog := p.log
	for i := range 2 {
		subLog := p.log.With(
			slog.String("iteration", fmt.Sprintf("%d", i+1)),
			slog.String("child_name", RandomStr()),
		)
		p.log = subLog
		if err := p.DoSomethingAndLogIt(); err != nil {
			p.log.Error(fmt.Sprintf("error in child: %s", err.Error()))
		}
	}
	p.log = previousLog // restore the previous log context

	// if i get outside the loop, this may be being logged, so
	// it's a great idea to keep logger in struct
	slog.Info("request processed successfully")
}

func (p *Process) DoSomethingAndLogIt() error {
	p.log.Info("no probs here, all good")
	return nil
}

func LogSlogInMemoryTest() (events.APIGatewayProxyResponse, error) {
	slog.SetDefault(NewJSONHandler())
	process := NewProcess()
	process.ReceiveRequest()
	return events.APIGatewayProxyResponse{StatusCode: http.StatusOK}, nil
}

func LogSlogExample() (events.APIGatewayProxyResponse, error) {
	// logger := NewTextHandler()
	// logger := NewJSONHandler()
	logger := NewJSONHandlerWithOptions()

	// Set a logger variable to be used by the slog package
	slog.SetDefault(logger)

	// Levels
	PrintLevels()

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
