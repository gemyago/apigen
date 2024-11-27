package cmd

import (
	"context"
	"fmt"
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

func NewGenerator() Generator {
	return func(ctx context.Context, params GeneratorParams) error {
		fmt.Println("Generator is not implemented", params)
		return nil
	}
}
