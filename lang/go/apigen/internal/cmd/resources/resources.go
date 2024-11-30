package resources

import (
	"bufio"
	"embed"
	"errors"
	"io/fs"
	"strings"
)

//go:generate cp -r ../../../../../../.versions ./.versions

//go:embed .versions
var resources embed.FS

const versionLineSegmentsCount = 2

type MetadataReader struct {
	fs fs.ReadFileFS
}

// ReadAppVersion reads the APP_VERSION value from embedded .versions file.
func (r *MetadataReader) ReadAppVersion() (string, error) {
	data, err := r.fs.ReadFile(".versions")
	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "APP_VERSION:") {
			parts := strings.Split(line, ":")
			if len(parts) == versionLineSegmentsCount {
				return strings.TrimSpace(parts[1]), nil
			}
		}
	}

	return "", errors.New("APP_VERSION version not found")
}

// ReadOpenapiGeneratorCliVersion reads the OPENAPI_GENERATOR_CLI value from embedded .versions file.
func (r *MetadataReader) ReadOpenapiGeneratorCliVersion() (string, error) {
	data, err := r.fs.ReadFile(".versions")
	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "OPENAPI_GENERATOR_CLI:") {
			parts := strings.Split(line, ":")
			if len(parts) == versionLineSegmentsCount {
				return strings.TrimSpace(parts[1]), nil
			}
		}
	}

	return "", errors.New("OPENAPI_GENERATOR_CLI version not found")
}

func NewMetadataReader() *MetadataReader {
	return &MetadataReader{fs: resources}
}
