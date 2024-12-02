package apigen

import (
	"os"

	"github.com/gemyago/apigen/lang/go/apigen/internal/cmd"
)

func Execute() { // coverage-ignore
	rootCmd := cmd.NewRootCmd()
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
