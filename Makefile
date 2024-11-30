.SECONDEXPANSION:
.SECONDARY:

# No built-in rules and variables. Makes debugging easier.
MAKEFLAGS += --no-builtin-rules
MAKEFLAGS += --no-builtin-variables
.SUFFIXES:

UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S), Darwin)
    sed_i := sed -i ''
else
    sed_i := sed -i
endif

tmp=./tmp
bin=bin

app_version=$(shell sed -n -r 's/APP_VERSION: (.+)/\1/p' .versions)

cli_version=$(shell sed -n -r 's/OPENAPI_GENERATOR_CLI: (.+)/\1/p' .versions)
cli_url=https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/$(cli_version)/openapi-generator-cli-$(cli_version).jar
cli_jar=bin/openapi-generator-cli-$(cli_version).jar

maven_version=$(shell sed -n -r 's/MAVEN: (.+)/\1/p' .versions)
maven_dir_name=apache-maven-$(maven_version)
maven_url=https://dlcdn.apache.org/maven/maven-3/$(maven_version)/binaries/$(maven_dir_name)-bin.tar.gz
maven_archive=$(tmp)/$(maven_dir_name)-bin.tar.gz
mvn=$(bin)/apache-maven/bin/mvn

golang_tests_cover_dir=tests/golang/.cover
golang_tests_cover_profile=${golang_tests_cover_dir}/profile.out
golang_tests_cover_html=${golang_tests_cover_dir}/coverage.html

golang_server_jar=generators/go-apigen-server/target/server.jar

$(bin):
	mkdir -p $@

$(tmp):
	mkdir -p $@

$(cli_jar): $(bin)
	curl -L $(cli_url) -o $@

$(maven_archive): $(tmp)
	curl -L $(maven_url) -o $@

$(mvn): $(maven_archive) $(bin)
	rm -r -f $(tmp)/$(maven_dir_name) $(bin)/apache-maven
	tar -C $(tmp) -xf $(maven_archive)
	mv $(tmp)/$(maven_dir_name)/ $(bin)/apache-maven
	touch $@

.PHONY: deps
deps: $(cli_jar) $(mvn)
	go install github.com/mitranim/gow@latest

.PHONY: cli
cli:
		@echo java -jar $(cli_jar)

.PHONY: clean
clean:
	rm -r -f $(tmp) $(bin)

current_make := @make -f $(firstword $(MAKEFILE_LIST))	

