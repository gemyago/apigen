package cmd

import (
	"context"
	"errors"
	"io/fs"
	"log/slog"
	"os"
	"os/exec"

	"github.com/gemyago/apigen/lang/go/apigen/internal/cmd/resources"
	"github.com/spf13/cobra"
)

const expectedArgsCount = 2

func newNoopableSupportFilesInstaller(
	rootLogger *slog.Logger,
	noop bool,
	installer SupportFilesInstaller,
) SupportFilesInstaller {
	if noop {
		return func(ctx context.Context, params SupportFilesInstallerParams) (SupportingFilesInstallResult, error) {
			rootLogger.InfoContext(ctx, "NOOP: Installing support files", slog.Any("params", params))
			return SupportingFilesInstallResult{}, nil
		}
	}
	return installer
}

func newNoopableGeneratorInvoker(
	rootLogger *slog.Logger,
	noop bool,
	invoker GeneratorInvoker,
) GeneratorInvoker {
	if noop {
		return func(ctx context.Context, params GeneratorInvokerParams) error {
			rootLogger.InfoContext(ctx, "NOOP: Invoking generator", slog.Any("params", params))
			return nil
		}
	}
	return invoker
}

func NewRootCmd(embeddedRootFS fs.ReadFileFS) *cobra.Command {
	var params GeneratorParams
	verbose := false
	noop := false
	jsonLogs := false

	var generator Generator

	cmd := &cobra.Command{
		Use:   "apigengo [input] [output]",
		Short: "Generate HTTP layer from OpenAPI spec",
		Args:  cobra.ExactArgs(expectedArgsCount),
		PreRunE: func(_ *cobra.Command, _ []string) error {
			level := slog.LevelWarn
			if verbose {
				level = slog.LevelDebug
			}

			var logHandler slog.Handler
			if jsonLogs {
				logHandler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
					Level: level,
				})
			} else {
				logHandler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
					Level: level,
				})
			}

			rootLogger := slog.New(logHandler)

			cwdFS, ok := os.DirFS(".").(fs.ReadFileFS)
			if !ok {
				return errors.New("failed to get current working directory filesystem")
			}

			generator = NewGenerator(GeneratorDeps{
				RootLogger:  rootLogger,
				OsChdirFunc: os.Chdir,
				SupportFilesInstaller: newNoopableSupportFilesInstaller(
					rootLogger,
					noop,
					NewSupportFilesInstaller(SupportFilesInstallerDeps{
						RootLogger: rootLogger,
						Downloader: NewResourceDownloader(),
						CwdFS:      cwdFS,
					}),
				),
				MetadataReader: resources.NewMetadataReader(embeddedRootFS),
				GeneratorInvoker: newNoopableGeneratorInvoker(
					rootLogger,
					noop,
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
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			params.input = args[0]
			params.output = args[1]

			return generator(cmd.Context(), params)
		},
	}

	pFlags := cmd.PersistentFlags()
	pFlags.StringVarP(&params.chdir, "chdir", "c", "",
		"Change working directory before running the generator")
	pFlags.StringVar(&params.supportDir, "support-dir", "",
		"Path to the directory where supporting files will be placed. "+
			"Defaults to the <output>/.apigen")
	pFlags.StringVar(&params.oagCliVersion, "oag-cli-version", "", "Alternative version of openapi-generator-cli")
	pFlags.StringVar(&params.oagCliLocation, "oag-cli-location", "",
		"Alternative location of the openapi-generator-cli jar file "+
			"(can be file or url)")
	pFlags.StringVar(&params.appVersion, "app-version", "", "Alternative application version. Impacts plugins locations.")
	pFlags.StringVar(&params.serverGeneratorLocation, "server-generator-location", "",
		"Alternative location of the server generator plugin jar file "+
			"(can be file or url)")
	pFlags.BoolVar(&verbose, "verbose", verbose, "Enable verbose output")
	pFlags.BoolVar(&noop, "noop", noop, "Enable no-op mode")
	cmd.PersistentFlags().BoolVar(
		&jsonLogs,
		"json-logs",
		false,
		"Indicates if logs should be in JSON format or text (default)",
	)
	return cmd
}
