package cmd

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"golang.org/x/net/context/ctxhttp"
)

type GeneratorParams struct {
	input             string
	output            string
	supportDir        string
	oagCliVersion     string
	oagCliLocation    string
	generatorVersion  string
	generatorLocation string
}

type Generator func(ctx context.Context, params GeneratorParams) error

type ResourceDownloader func(ctx context.Context, src, dst string) error

func NewResourceDownloader() ResourceDownloader {
	return func(ctx context.Context, src, dst string) error {
		res, err := ctxhttp.Get(ctx, http.DefaultClient, src)
		if err != nil {
			return fmt.Errorf("failed to download %s: %w", src, err)
		}
		defer res.Body.Close()
		// need to copy res body to dst file path
		dstFile, err := os.Create(dst)
		if err != nil {
			return fmt.Errorf("failed to create file %s: %w", dst, err)
		}
		defer dstFile.Close()
		if _, err = io.Copy(dstFile, res.Body); err != nil {
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