# Collect all files that are not in .openapi-generator/FILES (& some extras)
define openapi_generator_removed_files
	$(addprefix .removed-generated-files/$*/, $(filter-out \
		$(patsubst $*/%, %, $(wildcard $*/.openapi-generator/*)) .openapi-generator-ignore $(shell cat $*/.openapi-generator/FILES), \
		$(shell find $* -type f -name "*" | sed 's#$*/##g')\
	))
endef

.removed-generated-files/%:
	rm $*

# Remove files that are no longer included in the .openapi-generator/FILES
%/.openapi-generator/REMOVED_FILES:
	$(if $(strip $(openapi_generator_removed_files)),$(current_make) $(openapi_generator_removed_files), $(NOOP))

# Update pom.xml (version) with app_version
generators/go-apigen-server/pom.xml: .versions
	mvn -f generators/go-apigen-server/pom.xml versions:set -DnewVersion=$(app_version)

$(golang_server_jar): $(shell find generators/go-apigen-server/src/main -type f)
	mvn -f generators/go-apigen-server/pom.xml package
	touch $@

examples/go-apigen-server/pkg/api/http/v1routes: $(golang_server_jar) examples/petstore.yaml
	java -cp $(cli_jar):$(golang_server_jar) \
		org.openapitools.codegen.OpenAPIGenerator generate \
		-g go-apigen-server \
		-i examples/petstore.yaml \
		-o $@
	$(current_make) $@/.openapi-generator/REMOVED_FILES
	touch $@

examples/go-apigen-server: examples/go-apigen-server/pkg/api/http/v1routes

# generatedCodeComment set to empty to allow linter to lint generated code.
tests/golang/routes: tests/openapi/openapi.yaml tests/openapi/*/*.yaml $(golang_server_jar)
	mkdir -p $@
	java -cp $(cli_jar):$(golang_server_jar) \
		org.openapitools.codegen.OpenAPIGenerator generate \
		-g go-apigen-server \
		--additional-properties generatedCodeComment="" \
		-i $< \
		-o $@
	$(current_make) $@/.openapi-generator/REMOVED_FILES
	touch $@

generate/golang: examples/go-apigen-server tests/golang/routes

bin/golangci-lint: ./.golangci-version
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s $(shell cat $^)

.PHONY: lint/golang
lint/golang: bin/golangci-lint
	cd ./tests/golang && ../../bin/golangci-lint run --config ../../.golangci.yml
	cd ./examples/go-apigen-server && ../../bin/golangci-lint run --config ../../.golangci.yml

.PHONY: lint
lint: lint/golang

go_path=$(shell go env GOPATH)
go-test-coverage=$(go_path)/bin/go-test-coverage

$(go-test-coverage):
	go install github.com/vladopajic/go-test-coverage/v2@latest

$(golang_tests_cover_dir):
	mkdir -p $(golang_tests_cover_dir)

.PHONY: tests/golang
tests/golang: $(golang_tests_cover_dir) $(go-test-coverage)
	TZ=US/Alaska go test -shuffle=on -failfast -coverpkg=./tests/golang/... -coverprofile=$(golang_tests_cover_profile) -covermode=atomic ./tests/golang/...
	go tool cover -html=$(golang_tests_cover_profile) -o $(golang_tests_cover_html)
	@echo "Test coverage report: $(shell realpath $(golang_tests_cover_html))"
	$(go-test-coverage) --badge-file-name $(golang_tests_cover_dir)/coverage.svg --config tests/golang/.testcoverage.yaml --profile $(golang_tests_cover_profile)

$(golang_tests_cover_dir)/coverage.%.blob-sha:
	@gh api \
		--method GET \
		-H "Accept: application/vnd.github+json" \
		-H "X-GitHub-Api-Version: 2022-11-28" \
		/repos/gemyago/apigen/contents/coverage/golang-coverage.$*?ref=test-artifacts \
		| jq -jr '.sha' > $@

$(golang_tests_cover_dir)/coverage.%.gh-cli-body.json: $(golang_tests_cover_dir)/coverage.% $(golang_tests_cover_dir)/coverage.%.blob-sha
	@echo "{" > $@
	@echo "\"branch\": \"test-artifacts\"," >> $@
	@printf "\"sha\": \"">> $@
	@cat $(golang_tests_cover_dir)/coverage.$*.blob-sha >> $@
	@printf "\",\n">> $@
	@echo "\"message\": \"Updating golang coverage.$*\",">> $@
	@printf "\"content\": \"">> $@
	@base64 -i $< | tr -d '\n' >> $@
	@printf "\"\n}">> $@

.PHONY: tests/golang/push-test-artifacts
tests/golang/push-test-artifacts: $(golang_tests_cover_dir)/coverage.svg.gh-cli-body.json $(golang_tests_cover_dir)/coverage.html.gh-cli-body.json
	@gh api \
		--method PUT \
		-H "Accept: application/vnd.github+json" \
		-H "X-GitHub-Api-Version: 2022-11-28" \
		/repos/gemyago/apigen/contents/coverage/golang-coverage.svg \
		--input $(golang_tests_cover_dir)/coverage.svg.gh-cli-body.json
	@gh api \
		--method PUT \
		-H "Accept: application/vnd.github+json" \
		-H "X-GitHub-Api-Version: 2022-11-28" \
		/repos/gemyago/apigen/contents/coverage/golang-coverage.html \
		--input $(golang_tests_cover_dir)/coverage.html.gh-cli-body.json

.PHONY: tests
tests: tests/golang

.PHONY: release
release: $(golang_server_jar) $(tmp)
	gh release create $(shell ./scripts/release.sh resolve-release-tag) \
		--generate-notes \
		--latest=false \
		--draft \
		$(golang_server_jar) > tmp/release-url.txt
