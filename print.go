package main

import "log/slog"

func PrintLevels() {
	slog.Info("Hi! I'm an info message")
	slog.Debug("Hi! I'm a debug message")
	slog.Warn("Hi! I'm a warning message")
	slog.Error("Hi! I'm an error message")
}
