package cmd

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/gemyago/apigen/lang/go/apigen/internal/cmd/resources"
)

type GeneratorParams struct {
	input      string
	output     string
	supportDir string

	// oagCliVersion is a version of the required openapi-generator-cli
	oagCliVersion string

	// oagCliLocation is a path to the source openapi-generator-cli jar file
	oagCliLocation string

	// generatorVersion is a version of the required generator plugin
	generatorVersion string

	// generatorLocation is a path to the source generator plugin jar file
	generatorLocation string

	// verboseLogging enables logging of the generator
	verboseLogging bool
}

type Generator func(ctx context.Context, params GeneratorParams) error

type GeneratorDeps struct {
	RootLogger *slog.Logger
	SupportFilesInstaller
	MetadataReader *resources.MetadataReader
	GeneratorInvoker
	Downloader ResourceDownloader
}

func NewGenerator(deps GeneratorDeps) Generator {
	return func(ctx context.Context, params GeneratorParams) error {
		if params.oagCliVersion == "" {
			ver, err := deps.MetadataReader.ReadOpenapiGeneratorCliVersion()
			if err != nil {
				return fmt.Errorf("failed to read openapi-generator-cli version: %w", err)
			}
			params.oagCliVersion = ver
		}

		if params.oagCliLocation == "" {
			params.oagCliLocation = fmt.Sprintf(
				"https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/%s/openapi-generator-cli-%s.jar",
				params.oagCliVersion,
				params.oagCliVersion,
			)
		}

		if params.generatorVersion == "" {
			ver, err := deps.MetadataReader.ReadPluginVersion()
			if err != nil {
				return fmt.Errorf("failed to read generator version: %w", err)
			}
			params.generatorVersion = ver
		}

		if params.generatorLocation == "" {
			params.generatorLocation = fmt.Sprintf(
				"https://github.com/gemyago/apigen/releases/download/%s/server.jar",
				params.generatorVersion,
			)
		}

		installResult, err := deps.SupportFilesInstaller(ctx, SupportFilesInstallerParams{
			SupportDir:                    params.supportDir,
			OagSourceVersion:              params.oagCliVersion,
			OagSourceLocation:             params.oagCliLocation,
			ServerGeneratorSourceVersion:  params.generatorVersion,
			ServerGeneratorSourceLocation: params.generatorLocation,
		})
		if err != nil {
			return fmt.Errorf("failed to install support files: %w", err)
		}

		if err = deps.GeneratorInvoker(ctx, GeneratorInvokerParams{
			Input:             params.input,
			Output:            params.output,
			OagCliLocation:    installResult.OagLocation,
			GeneratorLocation: installResult.ServerGeneratorLocation,
		}); err != nil {
			return fmt.Errorf("generator failed: %w", err)
		}

		return nil
	}
}
