package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const expectedArgsCount = 2

func NewRootCmd(generator Generator) *cobra.Command {
	var params GeneratorParams

	cmd := &cobra.Command{
		Use:   "apigengo [input] [output]",
		Short: "Generate HTTP layer from OpenAPI spec",
		Args:  cobra.ExactArgs(expectedArgsCount),
		RunE: func(cmd *cobra.Command, args []string) error {
			params.input = args[0]
			params.output = args[1]

			if err := generator(cmd.Context(), params); err != nil {
				return fmt.Errorf("generator failed: %w", err)
			}
			return nil
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
	pFlags.StringVar(&params.generatorVersion, "generator-version", "", "Alternative version of the generator plugin")
	pFlags.StringVar(&params.generatorLocation, "generator-location", "",
		"Alternative location of the generator plugin jar file "+
			"(can be file or url)")
	return cmd
}
