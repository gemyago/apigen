package cmd

import (
	"io/fs"
	"log/slog"
	"os"
	"os/exec"

	"github.com/gemyago/apigen/lang/go/apigen/internal/cmd/resources"
	"github.com/spf13/cobra"
)

const expectedArgsCount = 2

func NewRootCmd() *cobra.Command {
	var params GeneratorParams
	verbose := false

	var generator Generator

	cmd := &cobra.Command{
		Use:   "apigengo [input] [output]",
		Short: "Generate HTTP layer from OpenAPI spec",
		Args:  cobra.ExactArgs(expectedArgsCount),
		PreRun: func(_ *cobra.Command, _ []string) {
			level := slog.LevelWarn
			if verbose {
				level = slog.LevelDebug
			}
			rootLogger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
				Level: level,
			}))

			generator = NewGenerator(GeneratorDeps{
				RootLogger: rootLogger,
				SupportFilesInstaller: NewSupportFilesInstaller(SupportFilesInstallerDeps{
					RootLogger: rootLogger,
					Downloader: NewResourceDownloader(),
					RootFS:     os.DirFS("/").(fs.ReadFileFS), //TODO: without the cast?
				}),
				MetadataReader: resources.NewMetadataReader(),
				GeneratorInvoker: NewGeneratorInvoker(GeneratorInvokerDeps{
					RootLogger: rootLogger,
					StdOut:     os.Stdout,
					StdErr:     os.Stderr,
					OsExecutableCmdFactoryFunc: func(name string, arg ...string) OsExecutableCmd {
						return exec.Command(name, arg...)
					},
				}),
			})
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			params.input = args[0]
			params.output = args[1]

			return generator(cmd.Context(), params)
		},
	}

	pFlags := cmd.PersistentFlags()
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
	return cmd
}
