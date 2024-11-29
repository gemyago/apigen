package cmd

import (
	"io"
	"log/slog"
)

func NewNullLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(io.Discard, nil))
}
