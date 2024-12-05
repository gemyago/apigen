module github.com/gemyago/apigen

go 1.23

require (
	github.com/go-faker/faker/v4 v4.5.0
	github.com/samber/lo v1.47.0
	github.com/spf13/cobra v1.8.1
	github.com/stretchr/testify v1.10.0
	golang.org/x/net v0.31.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/text v0.20.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

retract (
	v1.2.3 // published while testing
	v1.2.3-rc // published while testing
)
