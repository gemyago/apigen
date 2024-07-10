# apigen
API Layer Generator from the OpenAPI specification.

This project is using the [OpenAPI Generator](https://github.com/openapitools/openapi-generator) and provides a set of plugins for it. Primary target is a golang, however other languages to be considered in the future.

Project status: **pre-alpha**.

## TODOs

* authentication
* per route middleware
* allow handling as plain http handler
* nullable types

## Supported OpenAPI features

Numeric data types
|type|format|in|minimum|maximum|
|----|----|----|----|----|
|number|-|query,path|&check;|&check;|

## Contributing

Please have the following tools installed:
* [direnv](https://github.com/direnv/direnv)
* [gobrew](https://github.com/kevincobain2000/gobrew#install-or-update)
* make
* Java 11 runtime at a minimum

Follow [Customization](https://openapi-generator.tech/docs/customization) docs of openapi generator cli.

Build particular generator:
```
mvn -f generators/go-apigen-server/ package
```