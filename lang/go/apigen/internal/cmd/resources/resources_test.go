package resources

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadPluginVersion(t *testing.T) {
	t.Run("should read the version", func(t *testing.T) {
		ver, err := ReadPluginVersion()
		require.NoError(t, err)
		assert.NotEmpty(t, ver)
	})
}
