package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"path"
	"path/filepath"
	"testing"
	"testing/fstest"

	"github.com/go-faker/faker/v4"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSupportFilesInstaller(t *testing.T) {
	makeRandomGeneratorParams := func(t *testing.T) SupportFilesInstallerParams {
		// Can not use t.TempDir() because MapFS does not support absolute paths
		// so provisioning temp dir manually
		require.NoError(t, os.MkdirAll("tmp", 0o755))
		tmpDir, err := os.MkdirTemp("tmp", "TestSupportFilesInstaller-root-") //nolint:usetesting // we need to parametrize it
		require.NoError(t, err)
		t.Cleanup(func() {
			assert.NoError(t, os.RemoveAll(tmpDir))
		})

		supportDir := path.Join(tmpDir, "support", faker.Word())

		return SupportFilesInstallerParams{
			SupportDir:                    supportDir,
			OagSourceVersion:              "1.2.3-" + faker.Word(),
			OagSourceLocation:             "file://" + path.Join(tmpDir, "oag-cli-"+faker.Word()+".jar"),
			AppVersion:                    "3.2.1-" + faker.Word(),
			ServerGeneratorSourceLocation: "file://" + path.Join(tmpDir, "generator-"+faker.Word()+".jar"),
		}
	}

	t.Run("should do nothing if generator and cli are already installed", func(t *testing.T) {
		params := makeRandomGeneratorParams(t)
		metadata := SupportFilesMetadata{
			OagSourceVersion:        params.OagSourceVersion,
			OagSourceLocation:       params.OagSourceLocation,
			AppVersion:              params.AppVersion,
			GeneratorSourceLocation: params.ServerGeneratorSourceLocation,
		}
		metadataFile := filepath.Join(params.SupportDir, "metadata.json")
		oagCliFile := filepath.Join(params.SupportDir, "openapi-generator-cli.jar")
		serverGeneratorFile := filepath.Join(params.SupportDir, "server-generator.jar")
		mockFS := fstest.MapFS{
			metadataFile:        {Data: lo.Must(json.Marshal(metadata))},
			oagCliFile:          {Data: []byte{}},
			serverGeneratorFile: {Data: []byte{}},
		}
		installer := NewSupportFilesInstaller(SupportFilesInstallerDeps{
			RootLogger: TestRootLogger,
			CwdFS:      mockFS,
			Downloader: func(_ context.Context, _, _ string) error {
				require.Fail(t, "downloader should not be called")
				return nil
			},
		})
		res, err := installer(t.Context(), params)
		require.NoError(t, err)
		assert.Equal(t, SupportingFilesInstallResult{
			OagLocation:             oagCliFile,
			ServerGeneratorLocation: serverGeneratorFile,
		}, res)
	})

	t.Run("should download all files on fresh environment", func(t *testing.T) {
		params := makeRandomGeneratorParams(t)
		downloaderCalls := [][]string{}
		installer := NewSupportFilesInstaller(SupportFilesInstallerDeps{
			RootLogger: TestRootLogger,
			CwdFS:      fstest.MapFS{},
			Downloader: func(_ context.Context, source, target string) error {
				downloaderCalls = append(downloaderCalls, []string{source, target})
				return nil
			},
		})
		res, err := installer(t.Context(), params)
		require.NoError(t, err)

		supportDirStat, err := os.Stat(params.SupportDir)
		require.NoError(t, err)
		assert.True(t, supportDirStat.IsDir())

		gitIgnoreFile, err := os.ReadFile(filepath.Join(params.SupportDir, ".gitignore"))
		require.NoError(t, err)
		assert.Equal(t, []byte("*\n"), gitIgnoreFile)

		assert.Equal(t, SupportingFilesInstallResult{
			OagLocation:             path.Join(params.SupportDir, "openapi-generator-cli.jar"),
			ServerGeneratorLocation: path.Join(params.SupportDir, "server-generator.jar"),
		}, res)
		require.Len(t, downloaderCalls, 2)
		assert.Equal(t, []string{
			params.OagSourceLocation,
			path.Join(params.SupportDir, "openapi-generator-cli.jar"),
		}, downloaderCalls[0])
		assert.Equal(t, []string{
			params.ServerGeneratorSourceLocation,
			path.Join(params.SupportDir, "server-generator.jar"),
		}, downloaderCalls[1])
		gotUpdatedMetadata := SupportFilesMetadata{}
		metadataFile := path.Join(params.SupportDir, "metadata.json")
		require.NoError(t, json.Unmarshal(lo.Must(os.ReadFile(metadataFile)), &gotUpdatedMetadata))
		assert.Equal(t, params.OagSourceLocation, gotUpdatedMetadata.OagSourceLocation)
		assert.Equal(t, params.ServerGeneratorSourceLocation, gotUpdatedMetadata.GeneratorSourceLocation)
	})

	t.Run("should download if actual files are missing but not metadata", func(t *testing.T) {
		params := makeRandomGeneratorParams(t)
		metadata := SupportFilesMetadata{
			OagSourceVersion:        params.OagSourceVersion,
			OagSourceLocation:       params.OagSourceLocation,
			AppVersion:              params.AppVersion,
			GeneratorSourceLocation: params.ServerGeneratorSourceLocation,
		}
		metadataFile := path.Join(params.SupportDir, "metadata.json")
		mockFS := fstest.MapFS{
			metadataFile: &fstest.MapFile{
				Data: lo.Must(json.Marshal(metadata)),
			},
		}
		downloaderCalls := [][]string{}
		installer := NewSupportFilesInstaller(SupportFilesInstallerDeps{
			RootLogger: TestRootLogger,
			CwdFS:      mockFS,
			Downloader: func(_ context.Context, source, target string) error {
				downloaderCalls = append(downloaderCalls, []string{source, target})
				return nil
			},
		})
		res, err := installer(t.Context(), params)
		require.NoError(t, err)
		assert.Equal(t, SupportingFilesInstallResult{
			OagLocation:             path.Join(params.SupportDir, "openapi-generator-cli.jar"),
			ServerGeneratorLocation: path.Join(params.SupportDir, "server-generator.jar"),
		}, res)
		require.Len(t, downloaderCalls, 2)
		assert.Equal(t, []string{
			params.OagSourceLocation,
			path.Join(params.SupportDir, "openapi-generator-cli.jar"),
		}, downloaderCalls[0])
		assert.Equal(t, []string{
			params.ServerGeneratorSourceLocation,
			path.Join(params.SupportDir, "server-generator.jar"),
		}, downloaderCalls[1])
	})

	t.Run("should download if metadata versions are different", func(t *testing.T) {
		params := makeRandomGeneratorParams(t)
		metadata := SupportFilesMetadata{
			OagSourceVersion:        params.OagSourceVersion + "-" + faker.Word(),
			OagSourceLocation:       params.OagSourceLocation,
			AppVersion:              params.AppVersion + "-" + faker.Word(),
			GeneratorSourceLocation: params.ServerGeneratorSourceLocation,
		}
		metadataFile := path.Join(params.SupportDir, "metadata.json")
		mockFS := fstest.MapFS{
			metadataFile: &fstest.MapFile{
				Data: lo.Must(json.Marshal(metadata)),
			},
		}
		downloaderCalls := [][]string{}
		installer := NewSupportFilesInstaller(SupportFilesInstallerDeps{
			RootLogger: TestRootLogger,
			CwdFS:      mockFS,
			Downloader: func(_ context.Context, source, target string) error {
				downloaderCalls = append(downloaderCalls, []string{source, target})
				return nil
			},
		})
		res, err := installer(t.Context(), params)
		require.NoError(t, err)
		assert.Equal(t, SupportingFilesInstallResult{
			OagLocation:             path.Join(params.SupportDir, "openapi-generator-cli.jar"),
			ServerGeneratorLocation: path.Join(params.SupportDir, "server-generator.jar"),
		}, res)
		require.Len(t, downloaderCalls, 2)
		assert.Equal(t, []string{
			params.OagSourceLocation,
			path.Join(params.SupportDir, "openapi-generator-cli.jar"),
		}, downloaderCalls[0])
		assert.Equal(t, []string{
			params.ServerGeneratorSourceLocation,
			path.Join(params.SupportDir, "server-generator.jar"),
		}, downloaderCalls[1])
	})

	t.Run("should fail if failed to read metadata", func(t *testing.T) {
		params := makeRandomGeneratorParams(t)
		metadataFile := path.Join(params.SupportDir, "metadata.json")
		mockFS := fstest.MapFS{
			metadataFile: &fstest.MapFile{
				Data: []byte("invalid json"),
			},
		}
		installer := NewSupportFilesInstaller(SupportFilesInstallerDeps{
			RootLogger: TestRootLogger,
			CwdFS:      mockFS,
			Downloader: func(_ context.Context, _, _ string) error {
				require.Fail(t, "downloader should not be called")
				return nil
			},
		})
		_, err := installer(t.Context(), params)
		require.Error(t, err)
	})

	t.Run("should fail if failed to download oag-cli", func(t *testing.T) {
		params := makeRandomGeneratorParams(t)
		wantErr := errors.New(faker.Sentence())
		installer := NewSupportFilesInstaller(SupportFilesInstallerDeps{
			RootLogger: TestRootLogger,
			CwdFS:      fstest.MapFS{},
			Downloader: func(_ context.Context, _, _ string) error {
				return wantErr
			},
		})
		_, err := installer(t.Context(), params)
		require.Error(t, err)
	})

	t.Run("should fail if failed to download server-generator", func(t *testing.T) {
		params := makeRandomGeneratorParams(t)
		wantErr := errors.New(faker.Sentence())
		installer := NewSupportFilesInstaller(SupportFilesInstallerDeps{
			RootLogger: TestRootLogger,
			CwdFS:      fstest.MapFS{},
			Downloader: func(_ context.Context, source, _ string) error {
				if source == params.OagSourceLocation {
					return nil
				}
				return wantErr
			},
		})
		_, err := installer(t.Context(), params)
		require.Error(t, err)
	})
}
