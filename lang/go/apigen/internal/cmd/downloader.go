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
