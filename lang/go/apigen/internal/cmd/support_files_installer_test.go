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
		require.NoError(t, os.MkdirAll(supportDir, 0755))
		return SupportFilesInstallerParams{
			SupportDir:                    supportDir,
			OagSourceLocation:             "file://" + path.Join(t.TempDir(), "oag-cli-"+faker.Word()+".jar"),
			ServerGeneratorSourceLocation: "file://" + path.Join(t.TempDir(), "generator-"+faker.Word()+".jar"),
		}
	}

	t.Run("should do nothing if generator and cli are already installed", func(t *testing.T) {
		params := makeRandomGeneratorParams(t)
		metadata := SupportFilesMetadata{
			OagSourceLocation:       params.OagSourceLocation,
			GeneratorSourceLocation: params.ServerGeneratorSourceLocation,
		}
		metadataFile := filepath.Join(params.SupportDir, "metadata.json")
		oagCliFile := filepath.Join(params.SupportDir, "openapi-generator-cli.jar")
		serverGeneratorFile := filepath.Join(params.SupportDir, "server-generator.jar")
		mockFS := fstest.MapFS{
			metadataFile[1:]:        {Data: lo.Must(json.Marshal(metadata))},
			oagCliFile[1:]:          {},
			serverGeneratorFile[1:]: {},
		}
		installer := NewSupportFilesInstaller(SupportFilesInstallerDeps{
			RootFS: mockFS,
		})
		err := installer(context.Background(), params)
		require.NoError(t, err)
	})

	t.Run("should download all files on fresh environment", func(t *testing.T) {
		params := makeRandomGeneratorParams(t)
		downloaderCalls := [][]string{}
		installer := NewSupportFilesInstaller(SupportFilesInstallerDeps{
			RootFS: fstest.MapFS{},
			Downloader: func(_ context.Context, source, target string) error {
				downloaderCalls = append(downloaderCalls, []string{source, target})
				return nil
			},
		})
		err := installer(context.Background(), params)
		require.NoError(t, err)
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
			OagSourceLocation:       params.OagSourceLocation,
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
			RootFS: mockFS,
			Downloader: func(_ context.Context, source, target string) error {
				downloaderCalls = append(downloaderCalls, []string{source, target})
				return nil
			},
		})
		err := installer(context.Background(), params)
		require.NoError(t, err)
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

	t.Run("should download if metadata locations are different", func(t *testing.T) {
		params := makeRandomGeneratorParams(t)
		metadata := SupportFilesMetadata{
			OagSourceLocation:       "file://" + path.Join(t.TempDir(), "oag-cli-"+faker.Word()+".jar"),
			GeneratorSourceLocation: "file://" + path.Join(t.TempDir(), "generator-"+faker.Word()+".jar"),
		}
		metadataFile := path.Join(params.SupportDir, "metadata.json")
		mockFS := fstest.MapFS{
			metadataFile: &fstest.MapFile{
				Data: lo.Must(json.Marshal(metadata)),
			},
		}
		downloaderCalls := [][]string{}
		installer := NewSupportFilesInstaller(SupportFilesInstallerDeps{
			RootFS: mockFS,
			Downloader: func(_ context.Context, source, target string) error {
				downloaderCalls = append(downloaderCalls, []string{source, target})
				return nil
			},
		})
		err := installer(context.Background(), params)
		require.NoError(t, err)
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
