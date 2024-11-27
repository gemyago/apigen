package main

import (
	"os"

	"github.com/gemyago/apigen/lang/go/apigen/internal/cmd"
)

func main() {
	generator := cmd.NewGenerator()
	rootCmd := cmd.NewRootCmd(generator)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
