package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path"
)

type SupportFilesMetadata struct {
	OagSourceVersion  string `json:"oagSourceVersion"`
	OagSourceLocation string `json:"oagSourceLocation"`

	GeneratorSourceVersion  string `json:"serverGeneratorSourceVersion"`
	GeneratorSourceLocation string `json:"serverGeneratorSourceLocation"`
}

type SupportFilesInstallerParams struct {
	SupportDir                    string
	OagSourceVersion              string
	OagSourceLocation             string
	ServerGeneratorSourceVersion  string
	ServerGeneratorSourceLocation string
}

type SupportFilesInstaller func(ctx context.Context, params SupportFilesInstallerParams) error

type SupportFilesInstallerDeps struct {
	Downloader ResourceDownloader
	RootFS     fs.ReadFileFS
}

func readSupportFilesMetadata(rootFS fs.ReadFileFS, metadataFile string) (SupportFilesMetadata, error) {
	var metadata SupportFilesMetadata
	if data, err := rootFS.ReadFile(metadataFile[1:]); err == nil {
		if err = json.Unmarshal(data, &metadata); err != nil {
			return SupportFilesMetadata{}, fmt.Errorf("failed to unmarshal metadata: %w", err)
		}
	} else if !os.IsNotExist(err) {
		return SupportFilesMetadata{}, fmt.Errorf("failed to read metadata: %w", err)
	}
	return metadata, nil
}

func fsFileExists(fs fs.FS, fullPath string) bool {
	file, err := fs.Open(fullPath[1:])
	if err != nil {
		return false
	}
	defer file.Close()
	return true
}

func NewSupportFilesInstaller(deps SupportFilesInstallerDeps) SupportFilesInstaller {
	return func(ctx context.Context, params SupportFilesInstallerParams) error {
		metadataFile := path.Join(params.SupportDir, "metadata.json")
		metadata, err := readSupportFilesMetadata(deps.RootFS, metadataFile)
		if err != nil {
			return err
		}

		oagDestination := path.Join(params.SupportDir, "openapi-generator-cli.jar")
		metadataChanged := false
		if !fsFileExists(deps.RootFS, oagDestination) || metadata.OagSourceVersion != params.OagSourceVersion {
			if err = deps.Downloader(
				ctx,
				params.OagSourceLocation,
				oagDestination,
			); err != nil {
				return fmt.Errorf("failed to download openapi-generator-cli: %w", err)
			}
			metadata.OagSourceLocation = params.OagSourceLocation
			metadata.OagSourceVersion = params.OagSourceVersion
			metadataChanged = true
		}

		serverGeneratorDestination := path.Join(params.SupportDir, "server-generator.jar")
		if !fsFileExists(deps.RootFS, serverGeneratorDestination) ||
			metadata.GeneratorSourceVersion != params.ServerGeneratorSourceVersion {
			if err = deps.Downloader(
				ctx,
				params.ServerGeneratorSourceLocation,
				serverGeneratorDestination,
			); err != nil {
				return fmt.Errorf("failed to download server-generator: %w", err)
			}
			metadata.GeneratorSourceLocation = params.ServerGeneratorSourceLocation
			metadata.GeneratorSourceVersion = params.ServerGeneratorSourceVersion
			metadataChanged = true
		}

		if metadataChanged {
			var metadataOutput *os.File
			metadataOutput, err = os.Create(metadataFile)
			if err != nil {
				return fmt.Errorf("failed to open metadata file for writing: %w", err)
			}
			defer metadataOutput.Close()
			if err = json.NewEncoder(metadataOutput).Encode(metadata); err != nil {
				return fmt.Errorf("failed to write metadata: %w", err)
			}
		}

		return nil
	}
}
