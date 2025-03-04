package cmd

import (
	"testing"
	"testing/fstest"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"
)

func TestCommands(t *testing.T) {
	t.Run("should run server command in noop mode", func(t *testing.T) {
		rootCmd := WireUpCommands(fstest.MapFS{
			".versions": &fstest.MapFile{
				Data: []byte("OPENAPI_GENERATOR_CLI: " + faker.Word() + "\nAPP_VERSION: " + faker.Word()),
			},
		})
		rootCmd.SetArgs([]string{"server", faker.URL(), faker.URL(), "--noop"})

		err := rootCmd.Execute()
		require.NoError(t, err)
	})
}
