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

type SupportingFilesInstallResult struct {
	OagLocation             string
	ServerGeneratorLocation string
}

type SupportFilesInstaller func(
	ctx context.Context,
	params SupportFilesInstallerParams,
) (SupportingFilesInstallResult, error)

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

type downloadSupportFileIfRequiredParams struct {
	sourceLocation   string
	destinationPath  string
	sourceVersion    string
	metadataVersion  *string
	metadataLocation *string
	metadataChanged  *bool
}

func downloadSupportFileIfRequired(
	ctx context.Context,
	deps SupportFilesInstallerDeps,
	params downloadSupportFileIfRequiredParams,
) error {
	fileExists := true
	file, err := deps.RootFS.Open(params.destinationPath[1:])
	if err != nil {
		fileExists = false
	} else {
		defer file.Close()
	}

	if !fileExists || params.sourceVersion != *params.metadataVersion {
		if err = deps.Downloader(
			ctx,
			params.sourceLocation,
			params.destinationPath,
		); err != nil {
			return fmt.Errorf("failed to download openapi-generator-cli: %w", err)
		}
		*params.metadataLocation = params.sourceLocation
		*params.metadataVersion = params.sourceVersion
		*params.metadataChanged = true
	}
	return nil
}

func NewSupportFilesInstaller(deps SupportFilesInstallerDeps) SupportFilesInstaller {
	return func(ctx context.Context, params SupportFilesInstallerParams) (SupportingFilesInstallResult, error) {
		var emptyResult SupportingFilesInstallResult
		metadataFile := path.Join(params.SupportDir, "metadata.json")
		metadata, err := readSupportFilesMetadata(deps.RootFS, metadataFile)
		if err != nil {
			return emptyResult, err
		}

		metadataChanged := false
		oagDestination := path.Join(params.SupportDir, "openapi-generator-cli.jar")
		if err = downloadSupportFileIfRequired(
			ctx,
			deps,
			downloadSupportFileIfRequiredParams{
				sourceLocation:   params.OagSourceLocation,
				destinationPath:  oagDestination,
				sourceVersion:    params.OagSourceVersion,
				metadataVersion:  &metadata.OagSourceVersion,
				metadataLocation: &metadata.OagSourceLocation,
				metadataChanged:  &metadataChanged,
			},
		); err != nil {
			return emptyResult, err
		}

		serverGeneratorDestination := path.Join(params.SupportDir, "server-generator.jar")
		if err = downloadSupportFileIfRequired(
			ctx,
			deps,
			downloadSupportFileIfRequiredParams{
				sourceLocation:   params.ServerGeneratorSourceLocation,
				destinationPath:  serverGeneratorDestination,
				sourceVersion:    params.ServerGeneratorSourceVersion,
				metadataVersion:  &metadata.GeneratorSourceVersion,
				metadataLocation: &metadata.GeneratorSourceLocation,
				metadataChanged:  &metadataChanged,
			},
		); err != nil {
			return emptyResult, err
		}

		if metadataChanged {
			var metadataOutput *os.File
			metadataOutput, err = os.Create(metadataFile)
			if err != nil {
				return emptyResult, fmt.Errorf("failed to open metadata file for writing: %w", err)
			}
			defer metadataOutput.Close()
			if err = json.NewEncoder(metadataOutput).Encode(metadata); err != nil {
				return emptyResult, fmt.Errorf("failed to write metadata: %w", err)
			}
		}

		return SupportingFilesInstallResult{
			OagLocation:             oagDestination,
			ServerGeneratorLocation: serverGeneratorDestination,
		}, nil
	}
}
