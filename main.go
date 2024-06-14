package main

import (
	"log/slog"
	"os"
)

func main() {
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
	Status   bool    `json:"status"`
	Floating float64 `json:"floating"`
	Hidden   string  `json:"hidden"`
}

func (m MyType) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("name", m.Name),
		slog.Int("number", m.Number),
		slog.Bool("status", m.Status),
		slog.Float64("floating", m.Floating),
		slog.String("hidden", "*REDACTED*"), // that's interesting...
	)
}

func withType() {
	myType := MyType{
		Name:     "Manuel Gomes",
		Number:   42,
		Status:   true,
		Floating: 3.14,
		Hidden:   "should_be_hidden",
	}

	slog.Info("log custom type data", "my_type", myType)
}
