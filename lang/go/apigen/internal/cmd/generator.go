package cmd

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"regexp"
)

type GeneratorParams struct {
	input      string
	output     string
	chdir      string
	supportDir string

	// oagCliVersion is a version of the required openapi-generator-cli
	oagCliVersion string

	// oagCliLocation is a path to the source openapi-generator-cli jar file
	oagCliLocation string

	// appVersion is a version of the required generator plugin
	appVersion string

	// serverGeneratorLocation is a path to the source generator plugin jar file
	serverGeneratorLocation string
}

type Generator func(ctx context.Context, params GeneratorParams) error

type metadataReader interface {
	ReadOpenapiGeneratorCliVersion() (string, error)
	ReadAppVersion() (string, error)
}

type GeneratorDeps struct {
	RootLogger *slog.Logger
	SupportFilesInstaller
	MetadataReader metadataReader
	GeneratorInvoker
	OsChdirFunc func(dir string) error
}

func NewGenerator(deps GeneratorDeps) Generator {
	logger := deps.RootLogger.WithGroup("generator")
	semverPattern := regexp.MustCompile(`^(\d+\.\d+\.\d+)(.*)$`)
	return func(ctx context.Context, params GeneratorParams) error {
		if params.chdir != "" {
			if err := deps.OsChdirFunc(params.chdir); err != nil {
				return fmt.Errorf("failed to change working directory: %w", err)
			}
		}

		if params.supportDir == "" {
			params.supportDir = params.output + "/.apigen"
			pwd, _ := os.Getwd() // we don't care here if it fails, need it for logging purposes
			logger.InfoContext(ctx,
				"Using default support directory",
				slog.String("supportDir", params.supportDir),
				slog.String("pwd", pwd),
			)
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

		if params.appVersion == "" {
			ver, err := deps.MetadataReader.ReadAppVersion()
			if err != nil {
				return fmt.Errorf("failed to read generator version: %w", err)
			}
			// if ver follows semver, prefix it with "v"
			if semverPattern.MatchString(ver) {
				ver = "v" + ver
			}
			params.appVersion = ver
		}

		if params.serverGeneratorLocation == "" {
			params.serverGeneratorLocation = fmt.Sprintf(
				"https://github.com/gemyago/apigen/releases/download/%s/server.jar",
				params.appVersion,
			)
			logger.InfoContext(ctx,
				"Using default generator location",
				slog.String("generatorLocation", params.serverGeneratorLocation),
			)
		}

		installResult, err := deps.SupportFilesInstaller(ctx, SupportFilesInstallerParams{
			SupportDir:                    params.supportDir,
			OagSourceVersion:              params.oagCliVersion,
			OagSourceLocation:             params.oagCliLocation,
			AppVersion:                    params.appVersion,
			ServerGeneratorSourceLocation: params.serverGeneratorLocation,
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
