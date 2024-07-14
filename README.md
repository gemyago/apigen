# apigen
API Layer Generator from the OpenAPI specification.

This project is using the [OpenAPI Generator](https://github.com/openapitools/openapi-generator) and provides a set of plugins for it. Primary target is a golang, however other languages to be considered in the future.

Project status: **pre-alpha**.

## TODOs

### Golang
* authentication
* per route middleware
* allow handling as plain http handler
* nullable types
* polymorphic models
* enums

### Typescript

Set of features compatible with golang

## Supported OpenAPI features

Numeric data types
|type|format|in|minimum|maximum|
|----|----|----|----|----|
|number|-|query,path|&check;|&check;|
|number|float|query,path|&check;|&check;|
|number|double|query,path|&check;|&check;|
|integer|-|query,path|&check;|&check;|
|integer|int32|query,path|&check;|&check;|
|integer|int64|query,path|&check;|&check;|

## Contributing

Please have the following tools installed:
* [direnv](https://github.com/direnv/direnv)
* [gobrew](https://github.com/kevincobain2000/gobrew#install-or-update)
* make
* Java 11 runtime at a minimum

Review [Customization](https://openapi-generator.tech/docs/customization) docs of openapi generator cli.

Build particular generator:
```
mvn -f generators/go-apigen-server/ package
```

Generate particular generator (will build if needed):
```
make generate/golang
```

Run tests
```
make tests
```