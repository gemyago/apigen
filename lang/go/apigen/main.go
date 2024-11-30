package main

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/gemyago/apigen/lang/go/apigen/internal/cmd"
)

func main() {
	bi, ok := debug.ReadBuildInfo()
	fmt.Println("bi", bi, "ok", ok)

	rootCmd := cmd.NewRootCmd()
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
