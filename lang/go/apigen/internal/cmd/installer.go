package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path"
)

type SupportFilesMetadata struct {
	AppVersion string `json:"appVersion"`

	OagSourceVersion  string `json:"oagSourceVersion"`
	OagSourceLocation string `json:"oagSourceLocation"`

	GeneratorSourceLocation string `json:"serverGeneratorSourceLocation"`
}

type SupportFilesInstallerParams struct {
	SupportDir                    string
	OagSourceVersion              string
	OagSourceLocation             string
	AppVersion                    string
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
	RootLogger *slog.Logger
	Downloader ResourceDownloader
}

func readSupportFilesMetadata(
	ctx context.Context,
	logger *slog.Logger,
	metadataFile string,
) (SupportFilesMetadata, error) {
	var metadata SupportFilesMetadata
	data, err := os.ReadFile(metadataFile)
	if err == nil {
		if err = json.Unmarshal(data, &metadata); err != nil {
			return SupportFilesMetadata{}, fmt.Errorf("failed to unmarshal metadata: %w", err)
		}
		return metadata, nil
	}

	if !os.IsNotExist(err) {
		return SupportFilesMetadata{}, fmt.Errorf("failed to read metadata: %w", err)
	}
	logger.DebugContext(ctx, "Support metadata file not found", slog.String("file", metadataFile))
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
	logger *slog.Logger,
	deps SupportFilesInstallerDeps,
	params downloadSupportFileIfRequiredParams,
) error {
	fileExists := true
	file, err := os.Open(params.destinationPath)
	if err != nil {
		fileExists = false
	} else {
		defer file.Close()
	}

	if !fileExists ||
		params.sourceVersion != *params.metadataVersion ||
		params.sourceLocation != *params.metadataLocation {
		logger.InfoContext(ctx, "Downloading support file",
			slog.String("metadataVersion", *params.metadataVersion),
			slog.String("metadataLocation", *params.metadataLocation),
			slog.String("sourceVersion", params.sourceVersion),
			slog.String("sourceLocation", params.sourceLocation),
			slog.String("destination", params.destinationPath),
		)
		if err = deps.Downloader(
			ctx,
			params.sourceLocation,
			params.destinationPath,
		); err != nil {
			return fmt.Errorf("failed to download support file %s: %w", params.sourceLocation, err)
		}
		*params.metadataLocation = params.sourceLocation
		*params.metadataVersion = params.sourceVersion
		*params.metadataChanged = true
	}
	return nil
}

const supportDirGitIgnoreContents = `# Contents of this directory are produced by apigen tool and used to 
# hold tool related supporting files. Generally you don't need to
# commit these files to your repository. They are required at dev time only (to run the generator). 
# However if for some reason you want to keep them in git, you are free to do so. You can also remove
# this file in order to have a git track on this directory.
*
`

func ensureSupportDir(logger *slog.Logger, dirName string) error {
	_, err := os.Stat(dirName)
	if err == nil {
		return nil
	}
	if !os.IsNotExist(err) {
		return fmt.Errorf("failed to open support directory: %w", err)
	}

	logger.Info("Creating support directory", slog.String("dir", dirName))
	if err = os.MkdirAll(dirName, 0755); err != nil {
		return fmt.Errorf("failed to create support directory: %w", err)
	}

	gitIgnorePath := path.Join(dirName, ".gitignore")
	logger.Info("Creating .gitignore file", slog.String("path", gitIgnorePath))
	if err = os.WriteFile(gitIgnorePath, []byte(supportDirGitIgnoreContents), 0600); err != nil {
		return fmt.Errorf("failed to create .gitignore file: %w", err)
	}

	return nil
}

func NewSupportFilesInstaller(deps SupportFilesInstallerDeps) SupportFilesInstaller {
	logger := deps.RootLogger.WithGroup("support-files-installer")
	return func(ctx context.Context, params SupportFilesInstallerParams) (SupportingFilesInstallResult, error) {
		var emptyResult SupportingFilesInstallResult

		if err := ensureSupportDir(logger, params.SupportDir); err != nil {
			return emptyResult, err
		}

		metadataFile := path.Join(params.SupportDir, "metadata.json")
		metadata, err := readSupportFilesMetadata(ctx, logger, metadataFile)
		if err != nil {
			return emptyResult, err
		}

		metadataChanged := false
		oagDestination := path.Join(params.SupportDir, "openapi-generator-cli.jar")
		if err = downloadSupportFileIfRequired(
			ctx,
			logger,
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
			logger,
			deps,
			downloadSupportFileIfRequiredParams{
				sourceLocation:   params.ServerGeneratorSourceLocation,
				destinationPath:  serverGeneratorDestination,
				sourceVersion:    params.AppVersion,
				metadataVersion:  &metadata.AppVersion,
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
			encoder := json.NewEncoder(metadataOutput)
			encoder.SetIndent("", "  ")
			if err = encoder.Encode(metadata); err != nil {
				return emptyResult, fmt.Errorf("failed to write metadata: %w", err)
			}
		}

		return SupportingFilesInstallResult{
			OagLocation:             oagDestination,
			ServerGeneratorLocation: serverGeneratorDestination,
		}, nil
	}
}
