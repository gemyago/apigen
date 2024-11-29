package cmd

import (
	"context"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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
			SupportFilesInstaller: func(
				_ context.Context,
				installerParams SupportFilesInstallerParams,
			) (SupportingFilesInstallResult, error) {
				assert.Equal(t, installerParams.SupportDir, params.supportDir)
				assert.Equal(t, installerParams.OagSourceVersion, params.oagCliVersion)
				assert.Equal(t, installerParams.OagSourceLocation, params.oagCliLocation)
				assert.Equal(t, installerParams.ServerGeneratorSourceVersion, params.generatorVersion)
				assert.Equal(t, installerParams.ServerGeneratorSourceLocation, params.generatorLocation)
				installerInvoked = true

				return installResult, nil
			},
			GeneratorInvoker: func(_ context.Context, invokerParams GeneratorInvokerParams) error {
				assert.Equal(t, invokerParams.Input, params.input)
				assert.Equal(t, invokerParams.Output, params.output)
				assert.Equal(t, invokerParams.OagCliLocation, installResult.OagLocation)
				assert.Equal(t, invokerParams.GeneratorLocation, installResult.ServerGeneratorLocation)
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
