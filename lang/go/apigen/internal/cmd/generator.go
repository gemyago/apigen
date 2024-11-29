package cmd

import (
	"context"
	"fmt"
	"log/slog"
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
}

type Generator func(ctx context.Context, params GeneratorParams) error

type metadataReader interface {
	ReadOpenapiGeneratorCliVersion() (string, error)
	ReadPluginVersion() (string, error)
}

type GeneratorDeps struct {
	RootLogger *slog.Logger
	SupportFilesInstaller
	MetadataReader metadataReader
	GeneratorInvoker
}

func NewGenerator(deps GeneratorDeps) Generator {
	logger := deps.RootLogger.WithGroup("generator")
	return func(ctx context.Context, params GeneratorParams) error {
		if params.supportDir == "" {
			params.supportDir = params.output + "/.apigen"
			logger.InfoContext(ctx, "Using default support directory", slog.String("supportDir", params.supportDir))
		}

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
			logger.InfoContext(ctx,
				"Using default openapi-generator-cli location",
				slog.String("oagCliLocation", params.oagCliLocation),
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
			logger.InfoContext(ctx,
				"Using default generator location",
				slog.String("generatorLocation", params.generatorLocation),
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
