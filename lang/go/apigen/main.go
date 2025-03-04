package apigen

import (
	"io/fs"
	"os"

	"github.com/gemyago/apigen/lang/go/apigen/internal/cmd"
)

func Execute(embeddedRootFS fs.ReadFileFS) { // coverage-ignore
	rootCmd := cmd.WireUpCommands(embeddedRootFS)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
