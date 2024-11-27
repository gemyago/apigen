package cmd

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResourceDownloader(t *testing.T) {
	t.Run("should download the resource to a given location", func(t *testing.T) {
		outputDir := t.TempDir()
		outputFile := path.Join(outputDir, faker.Word()+".txt")

		sourceContent := faker.Sentence()
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			lo.Must(w.Write([]byte(sourceContent)))
		}))
		defer srv.Close()

		downloader := NewResourceDownloader()

		err := downloader(context.Background(), srv.URL, outputFile)
		require.NoError(t, err)

		content, err := os.ReadFile(outputFile)
		require.NoError(t, err)
		assert.Equal(t, sourceContent, string(content))
	})
}
