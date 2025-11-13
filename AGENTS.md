# AGENTS.md

> This file is auto-generated and should be kept in sync with the project's build processes, dependencies, and conventions.
> Update this file whenever you modify the Makefile, CI workflows, or core development practices.

## Project Overview

This is a code generator that creates HTTP API layer code for Go projects from OpenAPI specifications. The project is a monorepo containing:
- `/generators/go-apigen-server` - Java-based OpenAPI Generator plugin (Maven project)
- `/lang/go/apigen` - Go CLI tool for invoking the generator
- `/examples` - Generated example projects (petstore, ping)
- `/tests/golang` - Comprehensive test suite using generated code

## Setup Commands

Install required dependencies and tools:
```bash
make deps
```

This installs all the necessary dependencies required to work locally.

## Build Commands

Build the Java-based generator plugin:
```bash
mvn -f generators/go-apigen-server/pom.xml package
```

The output JAR is at `generators/go-apigen-server/target/server.jar`.

Regenerate all example code and test routes:
```bash
make generate/golang
```

This regenerates code in `examples/` and `tests/golang/routes` using the latest generator.

## Test Commands

Run full test suite with coverage:
```bash
make tests
```

Run only generated routes tests:
```bash
make tests/golang-generated-routes
```

Run only apigen CLI tests:
```bash
make tests/golang-apigen
```

Run tests for a specific package:
```bash
cd tests/golang && go test -v ./controllers/...
```

## Linting

Run all linters:
```bash
make lint
```

This runs `golangci-lint` on:
- `lang/go/apigen`
- `tests/golang`
- All example projects

Configuration is in `.golangci.yml` with strict settings (cyclomatic complexity max 30, exhaustive switch checks, etc.).

## Code Style

- **Go Version**: 1.24+
- **Linter**: golangci-lint with comprehensive rules (see `.golangci.yml`)
- **Testing**: Use `testify/assert` for assertions
- **Error Handling**: Check all errors (errcheck enabled)
- **Type Assertions**: Must check type assertions (check-type-assertions enabled)
- **Exhaustiveness**: Switch statements and map iterations must be exhaustive

Key conventions:
- Generated code should be committed to the repository
- Test coverage targets are defined in `.testcoverage.yaml` files
- All shell commands in Makefile must be POSIX-compatible

## Development Workflow

1. Make changes to generator templates in `generators/go-apigen-server/src/main/`
2. Rebuild generator: `mvn -f generators/go-apigen-server/pom.xml package`
3. Regenerate test code: `make generate/golang`
4. Run tests: `make tests`
5. Run linter: `make lint`

For changes to the Go CLI (`lang/go/apigen`):
1. Make changes to Go source files
2. Run tests: `cd lang/go/apigen && go test ./...`
3. Test integration: `make generate/golang`

## Git Workflow

- **Main branch**: `main`
- **Feature branches**: Create from `main`, use descriptive names
- **Release branches**: Format `release/X.Y.Z`
- **Current branch**: `feature/refine`

CI runs on all PRs to `main` and executes `make lint` and `make tests`.

## Release Process

Release process is detailed in `.ai/instructions/create-release.md`. Follow this instruction if user requests to create a new release.

## Architecture Notes

**Generator Structure**:
- Java generator extends OpenAPI Generator's `AbstractGoCodegen`
- Templates use Mustache syntax
- Generator produces Go code with no runtime dependencies

**Generated Code Pattern**:
- Models in `models/` package
- Handlers in `handlers/` package  
- Controllers implement handler interfaces
- Root handler manages route registration

**Separation of Layers** (app-layer examples):
- Models generated to `internal/app/models`
- HTTP routes generated to `internal/api/http/routes`
- Use `--global-property=models` and `--global-property=apis` flags

## Key Files and Paths

- `Makefile` - All build, test, and generate commands
- `.versions` - Version specifications for dependencies
- `generators/go-apigen-server/pom.xml` - Maven project for generator
- `lang/go/apigen/main.go` - Go CLI entry point
- `.golangci.yml` - Linter configuration
- `.github/workflows/test.yml` - CI pipeline

## Common Tasks

Generate code from OpenAPI spec:
```bash
go run github.com/gemyago/apigen server ./path/to/spec.yaml ./output/dir
```

Or using the Makefile pattern for examples:
```bash
make examples/petstore-server-go
```

Update generator version:
1. Update version in `.versions` and `generators/go-apigen-server/pom.xml`
2. Run `make deps` and `make generate/golang`
3. Review changes and run tests

## Testing Philosophy

- Generated code must have extensive test coverage
- Tests validate parsing, validation, serialization of all OpenAPI types
- Integration tests use real HTTP handlers (not mocked)
- Mock external dependencies in controller tests, not generated handlers

## Security Considerations

- Generated code validates all inputs per OpenAPI spec constraints
- No code execution during repository analysis
- All dependencies are pinned in `go.mod` and `.versions`
- Coverage reports published to separate `test-artifacts` branch

### Non coding task completion protocol

Examples of non coding tasks are:
- Investigation of a problem
- Documentation updates
- Script updates

Generally anything that is not related to code changes.

The completion protocol for non coding task is established by the user.

## Coding Task Completion Protocol

This protocol is used to complete a coding task (e.g the one that resulted in changing any code files). For non coding tasks, use the non coding task completion protocol.

If you changed any code then **always** perform the completion protocol below:
1. Lint status: Run `make lint` and confirm no errors
2. Test status: Run `make test` and confirm no failures; coverage: XX.XX% (meets threshold)
3. Make sure AGENTS.md are in sync if updated commands, workflows, or architecture.

Report the result to the user. The **only** acceptable result is "All tests pass and no lint errors". Failing tests or linting errors means the task **is not complete**. Any failure in the above steps MUST be resolved prior to task completion.

Report task completion:
- Lint: no errors / fixed all errors
- Tests: all passing, coverage XX.XX%
- AGENTS.md: updated to reflect changes (if any) / no changes needed