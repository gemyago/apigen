# apigen

[![Test](https://github.com/gemyago/apigen/actions/workflows/test.yml/badge.svg)](https://github.com/gemyago/apigen/actions/workflows/test.yml)

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

Some language specific features may be challenging (if possible) to implement correctly. The [Language specific caveats](#language-specific-caveats) summarises various implementation details.

###  Data types

#### All types
|type|required|nullable|
|----|----|----|
|string|&check;|&check;|
|number/integer|&check;|&check;|
|boolean|&check;|&check;|
|object|&check;|&check;|

#### Strings
|format|in|minLength|maxLength|pattern|
|----|----|----|----|----|
|none or custom|query,path,body|&check;|&check;|&check;|
|date|query,path,body|-|-|-|
|date-time|query,path,body|-|-|-|
|byte|query,path,body|&check;|&check;|-|

#### Numeric data types
|type|format|in|minimum|maximum|
|----|----|----|----|----|
|number|-|query,path|&check;|&check;|
|number|float|query,path,body|&check;|&check;|
|number|double|query,path,body|&check;|&check;|
|integer|-|query,path,body|&check;|&check;|
|integer|int32|query,path,body|&check;|&check;|
|integer|int64|query,path,body|&check;|&check;|

#### Boolean
|type|in|supported?|
|----|----|----|
|boolean|query,path,body|&check;|

#### Object

Objects are only supported in request body of `application/json` content type.

### Language specific caveats

Golang:
* `date` type in request **body** is parsed as time.Time RFC3339Nano
* `required` in request **body** has the following constraints:
  * The `required` check on booleans in request body is not performed
  * For simple data types - will validate if the field is non default
  * For objects - optional and nullable are synonyms. This means that required validation will only work for non nullable fields.

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