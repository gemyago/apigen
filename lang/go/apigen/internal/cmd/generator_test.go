package cmd

import (
	"context"
	"fmt"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockResourceMetadataReader struct {
	oagCliVersion    string
	oagPluginVersion string
}

func (m *mockResourceMetadataReader) ReadOpenapiGeneratorCliVersion() (string, error) {
	return m.oagCliVersion, nil
}

func (m *mockResourceMetadataReader) ReadPluginVersion() (string, error) {
	return m.oagPluginVersion, nil
}

func TestGenerator(t *testing.T) {
	t.Run("should install provided support files and invoke generator", func(t *testing.T) {
		params := GeneratorParams{
			input:      faker.URL(),
			output:     faker.URL(),
			supportDir: faker.URL(),

			oagCliVersion:  "1.2.3-" + faker.Word(),
			oagCliLocation: faker.URL(),

			generatorVersion:  "4.5.6-" + faker.Word(),
			generatorLocation: faker.URL(),
		}

		installResult := SupportingFilesInstallResult{
			OagLocation:             faker.URL(),
			ServerGeneratorLocation: faker.URL(),
		}

		installerInvoked := false
		generatorInvoked := false
		generator := NewGenerator(GeneratorDeps{
			RootLogger: DiscardLogger,
			SupportFilesInstaller: func(
				_ context.Context,
				installerParams SupportFilesInstallerParams,
			) (SupportingFilesInstallResult, error) {
				assert.Equal(t, params.supportDir, installerParams.SupportDir)
				assert.Equal(t, params.oagCliVersion, installerParams.OagSourceVersion)
				assert.Equal(t, params.oagCliLocation, installerParams.OagSourceLocation)
				assert.Equal(t, params.generatorVersion, installerParams.ServerGeneratorSourceVersion)
				assert.Equal(t, params.generatorLocation, installerParams.ServerGeneratorSourceLocation)
				installerInvoked = true

				return installResult, nil
			},
			GeneratorInvoker: func(_ context.Context, invokerParams GeneratorInvokerParams) error {
				assert.Equal(t, params.input, invokerParams.Input)
				assert.Equal(t, params.output, invokerParams.Output)
				assert.Equal(t, installResult.OagLocation, invokerParams.OagCliLocation)
				assert.Equal(t, installResult.ServerGeneratorLocation, invokerParams.GeneratorLocation)
				generatorInvoked = true
				return nil
			},
		})

		err := generator(context.Background(), params)
		require.NoError(t, err)
		assert.True(t, installerInvoked)
		assert.True(t, generatorInvoked)
	})

	t.Run("should use predefined versions and locations if not provided", func(t *testing.T) {
		params := GeneratorParams{
			input:      faker.URL(),
			output:     faker.URL(),
			supportDir: faker.URL(),
		}

		installResult := SupportingFilesInstallResult{
			OagLocation:             faker.URL(),
			ServerGeneratorLocation: faker.URL(),
		}

		installerInvoked := false
		generatorInvoked := false

		mockMetadataReader := mockResourceMetadataReader{
			oagCliVersion:    "1.2.3-" + faker.Word(),
			oagPluginVersion: "4.5.6-" + faker.Word(),
		}
		generator := NewGenerator(GeneratorDeps{
			RootLogger:     DiscardLogger,
			MetadataReader: &mockMetadataReader,
			SupportFilesInstaller: func(
				_ context.Context,
				installerParams SupportFilesInstallerParams,
			) (SupportingFilesInstallResult, error) {
				installerInvoked = true
				assert.Equal(t, params.supportDir, installerParams.SupportDir)
				assert.Equal(t, mockMetadataReader.oagCliVersion, installerParams.OagSourceVersion)
				assert.Equal(t, fmt.Sprintf(
					"https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/%s/openapi-generator-cli-%s.jar",
					mockMetadataReader.oagCliVersion,
					mockMetadataReader.oagCliVersion,
				), installerParams.OagSourceLocation)
				assert.Equal(t, mockMetadataReader.oagPluginVersion, installerParams.ServerGeneratorSourceVersion)
				assert.Equal(t, fmt.Sprintf(
					"https://github.com/gemyago/apigen/releases/download/%s/server.jar",
					mockMetadataReader.oagPluginVersion,
				), installerParams.ServerGeneratorSourceLocation)

				return installResult, nil
			},
			GeneratorInvoker: func(_ context.Context, invokerParams GeneratorInvokerParams) error {
				generatorInvoked = true
				assert.Equal(t, params.input, invokerParams.Input)
				assert.Equal(t, params.output, invokerParams.Output)
				assert.Equal(t, installResult.OagLocation, invokerParams.OagCliLocation)
				assert.Equal(t, installResult.ServerGeneratorLocation, invokerParams.GeneratorLocation)
				return nil
			},
		})

		err := generator(context.Background(), params)
		require.NoError(t, err)
		assert.True(t, installerInvoked)
		assert.True(t, generatorInvoked)
	})

	t.Run("should use output relative support path if not provided", func(t *testing.T) {
		params := GeneratorParams{
			input:      faker.URL(),
			output:     faker.URL(),
			supportDir: "",

			oagCliVersion:  "1.2.3-" + faker.Word(),
			oagCliLocation: faker.URL(),

			generatorVersion:  "4.5.6-" + faker.Word(),
			generatorLocation: faker.URL(),
		}

		installResult := SupportingFilesInstallResult{
			OagLocation:             faker.URL(),
			ServerGeneratorLocation: faker.URL(),
		}

		installerInvoked := false
		generatorInvoked := false
		generator := NewGenerator(GeneratorDeps{
			RootLogger: DiscardLogger,
			SupportFilesInstaller: func(
				_ context.Context,
				installerParams SupportFilesInstallerParams,
			) (SupportingFilesInstallResult, error) {
				installerInvoked = true
				assert.Equal(t, params.output+"/.apigen", installerParams.SupportDir)
				return installResult, nil
			},
			GeneratorInvoker: func(_ context.Context, _ GeneratorInvokerParams) error {
				generatorInvoked = true
				return nil
			},
		})

		err := generator(context.Background(), params)
		require.NoError(t, err)
		assert.True(t, installerInvoked)
		assert.True(t, generatorInvoked)
	})
}
