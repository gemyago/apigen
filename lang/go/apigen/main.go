package main

import (
	"io/fs"
	"os"
	"os/exec"

	"github.com/gemyago/apigen/lang/go/apigen/internal/cmd"
	"github.com/gemyago/apigen/lang/go/apigen/internal/cmd/resources"
)

func main() {
	generator := cmd.NewGenerator(cmd.GeneratorDeps{
		SupportFilesInstaller: cmd.NewSupportFilesInstaller(cmd.SupportFilesInstallerDeps{
			Downloader: cmd.NewResourceDownloader(),
			RootFS:     os.DirFS("/").(fs.ReadFileFS), //TODO: without the cast?
		}),
		MetadataReader: resources.NewMetadataReader(),
		GeneratorInvoker: cmd.NewGeneratorInvoker(cmd.GeneratorInvokerDeps{
			StdOut: os.Stdout,
			StdErr: os.Stderr,
			OsExecutableCmdFactoryFunc: func(name string, arg ...string) cmd.OsExecutableCmd {
				return exec.Command(name, arg...)
			},
		}),
	})
	rootCmd := cmd.NewRootCmd(generator)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
