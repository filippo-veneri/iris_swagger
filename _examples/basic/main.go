package main

import (
	swagger "github.com/filippo-veneri/iris_swagger/v12"
	"github.com/kataras/iris/v12"

	"github.com/filippo-veneri/iris_swagger/v12/_examples/basic/api"
	"github.com/filippo-veneri/iris_swagger/v12/swaggerFiles"

	_ "github.com/filippo-veneri/iris_swagger/v12/_examples/basic/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /v2
func main() {
	app := iris.New()

	// To enable compression:
	// app.Use(iris.Compression)

	app.Get("/v2/testapi/get-string-by-int/{some_id:int}", api.GetStringByInt)
	app.Get("/v2/testapi/get-struct-array-by-string/{some_id:int}", api.GetStructArrayByString)

	// Configure the swagger UI page.
	swaggerUI := swagger.Handler(swaggerFiles.Handler,
		swagger.URL("/swagger/swagger.json"), // The url pointing to API definition.
		swagger.DeepLinking(true),
		swagger.Prefix("/swagger"), // Must match the registration prefix
	)
	// Register on http://localhost:8080/swagger
	app.Get("/swagger", swaggerUI)
	// And http://localhost:8080/swagger/index.html, *.js, *.css and e.t.c.
	app.Get("/swagger/{any:path}", swaggerUI)

	app.Listen(":8080")
}
