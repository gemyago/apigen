package http

// Below will use server generator that was previously built in this project
// and referencing pre-built server generator module
// In real world scenario it can be just: //go:generate go run github.com/gemyago/apigen ./routes.yaml ./routes

//go:generate -command petstore-apigen go run ../../../../../ ./../../../../petstore.yaml --server-generator-location ../../../../../generators/go-apigen-server/target/server.jar
//go:generate petstore-apigen ./routes
