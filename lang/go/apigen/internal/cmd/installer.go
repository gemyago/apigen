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

func downloadSupportFileIfRequired(
	ctx context.Context,
	deps SupportFilesInstallerDeps,
	sourceLocation string,
	destinationPath string,
	sourceVersion string,
	metadataVersion *string,
	metadataLocation *string,
	metadataChanged *bool,
) error {
	fileExists := true
	file, err := deps.RootFS.Open(destinationPath[1:])
	if err != nil {
		fileExists = false
	} else {
		defer file.Close()
	}

	if !fileExists || sourceVersion != *metadataVersion {
		if err = deps.Downloader(
			ctx,
			sourceLocation,
			destinationPath,
		); err != nil {
			return fmt.Errorf("failed to download openapi-generator-cli: %w", err)
		}
		*metadataLocation = sourceLocation
		*metadataVersion = sourceVersion
		*metadataChanged = true
	}
	return nil
}

func NewSupportFilesInstaller(deps SupportFilesInstallerDeps) SupportFilesInstaller {
	return func(ctx context.Context, params SupportFilesInstallerParams) error {
		metadataFile := path.Join(params.SupportDir, "metadata.json")
		metadata, err := readSupportFilesMetadata(deps.RootFS, metadataFile)
		if err != nil {
			return err
		}

		metadataChanged := false
		oagDestination := path.Join(params.SupportDir, "openapi-generator-cli.jar")
		if err = downloadSupportFileIfRequired(
			ctx,
			deps,
			params.OagSourceLocation,
			oagDestination,
			params.OagSourceVersion,
			&metadata.OagSourceVersion,
			&metadata.OagSourceLocation,
			&metadataChanged,
		); err != nil {
			return err
		}

		serverGeneratorDestination := path.Join(params.SupportDir, "server-generator.jar")
		if err = downloadSupportFileIfRequired(
			ctx,
			deps,
			params.ServerGeneratorSourceLocation,
			serverGeneratorDestination,
			params.ServerGeneratorSourceVersion,
			&metadata.GeneratorSourceVersion,
			&metadata.GeneratorSourceLocation,
			&metadataChanged,
		); err != nil {
			return err
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
