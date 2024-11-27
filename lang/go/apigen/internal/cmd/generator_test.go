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
	t.Run("should download the resource from a given url", func(t *testing.T) {
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

	t.Run("should copy the resource if it's not an url", func(t *testing.T) {
		outputDir := t.TempDir()
		outputFile := path.Join(outputDir, faker.Word()+".txt")

		sourceContent := faker.Sentence()
		sourceFile := path.Join(outputDir, faker.Word()+".txt")
		require.NoError(t, os.WriteFile(sourceFile, []byte(sourceContent), 0644))

		downloader := NewResourceDownloader()

		err := downloader(context.Background(), sourceFile, outputFile)
		require.NoError(t, err)

		content, err := os.ReadFile(outputFile)
		require.NoError(t, err)
		assert.Equal(t, sourceContent, string(content))
	})

	t.Run("should copy the resource if it's a file url", func(t *testing.T) {
		outputDir := t.TempDir()
		outputFile := path.Join(outputDir, faker.Word()+".txt")

		sourceContent := faker.Sentence()
		sourceFile := path.Join(outputDir, faker.Word()+".txt")
		require.NoError(t, os.WriteFile(sourceFile, []byte(sourceContent), 0644))

		downloader := NewResourceDownloader()

		err := downloader(context.Background(), "file://"+sourceFile, outputFile)
		require.NoError(t, err)

		content, err := os.ReadFile(outputFile)
		require.NoError(t, err)
		assert.Equal(t, sourceContent, string(content))
	})
}
