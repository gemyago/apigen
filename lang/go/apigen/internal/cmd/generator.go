package cmd

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/context/ctxhttp"
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

type ResourceDownloader func(ctx context.Context, src, dst string) error

func NewResourceDownloader() ResourceDownloader {
	return func(ctx context.Context, src, dst string) error {
		var srcStream io.ReadCloser
		if strings.Index(src, "http") == 0 {
			res, err := ctxhttp.Get(ctx, http.DefaultClient, src)
			if err != nil {
				return fmt.Errorf("failed to download %s: %w", src, err)
			}
			defer res.Body.Close()
			srcStream = res.Body
		} else {
			if strings.Index(src, "file") == 0 {
				src = src[7:] // remove "file://"
			}
			file, err := os.Open(src)
			if err != nil {
				return fmt.Errorf("failed to open source file %s: %w", src, err)
			}
			defer file.Close()
			srcStream = file
		}
		dstFile, err := os.Create(dst)
		if err != nil {
			return fmt.Errorf("failed to create file %s: %w", dst, err)
		}
		defer dstFile.Close()
		if _, err = io.Copy(dstFile, srcStream); err != nil {
			return fmt.Errorf("failed to write to file %s: %w", dst, err)
		}
		return nil
	}
}

type GeneratorDeps struct {
	Downloader ResourceDownloader
}

func NewGenerator() Generator {
	return func(ctx context.Context, params GeneratorParams) error {
		fmt.Println("Generator is not implemented", params)
		return nil
	}
}
