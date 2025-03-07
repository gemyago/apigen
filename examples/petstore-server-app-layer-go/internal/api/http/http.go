package http

// Below will generate routes and models in separate packages.
// Useful in scenarios where you want to have
// application layer and api layer fully separated and prefer
// to have models on the application layer side.

//nolint:lll // hard to keep it short
// Generate directives here will use the server generator that was previously built in this project
// In real world scenario `go:generate` directives can be just:
// - go:generate go run github.com/gemyago/apigen server ./models.yaml ./models --global-property models
// - go:generate go run github.com/gemyago/apigen server ./routes.yaml ./routes --global-property apis --model-package "<full-models-package>"

//go:generate -command petstore-apigen go run github.com/gemyago/apigen server ./../../../../petstore.yaml --server-generator-location ../../../../../generators/go-apigen-server/target/server.jar --verbose
//go:generate petstore-apigen ../../app/models --global-property models
//go:generate petstore-apigen ./routes --global-property apis --model-package "github.com/gemyago/apigen/examples/petstore-server-app-layer-go/internal/app/models"
