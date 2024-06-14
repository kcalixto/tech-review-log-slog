package main

import (
	"log/slog"
	"runtime"
)

func PrintWithInlineFields() {
	slog.Info("OS Info - 1", "os", runtime.GOOS, "arch", runtime.GOARCH)
}

func PrintWithSlogGroup() {
	slog.Info("OS Info - 2", slog.Group("runtime",
		slog.String("os", runtime.GOOS),
		slog.String("arch", runtime.GOARCH),
	))
}

func PrintWithSlogWith() {
	slog.With(
		slog.Group("runtime",
			slog.String("os", runtime.GOOS),
			slog.String("arch", runtime.GOARCH),
		),
	).Info("OS Info - 3")
}

func AddToDefaultLogStructure(field, value string) {
	slog.SetDefault(
		slog.Default().With(
			slog.String(field, value),
		),
	)
}
