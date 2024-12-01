package cmd

import (
	"context"
	"io"
	"strings"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockOsExecutableCmd struct {
	runCallsCount int
	stdOut        string
	stdErr        string
}

func (m *mockOsExecutableCmd) Run() error {
	m.runCallsCount++
	return nil
}

func (m *mockOsExecutableCmd) StderrPipe() (io.ReadCloser, error) {
	return io.NopCloser(strings.NewReader(m.stdErr)), nil
}

func (m *mockOsExecutableCmd) StdoutPipe() (io.ReadCloser, error) {
	return io.NopCloser(strings.NewReader(m.stdOut)), nil
}

func TestGeneratorInvoker(t *testing.T) {
	t.Run("should invoke the generator and pipe stdout and stderr", func(t *testing.T) {
		/*
			java -cp $(cli_jar):$(golang_server_jar) \
				org.openapitools.codegen.OpenAPIGenerator generate \
				-g go-apigen-server \
				-i $< \
				-o $@
		*/
		wantStdOut := faker.Sentence()
		wantStdErr := faker.Sentence()

		oagCliLocation := faker.URL()
		generatorLocation := faker.URL()
		inputLocation := faker.URL()
		outputLocation := faker.URL()

		mockStdOut := &strings.Builder{}
		mockStdErr := &strings.Builder{}
		wantCmd := &mockOsExecutableCmd{
			stdOut: wantStdOut,
			stdErr: wantStdErr,
		}

		invoker := NewGeneratorInvoker(GeneratorInvokerDeps{
			RootLogger: DiscardLogger,
			StdOut:     mockStdOut,
			StdErr:     mockStdErr,
			OsExecutableCmdFactoryFunc: func(name string, arg ...string) OsExecutableCmd {
				assert.Equal(t, "java", name)
				assert.Equal(t, []string{
					"-cp",
					oagCliLocation + ":" + generatorLocation,
					"org.openapitools.codegen.OpenAPIGenerator",
					"generate",
					"-g",
					"go-apigen-server",
					"-i",
					inputLocation,
					"-o",
					outputLocation,
				}, arg)
				return wantCmd
			},
		})

		err := invoker(context.Background(), GeneratorInvokerParams{
			Input:             inputLocation,
			Output:            outputLocation,
			OagCliLocation:    oagCliLocation,
			GeneratorLocation: generatorLocation,
		})
		require.NoError(t, err)

		assert.Equal(t, wantStdOut, mockStdOut.String())
		assert.Equal(t, wantStdErr, mockStdErr.String())
		assert.Equal(t, 1, wantCmd.runCallsCount)
	})
}
