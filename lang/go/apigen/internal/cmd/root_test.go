package cmd

import (
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"
)

func TestRootCmd(t *testing.T) {
	t.Run("should run root command in noop mode", func(t *testing.T) {
		rootCmd := NewRootCmd()
		rootCmd.SetArgs([]string{faker.URL(), faker.URL(), "--noop"})
		err := rootCmd.Execute()
		require.NoError(t, err)
	})
}
