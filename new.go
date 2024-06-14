package main

import (
	"log/slog"
	"os"
	"strings"
)

// slog.NewTextHandler(io.Writer, *slog.HandlerOptions)
//
//	type Writer interface {
//			Write(p []byte) (n int, err error)
//	}
func NewTextHandler() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))
}

// slog.NewJSONHandler(io.Writer, *slog.HandlerOptions)
//
//	type Writer interface {
//			Write(p []byte) (n int, err error)
//	}
func NewJSONHandler() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
}

// slog.NewJSONHandler(io.Writer, *slog.HandlerOptions)
func NewJSONHandlerWithOptions() *slog.Logger {
	return slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true, // -> logs the source code location, with the file name and line number
			Level: func() slog.Level {
				if os.Getenv("DEBUG") == "true" {
					return slog.LevelDebug
				}
				return slog.LevelInfo
			}(),
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				switch a.Key {
				case "msg":
					return slog.String("message", a.Value.String())
				case "level":
					return slog.String(a.Key, strings.ToLower(a.Value.String()))
				}

				return a
			},
		}).
			WithAttrs(
				[]slog.Attr{
					slog.String("app", "tech-review"),
				},
			),
	)
}
