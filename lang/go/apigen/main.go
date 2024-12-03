package apigen

import (
	"io/fs"
	"os"

	"github.com/gemyago/apigen/lang/go/apigen/internal/cmd"
)

func Execute(rootFS fs.ReadFileFS) { // coverage-ignore
	rootCmd := cmd.NewRootCmd(rootFS)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
