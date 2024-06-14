package main

import (
	"log/slog"
	"net/http"
	"os"
	"strings"

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
	logger := slog.New(
		slog.NewJSONHandler(
			os.Stdout, // acho q da pra usar o datalake aqui ou fazer log em batch preencheno um buffer
			&slog.HandlerOptions{
				AddSource: false, // mto pika
				Level:     slog.LevelDebug,
			},
		).WithAttrs(
			[]slog.Attr{
				slog.String("app", "myapp"),
			},
		),
	)

	slog.SetDefault(logger)
	slog.With(slog.String("specific", "info-with-with")).Info("Hello, World!")

	// adds attributes to the logger

	// slog.Default().With(slog.String("default", "value")) // not works

	// slog.SetDefault(slog.Default().With(slog.String("default", "value"))) // works

	// logger.With(slog.String("default", "value"))
	// logger.Info("test") // not works

	fn()
	level()
	withType()

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
	}, nil
}

func fn() {
	slog.Info(
		"Hello, World!",
		slog.String("key", "value"),
		slog.Group("OS Info",
			slog.String("arch", "x86_64"),
			slog.String("os", "linux"),
		))

	slog.Error("ayoo!", "error", "something went wrong")

}

func level() {
	customLevel := new(slog.LevelVar)
	customLevel.Set(slog.LevelDebug)

	l := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: customLevel,
	}))

	// hmm, estranho
	l.Debug("hmmmmmmmm")
	l.Info("hmmmmmmmm")
}

type MyType struct {
	Name     string  `json:"name"`
	Number   int     `json:"number"`
	Flag     bool    `json:"flag"`
	Floating float64 `json:"floating"`
	Emoji    Emoji   `json:"emoji"`
	Hidden   string  `json:"hidden"`
}

type Emoji struct {
	emoji string
}

func (e Emoji) JustOneMore() (all string) {
	return strings.Repeat(e.emoji, 5)
}

func NewSkullEmoji() Emoji {
	return Emoji{emoji: "💀"}
}

func EmojiAttr(key string, emoji Emoji) slog.Attr {
	return slog.String("emoji", emoji.JustOneMore())
}

func (m MyType) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("name", m.Name),
		slog.Int("number", m.Number),
		slog.Bool("flag", m.Flag),
		slog.Float64("floating", m.Floating),
		EmojiAttr("emoji", m.Emoji),
		slog.String("hidden", "*REDACTED*"), // that's interesting...
	)
}

func withType() {
	myType := MyType{
		Name:     "Ca Calixto lixto",
		Number:   42,
		Flag:     true,
		Floating: 3.14,
		Emoji:    NewSkullEmoji(),
		Hidden:   "his favorite number is 7 (but don't tell anyone)",
	}

	slog.Info("log custom type data", "my_type", myType)
}
