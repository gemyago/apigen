package cmd

import (
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "apigengo",
		Short: "Generate HTTP layer from OpenAPI spec",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	pFlags := cmd.PersistentFlags()
	pFlags.StringP("spec", "s", "", "Path to OpenAPI spec file")
	pFlags.String("oag-cli-version", "", "Alternative version of openapi-generator-cli")
	pFlags.String("oag-cli-path", "", "Alternative path to the openapi-generator-cli jar file")
	pFlags.String("generator-version", "", "Alternative version of the generator plugin")
	pFlags.String("generator-path", "", "Alternative path of the generator plugin")
	pFlags.StringP("input", "i", "", "Path to input spec")
	pFlags.StringP("output", "o", "", "Folder to write generated files to")
	return cmd
}
