package http

// Below will generate routes and models in separate packages.
// Useful in scenarios where you want to have
// application layer and api layer fully separated and prefer
// to have models on the application layer side.
//go:generate -command petstore-apigen go run github.com/gemyago/apigen server ./../../../../petstore.yaml --server-generator-location ../../../../../generators/go-apigen-server/target/server.jar --verbose
//go:generate petstore-apigen ./routes --global-property apis --model-package "github.com/gemyago/apigen/examples/petstore-server-app-layer-go/internal/app/models"
//go:generate petstore-apigen ../../app/models --global-property models
