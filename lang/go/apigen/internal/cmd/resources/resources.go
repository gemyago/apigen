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

//go:embed go-apigen-server.xml
var goApigenServer embed.FS

//go:embed .versions
var versions embed.FS

type mavenProject struct {
	Version string `xml:"version"`
}

func readPluginVersion(fs fs.ReadFileFS) (string, error) {
	data, err := fs.ReadFile("go-apigen-server.xml")
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

func ReadPluginVersion() (string, error) {
	return readPluginVersion(goApigenServer)
}

func readOpenapiGeneratorCliVersion(fs fs.ReadFileFS) (string, error) {
	data, err := fs.ReadFile(".versions")
	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "OPENAPI_GENERATOR_CLI:") {
			parts := strings.Split(line, ":")
			if len(parts) == 2 {
				return strings.TrimSpace(parts[1]), nil
			}
		}
	}

	return "", errors.New("OPENAPI_GENERATOR_CLI version not found")
}

func ReadOpenapiGeneratorCliVersion() (string, error) {
	return readOpenapiGeneratorCliVersion(versions)
}
