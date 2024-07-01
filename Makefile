tmp=./tmp

bin=bin

cli_version=$(shell sed -r 's/OPENAPI_GENERATOR_CLI: //' .versions)
cli_url=https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/$(cli_version)/openapi-generator-cli-$(cli_version).jar

$(bin):
	mkdir -p $@

bin/openapi-generator-cli-$(cli_version).jar: $(bin)
	curl -L $(cli_url) -o $@

.PHONY: deps
deps: bin/openapi-generator-cli-$(cli_version).jar