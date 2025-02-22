package resources

import (
	"testing"
	"testing/fstest"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMetadataReader(t *testing.T) {
	t.Run("ReadOpenapiGeneratorCliVersion", func(t *testing.T) {
		t.Run("should read the OPENAPI_GENERATOR_CLI version", func(t *testing.T) {
			wantVer := faker.Word()
			mockFs := fstest.MapFS{
				".versions": &fstest.MapFile{
					Data: []byte("OPENAPI_GENERATOR_CLI: " + wantVer + "\nMAVEN: 3.8.8"),
				},
			}
			reader := MetadataReader{embeddedRootFS: mockFs}
			ver, err := reader.ReadOpenapiGeneratorCliVersion()
			require.NoError(t, err)
			assert.Equal(t, wantVer, ver)
		})

		t.Run("should fail if no such file", func(t *testing.T) {
			mockFs := fstest.MapFS{}
			reader := MetadataReader{embeddedRootFS: mockFs}
			_, err := reader.ReadOpenapiGeneratorCliVersion()
			require.Error(t, err)
		})

		t.Run("should fail if the version is not found", func(t *testing.T) {
			mockFs := fstest.MapFS{
				".versions": &fstest.MapFile{
					Data: []byte("MAVEN: 3.8.8"),
				},
			}
			reader := MetadataReader{embeddedRootFS: mockFs}
			_, err := reader.ReadOpenapiGeneratorCliVersion()
			require.Error(t, err)
		})
	})

	t.Run("ReadAppVersion", func(t *testing.T) {
		t.Run("should read the APP_VERSION version", func(t *testing.T) {
			wantVer := faker.Word()
			mockFs := fstest.MapFS{
				".versions": &fstest.MapFile{
					Data: []byte("APP_VERSION: " + wantVer + "\nMAVEN: 3.8.8"),
				},
			}
			reader := MetadataReader{embeddedRootFS: mockFs}
			ver, err := reader.ReadAppVersion()
			require.NoError(t, err)
			assert.Equal(t, wantVer, ver)
		})

		t.Run("should fail if no such file", func(t *testing.T) {
			mockFs := fstest.MapFS{}
			reader := MetadataReader{embeddedRootFS: mockFs}
			_, err := reader.ReadAppVersion()
			require.Error(t, err)
		})

		t.Run("should fail if the version is not found", func(t *testing.T) {
			mockFs := fstest.MapFS{
				".versions": &fstest.MapFile{
					Data: []byte("MAVEN: 3.8.8"),
				},
			}
			reader := MetadataReader{embeddedRootFS: mockFs}
			_, err := reader.ReadAppVersion()
			require.Error(t, err)
		})
	})
}
