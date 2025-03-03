package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockResourceMetadataReader struct {
	oagCliVersion        string
	oagCliVersionReadErr error
	appVersion           string
	appVersionReadErr    error
}

func (m *mockResourceMetadataReader) ReadOpenapiGeneratorCliVersion() (string, error) {
	return m.oagCliVersion, m.oagCliVersionReadErr
}

func (m *mockResourceMetadataReader) ReadAppVersion() (string, error) {
	return m.appVersion, m.appVersionReadErr
}

func TestGenerator(t *testing.T) {
	t.Run("should install provided support files and invoke generator", func(t *testing.T) {
		params := GeneratorParams{
			input:      faker.URL(),
			output:     faker.URL(),
			supportDir: faker.URL(),

			oagCliVersion:  "1.2.3-" + faker.Word(),
			oagCliLocation: faker.URL(),

			appVersion:              "4.5.6-" + faker.Word(),
			serverGeneratorLocation: faker.URL(),
		}

		installResult := SupportingFilesInstallResult{
			OagLocation:             faker.URL(),
			ServerGeneratorLocation: faker.URL(),
		}

		installerInvoked := false
		generatorInvoked := false
		generator := NewGenerator(GeneratorDeps{
			RootLogger: TestRootLogger,
			SupportFilesInstaller: func(
				_ context.Context,
				installerParams SupportFilesInstallerParams,
			) (SupportingFilesInstallResult, error) {
				assert.Equal(t, params.supportDir, installerParams.SupportDir)
				assert.Equal(t, params.oagCliVersion, installerParams.OagSourceVersion)
				assert.Equal(t, params.oagCliLocation, installerParams.OagSourceLocation)
				assert.Equal(t, params.appVersion, installerParams.AppVersion)
				assert.Equal(t, params.serverGeneratorLocation, installerParams.ServerGeneratorSourceLocation)
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

		err := generator(t.Context(), params)
		require.NoError(t, err)
		assert.True(t, installerInvoked)
		assert.True(t, generatorInvoked)
	})

	t.Run("should chdir to a given directory", func(t *testing.T) {
		params := GeneratorParams{
			input:      faker.URL(),
			output:     faker.URL(),
			supportDir: faker.URL(),
			chdir:      faker.URL(),

			oagCliVersion:  "1.2.3-" + faker.Word(),
			oagCliLocation: faker.URL(),

			appVersion:              "4.5.6-" + faker.Word(),
			serverGeneratorLocation: faker.URL(),
		}

		installResult := SupportingFilesInstallResult{
			OagLocation:             faker.URL(),
			ServerGeneratorLocation: faker.URL(),
		}

		invocationOrder := make([]int, 3)
		generator := NewGenerator(GeneratorDeps{
			RootLogger: TestRootLogger,
			OsChdirFunc: func(dir string) error {
				assert.Equal(t, params.chdir, dir)
				invocationOrder[0] = 1
				return nil
			},
			SupportFilesInstaller: func(
				_ context.Context,
				_ SupportFilesInstallerParams,
			) (SupportingFilesInstallResult, error) {
				invocationOrder[1] = 2
				return installResult, nil
			},
			GeneratorInvoker: func(_ context.Context, _ GeneratorInvokerParams) error {
				invocationOrder[2] = 3
				return nil
			},
		})

		err := generator(t.Context(), params)
		require.NoError(t, err)
		assert.Equal(t, []int{1, 2, 3}, invocationOrder)
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
			oagCliVersion: "1-" + faker.Word(),
			appVersion:    "4-" + faker.Word(),
		}
		generator := NewGenerator(GeneratorDeps{
			RootLogger:     TestRootLogger,
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
				assert.Equal(t, mockMetadataReader.appVersion, installerParams.AppVersion)
				assert.Equal(t, fmt.Sprintf(
					"https://github.com/gemyago/apigen/releases/download/%s/server.jar",
					mockMetadataReader.appVersion,
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

		err := generator(t.Context(), params)
		require.NoError(t, err)
		assert.True(t, installerInvoked)
		assert.True(t, generatorInvoked)
	})

	t.Run("should prefix predefined app version with v if it follows semver pattern", func(t *testing.T) {
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
			oagCliVersion: "1-" + faker.Word(),
			appVersion:    "1.2.3-" + faker.Word(),
		}
		generator := NewGenerator(GeneratorDeps{
			RootLogger:     TestRootLogger,
			MetadataReader: &mockMetadataReader,
			SupportFilesInstaller: func(
				_ context.Context,
				installerParams SupportFilesInstallerParams,
			) (SupportingFilesInstallResult, error) {
				installerInvoked = true
				assert.Equal(t, "v"+mockMetadataReader.appVersion, installerParams.AppVersion)
				assert.Equal(t, fmt.Sprintf(
					"https://github.com/gemyago/apigen/releases/download/v%s/server.jar",
					mockMetadataReader.appVersion,
				), installerParams.ServerGeneratorSourceLocation)

				return installResult, nil
			},
			GeneratorInvoker: func(_ context.Context, _ GeneratorInvokerParams) error {
				generatorInvoked = true
				return nil
			},
		})

		err := generator(t.Context(), params)
		require.NoError(t, err)
		assert.True(t, installerInvoked)
		assert.True(t, generatorInvoked)
	})

	t.Run("should use project dir as support path if not provided", func(t *testing.T) {
		params := GeneratorParams{
			input:      faker.URL(),
			output:     faker.URL(),
			supportDir: "",

			oagCliVersion:  "1.2.3-" + faker.Word(),
			oagCliLocation: faker.URL(),

			appVersion:              "4.5.6-" + faker.Word(),
			serverGeneratorLocation: faker.URL(),
		}

		installResult := SupportingFilesInstallResult{
			OagLocation:             faker.URL(),
			ServerGeneratorLocation: faker.URL(),
		}

		wantProjectDir := faker.URL()
		installerInvoked := false
		generatorInvoked := false
		generator := NewGenerator(GeneratorDeps{
			RootLogger: TestRootLogger,
			projectDirLocator: func(output string) (string, error) {
				assert.Equal(t, params.output, output)
				return wantProjectDir, nil
			},
			SupportFilesInstaller: func(
				_ context.Context,
				installerParams SupportFilesInstallerParams,
			) (SupportingFilesInstallResult, error) {
				installerInvoked = true
				assert.Equal(t, path.Join(wantProjectDir, ".apigen"), installerParams.SupportDir)
				return installResult, nil
			},
			GeneratorInvoker: func(_ context.Context, _ GeneratorInvokerParams) error {
				generatorInvoked = true
				return nil
			},
		})

		err := generator(t.Context(), params)
		require.NoError(t, err)
		assert.True(t, installerInvoked)
		assert.True(t, generatorInvoked)
	})

	t.Run("should return error if failed to read openapi-generator-cli version", func(t *testing.T) {
		params := GeneratorParams{
			input:      faker.URL(),
			output:     faker.URL(),
			supportDir: faker.URL(),
		}

		wantErr := errors.New(faker.Sentence())
		mockMetadataReader := mockResourceMetadataReader{
			oagCliVersionReadErr: wantErr,
		}
		generator := NewGenerator(GeneratorDeps{
			RootLogger:     TestRootLogger,
			MetadataReader: &mockMetadataReader,
		})

		err := generator(t.Context(), params)
		require.ErrorIs(t, err, wantErr)
	})

	t.Run("should return error if failed to read app version", func(t *testing.T) {
		params := GeneratorParams{
			input:      faker.URL(),
			output:     faker.URL(),
			supportDir: faker.URL(),
		}

		wantErr := errors.New(faker.Sentence())
		mockMetadataReader := mockResourceMetadataReader{
			appVersionReadErr: wantErr,
		}
		generator := NewGenerator(GeneratorDeps{
			RootLogger:     TestRootLogger,
			MetadataReader: &mockMetadataReader,
		})

		err := generator(t.Context(), params)
		require.ErrorIs(t, err, wantErr)
	})

	t.Run("should return error if failed to install support files", func(t *testing.T) {
		params := GeneratorParams{
			input:      faker.URL(),
			output:     faker.URL(),
			supportDir: faker.URL(),
		}

		wantErr := errors.New(faker.Sentence())
		generator := NewGenerator(GeneratorDeps{
			RootLogger:     TestRootLogger,
			MetadataReader: &mockResourceMetadataReader{},
			SupportFilesInstaller: func(_ context.Context, _ SupportFilesInstallerParams) (SupportingFilesInstallResult, error) {
				return SupportingFilesInstallResult{}, wantErr
			},
		})

		err := generator(t.Context(), params)
		require.ErrorIs(t, err, wantErr)
	})

	t.Run("should return error if failed to invoke generator", func(t *testing.T) {
		params := GeneratorParams{
			input:      faker.URL(),
			output:     faker.URL(),
			supportDir: faker.URL(),
		}

		wantErr := errors.New(faker.Sentence())
		generator := NewGenerator(GeneratorDeps{
			RootLogger:     TestRootLogger,
			MetadataReader: &mockResourceMetadataReader{},
			SupportFilesInstaller: func(_ context.Context, _ SupportFilesInstallerParams) (SupportingFilesInstallResult, error) {
				return SupportingFilesInstallResult{}, nil
			},
			GeneratorInvoker: func(_ context.Context, _ GeneratorInvokerParams) error {
				return wantErr
			},
		})

		err := generator(t.Context(), params)
		require.ErrorIs(t, err, wantErr)
	})
}

func Test_locateDefaultSupportDir(t *testing.T) {
	createGoMod := func(dir string) error {
		return os.WriteFile(path.Join(dir, "go.mod"), []byte("module test\n"), 0600)
	}

	t.Run("should return output dir if it has go.mod", func(t *testing.T) {
		projectDir := t.TempDir()
		require.NoError(t, createGoMod(projectDir))

		defaultSupportDir, err := locateProjectDir(projectDir)
		require.NoError(t, err)
		assert.Equal(t, projectDir, defaultSupportDir)
	})

	t.Run("should return parent dir if it has go.mod", func(t *testing.T) {
		projectDir := t.TempDir()
		require.NoError(t, createGoMod(projectDir))

		subDir := path.Join(projectDir, faker.Word())
		require.NoError(t, os.Mkdir(subDir, 0755))

		defaultSupportDir, err := locateProjectDir(subDir)
		require.NoError(t, err)
		assert.Equal(t, projectDir, defaultSupportDir)
	})

	t.Run("should return nearest parent dir if it has go.mod", func(t *testing.T) {
		projectDir := t.TempDir()
		require.NoError(t, createGoMod(projectDir))

		lvl1 := path.Join(projectDir, faker.Word())
		require.NoError(t, os.Mkdir(lvl1, 0755))

		lvl2 := path.Join(lvl1, faker.Word())
		require.NoError(t, os.Mkdir(lvl2, 0755))

		lvl3 := path.Join(lvl2, faker.Word())
		require.NoError(t, os.Mkdir(lvl3, 0755))

		defaultSupportDir, err := locateProjectDir(lvl3)
		require.NoError(t, err)
		assert.Equal(t, projectDir, defaultSupportDir)
	})

	t.Run("should fail if no parent dir has go.mod", func(t *testing.T) {
		dir := t.TempDir()

		defaultSupportDir, err := locateProjectDir(dir)
		require.Error(t, err)
		assert.Empty(t, defaultSupportDir)
	})
}
