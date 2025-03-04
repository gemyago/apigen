package cmd

import (
	"io/fs"

	"github.com/spf13/cobra"
)

const generateServerExpectedArgsCount = 2

func newGenerateServerCmd(
	commonParams *commandsCommonParams,
	embeddedRootFS fs.ReadFileFS,
) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server [input] [output]",
		Short: "Generate server code (e.g handlers, models) from OpenAPI spec",
	}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(generateServerExpectedArgsCount)(cmd, args); err != nil {
			return err
		}
		generator := wireUpGenerator(wireUpGeneratorParams{
			jsonLogs:       commonParams.jsonLogs,
			verbose:        commonParams.verbose,
			noop:           commonParams.noop,
			embeddedRootFS: embeddedRootFS,
		})
		return generator(cmd.Context(), GeneratorParams{
			input:                   args[0],
			output:                  args[1],
			chdir:                   commonParams.chdir,
			supportDir:              commonParams.supportDir,
			oagCliVersion:           commonParams.oagCliVersion,
			oagCliLocation:          commonParams.oagCliLocation,
			appVersion:              commonParams.appVersion,
			serverGeneratorLocation: commonParams.serverGeneratorLocation,
			generatorName:           "go-apigen-server",
		})
	}
	return cmd
}
