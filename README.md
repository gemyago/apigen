# apigen - Frictionless API generator

[![Test](https://github.com/gemyago/apigen/actions/workflows/test.yml/badge.svg)](https://github.com/gemyago/apigen/actions/workflows/test.yml)
[![Golang Coverage](https://raw.githubusercontent.com/gemyago/apigen/test-artifacts/coverage/golang-coverage.svg)](https://htmlpreview.github.io/?https://raw.githubusercontent.com/gemyago/apigen/test-artifacts/coverage/golang-coverage.html)

HTTP API Layer Generator for the Go (golang) projects. Write less boilerplate code, focus on the business logic.

Features:
* OpenAPI first approach. Write the spec and generate the code.
* No runtime dependencies. Generated code is self-contained.
* No reflection. Code to parse and validate requests is fully generated.
* Framework agnostic and [http.Handler](https://pkg.go.dev/net/http#Handler) compatible. Use any http router or middleware.

## Table of Contents

- [Getting Started](#getting-started)
- [Basic Concepts](#basic-concepts)
- [Controllers in depth](#controllers-in-depth)
- [Supported OpenAPI features](#supported-openapi-features)

## Getting Starteds

Generated code requires Go 1.24 or higher.

Install `apigen` cli tool:
```bash
go install github.com/gemyago/apigen
```

Define the OpenAPI spec somewhere in your project. For example: `internal/api/http/v1routes.yaml`. You can use below as a starting point:
```yaml
openapi: "3.0.0"
info:
  version: 1.0.0
  title: Minimalistic API definition
paths:
  /ping:
    get:
      operationId: ping
      tags:
        - ping
      parameters:
        - name: message
          in: query
          required: false
          schema:
            type: string
      responses:
        '200':
          description: Request succeeded
          content:
            application/json:
              schema:
                type: object
                properties:
                  message:
                    type: string
```

Add a golang file with generation instructions. For example: `internal/api/http/v1routes.go`:
```go
//go:generate go run apigen ./v1routes.yaml ./v1routes
```

Run the generation:
```bash
go generate ./internal/api/http
```
The above will generate the code in the `internal/api/http/v1routes` folder. Commit the generated code to the repository. Some notes on generated folders structure:
* `handlers` contains required types and interfaces to handle requests.
* `models` contains data types used in the API.

Declare controller that implements the generated interface, for example:
```go
type pingController struct{}

func (c *pingController) Ping(b handlers.HandlerBuilder[
	*handlers.PingPingRequest,
	*models.Ping200Response,
]) http.Handler {
	return b.HandleWith(
		func(_ context.Context, req *handlers.PingPingRequest) (*models.Ping200Response, error) {
			message := req.Message
			if message == "" {
				message = "pong"
			}
			return &models.Ping200Response{Message: message}, nil
		},
	)
}
```

Define router adapter. For example `http.ServeMux` adapter may look like this:
```go
type httpRouter http.ServeMux

func (*httpRouter) PathValue(r *http.Request, paramName string) string {
	return r.PathValue(paramName)
}

func (router *httpRouter) HandleRoute(method, pathPattern string, h http.Handler) {
	(*http.ServeMux)(router).Handle(method+" "+pathPattern, h)
}

func (router *httpRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	(*http.ServeMux)(router).ServeHTTP(w, r)
}
```

Wire everything together:
```go
rootHandler := handlers.NewRootHandler((*httpRouter)(http.NewServeMux()))
rootHandler.RegisterPingRoutes(&pingController{})

const readHeaderTimeout = 5 * time.Second
srv := &http.Server{
	Addr:              "[::]:8080",
	ReadHeaderTimeout: readHeaderTimeout,
	Handler:           rootHandler,
}
log.Println("Starting server on port: 8080")
if err := srv.ListenAndServe(); err != nil {
	panic(err)
}
```

Fully functional example based on the above steps can be found [here](./examples/ping-server-go). More advanced example can be found [here](./examples/petstore-server-go).

## Basic Concepts

Generated code expects you to provide a controller that implements the generated interface. The controller is an adapter between the generated code and your business logic. Generated code will parse the request, validate parameters and call the corresponding controller method in a type-safe manner.

The controller is generated based on the `tags` in the OpenAPI spec. Prefer defining a single tag per operation. You can tag multiple operations with a same tag in order to group them under the same generated controller. Single OpenAPI spec can define as many tags (controllers) as needed.

## Controllers in depth

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

Get deps installed:
```
make deps
```

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

### Updating OpenAPI generator version

Specify a new version in the [.versions](.versions) and in the `pom.xml` of reach generator (`openapi-generator-version` property).

Regenerate the code:
```
make deps
make generate/golang
```

Review and commit changes. Run tests.

## Releasing

1. Create a release branch with the version number (e.g. `release/0.1.0`)
1. Update the version in the [.versions](.versions) file
1. Commit the changes and create the PR
1. Once the PR is merged, a release will be created automatically by the CI
