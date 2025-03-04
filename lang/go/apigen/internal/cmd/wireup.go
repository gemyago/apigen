package cmd

import (
	"io/fs"
	"log/slog"
	"os"
	"os/exec"

	"github.com/gemyago/apigen/lang/go/apigen/internal/cmd/resources"
)

type wireUpGeneratorParams struct {
	jsonLogs       bool
	verbose        bool
	noop           bool
	embeddedRootFS fs.ReadFileFS
}

func wireUpGenerator(params wireUpGeneratorParams) Generator {
	level := slog.LevelWarn
	if params.verbose {
		level = slog.LevelDebug
	}
	var logHandler slog.Handler
	if params.jsonLogs {
		logHandler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: level,
		})
	} else {
		logHandler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: level,
		})
	}

	rootLogger := slog.New(logHandler)

	return NewGenerator(GeneratorDeps{
		RootLogger:  rootLogger,
		OsChdirFunc: os.Chdir,
		SupportFilesInstaller: newNoopableSupportFilesInstaller(
			rootLogger,
			params.noop,
			NewSupportFilesInstaller(SupportFilesInstallerDeps{
				RootLogger: rootLogger,
				Downloader: NewResourceDownloader(),
			}),
		),
		MetadataReader: resources.NewMetadataReader(params.embeddedRootFS),
		GeneratorInvoker: newNoopableGeneratorInvoker(
			rootLogger,
			params.noop,
			NewGeneratorInvoker(GeneratorInvokerDeps{
				RootLogger: rootLogger,
				StdOut:     os.Stdout,
				StdErr:     os.Stderr,
				OsExecutableCmdFactoryFunc: func(name string, arg ...string) OsExecutableCmd {
					return exec.Command(name, arg...)
				},
			}),
		),
	})
}
