package cmd

import (
	"context"
	"encoding/json"
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
		supportDir := path.Join(t.TempDir(), "support", faker.Word())
		return SupportFilesInstallerParams{
			SupportDir:                    supportDir,
			OagSourceVersion:              "1.2.3-" + faker.Word(),
			OagSourceLocation:             "file://" + path.Join(t.TempDir(), "oag-cli-"+faker.Word()+".jar"),
			ServerGeneratorSourceVersion:  "3.2.1-" + faker.Word(),
			ServerGeneratorSourceLocation: "file://" + path.Join(t.TempDir(), "generator-"+faker.Word()+".jar"),
		}
	}

	t.Run("should do nothing if generator and cli are already installed", func(t *testing.T) {
		params := makeRandomGeneratorParams(t)
		metadata := SupportFilesMetadata{
			OagSourceVersion:        params.OagSourceVersion,
			OagSourceLocation:       params.OagSourceLocation,
			GeneratorSourceVersion:  params.ServerGeneratorSourceVersion,
			GeneratorSourceLocation: params.ServerGeneratorSourceLocation,
		}
		metadataFile := filepath.Join(params.SupportDir, "metadata.json")
		oagCliFile := filepath.Join(params.SupportDir, "openapi-generator-cli.jar")
		serverGeneratorFile := filepath.Join(params.SupportDir, "server-generator.jar")
		mockFS := fstest.MapFS{
			metadataFile[1:]:        {Data: lo.Must(json.Marshal(metadata))},
			oagCliFile[1:]:          {Data: []byte{}},
			serverGeneratorFile[1:]: {Data: []byte{}},
		}
		installer := NewSupportFilesInstaller(SupportFilesInstallerDeps{
			RootLogger: DiscardLogger,
			RootFS:     mockFS,
		})
		res, err := installer(context.Background(), params)
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
			RootLogger: DiscardLogger,
			RootFS:     fstest.MapFS{},
			Downloader: func(_ context.Context, source, target string) error {
				downloaderCalls = append(downloaderCalls, []string{source, target})
				return nil
			},
		})
		res, err := installer(context.Background(), params)
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
			GeneratorSourceVersion:  params.ServerGeneratorSourceVersion,
			GeneratorSourceLocation: params.ServerGeneratorSourceLocation,
		}
		metadataFile := path.Join(params.SupportDir, "metadata.json")
		mockFS := fstest.MapFS{
			metadataFile[1:]: &fstest.MapFile{
				Data: lo.Must(json.Marshal(metadata)),
			},
		}
		downloaderCalls := [][]string{}
		installer := NewSupportFilesInstaller(SupportFilesInstallerDeps{
			RootLogger: DiscardLogger,
			RootFS:     mockFS,
			Downloader: func(_ context.Context, source, target string) error {
				downloaderCalls = append(downloaderCalls, []string{source, target})
				return nil
			},
		})
		res, err := installer(context.Background(), params)
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
			GeneratorSourceVersion:  params.ServerGeneratorSourceVersion + "-" + faker.Word(),
			GeneratorSourceLocation: params.ServerGeneratorSourceLocation,
		}
		metadataFile := path.Join(params.SupportDir, "metadata.json")
		mockFS := fstest.MapFS{
			metadataFile[1:]: &fstest.MapFile{
				Data: lo.Must(json.Marshal(metadata)),
			},
		}
		downloaderCalls := [][]string{}
		installer := NewSupportFilesInstaller(SupportFilesInstallerDeps{
			RootLogger: DiscardLogger,
			RootFS:     mockFS,
			Downloader: func(_ context.Context, source, target string) error {
				downloaderCalls = append(downloaderCalls, []string{source, target})
				return nil
			},
		})
		res, err := installer(context.Background(), params)
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
}
