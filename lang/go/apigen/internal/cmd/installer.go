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
	OagSourceLocation       string `json:"oagSourceLocation"`
	GeneratorSourceLocation string `json:"serverGeneratorSourceLocation"`
}

type SupportFilesInstallerParams struct {
	SupportDir                    string
	OagSourceLocation             string
	ServerGeneratorSourceLocation string
}

type SupportFilesInstaller func(ctx context.Context, params SupportFilesInstallerParams) error

type SupportFilesInstallerDeps struct {
	Downloader ResourceDownloader
	RootFS     fs.ReadFileFS
}

func NewSupportFilesInstaller(deps SupportFilesInstallerDeps) SupportFilesInstaller {
	return func(ctx context.Context, params SupportFilesInstallerParams) error {
		metadataFile := path.Join(params.SupportDir, "metadata.json")
		var metadata SupportFilesMetadata
		if data, err := deps.RootFS.ReadFile(metadataFile[1:]); err == nil {
			if err = json.Unmarshal(data, &metadata); err != nil {
				return fmt.Errorf("failed to unmarshal metadata: %w", err)
			}
		} else if !os.IsNotExist(err) {
			return fmt.Errorf("failed to read metadata: %w", err)
		}

		metadataChanged := false
		if metadata.OagSourceLocation != params.OagSourceLocation {
			if err := deps.Downloader(
				ctx,
				params.OagSourceLocation,
				path.Join(params.SupportDir, "openapi-generator-cli.jar"),
			); err != nil {
				return fmt.Errorf("failed to download openapi-generator-cli: %w", err)
			}
			metadata.OagSourceLocation = params.OagSourceLocation
			metadataChanged = true
		}

		if metadata.GeneratorSourceLocation != params.ServerGeneratorSourceLocation {
			if err := deps.Downloader(
				ctx,
				params.ServerGeneratorSourceLocation,
				path.Join(params.SupportDir, "server-generator.jar"),
			); err != nil {
				return fmt.Errorf("failed to download server-generator: %w", err)
			}
			metadata.GeneratorSourceLocation = params.ServerGeneratorSourceLocation
			metadataChanged = true
		}

		if metadataChanged {
			metadataOutput, err := os.Create(metadataFile)
			if err != nil {
				return fmt.Errorf("failed to open metadata file: %w", err)
			}
			defer metadataOutput.Close()
			if err = json.NewEncoder(metadataOutput).Encode(metadata); err != nil {
				return fmt.Errorf("failed to write metadata: %w", err)
			}
		}

		return nil
	}
}
