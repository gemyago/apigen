package resources

import (
	"bufio"
	"embed"
	"encoding/xml"
	"errors"
	"io/fs"
	"strings"
)

//go:generate cp -r ../../../../../../generators/go-apigen-server/pom.xml ./go-apigen-server.xml
//go:generate cp -r ../../../../../../.versions ./.versions

//go:embed go-apigen-server.xml .versions
var resources embed.FS

const versionLineSegmentsCount = 2

type mavenProject struct {
	Version string `xml:"version"`
}

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

// ReadServerPluginVersion reads the version of the server plugin from the embedded go-apigen-server.xml file.
func (r *MetadataReader) ReadServerPluginVersion() (string, error) {
	data, err := r.fs.ReadFile("go-apigen-server.xml")
	if err != nil {
		return "", err
	}

	var project mavenProject
	if err = xml.Unmarshal(data, &project); err != nil {
		return "", err
	}

	if project.Version == "" {
		return "", errors.New("version not found")
	}

	return project.Version, nil
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
