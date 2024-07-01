# apigen
API Layer Generator from the OpenAPI specification.

This project is using the [OpenAPI Generator](https://github.com/openapitools/openapi-generator) and provides a set of plugins for it. Primary target is a golang, however other languages to be considered in the future.

Project status: **pre-alpha**.

## Contributing

Please have the following tools installed:
* [direnv](https://github.com/direnv/direnv)
* make
* Java 11 runtime at a minimum

Follow [Customization](https://openapi-generator.tech/docs/customization) docs of openapi generator cli.

Build particular generator:
```
mvn -f generators/src/go-apigen-server/ package
```