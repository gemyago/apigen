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
		ver, err := readPluginVersion(mockFs)
		require.NoError(t, err)
		assert.Equal(t, wantVer, ver)
	})

	t.Run("should read the version from the embedded file", func(t *testing.T) {
		ver, err := ReadPluginVersion()
		require.NoError(t, err)
		assert.NotEmpty(t, ver)
	})

	t.Run("should fail if no such file", func(t *testing.T) {
		mockFs := fstest.MapFS{}
		_, err := readPluginVersion(mockFs)
		require.Error(t, err)
	})

	t.Run("should fail if the file is not xml", func(t *testing.T) {
		mockFs := fstest.MapFS{
			"go-apigen-server.xml": &fstest.MapFile{
				Data: []byte("not an xml"),
			},
		}
		_, err := readPluginVersion(mockFs)
		require.Error(t, err)
	})

	t.Run("should fail if the version is not found", func(t *testing.T) {
		mockFs := fstest.MapFS{
			"go-apigen-server.xml": &fstest.MapFile{
				Data: []byte(`<project></project>`),
			},
		}
		_, err := readPluginVersion(mockFs)
		require.Error(t, err)
	})

	t.Run("should fail if the version is empty", func(t *testing.T) {
		mockFs := fstest.MapFS{
			"go-apigen-server.xml": &fstest.MapFile{
				Data: []byte(`<project><version></version></project>`),
			},
		}
		_, err := readPluginVersion(mockFs)
		require.Error(t, err)
	})
}
