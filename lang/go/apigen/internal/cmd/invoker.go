package cmd

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os/exec"
	"sync"
)

type GeneratorInvokerParams struct {
	// Input is a path to the Input OpenAPI specification file
	Input string

	// Output is a path to the Output directory where the generated code will be placed
	Output string

	// OagCliLocation is a path to the source openapi-generator-cli jar file
	OagCliLocation string

	// GeneratorLocation is a path to the source generator plugin jar file
	GeneratorLocation string

	// GeneratorName is a name of the generator to use
	GeneratorName string

	// ExtraArgs is a list of additional arguments to pass to the generator
	ExtraArgs []string
}

type OsExecutableCmd interface {
	Run() error
	StderrPipe() (io.ReadCloser, error)
	StdoutPipe() (io.ReadCloser, error)
}

var _ OsExecutableCmd = (*exec.Cmd)(nil)

type GeneratorInvokerDeps struct {
	RootLogger                 *slog.Logger
	StdOut                     io.Writer
	StdErr                     io.Writer
	OsExecutableCmdFactoryFunc func(name string, arg ...string) OsExecutableCmd
}

type GeneratorInvoker func(ctx context.Context, params GeneratorInvokerParams) error

func NewGeneratorInvoker(
	deps GeneratorInvokerDeps,
) GeneratorInvoker {
	logger := deps.RootLogger.WithGroup("generator-invoker")
	return func(_ context.Context, params GeneratorInvokerParams) error {
		args := []string{
			"-cp",
			params.OagCliLocation + ":" + params.GeneratorLocation,
			"org.openapitools.codegen.OpenAPIGenerator",
			"generate",
			"-g",
			params.GeneratorName,
			"-i",
			params.Input,
			"-o",
			params.Output,
		}
		args = append(args, params.ExtraArgs...)
		logger.Info("Invoking java program", slog.String("args", fmt.Sprintf("%v", args)))
		cmd := deps.OsExecutableCmdFactoryFunc(
			"java",
			args...,
		)

		stdOut, err := cmd.StdoutPipe()
		if err != nil {
			return fmt.Errorf("failed to get stdout pipe: %w", err)
		}
		defer stdOut.Close()

		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			if _, pipeErr := io.Copy(deps.StdOut, stdOut); pipeErr != nil {
				fmt.Fprintf(deps.StdErr, "failed to copy stdout: %v", pipeErr)
			}
			wg.Done()
		}()

		stdErr, err := cmd.StderrPipe()
		if err != nil {
			return fmt.Errorf("failed to get stderr pipe: %w", err)
		}
		defer stdErr.Close()

		wg.Add(1)
		go func() {
			if _, pipeErr := io.Copy(deps.StdErr, stdErr); pipeErr != nil {
				fmt.Fprintf(deps.StdErr, "failed to copy stderr: %v", pipeErr)
			}
			wg.Done()
		}()

		if err = cmd.Run(); err != nil {
			return err
		}

		wg.Wait()

		return nil
	}
}
