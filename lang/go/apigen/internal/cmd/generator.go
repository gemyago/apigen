package cmd

import (
	"context"
	"errors"
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
	OsChdirFunc              func(dir string) error
	defaultSupportDirLocator func(output string) (string, error)
}

// The default support directory will be the one that has the nearest go.mod file, which is
// the root of the project.
func locateDefaultSupportDir(output string) (string, error) {
	return output + "/.apigen", nil
}

type generator struct {
	deps                     GeneratorDeps
	logger                   *slog.Logger
	defaultSupportDirLocator func(output string) (string, error)
	semverPattern            *regexp.Regexp
}

func (g generator) ensureSupportDir(ctx context.Context, output string, supportDir *string) error {
	if *supportDir == "" {
		val, err := g.defaultSupportDirLocator(output)
		if err != nil {
			return fmt.Errorf("failed to locate default support directory: %w", err)
		}
		*supportDir = val
		pwd, _ := os.Getwd() // we don't care here if it fails, need it for logging purposes
		g.logger.InfoContext(ctx,
			"Using default support directory",
			slog.String("supportDir", *supportDir),
			slog.String("pwd", pwd),
		)
	}
	return nil
}

func (g generator) ensureOagCliVersion(oagCliVersion *string) error {
	if *oagCliVersion == "" {
		ver, err := g.deps.MetadataReader.ReadOpenapiGeneratorCliVersion()
		if err != nil {
			return fmt.Errorf("failed to read openapi-generator-cli version: %w", err)
		}
		*oagCliVersion = ver
	}
	return nil
}

func (g generator) ensureOagCliLocation(ctx context.Context, oagCliVersion string, oagCliLocation *string) {
	if *oagCliLocation == "" {
		*oagCliLocation = fmt.Sprintf(
			"https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/%s/openapi-generator-cli-%s.jar",
			oagCliVersion,
			oagCliVersion,
		)
		g.logger.InfoContext(ctx,
			"Using default openapi-generator-cli location",
			slog.String("oagCliLocation", *oagCliLocation),
		)
	}
}

func (g generator) ensureAppVersion(appVersion *string) error {
	if *appVersion == "" {
		ver, err := g.deps.MetadataReader.ReadAppVersion()
		if err != nil {
			return fmt.Errorf("failed to read generator version: %w", err)
		}
		// if ver follows semver, prefix it with "v"
		if g.semverPattern.MatchString(ver) {
			ver = "v" + ver
		}
		*appVersion = ver
	}
	return nil
}

func (g generator) ensureServerGeneratorLocation(
	ctx context.Context,
	appVersion string,
	serverGeneratorLocation *string,
) {
	if *serverGeneratorLocation == "" {
		*serverGeneratorLocation = fmt.Sprintf(
			"https://github.com/gemyago/apigen/releases/download/%s/server.jar",
			appVersion,
		)
		g.logger.InfoContext(ctx,
			"Using default generator location",
			slog.String("generatorLocation", *serverGeneratorLocation),
		)
	}
}

func (g generator) invoke(ctx context.Context, params GeneratorParams) error {
	if params.chdir != "" {
		if err := g.deps.OsChdirFunc(params.chdir); err != nil {
			return fmt.Errorf("failed to change working directory: %w", err)
		}
	}

	if err := errors.Join(
		g.ensureSupportDir(ctx, params.output, &params.supportDir),
		g.ensureOagCliVersion(&params.oagCliVersion),
		g.ensureAppVersion(&params.appVersion),
	); err != nil {
		return err
	}

	g.ensureOagCliLocation(ctx, params.oagCliVersion, &params.oagCliLocation)
	g.ensureServerGeneratorLocation(ctx, params.appVersion, &params.serverGeneratorLocation)

	installResult, err := g.deps.SupportFilesInstaller(ctx, SupportFilesInstallerParams{
		SupportDir:                    params.supportDir,
		OagSourceVersion:              params.oagCliVersion,
		OagSourceLocation:             params.oagCliLocation,
		AppVersion:                    params.appVersion,
		ServerGeneratorSourceLocation: params.serverGeneratorLocation,
	})
	if err != nil {
		return fmt.Errorf("failed to install support files: %w", err)
	}

	if err = g.deps.GeneratorInvoker(ctx, GeneratorInvokerParams{
		Input:             params.input,
		Output:            params.output,
		OagCliLocation:    installResult.OagLocation,
		GeneratorLocation: installResult.ServerGeneratorLocation,
	}); err != nil {
		return fmt.Errorf("generator failed: %w", err)
	}

	return nil
}

func NewGenerator(deps GeneratorDeps) Generator {
	if deps.defaultSupportDirLocator == nil {
		deps.defaultSupportDirLocator = locateDefaultSupportDir
	}

	g := generator{
		deps:                     deps,
		logger:                   deps.RootLogger.WithGroup("generator"),
		defaultSupportDirLocator: deps.defaultSupportDirLocator,
		semverPattern:            regexp.MustCompile(`^(\d+\.\d+\.\d+)(.*)$`),
	}
	if g.defaultSupportDirLocator == nil {
		deps.defaultSupportDirLocator = g.defaultSupportDirLocator
	}
	return g.invoke
}
