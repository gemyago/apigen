# apigen - HTTP API Layer Generator

[![Test](https://github.com/gemyago/apigen/actions/workflows/test.yml/badge.svg)](https://github.com/gemyago/apigen/actions/workflows/test.yml)
[![Golang Coverage](https://raw.githubusercontent.com/gemyago/apigen/test-artifacts/coverage/golang-coverage.svg)](https://htmlpreview.github.io/?https://raw.githubusercontent.com/gemyago/apigen/test-artifacts/coverage/golang-coverage.html)

HTTP API Layer Generator for the Go (golang) projects. Write less boilerplate code, focus on the business logic.

Features:
* OpenAPI first approach. Write the spec and generate the code.
* No runtime dependencies. Generated code is self-contained.
* No reflection. Code to parse and validate requests is fully generated.
* Framework agnostic and [http.Handler](https://pkg.go.dev/net/http#Handler) compatible.

Project status: 
* The generated code is extensively tested and **production ready**
* The generator itself is in **beta** stage. This means that **minor breaking changes** in the generated code may occur.

## Table of Contents

- [Getting Started](#getting-started)
- [Basic Concepts](#basic-concepts)
- [Controllers in depth](#controllers-in-depth)
- [Root Handler](#root-handler)
  - [Router adapter](#router-adapter)
  - [Handling errors](#handling-errors)
- [Supported OpenAPI features](#supported-openapi-features)

## Getting Started

The only runtime dependency is a **go 1.24** or higher. The generator is a plugin for the [OpenAPI Generator](https://openapi-generator.tech/) which is Java based and requires Java 11 runtime at a minimum. The java and openapi-generator are only required to generate the code. The generated code is self-contained and does not have any runtime dependencies.

To get started, install `apigen` cli tool:
```bash
go install github.com/gemyago/apigen
```

Define the OpenAPI spec somewhere in your project. For example: `internal/api/http/routes.yaml`. You can use below as a starting point:
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

Add a golang file with generation instructions. For example: `internal/api/http/routes.go`:
```go
//go:generate go run github.com/gemyago/apigen ./routes.yaml ./routes
```

Run the generation:
```bash
go generate ./internal/api/http
```
The above will generate the code in the `internal/api/http/routes` folder. Commit the generated code to the repository.

Declare controller that implements the generated interface, for example:
```go
func pingHandler(_ context.Context, params *models.PingParams) (*models.Ping200Response, error) {
	message := params.Message
	if message == "" {
		message = "pong"
	}
	return &models.Ping200Response{Message: message}, nil
}

type pingController struct{}

func (c *pingController) Ping(b handlers.HandlerBuilder[
	*models.PingParams,
	*models.Ping200Response,
]) http.Handler {
	return b.HandleWith(pingHandler)
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

Fully functional example based on the above steps can be found [here](./examples/ping-server-go). More advanced examples: 
* [petstore-server](./examples/petstore-server-go) - example with more routes
* [petstore-server-app-layer](./examples/petstore-server-app-layer-go) - same routes as `petstore-server` but with separate application layer with models and controllers generated in a separate packages.

## Basic Concepts

Generated code expects you to provide a controller that implements the generated interface. The controller is an adapter between the generated code and your business logic. Generated code will parse the request, validate parameters and call the corresponding controller method in a type-safe manner.

The controller is generated based on the `tags` in the OpenAPI spec. Prefer defining a single tag per operation. You can have multiple operations with a same tag in order to group them under the same generated controller. Single OpenAPI spec can define as many tags (controllers) as needed. Please note that operationIDs in OpenAPI are global and should be unique across the spec. Please see the [Controllers in depth](#controllers-in-depth) section for more details.

The generated code also includes so called `RootHandler`. The root handler is a bridge between your router of choice and the generated code. Please see the [Root Handler](#root-handler) section for more details.

Typically you will need to import generated code from the following packages:
* `handlers` contains controller interfaces, root handler and other components handle requests.
* `models` contains data structures corresponding to schemas defined in the OpenAPI spec.

It is possible to generate models and controllers in separate packages. This is useful when you want to keep your application layer separate from the HTTP API layer. Please see the [Separate packages for models and controllers](#separate-packages-for-models-and-controllers) section for more details.

## Controllers in depth

Controller should implement a set of methods, each corresponding to an operation in the OpenAPI spec. The method signature is as follows:
```go
func (c *PetsController) GetPetByID(
	b handlers.HandlerBuilder[*models.GetPetByIDParams, *models.PetResponse],
) http.Handler {
	// Your implementation here
}
```

The `HandlerBuilder` allows you to create an actual http.Handler that will be used to process requests. In most simplest case you can implement the method in place. However in real-world scenarios you may want to extract the implementation to a separate component and keep your controller clean and declarative. It is not required to use the `HandlerBuilder` and you can return a `http.Handler` directly if your use-case requires it, this allows you to fully bypass the generated code and handle the request processing as required.

The `HandlerBuilder` has the following methods:
* `HandleWith` - will bind your application logic to the generated code. The handler function should have the following signature:
  ```go
  func(context.Context, *models.GetPetByIDParams) (*models.PetResponse, error)
  ```
  This would usually be the most typical way to implement the controller method. You can define the handler in place however it is advised have a separate component that implements the handler function. This approach will help you to keep your controller clean and declarative.
  
* `HandleWithHTTP` - similar to the above, but allows you to access the underlying http.Request and http.ResponseWriter. The handler function should have the following signature:
  ```go
  func(http.ResponseWriter, *http.Request, *models.GetPetByIDParams)  (*models.PetResponse, error)
  ```
  This method is useful when you need to access the underlying http request and response objects. For example, when you need to set custom headers or status codes.
  
  **Notes**: 
    * You may return response that will be automatically written to the response writer.
    * The generated code will not attempt to write to the response writer if you have already written to it.
    * You may still return an error. In this case the generated code will handle the error as explained in the [Handling errors](#handling-errors) section.

Due to Go language constraints there are several variations of the `HandlerBuilder`. The variations are:
* `NoResponseHandlerBuilder` - for operations that do not return a response. The handler function should have the following signature:
  ```go
  func(context.Context, *models.DeletePetParams) error
  ```
* `NoRequestHandlerBuilder` - for operations that do not have request parameters. The handler function should have the following signature:
  ```go
  func(context.Context) (*models.PetResponse, error)
  ```
* `NoRequestNoResponseHandlerBuilder` - for operations that do not have request parameters and do not return a response. The handler function should have the following signature:
  ```go
  func(context.Context) error
  ```
The generated code is type safe so you will catch mismatches at compile time.


## Root Handler

The root handler is an adapter that allows you to attach generated routes to a router of your choice. Once initialized, the root handler is a self contained `http.Handler` and can be used in any scenario where you would use a standard http handler.

The root handler will have `Register[Controller]Routes` methods generated for each controller of your APIs. The method will accept an instance of the controller and will register all routes defined in the OpenAPI spec. Example:
```go
rootHandler := handlers.NewRootHandler(routerAdapter)

// This will register all routes tagged with "pets" tag
rootHandler.RegisterPetsRoutes(&petsController{})

// This will register all routes tagged with "users" tag
rootHandler.RegisterUsersRoutes(&usersController{})
```

### Router adapter

In order to instantiate the root handler you need to provide router adapter instance that must implement the following interface:
```go
type httpRouter interface {
	// PathValue returns a named path parameter of a given name
	PathValue(r *http.Request, paramName string) string

	// HandleRoute register a given handler function to handle given route
	HandleRoute(method, pathPattern string, h http.Handler)

	// ServeHTTP is a standard http.Handler method
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}
```
The underlying router should support named path parameters and should be able to operate with standard `http.Handler` interface.

### Handling errors

Errors can occur at various stages of the request processing. The root handler includes a default implementation that operates as follows:
* Request parsing and validation - default implementation respond with 400 status code and log the error with `warn` level.
* Controller methods (actions) execution - default implementation respond with 500 status code and log the error with `error` level.
* Response serialization - default implementation will log the error with `error` level and will attempt to set 500 status code. Setting status code is not guaranteed as the response may have already been written.

You can customize the error handling behavior as well as set a custom logger implementation using the following options when initializing the root handler:
```go
NewRootHandler(routerAdapter, 
  // Set custom logger for the root handler. The default logger is slog.Default().
  // The provided logger must be compatible with slog interface.
  WithLogger(logger),

  // Set custom error handler for parsing and validation errors
  WithParsingErrorHandler(errorHandler),

  // Set custom error handler to process action execution errors
  WithErrorHandler(errorHandler),
)
```

## Separate packages for models and controllers

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

Build the generator:
```
mvn -f generators/go-apigen-server/ package
```

Regenerate test code (will build generator if needed):
```
make generate/golang
```

Run tests
```
make tests
```

### Updating OpenAPI generator version

Specify a new version in the [.versions](.versions) and in the `pom.xml` of the generator (`openapi-generator-version` property).

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
