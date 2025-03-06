package http

// Below will use server generator that was previously built in this project.
// In real world scenario it can be just: apigen <path-to-openapi.yaml> <output-dir>
//go:generate -command petstore-apigen go run github.com/gemyago/apigen server ./../../../../petstore.yaml --server-generator-location ../../../../../generators/go-apigen-server/target/server.jar --verbose
//go:generate petstore-apigen ./routes --global-property apis --model-package "github.com/gemyago/apigen/examples/petstore-server-app-layer-go/internal/app/models"
//go:generate petstore-apigen ../../app/models --global-property models
