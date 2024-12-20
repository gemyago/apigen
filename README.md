# apigen

[![Test](https://github.com/gemyago/apigen/actions/workflows/test.yml/badge.svg)](https://github.com/gemyago/apigen/actions/workflows/test.yml)
[![Golang Coverage](https://raw.githubusercontent.com/gemyago/apigen/test-artifacts/coverage/golang-coverage.svg)](https://htmlpreview.github.io/?https://raw.githubusercontent.com/gemyago/apigen/test-artifacts/coverage/golang-coverage.html)

API Layer Generator from the OpenAPI specification.

This project is using the [OpenAPI Generator](https://github.com/openapitools/openapi-generator) and provides a set of plugins for it. Primary target is a golang, however other languages to be considered in the future.

Project status: **alpha**.

## TODOs

### General
* Documentation on generated code structure and suggested patterns

### Golang
* enums
* simple types directly in body
* array types directly in body
* allow handling as plain http handler
* authentication
* per route middleware
* polymorphic models

### Typescript
Set of features compatible with golang

## Installation & Usage

### Golang projects

Install `apigen` cli tool:
```bash
go install github.com/gemyago/apigen
```

Define the OpenAPI spec somewhere in your project. For example: `internal/api/http/v1routes.yaml`.

Add a golang file with generation instructions. For example: `internal/api/http/v1routes.go`:
```go
//go:generate go run apigen ./v1routes.yaml ./v1routes
```

Run the generation:
```bash
go generate ./internal/api/http
```

Folders structure suggested above is not enforced and can be adjusted to your project needs.

## Supported OpenAPI features

Some language specific features may be challenging (if possible) to implement correctly. The [Language specific caveats](#language-specific-caveats) summarises various implementation details.

### Parameter Serialization

[OpenAPI Spec Reference](https://swagger.io/docs/specification/serialization/)

Supported serialization styles:
|parameter location|style|explode|example primitive type|example array type|object support|
|----|----|----|----|----|----|
|path|simple|false|`/users/5`|`/users/3,4,5`|-|
|query|form|true|`/users?id=5`|`/users?id=3&id=4&id=5`|-|

###  Data types

#### All types
|type|required|nullable|
|----|----|----|
|string|&check;|&check;|
|number/integer|&check;|&check;|
|boolean|&check;|&check;|
|object|&check;|&check;|
|array|&check;|&check;|

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

#### Objects

Objects are only supported in request body of `application/json` content type.

#### Arrays
|items type|in|minItems|maxItems|
|----|----|----|----|
|string|query,path,body|&check;|&check;|
|number|query,path,body|&check;|&check;|
|integer|query,path,body|&check;|&check;|
|boolean|query,path,body|&check;|&check;|
|object|body|&check;|&check;|


### Language specific caveats

Golang:
* `date` type in request **body** is parsed as time.Time RFC3339Nano
* `required` in request **body** has the following constraints:
  * The `required` check on booleans in request body is not performed
  * For simple data types - will validate if the field is non default
  * For objects - optional and nullable are synonyms. This means that required validation will only work for non nullable fields.
* `array` fields
  * `required` is equivalent to `minItems: 1`
  * `nullable` arrays are only supported in request body
  * nullable and optional array fields are synonyms

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

Generate code using a particular generator (will build if needed):
```
make generate/golang
```

Run tests
```
make tests
```

## Releasing

1. Create a release branch with the version number (e.g. `release/0.1.0`)
1. Update the version in the [.versions](.versions) file
1. Propagate the version changes:
    ```
    cd lang/go/apigen && go generate ./... && cd -
    ```
1. Commit the changes and create the PR
1. Once the PR is merged, a release will be created automatically by the CI
