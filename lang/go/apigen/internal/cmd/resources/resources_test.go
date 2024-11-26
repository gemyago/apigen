package resources

import (
	"testing"
	"testing/fstest"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadPluginVersion(t *testing.T) {
	t.Run("should read the version", func(t *testing.T) {
		wantVer := faker.Word()
		mockFs := fstest.MapFS{
			"go-apigen-server.xml": &fstest.MapFile{
				Data: []byte(`<project><version>` + wantVer + `</version></project>`),
			},
		}
		reader := ResourcesMetadataReader{fs: mockFs}
		ver, err := reader.ReadPluginVersion()
		require.NoError(t, err)
		assert.Equal(t, wantVer, ver)
	})

	t.Run("should read the version from the embedded file", func(t *testing.T) {
		reader := NewResourcesMetadataReader()
		ver, err := reader.ReadPluginVersion()
		require.NoError(t, err)
		assert.NotEmpty(t, ver)
	})

	t.Run("should fail if no such file", func(t *testing.T) {
		mockFs := fstest.MapFS{}
		reader := ResourcesMetadataReader{fs: mockFs}
		_, err := reader.ReadPluginVersion()
		require.Error(t, err)
	})

	t.Run("should fail if the file is not xml", func(t *testing.T) {
		mockFs := fstest.MapFS{
			"go-apigen-server.xml": &fstest.MapFile{
				Data: []byte("not an xml"),
			},
		}
		reader := ResourcesMetadataReader{fs: mockFs}
		_, err := reader.ReadPluginVersion()
		require.Error(t, err)
	})

	t.Run("should fail if the version is not found", func(t *testing.T) {
		mockFs := fstest.MapFS{
			"go-apigen-server.xml": &fstest.MapFile{
				Data: []byte(`<project></project>`),
			},
		}
		reader := ResourcesMetadataReader{fs: mockFs}
		_, err := reader.ReadPluginVersion()
		require.Error(t, err)
	})

	t.Run("should fail if the version is empty", func(t *testing.T) {
		mockFs := fstest.MapFS{
			"go-apigen-server.xml": &fstest.MapFile{
				Data: []byte(`<project><version></version></project>`),
			},
		}
		reader := ResourcesMetadataReader{fs: mockFs}
		_, err := reader.ReadPluginVersion()
		require.Error(t, err)
	})
}

func TestReadOpenapiGeneratorCliVersion(t *testing.T) {
	t.Run("should read the OPENAPI_GENERATOR_CLI version", func(t *testing.T) {
		wantVer := faker.Word()
		mockFs := fstest.MapFS{
			".versions": &fstest.MapFile{
				Data: []byte("OPENAPI_GENERATOR_CLI: " + wantVer + "\nMAVEN: 3.8.8"),
			},
		}
		reader := ResourcesMetadataReader{fs: mockFs}
		ver, err := reader.ReadOpenapiGeneratorCliVersion()
		require.NoError(t, err)
		assert.Equal(t, wantVer, ver)
	})

	t.Run("should read the version from the embedded file", func(t *testing.T) {
		reader := NewResourcesMetadataReader()
		ver, err := reader.ReadOpenapiGeneratorCliVersion()
		require.NoError(t, err)
		assert.NotEmpty(t, ver)
	})

	t.Run("should fail if no such file", func(t *testing.T) {
		mockFs := fstest.MapFS{}
		reader := ResourcesMetadataReader{fs: mockFs}
		_, err := reader.ReadOpenapiGeneratorCliVersion()
		require.Error(t, err)
	})

	t.Run("should fail if the version is not found", func(t *testing.T) {
		mockFs := fstest.MapFS{
			".versions": &fstest.MapFile{
				Data: []byte("MAVEN: 3.8.8"),
			},
		}
		reader := ResourcesMetadataReader{fs: mockFs}
		_, err := reader.ReadOpenapiGeneratorCliVersion()
		require.Error(t, err)
	})
}
