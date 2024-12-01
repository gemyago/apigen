package cmd

import (
	"io"
	"log/slog"
)

//nolint:gochecknoglobals // used in tests only so it's ok to have it global
var DiscardLogger = slog.New(slog.NewTextHandler(io.Discard, nil))
