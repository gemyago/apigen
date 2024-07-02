tmp=./tmp

bin=bin
tmp=tmp

cli_version=$(shell sed -n -r 's/OPENAPI_GENERATOR_CLI: (.+)/\1/p' .versions)
cli_url=https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/$(cli_version)/openapi-generator-cli-$(cli_version).jar
cli_jar=bin/openapi-generator-cli-$(cli_version).jar

maven_version=$(shell sed -n -r 's/MAVEN: (.+)/\1/p' .versions)
maven_dir_name=apache-maven-$(maven_version)
maven_url=https://dlcdn.apache.org/maven/maven-3/$(maven_version)/binaries/$(maven_dir_name)-bin.tar.gz
maven_archive=$(tmp)/$(maven_dir_name)-bin.tar.gz
mvn=$(bin)/apache-maven/bin/mvn

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

.PHONY: cli
cli:
		@echo java -jar $(cli_jar)

.PHONY: clean
clean:
	rm -r -f $(tmp) $(bin)

generators/go-apigen-server: $(shell find generators/go-apigen-server/src/main -type f)
	mvn -f generators/go-apigen-server/pom.xml package
	touch $@

examples/go-apigen-server: generators/go-apigen-server
	java -cp $(cli_jar):generators/go-apigen-server/target/go-apigen-server-openapi-generator-0.0.1.jar \
		org.openapitools.codegen.OpenAPIGenerator generate \
		-g go-apigen-server \
		-i examples/petstore.yaml \
		-o examples/go-apigen-server
	touch $@