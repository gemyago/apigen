package apigen

import (
	"io/fs"
	"os"

	"github.com/gemyago/apigen/lang/go/apigen/internal/cmd"
)

func Execute(embeddedRootFS fs.ReadFileFS) { // coverage-ignore
	commonParams := cmd.CommandsCommonParams{}
	rootCmd := cmd.NewRootCmd(&commonParams)
	rootCmd.AddCommand(
		cmd.NewGenerateServerCmd(commonParams, embeddedRootFS),
	)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
