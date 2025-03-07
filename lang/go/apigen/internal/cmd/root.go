package cmd

import (
	"context"
	"log/slog"
	"runtime/debug"

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

type commandsCommonParams struct {
	jsonLogs                bool
	verbose                 bool
	noop                    bool
	chdir                   string
	supportDir              string
	oagCliVersion           string
	oagCliLocation          string
	appVersion              string
	serverGeneratorLocation string
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

func newRootCmd(params *commandsCommonParams) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "apigen",
		Short: "Generate HTTP layer from OpenAPI spec",
		Args:  cobra.ExactArgs(expectedArgsCount),
	}

	buildInfo, ok := debug.ReadBuildInfo()
	if ok {
		cmd.Version = buildInfo.Main.Version
		cmd.SetVersionTemplate(`{{printf "%s" .Version}}` + "\n")
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
	pFlags.StringVar(&params.appVersion, "app-version", "", "Alternative apigen version. Impacts plugins locations.")
	pFlags.StringVar(&params.serverGeneratorLocation, "server-generator-location", "",
		"Alternative location of the server generator plugin jar file "+
			"(can be file or url)")
	pFlags.BoolVar(&params.verbose, "verbose", params.verbose, "Enable verbose output")
	pFlags.BoolVar(&params.noop, "noop", params.noop, "Enable no-op mode")
	cmd.PersistentFlags().BoolVar(
		&params.jsonLogs,
		"json-logs",
		false,
		"Indicates if logs should be in JSON format or text (default)",
	)
	return cmd
}
