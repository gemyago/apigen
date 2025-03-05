package http

// Below will use server generator that was previously built in this project.
// In real world scenario it can be just: apigen <path-to-openapi.yaml> <output-dir>
//go:generate go run github.com/gemyago/apigen server ./../../../../ping.yaml ./routes --server-generator-location ../../../../../generators/go-apigen-server/target/server.jar
