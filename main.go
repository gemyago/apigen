package main

import (
	"embed"

	"github.com/gemyago/apigen/lang/go/apigen"
)

//go:embed .versions
var embeddedRootFS embed.FS

func main() { // coverage-ignore
	apigen.Execute(embeddedRootFS)
}
