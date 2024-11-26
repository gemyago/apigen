package resources

import (
	"embed"
	"encoding/xml"
	"errors"
	"io/fs"
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
