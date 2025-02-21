package cmd

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
)

func openTestLogFile() *os.File {
	_, filename, _, _ := runtime.Caller(0) // Will be current file
	testFilePath := filepath.Join(filename, "..", "..", "..", "..", "..", "..", "golang-test.log")
	f, err := os.OpenFile(testFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o666)
	if err != nil {
		err = fmt.Errorf("fail to open log file %q for test logging: %w", testFilePath, err)
		panic(err)
	}
	return f
}

//nolint:gochecknoglobals // used in tests only so it's ok to have it global
var TestRootLogger = slog.New(slog.NewJSONHandler(openTestLogFile(), &slog.HandlerOptions{
	Level: slog.LevelDebug,
}))
