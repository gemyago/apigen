package main

import (
	"embed"

	"github.com/gemyago/apigen/lang/go/apigen"
)

//go:embed .versions
var rootFS embed.FS

func main() { // coverage-ignore
	apigen.Execute(rootFS)
}
